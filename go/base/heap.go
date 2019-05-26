package base

// 原地递归调整各层的大根堆
func down(target []int, from int, to int) {
	parent := from
	left := 2*from + 1
	right := 2*from + 2

	if left >= to || left < 0 { // < 0: int overflow
		return
	}
	// 比较父节点与左右叶子节点，取最大值的下标设为父节点下标
	if target[parent] < target[left] {
		parent = left
	}
	if right < to && target[parent] < target[right] {
		parent = right
	}
	// 只有父节点发生改变才会破坏大根堆结构，此时才需要继续调整下级大根堆
	if parent != from {
		swap(target, from, parent)
		down(target, parent, to)
	}
}

func NewBigHeap(target []int) {
	length := len(target)
	for i := length/2 - 1; i >= 0; i-- {
		down(target, i, length)
	}
}
