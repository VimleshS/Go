// +build OMIT

package main

import (
	"fmt"
	"net/http"
)

//http://stackoverflow.com/questions/28699582/how-get-multiselect-values-from-form-using-golang

func myHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Form submitted
		r.ParseForm() // Required if you don't call r.FormValue()
		fmt.Println(r.Form["new_data"])
	}
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":9090", nil)
}

// endmain OMIT

const html = `
<html><body>
<form action="process" method="post">
    <select id="new_data" name="new_data" class="tag-select chzn-done" multiple="" >
        <option value="1">111mm1</option>
        <option value="2">222mm2</option>
        <option value="3">012nx1</option>
    </select>
    <input type="Submit" value="Send" />
</form>
</body></html>
`
