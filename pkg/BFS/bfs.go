/*
BFS Breadth-First Search
Алгоритм поиска в ширину "Грокаем Алгоритмы" глава 6
*/
package bfs

import "container/list"

// Persons Упрощение типа map[string][]string
type Persons map[string][]string

// personIsSeller если последнаяя буква имени == m значит это продавец.
func personIsSeller(name string) bool {
	return (name[len(name)-1] == 'm')
}

// FindSellerBFS создает из карты persons очередь LIFO, и начиная с корня firstNode
// ищет ближайшего соседа с последней буквой 'm' в имени, возвращает name, true если есть такой
// иначе "", false
func FindSellerBFS(persons Persons, firstNode string) (string, bool) {
	// Создаем очередь
	FIFO := list.New()
	verifiedPersons := make(map[string]int)
	pushNodesToLifo := func(names []string) {
		for _, name := range names {
			FIFO.PushBack(name)
		}
	}
	pushNodesToLifo(persons[firstNode])

	for e := FIFO.Front(); e != nil; e = e.Next() {
		if name, ok := e.Value.(string); ok {
			if verifiedPersons[name] == 0 { // Если еще не проверен
				if personIsSeller(name) {
					return name, true
				} else {
					pushNodesToLifo(persons[name])
					verifiedPersons[name]++ // Уже проверен
				}
			}
		}
	}
	return "", false
}
