/*
	Жадные алгоритмы, Грокаем алгоритмы:10 глава.
	Задача найти минимальное количество необходимых радиостанций, которые покроют радиовещанием все штаты
*/

package greedyalgs

type States []string
type Radio map[string]States

/*
 	getRadioMap получает срез станций, покрывающих все штаты
	states - мапа со штатами
	stations - срез радиостанций
*/
func greedyAlgorithm(statesNeeded States, stations map[string][]string) []string {
	finalStations := make([]string, 0)      // финальные станции
	unservedStates := make(map[string]bool) // Необслуженные штаты

	// Изначально все штаты не обслужены
	for _, state := range statesNeeded {
		unservedStates[state] = true
	}

	for len(unservedStates) > 0 {
		foundCoverage := false                   // Флаг, чтобы выйти из цикла если нет покрытия
		bestStation := ""                        // Станция, которая покрывает наибольшее количество штатов
		statesCovered := make(map[string]bool)   // покрытые штаты
		for station, regions := range stations { // получаем название текущей станции и ее охват
			currentlyCovered := make(map[string]bool)
			for _, region := range regions { // смотрим по каждому штату их текущей станции
				if unservedStates[region] { // Если штат еще не обслужен
					currentlyCovered[region] = true // Добавляем штат в текущий охват
				}
			}
			if len(currentlyCovered) > len(statesCovered) { // Если текущий охват станции больше чем общий охват
				bestStation = station            // Это лучшая станция из рассмаотренных
				statesCovered = currentlyCovered // обновляем охват до текущей
				foundCoverage = true
			}
		}
		if foundCoverage {
			// Добавляем лучшую станцию в список
			finalStations = append(finalStations, bestStation)
			// исключаем покрытые штаты из списка необслуженных
			for state := range statesCovered {
				delete(unservedStates, state)
			}
		} else {
			break // ни одна станция не смогла покрыть оставшиеся штаты
		}
	}
	return finalStations
}
