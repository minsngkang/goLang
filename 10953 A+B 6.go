// https://www.acmicpc.net/problem/10953
// 첫째 줄에 테스트 케이스의 개수 T가 주어진다.각 테스트 케이스는 한 줄로 이루어져 있으며, 각 줄에 A와 B가 주어진다. A와 B는 콤마(,)로 구분되어 있다. (0 < A, B < 10)

package main

import (
"fmt"
"bufio"
"os"
"strings"
"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {

    defer writer.Flush()
	
	var T int; // 테스트 케이스의 갯수 받기
	var A string; // 통째로 한 줄을 읽어서 A라는 스트링으로 받자
	var X, Y int; // 덧셈하기 위한 내부 정수변수 2개 만들기
	
	fmt.Fscanln(reader, &T) //Fscan쓰면 에러나고 Fscanln쓰면 괜찮다. https://stackoverflow.com/questions/26126235/panic-runtime-error-index-out-of-range-in-go
  //Fscan: 공백기준으로 나눈다, 만약 테스트케이스 입력이 안들어왔으면 에러가 발생, Fscanln: 공백+개행문자 기준으로 나눔, 개행을 기대하는듯. 그래서 테스트케이스 입력이 성공적으로 완료됨
	for i:=0; i<T; i++ {
		fmt.Fscanln(reader, &A) // 여기서 Fscan도 되지만, 문맥적으로 Fscanln이 맞는듯.
		문자열:=strings.Split(A, ",") //스플릿한 슬라이스 전체를 '문자열'이라고 하자,

		X, _ = strconv.Atoi(문자열[0]) // 해당 값을 X, Y로 받고, 에러변수는 _로 받아서 넘긴다. // 문자열슬라이스에 담긴 값을 진짜 숫자로 캐스팅함
		Y, _ = strconv.Atoi(문자열[1])

		// fmt.Fprintln(writer, strconv.Atoi(문자열[0])+strconv.Atoi(문자열[1]) ) //성공적으로 뜨는 것을 확인

		fmt.Fprintln(writer, X+Y)
	}
}

