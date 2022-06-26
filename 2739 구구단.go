// https://www.acmicpc.net/problem/2739
// N입력받고 N단 찍기 2 * 1 = 2


package main

import (
"fmt"
"bufio"
"os"
"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {

    defer writer.Flush()
	var N int; // 단수
	fmt.Fscanln(reader, &N);
	for i:=1; i<10; i++ { // 1, 2, 3, 4, ... 9
		fmt.Fprintln(writer, N, "*", strconv.Itoa(i), "=", strconv.Itoa(i*N))
	}
}
