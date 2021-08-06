package main

import "testing"

func TestHasDuplicate_hasDuplicate(t1 *testing.T) {
	type args struct {
		someArray []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check if array doesn't contain duplicates",
			args: args{[]int{1, 2, 3}},
			want: false,
		},
		{
			name: "check if array contains duplicates",
			args: args{[]int{1, 2, 2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &HasDuplicate{}
			if got := t.hasDuplicate(tt.args.someArray); got != tt.want {
				t1.Errorf("magic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasDuplicate_contains(t1 *testing.T) {
	type args struct {
		someArray   []int
		searchValue int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check if array doesn't contain value",
			args: args{[]int{1, 2, 3}, 4},
			want: false,
		},
		{
			name: "check if array contains value",
			args: args{[]int{1, 2, 3}, 3},
			want: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &HasDuplicate{}
			if got := t.contains(tt.args.someArray, tt.args.searchValue); got != tt.want {
				t1.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasDuplicate_Run(t1 *testing.T) {
	type args struct {
		someArray []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check if Run returns 'array [1 2 2] has duplicates'",
			args: args{[]int{1, 2, 2}},
			want: "array [1 2 2] has duplicates",
		},
		{
			name: "check if Run returns 'array [1 2 3] has not duplicates'",
			args: args{[]int{1, 2, 3}},
			want: "array [1 2 3] has not duplicates",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &HasDuplicate{}
			if got := t.Run(tt.args.someArray); got != tt.want {
				t1.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
