package internal

import (
	"os"
	"testing"
)

func TestGetStencilPath(t *testing.T) {
	tests := []struct {
		name    string
		prep    func()
		want    string
		wantErr bool
	}{
		{
			name: "Case 1 - Stencil path is set",
			prep: func() {
				os.Setenv("STENCIL", "/home/$USER/.stencil/")
			},
			want:    "/home/$USER/.stencil/",
			wantErr: false,
		},
		{
			name: "Case 2 - Stencil path is not set",
			prep: func() {
				os.Setenv("STENCIL", "")
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.prep()

			got, err := GetStencilPath()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetStencilPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("GetStencilPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
