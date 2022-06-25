// https://www.acmicpc.net/problem/10952
// 입력은 여러 개의 테스트 케이스로 이루어져 있다.  테스트 케이스는 한 줄로 이루어져 있으며, 각 줄에 A와 B가 주어진다. (0 < A, B < 10) . 입력의 마지막에는 0 두 개가 들어온다.

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

	var A, B int;
	fmt.Fscan(reader, &A, &B)
    
	if A !=0 || B != 0 { // 0 0이 뜨지 않았다면 계쏙 루핑하세요. 0 0이면 끝인가 ㅇㅇ;
		fmt.Fprintln(writer, A+B)
		main()
	}
}
