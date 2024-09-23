package reco

import "testing"

func TestIsMobile(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid",
			args: args{
				s: "ssssssss",
			},
			want: false,
		},
		{
			name: "valid",
			args: args{
				s: "15699887762",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMobile(tt.args.s); got != tt.want {
				t.Errorf("IsMobile() = %v, want %v", got, tt.want)
			}
		})
	}
}
