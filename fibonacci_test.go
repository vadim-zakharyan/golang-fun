package main

import "testing"

func TestFibonacci_magic(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		tr   *Fibonacci
		args args
		want int
	}{
		{
			name: "Fibonacci shall return 9227465 for 35",
			args: args{number: 35},
			want: 9227465,
		},
		{
			name: "Fibonacci shall return 1 for 1",
			args: args{number: 1},
			want: 1,
		},
		{
			name: "Fibonacci shall return 1 for 2",
			args: args{number: 2},
			want: 1,
		},
		{
			name: "Fibonacci shall return 2 for 3",
			args: args{number: 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Fibonacci{}
			if got := tr.magic(tt.args.number); got != tt.want {
				t.Errorf("Fibonacci.magic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacci_Run(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Fibonacci prints correct output",
			args: args{number: 35},
			want: "Fibonacci for 35:\t9227465",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Fibonacci{}
			if got := tr.Run(tt.args.number); got != tt.want {
				t.Errorf("Fibonacci.Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
