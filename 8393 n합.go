// https://www.acmicpc.net/problem/8393
// n이 주어졌을 때, 1부터 n까지 합을 구하는 프로그램을 작성하시오.

package main

import (
"fmt"
"bufio"
"os"
"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)


func summa(n int) int { //함수 선언하는 형식 (n int) <== "파라미터 셋팅" 이렇게 하고, 괄호밖 int는 리턴타입을 명시
	if n == 0 {
		return 0
	} else if n== 1 {
		return 1
	}
	return n + summa(n-1) // 큰 문제는 그것보다 작게 쪼갤수 있다. < 전체 = 대가리 + 몸통 > ... 이렇게. 
	// 점화식 수식으로 나타내면 S_n = S_n-1 + a_n. 이 경우 a_n = N
}

func main() {
    defer writer.Flush()
	var N int;
	fmt.Fscanln(reader, &N)
	fmt.Fprintln(writer, strconv.Itoa(summa(N)))
}
