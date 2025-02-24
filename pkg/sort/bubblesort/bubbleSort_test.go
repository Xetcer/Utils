package bubblesort

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		name   string
		tstArr []int
		want   []int
	}{
		{name: "First test", tstArr: []int{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Sort(test.tstArr)
			for i, value := range test.tstArr {
				if test.want[i] != value {
					t.Errorf("Sort(), want %d, got %d", test.want[i], value)
				}
			}
		})
	}
}
