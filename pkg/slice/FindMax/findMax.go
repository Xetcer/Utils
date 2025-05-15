/*
Рекурсивный поиск максимума в массиве "Грокаем Алгоритмы" глава 4 упражнение 4.3
*/
package findmax

// FindMax рекурсивный поиск максимального элемента в массиве
func FindMax(s []int) int {
	// базовый случай
	switch len(s) {
	case 0:
		return 0 // базовый случай
	case 1:
		return s[0] // базовый случай
	default: // рекурсивный случай
		{
			findedMax := FindMax(s[1:])
			if s[0] >= findedMax {
				return s[0]
			} else {
				return findedMax
			}
		}
	}

}
