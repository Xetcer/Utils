// go test -coverprofile=c
// go tool cover -html=c
package treesort

import (
	"fmt"
	"testing"
)

// {9,4,12,6,1} {1,4,6,9,12}
func TestSort(t *testing.T) {
	var testStruct = []struct {
		testName string
		initArr  []int
		wantArr  []int
	}{
		{testName: "First test", initArr: []int{9, 4, 12, 6, 1}, wantArr: []int{1, 4, 6, 9, 12}},
		{testName: "Second test", initArr: []int{9, 4, 12, 6, 1, 13}, wantArr: []int{1, 4, 6, 9, 12, 13}},
		{testName: "Third test",
			initArr: []int{8, 9, 7, 1, 12, 15, 10, 11, 6, 17, 5, 14, 13, 3, 4, 2, 18, 16},
			wantArr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}},
	}
	for _, test := range testStruct {
		t.Run(test.testName, func(t *testing.T) {
			got := Sort(test.initArr)
			if len(got) != len(test.wantArr) {
				t.Errorf("Sort([]arr), got len= %d, want len=%d", len(got), len(test.wantArr))
			}
			for i := 0; i < len(got); i++ {
				if got[i] != test.wantArr[i] {
					t.Errorf("Sort([]arr), Got %d, want %d", got[i], test.wantArr[i])
				}
			}
			fmt.Println("Got:", got, "Want:", test.wantArr)
		})

	}
}
