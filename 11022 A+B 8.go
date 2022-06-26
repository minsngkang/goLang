// https://www.acmicpc.net/problem/11022
// 두 정수 A와 B를 입력받은 다음, A+B를 출력하는 프로그램을 작성하시오.
// 입력
// 첫째 줄에 테스트 케이스의 개수 T가 주어진다.
// 각 테스트 케이스는 한 줄로 이루어져 있으며, 각 줄에 A와 B가 주어진다. (0 < A, B < 10)
// 출력
// 각 테스트 케이스마다 "Case #x: A + B = C" 형식으로 출력한다. x는 테스트 케이스 번호이고 1부터 시작하며, C는 A+B이다.

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
	var T int;
	var A, B int;
	fmt.Fscanln(reader, &T)
	for i:=0; i<T; i++ {
		fmt.Fscanln(reader, &A, &B)
		fmt.Fprint(writer, "Case #" + strconv.Itoa(i+1) +": " )
		fmt.Fprintln(writer, A, "+", B, "=", A+B) // 공백을 기준으로 채워넣고, EOF에선 \n을 붙이기 ㅇㅇ!!
	}   
}

