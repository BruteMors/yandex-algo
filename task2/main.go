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

	sort := chooseSort(arr)

	fmt.Fprintln(out, sort)

}

func chooseSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := 0
		posMin := 0
		for j := i; j < len(arr); j++ {
			min = arr[i]
			posMin = i
			if arr[j] < min {
				min = arr[j]
				posMin = j
			}
		}
		arr[i], arr[posMin] = arr[posMin], arr[i]
	}
	return arr
}
