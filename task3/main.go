package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	numOfElem := 0

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	_, err := fmt.Fscanln(in, &numOfElem)
	if err != nil {
		return
	}

	arr := make([]int, numOfElem)
	for i := 0; i < numOfElem; i++ {
		var num int
		fmt.Fscan(in, &num)
		arr[i] = num
	}

	sort := insertSort(arr)

	fmt.Fprintln(out, sort)

}

// Сортировка вставками
// В сортировке вставками начинаем со второго элемента. Проверяем между собой второй элемент с первым и, если надо,
// меняем местами. Сравниваем следующую пару элементов и проверяем все пары до нее.
func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
