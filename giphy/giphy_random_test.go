package giphy

import "testing"

func TestGetRandomGif(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		have string
		want error
	}{
		{name: "A", have: "hello", want: nil},
		{name: "B", have: "", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetRandomGif(tt.have); err != tt.want {
				t.Errorf("GetRandomGif() = %v, want %v", err, tt.want)
			}
		})
	}
}
