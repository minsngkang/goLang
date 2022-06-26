// https://www.acmicpc.net/problem/11021
// 첫째 줄에 테스트 케이스의 개수 T가 주어진다.
// 각 테스트 케이스는 한 줄로 이루어져 있으며, 각 줄에 A와 B가 주어진다. (0 < A, B < 10)
// 각 테스트 케이스마다 "Case #x: "를 출력한 다음, A+B를 출력한다. 테스트 케이스 번호는 1부터 시작한다.

package main

import (
"fmt"
"bufio"
"os"
"strconv"
// "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout) 

func main() {

    defer writer.Flush()
	var T int;
	var A, B int;
	fmt.Fscanln(reader, &T)
	for i:=0; i<T; i++ {
		fmt.Fscanln(reader, &A, &B) // Fscanln은 \n나올때까지 읽는다.(공백은 쎄미 구분자가 되는 것) 즉 EOF까지 받는것 ㅇㅇ!!
		fmt.Fprint(writer, "Case #" + strconv.Itoa(i+1) +": ") // Fprint는 \n을 포함하지 않는다
		fmt.Fprintln(writer, A+B) // Fprintln은 \n을 포함한다 ㅇㅇ!!
	}   
}

