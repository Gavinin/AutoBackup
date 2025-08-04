package util

import "testing"

func TestNameFormat2DateFormat(t *testing.T) {
	type args struct {
		nameFormat string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "DateTime",
			args: args{
				nameFormat: "%Y%m%D_%H%M%S",
			},
			want: "20060102_150405",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NameFormat2DateFormat(tt.args.nameFormat); got != tt.want {
				t.Errorf("NameFormat2DateFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
