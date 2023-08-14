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

	sort := findMaxSumArray(arr)

	fmt.Fprintln(out, sort)

}

// Подмассив с максимальной суммой
func findMaxSumArray(arr []int) []int {
	outArr := make([]int, 0, len(arr))
	sum := 0
	for _, elem := range arr {
		if sum+elem > 0 {
			outArr = append(outArr, sum+elem)
			sum = sum + elem
		} else {
			outArr = append(outArr, 0)
			sum = 0
		}
	}

	maxElem := 0
	posMaxElem := 0
	posLastZero := 0
	for i, elem := range outArr {
		if elem > maxElem {
			maxElem = elem
			posMaxElem = i
		}
		if elem == 0 && posLastZero < posMaxElem {
			posLastZero = i
		}
	}
	return arr[posLastZero+1 : posMaxElem+1]
}
