package leetcode

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "subset does not fit",
			args: args{
				s: "large",
				t: "lar",
			},
			want: false,
		},
		{
			name: "straight subset",
			args: args{
				s: "lar",
				t: "large",
			},
			want: true,
		},
		{
			name: "not a straight subset",
			args: args{
				s: "lat",
				t: "large",
			},
			want: false,
		},
		{
			name: "skipping subset",
			args: args{
				s: "lre",
				t: "large",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
