package src

import "testing"

func TestPrintAsciiArt(t *testing.T) {
	type args struct {
		inputString string
		substr      string
		color       string
		asciiMap    map[int][]string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Invalid string",args{"hello","hello world","red",map[int][]string{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintAsciiArt(tt.args.inputString, tt.args.substr, tt.args.color, tt.args.asciiMap)
		})
	}
}
