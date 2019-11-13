package giphy

import "testing"

func TestGetRandom(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		have string
		want error
	}{
		{name: "withParam", have: "hello", want: nil},
		{name: "withoutParam", have: "", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetRandom(tt.have); err != tt.want {
				t.Errorf("GetRandom() = %v, want %v", err, tt.want)
			}
		})
	}
}
