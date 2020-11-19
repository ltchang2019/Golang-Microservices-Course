package utils

func BubbleSort(elems []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elems)-1; i++ {
			if elems[i] > elems[i+1] {
				elems[i], elems[i+1] = elems[i+1], elems[i]
				keepRunning = true
			}
		}
	}
}
