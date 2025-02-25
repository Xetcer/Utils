package quicksort

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		name   string
		tstArr []int
		want   []int
	}{
		{name: "First test", tstArr: []int{0, -1, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, want: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sort(test.tstArr)
			if len(got) != len(test.want) {
				t.Fatalf("Sort() got len %d, want len %d", len(got), len(test.want))
			}
			for i := 0; i < len(got); i++ {
				if got[i] != test.want[i] {
					t.Errorf("Sort() want %d got %d", test.want[i], got[i])
				}
			}
		})
	}
}
