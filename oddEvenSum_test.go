package main

import (
	"testing"
)

func TestOddEvenSum_even(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		tr   OddEvenSum
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
			tr := OddEvenSum{}
			if got := tr.even(tt.args.val); got != tt.want {
				t.Errorf("OddEvenSum.even() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOddEvenSum_magic(t *testing.T) {
	type args struct {
		numberList []int
	}
	tests := []struct {
		name        string
		tr          *OddEvenSum
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
			tr := &OddEvenSum{}
			gotSumEven, gotSumOdd := tr.magic(tt.args.numberList)
			if gotSumEven != tt.wantSumEven {
				t.Errorf("OddEvenSum.magic() gotSumEven = %v, want %v", gotSumEven, tt.wantSumEven)
			}
			if gotSumOdd != tt.wantSumOdd {
				t.Errorf("OddEvenSum.magic() gotSumOdd = %v, want %v", gotSumOdd, tt.wantSumOdd)
			}
		})
	}
}

func TestOddEvenSum_Run(t *testing.T) {
	type args struct {
		numberList []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: "For given slice : \t[1 2 3 4 5 6 7 8 9 10]\n\tSum of evens: \t30\n\tSum of odds: \t25",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			tr := &OddEvenSum{}
			if got := tr.Run(tt.args.numberList); got != tt.want {
				t1.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
