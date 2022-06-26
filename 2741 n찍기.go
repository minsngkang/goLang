// https://www.acmicpc.net/problem/2741
// 자연수 N이 주어졌을 때, 1부터 N까지 한 줄에 하나씩 출력하는 프로그램을 작성하시오.

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
	var N int;
	fmt.Fscanln(reader, &N);
	for i:=0; i<N; i++ {
		fmt.Fprintln(writer, i+1)
	}
}

