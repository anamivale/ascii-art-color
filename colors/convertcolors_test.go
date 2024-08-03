package colors

import "testing"

func TestPrintColored(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		want2   int
		wantErr bool
	}{
		{"Valid color name",args{"red"},255,0,0,false},
		{"Invalid color name",args{"redy"},0,0,0,true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := PrintColored(tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintColored() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrintColored() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PrintColored() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("PrintColored() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestHexTo256(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		want2   int
		wantErr bool
	}{
		{"Invalid color",args{"#09999999"},0,0,0,true},
		{"Valid color",args{"#f00000"},240,0,0,false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := HexTo256(tt.args.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexTo256() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HexTo256() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HexTo256() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("HexTo256() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestRgbTo256ColorCode(t *testing.T) {
	type args struct {
		rgb string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		want2   int
		wantErr bool
	}{
		{"Valid RGB input",args{"rgb(255,0,0)"},255,0,0,false},
		{"Invalid RGB input",args{"rgb(260,0,0)"},0,0,0,true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := RgbTo256ColorCode(tt.args.rgb)
			if (err != nil) != tt.wantErr {
				t.Errorf("RgbTo256ColorCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RgbTo256ColorCode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RgbTo256ColorCode() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("RgbTo256ColorCode() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
