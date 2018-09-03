package arrayutil

import (
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"sum1", args{[]int{1, 2, 3}}, 6},
		{"sum2", args{[]int{2, 2, 3}}, 7},
		{"sum3", args{[]int{3, 2, 3}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"mul1", args{[]int{1, 2, 3}}, 6},
		{"mul2", args{[]int{2, 2, 3}}, 12},
		{"mul3", args{[]int{3, 2, 3}}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.args.a); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
