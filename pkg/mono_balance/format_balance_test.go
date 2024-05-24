package mono_balance

import "testing"

func Test_formatBalance(t *testing.T) {
	type args struct {
		balance int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Standard balance",
			args: args{balance: 123456},
			want: "1234,56",
		},
		{
			name: "Two digits balance",
			args: args{balance: 56},
			want: "0,56",
		},
		{
			name: "One digit balance",
			args: args{balance: 5},
			want: "0,05",
		},
		{
			name: "Zero balance",
			args: args{balance: 0},
			want: "0,00",
		},
		{
			name: "High balance",
			args: args{balance: 1000000},
			want: "10000,00",
		},
		{
			name: "Three digit balance",
			args: args{balance: 123},
			want: "1,23",
		},
		{
			name: "Negative balance",
			args: args{balance: -12345},
			want: "-123,45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatBalance(tt.args.balance); got != tt.want {
				t.Errorf("formatBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
