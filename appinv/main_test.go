package main

import (
	"os"
	"testing"
)

// func Test_main(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 		{name: "test"},
// 	}
// 	os.Args = []string{"-V disable", "e"}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			main()
// 		})
// 	}
// }

func BenchmarkApp(b *testing.B) {
	b.ResetTimer()
	var i uint32
	os.Args = []string{"-V disable", "e"}
	for i = 0; i < uint32(b.N); i++ {
		main()
	}
}
