package secret

import "testing"

func TestApplyMask(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "More than 8 characters",
			args: args{input: "1234567890"},
			want: "1234**7890",
		},
		{
			name: "Between 5 and 8 characters",
			args: args{input: "1234567"},
			want: "*****67",
		},
		{
			name: "Exactly 4 characters",
			args: args{input: "1234"},
			want: "****",
		},
		{
			name: "Fewer than 4 characters",
			args: args{input: "123"},
			want: "*",
		},
		{
			name: "Empty string",
			args: args{input: ""},
			want: "*",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyMask(tt.args.input); got != tt.want {
				t.Errorf("ApplyMask() = %v, want %v", got, tt.want)
			}
		})
	}
}
