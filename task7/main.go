package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	str := ""
	_, err := fmt.Fscanln(in, &str)
	if err != nil {
		return
	}
	fmt.Println(isCompleted(str))
}

// Закрывающиеся скобки
func isCompleted(str string) bool {
	m := make(map[string]int, len(str))
	for _, ch := range str {
		m[string(ch)]++
	}
	if m["("] != m[")"] {
		return false
	}
	if m["{"] != m["}"] {
		return false
	}
	if m["["] != m["]"] {
		return false
	}
	return true
}
