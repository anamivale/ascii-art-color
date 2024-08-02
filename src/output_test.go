package src

import "testing"

func TestHandleFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Valid name of output file ", args{"file.txt"}, "/home/fgitonga/Downloads/gitea/ascii-art-output/src/file.txt", false},
		{"Invalid name of output file", args{"file"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handlefile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Handlefile() = %v, want %v", got, tt.want)
			}
		})
	}
}
