package main

import (
	"testing"
)

func Test_even(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Check if 2 is even",
			args: args{val: 2},
			want: true,
		},
		{
			name: "Check if 1 is not even",
			args: args{val: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := even(tt.args.val); got != tt.want {
				t.Errorf("even() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagic(t *testing.T) {
	type args struct {
		numberList []int
	}
	tests := []struct {
		name        string
		args        args
		wantSumEven int
		wantSumOdd  int
	}{
		{
			name:        "Check if 1..10 return even=30 and odd=25",
			args:        args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			wantSumEven: 30,
			wantSumOdd:  25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSumEven, gotSumOdd := magic(tt.args.numberList)
			if gotSumEven != tt.wantSumEven {
				t.Errorf("Magic() gotSumEven = %v, want %v", gotSumEven, tt.wantSumEven)
			}
			if gotSumOdd != tt.wantSumOdd {
				t.Errorf("Magic() gotSumOdd = %v, want %v", gotSumOdd, tt.wantSumOdd)
			}
		})
	}
}
