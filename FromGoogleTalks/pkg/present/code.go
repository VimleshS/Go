// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package present

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Is the playground available?
var PlayEnabled = false

// TOOD(adg): replace the PlayEnabled flag with something less spaghetti-like.
// Instead this will probably be determined by a template execution Context
// value that contains various global metadata required when rendering
// templates.

func init() {
	Register("code", parseCode)
	Register("play", parseCode)
}

type Code struct {
	Text template.HTML
	Play bool // runnable code
}

func (c Code) TemplateName() string { return "code" }

// The input line is a .code or .play entry with a file name and an optional HLfoo marker on the end.
// Anything between the file and HL (if any) is an address expression, which we treat as a string here.
// We pick off the HL first, for easy parsing.
var (
	highlightRE = regexp.MustCompile(`\s+HL([a-zA-Z0-9_]+)?$`)
	hlCommentRE = regexp.MustCompile(`(.+) // HL(.*)$`)
	codeRE      = regexp.MustCompile(`\.(code|play)\s+([^\s]+)(\s+)?(.*)?$`)
)

func parseCode(ctx *Context, sourceFile string, sourceLine int, cmd string) (Elem, error) {
	cmd = strings.TrimSpace(cmd)

	// Pull off the HL, if any, from the end of the input line.
	highlight := ""
	if hl := highlightRE.FindStringSubmatchIndex(cmd); len(hl) == 4 {
		highlight = cmd[hl[2]:hl[3]]
		cmd = cmd[:hl[2]-2]
	}

	// Parse the remaining command line.
	// Arguments:
	// args[0]: whole match
	// args[1]:  .code/.play
	// args[2]: file name
	// args[3]: space, if any, before optional address
	// args[4]: optional address
	args := codeRE.FindStringSubmatch(cmd)
	if len(args) != 5 {
		return nil, fmt.Errorf("%s:%d: syntax error for .code/.play invocation", sourceFile, sourceLine)
	}
	command, file, addr := args[1], args[2], strings.TrimSpace(args[4])
	play := command == "play" && PlayEnabled

	// Read in code file and (optionally) match address.
	filename := filepath.Join(filepath.Dir(sourceFile), file)
	textBytes, err := ctx.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s:%d: %v", sourceFile, sourceLine, err)
	}
	lo, hi, err := addrToByteRange(addr, 0, textBytes)
	if err != nil {
		return nil, fmt.Errorf("%s:%d: %v", sourceFile, sourceLine, err)
	}

	// Acme pattern matches can stop mid-line,
	// so run to end of line in both directions if not at line start/end.
	for lo > 0 && textBytes[lo-1] != '\n' {
		lo--
	}
	if hi > 0 {
		for hi < len(textBytes) && textBytes[hi-1] != '\n' {
			hi++
		}
	}

	lines := codeLines(textBytes, lo, hi)

	for i, line := range lines {
		// Replace tabs by spaces, which work better in HTML.
		line.L = strings.Replace(line.L, "\t", "    ", -1)

		// Highlight lines that end with "// HL[highlight]"
		// and strip the magic comment.
		if m := hlCommentRE.FindStringSubmatch(line.L); m != nil {
			line.L = m[1]
			line.HL = m[2] == highlight
		}

		lines[i] = line
	}

	data := &codeTemplateData{Lines: lines}

	// Include before and after in a hidden span for playground code.
	if play {
		data.Prefix = textBytes[:lo]
		data.Suffix = textBytes[hi:]
	}

	var buf bytes.Buffer
	if err := codeTemplate.Execute(&buf, data); err != nil {
		return nil, err
	}
	return Code{Text: template.HTML(buf.String()), Play: play}, nil
}

type codeTemplateData struct {
	Lines          []codeLine
	Prefix, Suffix []byte
}

var leadingSpaceRE = regexp.MustCompile(`^[ \t]*`)

var codeTemplate = template.Must(template.New("code").Funcs(template.FuncMap{
	"trimSpace":    strings.TrimSpace,
	"leadingSpace": leadingSpaceRE.FindString,
}).Parse(codeTemplateHTML))

const codeTemplateHTML = `
{{with .Prefix}}<pre style="display: none"><span>{{printf "%s" .}}</span></pre>{{end}}

<pre>{{range .Lines}}<span num="{{.N}}">{{/*
	*/}}{{if .HL}}{{leadingSpace .L}}<b>{{trimSpace .L}}</b>{{/*
	*/}}{{else}}{{.L}}{{end}}{{/*
*/}}</span>
{{end}}</pre>

{{with .Suffix}}<pre style="display: none"><span>{{printf "%s" .}}</span></pre>{{end}}
`

// codeLine represents a line of code extracted from a source file.
type codeLine struct {
	L  string // The line of code.
	N  int    // The line number from the source file.
	HL bool   // Whether the line should be highlighted.
}

// codeLines takes a source file and returns the lines that
// span the byte range specified by start and end.
// It discards lines that end in "OMIT".
func codeLines(src []byte, start, end int) (lines []codeLine) {
	startLine := 1
	for i, b := range src {
		if i == start {
			break
		}
		if b == '\n' {
			startLine++
		}
	}
	s := bufio.NewScanner(bytes.NewReader(src[start:end]))
	for n := startLine; s.Scan(); n++ {
		l := s.Text()
		if strings.HasSuffix(l, "OMIT") {
			continue
		}
		lines = append(lines, codeLine{L: l, N: n})
	}
	return
}

func parseArgs(name string, line int, args []string) (res []interface{}, err error) {
	res = make([]interface{}, len(args))
	for i, v := range args {
		if len(v) == 0 {
			return nil, fmt.Errorf("%s:%d bad code argument %q", name, line, v)
		}
		switch v[0] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("%s:%d bad code argument %q", name, line, v)
			}
			res[i] = n
		case '/':
			if len(v) < 2 || v[len(v)-1] != '/' {
				return nil, fmt.Errorf("%s:%d bad code argument %q", name, line, v)
			}
			res[i] = v
		case '$':
			res[i] = "$"
		default:
			return nil, fmt.Errorf("%s:%d bad code argument %q", name, line, v)
		}
	}
	return
}

// parseArg returns the integer or string value of the argument and tells which it is.
func parseArg(arg interface{}, max int) (ival int, sval string, isInt bool, err error) {
	switch n := arg.(type) {
	case int:
		if n <= 0 || n > max {
			return 0, "", false, fmt.Errorf("%d is out of range", n)
		}
		return n, "", true, nil
	case string:
		return 0, n, false, nil
	}
	return 0, "", false, fmt.Errorf("unrecognized argument %v type %T", arg, arg)
}

// oneLine returns the single line generated by a two-argument code invocation.
func oneLine(ctx *Context, file, text string, arg interface{}) (line, before, after string, err error) {
	contentBytes, err := ctx.ReadFile(file)
	if err != nil {
		return "", "", "", err
	}
	lines := strings.SplitAfter(string(contentBytes), "\n")
	lineNum, pattern, isInt, err := parseArg(arg, len(lines))
	if err != nil {
		return "", "", "", err
	}
	var n int
	if isInt {
		n = lineNum - 1
	} else {
		n, err = match(file, 0, lines, pattern)
		n -= 1
	}
	if err != nil {
		return "", "", "", err
	}
	return lines[n],
		strings.Join(lines[:n], ""),
		strings.Join(lines[n+1:], ""),
		nil
}

// multipleLines returns the text generated by a three-argument code invocation.
func multipleLines(ctx *Context, file string, arg1, arg2 interface{}) (line, before, after string, err error) {
	contentBytes, err := ctx.ReadFile(file)
	lines := strings.SplitAfter(string(contentBytes), "\n")
	if err != nil {
		return "", "", "", err
	}
	line1, pattern1, isInt1, err := parseArg(arg1, len(lines))
	if err != nil {
		return "", "", "", err
	}
	line2, pattern2, isInt2, err := parseArg(arg2, len(lines))
	if err != nil {
		return "", "", "", err
	}
	if !isInt1 {
		line1, err = match(file, 0, lines, pattern1)
	}
	if !isInt2 {
		line2, err = match(file, line1, lines, pattern2)
	} else if line2 < line1 {
		return "", "", "", fmt.Errorf("lines out of order for %q: %d %d", file, line1, line2)
	}
	if err != nil {
		return "", "", "", err
	}
	for k := line1 - 1; k < line2; k++ {
		if strings.HasSuffix(lines[k], "OMIT\n") {
			lines[k] = ""
		}
	}
	return strings.Join(lines[line1-1:line2], ""),
		strings.Join(lines[:line1-1], ""),
		strings.Join(lines[line2:], ""),
		nil
}

// match identifies the input line that matches the pattern in a code invocation.
// If start>0, match lines starting there rather than at the beginning.
// The return value is 1-indexed.
func match(file string, start int, lines []string, pattern string) (int, error) {
	// $ matches the end of the file.
	if pattern == "$" {
		if len(lines) == 0 {
			return 0, fmt.Errorf("%q: empty file", file)
		}
		return len(lines), nil
	}
	// /regexp/ matches the line that matches the regexp.
	if len(pattern) > 2 && pattern[0] == '/' && pattern[len(pattern)-1] == '/' {
		re, err := regexp.Compile(pattern[1 : len(pattern)-1])
		if err != nil {
			return 0, err
		}
		for i := start; i < len(lines); i++ {
			if re.MatchString(lines[i]) {
				return i + 1, nil
			}
		}
		return 0, fmt.Errorf("%s: no match for %#q", file, pattern)
	}
	return 0, fmt.Errorf("unrecognized pattern: %q", pattern)
}
