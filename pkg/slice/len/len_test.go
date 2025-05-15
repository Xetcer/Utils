package len

import "testing"

func TestLen(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{name: "Empty array", arr: make([]int, 0), want: 0},
		{name: "Array len = 1", arr: []int{1}, want: 1},
		{name: "Array len = 2", arr: []int{1, 2}, want: 2},
		{name: "Array len > 2", arr: []int{1, 2, 3, 4, 5, 6}, want: len([]int{1, 2, 3, 4, 5, 6})},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Len(test.arr)
			if got != test.want {
				t.Errorf("Len() want %d, got %d\r\n", test.want, got)
			}
		})
	}
}
