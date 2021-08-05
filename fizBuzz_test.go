package main

import (
	"reflect"
	"testing"
)

func TestFizBuzz_magic(t1 *testing.T) {
	type args struct {
		numberList []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{[]int{2, 3, 5, 15}},
			want: []string{"2", "Fizz", "Buzz", "FizzBuzz"},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &FizBuzz{}
			if gotResult := t.magic(tt.args.numberList); !reflect.DeepEqual(gotResult, tt.want) {
				t1.Errorf("magic() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}

func TestFizBuzz_Run(t1 *testing.T) {
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
			args: args{[]int{2, 3, 5, 15}},
			want: "2, Fizz, Buzz, FizzBuzz",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &FizBuzz{}
			if gotOut := t.Run(tt.args.numberList); gotOut != tt.want {
				t1.Errorf("Run() = %v, want %v", gotOut, tt.want)
			}
		})
	}
}
