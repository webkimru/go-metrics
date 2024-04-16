package utils

import "testing"

func TestGetFloat64ValueFromSting(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    float64
		wantErr bool
	}{
		{"valid float", "1.23", 1.23, false},
		{"invalid float", "none", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFloat64ValueFromSting(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFloat64ValueFromSting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFloat64ValueFromSting() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt64ValueFromSting(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    int64
		wantErr bool
	}{
		{"valid int", "123", 123, false},
		{"invalid int", "1.23", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInt64ValueFromSting(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt64ValueFromSting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt64ValueFromSting() got = %v, want %v", got, tt.want)
			}
		})
	}
}
