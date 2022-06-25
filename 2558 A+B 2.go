// https://www.acmicpc.net/problem/2558 두 정수 A와 B를 입력받은 다음, A+B를 출력하는 프로그램을 작성하시오. 2
// 첫째 줄에 A, 둘째 줄에 B가 주어진다. (0 < A, B < 10)


package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	var A, B int;
    defer writer.Flush() 

	fmt.Fscan(reader, &A,&B)
	fmt.Fprintln(writer, A+B) // reader로 A,B를 받아오고, 버퍼에 넣으며, A+B값을 만들어서 라이터로 보내고, 플러시한다.

}

