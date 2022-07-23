package arithmetic

import "testing"

type TestCase[T Val] struct {
	arg1 T
	arg2 T
	want T
}

func TestAdd(t *testing.T) {

	cases := []TestCase[int]{
		{2, 3, 5},
		{5, 5, 10},
		{-7, 6, -1},
	}

	for _, tc := range cases {
		got := Add(tc.arg1, tc.arg2)
		if tc.want != got {
			t.Errorf("TestAdd: Expected '%d', but got '%d'", tc.want, got)
		}
	}
}

func TestDiv(t *testing.T) {

	cases := []TestCase[int]{
		{6., 3., 2.},
		{5., 5., 1.},
		{-10., 2., -5.},
	}

	for _, tc := range cases {
		got := Div(tc.arg1, tc.arg2)
		if tc.want != got {
			t.Errorf("TestDiv: Expected '%d', but got '%d'", tc.want, got)
		}
	}
}

func TestMul(t *testing.T) {

	cases := []TestCase[int]{
		{7, 3, 21},
		{5, 5, 25},
		{-1, 6, -6},
	}

	for _, tc := range cases {
		got := Mul(tc.arg1, tc.arg2)
		if tc.want != got {
			t.Errorf("TestMul: Expected '%d', but got '%d'", tc.want, got)
		}
	}
}

func TestSub(t *testing.T) {

	cases := []TestCase[int]{
		{2, 3, -1},
		{5, 5, 0},
		{-7, -3, -4},
	}

	for _, tc := range cases {
		got := Sub(tc.arg1, tc.arg2)
		if tc.want != got {
			t.Errorf("TestSub: Expected '%d', but got '%d'", tc.want, got)
		}
	}
}
