package main

import "testing"
import "fmt"

//func TestAverage(t *testing.T) {
//	var v float64
//	v = Average([]float64{1, 2})
//	if v != 1.5 {
//		t.Error("Expected 1.5, got ", v)
//	}
//}

func BenchmarkAverage(b *testing.B) {

	//	var v float64
	//	v = Average([]float64{1, 2})
	//	if v != 1.5 {
	//		t.Error("Expected 1.5, got ", v)
	//	}
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
