package bfs

import (
	"testing"
)

func TestFindSellerBFS(t *testing.T) {
	pesonsMap := make(Persons)
	pesonsMap["firstNode"] = []string{"Alise", "Bob", "Kler"}
	pesonsMap["Bob"] = []string{"Anydg", "Peggy"}
	pesonsMap["Alise"] = []string{"Peggy"}
	pesonsMap["Kler"] = []string{"Tom", "Gonny"}
	pesonsMap["Anydg"] = []string{}
	pesonsMap["Peggy"] = []string{}
	pesonsMap["Tom"] = []string{}
	pesonsMap["Gonny"] = []string{}

	personLessMap := make(Persons)
	personLessMap["firstNode"] = []string{"Alise", "Bob"}
	personLessMap["Bob"] = []string{"Anydg", "Peggy"}
	personLessMap["Alise"] = []string{"Peggy"}
	personLessMap["Anydg"] = []string{}
	personLessMap["Peggy"] = []string{}
	tests := []struct {
		name string
		want struct {
			name   string
			result bool
		}
		pesrons Persons
	}{
		{name: "empty map", want: struct {
			name   string
			result bool
		}{name: "", result: false}, pesrons: make(Persons)},
		{
			name: "full map, with person", want: struct {
				name   string
				result bool
			}{
				name:   "Tom",
				result: true,
			}, pesrons: pesonsMap,
		},
		{
			name: "full map, withot person", want: struct {
				name   string
				result bool
			}{
				name:   "",
				result: false,
			}, pesrons: personLessMap,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got struct {
				name   string
				result bool
			} = struct {
				name   string
				result bool
			}{}
			got.name, got.result = FindSellerBFS(test.pesrons, "firstNode")
			if got.name != test.want.name || got.result != test.want.result {
				t.Errorf("FindSellerBFS() want %v, got %v\r\n", test.want, got)
			}
		})
	}
}
