package util

import "testing"

func TestSeparatePath(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name:  "test1",
			args:  args{s: "/aaa/bbb/asts.tar.gz"},
			want:  "/aaa/bbb",
			want1: "asts.tar.gz",
		}, {
			name:  "test2",
			args:  args{s: "/aaa/bbb/asts"},
			want:  "/aaa/bbb",
			want1: "asts",
		}, {
			name:  "test2",
			args:  args{s: "/aaa/bbb/asts/"},
			want:  "/aaa/bbb",
			want1: "asts",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SeparatePath(tt.args.s)
			if got != tt.want {
				t.Errorf("SeparatePath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SeparatePath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
