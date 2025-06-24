package greedyalgs

import (
	"sort"
	"strings"
	"testing"
)

// Функция сравнения срезов строк без учета порядка
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sortedA := make([]string, len(a))
	copy(sortedA, a)
	sort.Strings(sortedA)

	sortedB := make([]string, len(b))
	copy(sortedB, b)
	sort.Strings(sortedB)

	return strings.Join(sortedA, ",") == strings.Join(sortedB, ",")
}

// Тестовые кейсы
var testCases = []struct {
	name         string
	statesNeeded []string
	stations     map[string][]string
	expected     []string
}{
	{
		name:         "Базовый сценарий",
		statesNeeded: []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
		stations: map[string][]string{
			"kone":   {"id", "nv", "ut"},
			"ktwo":   {"wa", "id", "mt"},
			"kthree": {"or", "nv", "ca"},
			"kfour":  {"nv", "ut"},
			"kfive":  {"ca", "az"},
		},
		expected: []string{"kone", "ktwo", "kthree", "kfive"}, // Оптимальный набор станций
	},
	{
		name:         "Один штат",
		statesNeeded: []string{"mt"},
		stations: map[string][]string{
			"kone": {"mt"},
			"ktwo": {"wa"},
		},
		expected: []string{"kone"},
	},
	{
		name:         "Покрытие уже полным составом",
		statesNeeded: []string{"mt", "wa"},
		stations: map[string][]string{
			"kone": {"mt", "wa"},
		},
		expected: []string{"kone"},
	},
	{
		name:         "Частичное покрытие",
		statesNeeded: []string{"mt", "wa", "or"},
		stations: map[string][]string{
			"kone":   {"mt"},
			"ktwo":   {"wa"},
			"kthree": {"or"},
		},
		expected: []string{"kone", "ktwo", "kthree"},
	},
	{
		name:         "Без покрытия",
		statesNeeded: []string{"ny"},
		stations: map[string][]string{
			"kone": {"mt"},
			"ktwo": {"wa"},
		},
		expected: nil, // Никакие станции не покрывают штат "ny"
	},
}

// Сам тест
func TestGreedyAlgorithm(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := greedyAlgorithm(tc.statesNeeded, tc.stations)
			if !equalSlices(actual, tc.expected) {
				t.Errorf("Test '%s': expected %+v, got %+v", tc.name, tc.expected, actual)
			}
		})
	}
}
