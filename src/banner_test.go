package src

import (
	"reflect"
	"testing"
)

func TestGetBanner(t *testing.T) {
	type args struct {
		bannerName string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"Empty invalid name", args{"file"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBanner(tt.args.bannerName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBanner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBannerContent(t *testing.T) {
	type args struct {
		file []string
	}
	tests := []struct {
		name string
		args args
		want map[int][]string
	}{
		{
			"test simple slice string",
			args{[]string{"  _   | |  | |  | |  |_|  (_)            "}},
			map[int][]string{32: {"  _   | |  | |  | |  |_|  (_)            "}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadBannerContent(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBannerContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
