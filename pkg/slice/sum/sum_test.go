package sum

import "testing"

func TestSum(t *testing.T) {
	getStandartSum := func(arr []int) int {
		sum := 0
		for _, value := range arr {
			sum += value
		}
		return sum
	}
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{name: "Emptу slice", arr: make([]int, 0), want: 0},
		{name: "Slice len = 1", arr: []int{1}, want: 1},
		{name: "Slice len = 2", arr: []int{1, 2}, want: 3},
		{name: "Slice len > 2", arr: []int{1, 2, 3, 4, 5, 6}, want: getStandartSum([]int{1, 2, 3, 4, 5, 6})},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sum(test.arr)
			if got != test.want {
				t.Errorf("Sum() want %d, got %d\r\n", test.want, got)
			}
		})
	}

}
