package test

import (
    "fmt"
    "testing"
)


type Data struct {
	a string 
	b int
}

func Benchmark_Dev(b *testing.B) {
		data := []Data {}
		for i := 0; i < 1000; i++ {
			data = append(data, Data{a: fmt.Sprintf("data-%d", i), b: i+100})
		}

    fmt.Println("Benchmark_Dev")
    
    b.ResetTimer()
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        m := make(map[string]*Data, len(data))
        for _, d := range data {
        	//d := data[i]
        	//m[data[i].a] = &data[i]
        	m[d.a] = &d
        }
    }
}

func Benchmark_Mem(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s := "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"
		a := []byte(s)
		// If line below is uncommented it doubles the memory usage.
		//void(string(a))
		void(a)
	}
}


func void(interface{}) {}