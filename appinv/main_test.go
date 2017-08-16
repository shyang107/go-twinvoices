package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
	fmt.Printf("os.Args: %#v\n", os.Args[1:])
}

func BenchmarkApp(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	var i uint32
	for i = 0; i < uint32(b.N); i++ {
		_ = time.Now().Format("")
		main()
	}
	fmt.Printf("os.Args: %#v\n", os.Args[1:])
}
