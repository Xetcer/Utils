/*
// Моя реализация метода сортировки с помощью классического двоичного дерева.
https://apni.ru/article/8055-algoritmicheskaya-realizatsiya-drevesnoj-sort
https://habr.com/ru/companies/edison/articles/504012/
https://neerc.ifmo.ru/wiki/index.php?title=Теорема_о_нижней_оценке_для_сортировки_сравнениями
*/
package treesort

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func createNode(Value int) *TreeNode {
	return &TreeNode{Value: Value}
}

func (node *TreeNode) Insert(Value int) {
	if Value < node.Value {
		if node.Left == nil {
			node.Left = createNode(Value)
		} else {
			node.Left.Insert(Value)
		}
	} else {
		if node.Right == nil {
			node.Right = createNode(Value)
		} else {
			node.Right.Insert(Value)
		}
	}
}

func (node *TreeNode) TraverseByOrder() []int {
	if node == nil {
		return nil
	}
	arr := []int{}
	arr = append(arr, node.Left.TraverseByOrder()...)
	arr = append(arr, node.Value)
	arr = append(arr, node.Right.TraverseByOrder()...)
	return arr
}

func Sort(arr []int) []int {
	tree := createNode(arr[0])
	for _, value := range arr[1:] {
		tree.Insert(value)
	}
	return tree.TraverseByOrder()
}
