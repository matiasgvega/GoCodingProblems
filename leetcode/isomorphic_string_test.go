package leetcode

import (
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
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
			name: "simple both ways isomomorphic",
			args: args{s: "egg", t: "add"},
			want: true,
		},
		{
			name: "simple one way not isomorphic",
			args: args{s: "foo", t: "bar"},
			want: false,
		},
		{
			name: "simple one way not isomorphic reverse",
			args: args{s: "foo", t: "bar"},
			want: false,
		},
		{
			name: "more complex isomorphic",
			args: args{s: "paper", t: "title"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("strings has different lenght", func(t *testing.T) {
		got := isIsomorphic("uno", "1")
		want := false
		if got != want {
			t.Errorf("isIsomorphic() = %v, want %v", got, want)
		}
	})

	t.Run("any string is isomorphic with itself", func(t *testing.T) {
		got := isIsomorphic("uno", "uno")
		want := true
		if got != want {
			t.Errorf("isIsomorphic() = %v, want %v", got, want)
		}
	})

	oneWayIsomorphicTest := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple one way not isomorphic",
			args: args{s: "foo", t: "bar"},
			want: false,
		},
		{
			name: "simple one way not isomorphic reverse",
			args: args{s: "bar", t: "foo"},
			want: true,
		},
	}
	for _, tt := range oneWayIsomorphicTest {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneWayIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
