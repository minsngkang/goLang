// https://www.acmicpc.net/problem/10950 A+B 3
// 첫째 줄에 테스트 케이스의 개수 T가 주어진다. 각 테스트 케이스는 한 줄로 이루어져 있으며, 각 줄에 A와 B가 주어진다. (0 < A, B < 10)

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {

    defer writer.Flush()
    
	var T, A, B int;
	fmt.Fscan(reader, &T)

	for i:=0; i<T; i++ {
		fmt.Fscan(reader, &A, &B)
		fmt.Fprintln(writer, A+B)
	}
	// main()
}

