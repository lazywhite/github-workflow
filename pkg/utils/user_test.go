package utils

import "testing"

func TestGetUserName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "bob",
			want: "bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserName(); got != tt.want {
				t.Errorf("GetUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
