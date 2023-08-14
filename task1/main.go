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

	sort := bubbleSort(arr)

	fmt.Fprintln(out, sort)

}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
