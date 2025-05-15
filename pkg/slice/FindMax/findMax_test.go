package findmax

import "testing"

func TestFindMax(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{name: "Empty arr", arr: make([]int, 0), want: 0},
		{name: "Arr len = 1", arr: []int{1}, want: 1},
		{name: "Arr len = 2", arr: []int{1, 2}, want: 2},
		{name: "Arr len > 2", arr: []int{1, 2, 3, 4, 5}, want: 5},
		{name: "Arr len > 2, similar values", arr: []int{1, 1, 1, 1, 1}, want: 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FindMax(test.arr)
			if got != test.want {
				t.Errorf("FindMax() want %d, got %d\r\n", test.want, got)
			}
		})
	}
}
