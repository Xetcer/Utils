package deikstra

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDeikstra(t *testing.T) {
	testCases := []struct {
		name         string
		graph        Graph
		start        string
		end          string
		expectedPath []string
		expectedCost int
	}{
		{
			name: "Базовый случай",
			graph: Graph{
				"A": Node{"B": 6, "D": 1},
				"B": Node{"A": 6, "C": 5, "D": 2, "E": 2},
				"C": Node{"B": 5, "E": 5},
				"D": Node{"A": 1, "B": 2, "E": 1},
				"E": Node{"B": 2, "C": 5, "D": 1},
			},
			start:        "A",
			end:          "C",
			expectedPath: []string{"A", "D", "E", "C"},
			expectedCost: 7,
		},
		{
			name: "Без прямого пути",
			graph: Graph{
				"S": Node{"T": 10, "Y": 5},
				"T": Node{"X": 1, "Y": 2},
				"Y": Node{"T": 3, "X": 9, "Z": 2},
				"Z": Node{"S": 7, "X": 6},
				"X": Node{"Z": 4},
			},
			start:        "S",
			end:          "X",
			expectedPath: []string{"S", "Y", "T", "X"},
			expectedCost: 9, // Ошибка в изначальном тесте!
		},
		{
			name: "Минимальное кольцо",
			graph: Graph{
				"V": Node{"W": 10},
				"W": Node{"V": 10, "X": 1},
				"X": Node{"W": 1, "Y": 1},
				"Y": Node{"X": 1, "Z": 1},
				"Z": Node{"Y": 1, "V": 1},
			},
			start:        "V",
			end:          "Z",
			expectedPath: []string{"V", "W", "X", "Y", "Z"},
			expectedCost: 13,
		},
		{
			name: "Прямой путь отсутствует",
			graph: Graph{
				"P": Node{"Q": 1},
				"Q": Node{},
			},
			start:        "P",
			end:          "Q",
			expectedPath: []string{"P", "Q"},
			expectedCost: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path, cost := Deikstra(tc.graph, tc.start, tc.end)
			assert.Equal(t, path, tc.expectedPath, "Неверный путь")
			assert.Equal(t, cost, tc.expectedCost, "Неверная стоимость пути")
		})
	}
}
