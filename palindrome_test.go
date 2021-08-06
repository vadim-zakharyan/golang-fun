package main

import "testing"

func TestPalindrome_magic(t1 *testing.T) {
	type args struct {
		leftString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Check if palindrome magic works from 123 to 12321",
			args: args{leftString: "123"},
			want: "12321",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Palindrome{}
			if gotSymmetric := t.magic(tt.args.leftString); gotSymmetric != tt.want {
				t1.Errorf("magic() = %v, want %v", gotSymmetric, tt.want)
			}
		})
	}
}

func TestPalindrome_Run(t1 *testing.T) {
	type args struct {
		leftString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Check if Run returns formatted string",
			args: args{leftString: "123"},
			want: "Palindrome for '123':\t'12321'",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Palindrome{}
			if gotOut := t.Run(tt.args.leftString); gotOut != tt.want {
				t1.Errorf("Run() = %v, want %v", gotOut, tt.want)
			}
		})
	}
}
