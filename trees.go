package main

type Node struct {
	Value   int
	Left    *Node
	Right   *Node
	visited bool
}

func BinarySearch(arr []int, value int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == value {
			return mid
		}
		if arr[mid] > value {
			high = mid - 1
		}
		if arr[mid] < value {
			low = mid + 1
		}
	}
	return -1
}

func BFS(tree *Node) []int {
	result := []int{}
	queue := []*Node{tree}
	for len(queue) > 0 {
		node := queue[0]
		if node.visited {
			continue
		}
		result = append(result, node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		node.visited = true
		queue = queue[1:]

	}
	return result
}
