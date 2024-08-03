package src

import "testing"

func TestReadInput(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Valid input",args{"hello"},"hello",false},
		{"Invalid input",args{"hello✊"},"",true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadInput(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleArgs(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		want1   string
		want2   string
		want3   string
		wantErr bool
	}{
		{"","","","","",true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, err := HandleArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HandleArgs() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HandleArgs() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("HandleArgs() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("HandleArgs() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
