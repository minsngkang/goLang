/* https://www.acmicpc.net/problem/10818
N개의 정수가 주어진다. 이때, 최솟값과 최댓값을 구하는 프로그램을 작성하시오.
20 10 35 30 7 이런식으로 받는데, 얘는 공백이 있다...공백도 그냥 싹다 문자열에 넣어야하는데 어떻게 하나? ==> ReadString :: https://pkg.go.dev/bufio#Reader.ReadString
*/

package main

import (
"fmt"
"bufio"
"os"
"strings"
_"reflect"
"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {

	max, min := -9999999, 9999999

    defer writer.Flush() // 지연함수, 끝나기 막차에 플러시하라
	
	var T int; // 첫 줄에 주어지는 정수의 갯수 요약
	var input string; //
	
	fmt.Fscanln(reader, &T);
	// n, _ := fmt.Fscanln(reader, &input)

	
	input, _ = reader.ReadString('\n') // 리드스트링, 줄 자체를 싹다 스트링으로 읽기.(스캐너 갬성은 띄어쓰기들가면 구분되는 그런 느낌이었지)
	input = strings.TrimSuffix(input, "\n") // 다만 개행문자를 뒤에 포함하고 있어서 빼줘야한다. 개행문자 처리는 반드시 "쌍따표"
	// fmt.Fprintln(writer, input)

	문자열리스트 := strings.Split(input, " ") // :=는 붙여써야한다. : =이건 안돼.
	// :=는 다음에 익스프레션이 와야한다. 위처럼. 혹은 익스프레션이 아니라면 정수 혹은 단일값 하나만 ㅇㅇ;
	// 낯선변수 := "김고" // or 1 등 단일값
	// fmt.Println(낯선변수)
	// fmt.Println(reflect.TypeOf(낯선변수)) // 타입을 검사할 수 있다.

	var 실제입력갯수 int = 0;
	for i, v := range 문자열리스트 { // 임시로 이름표 붙이는 그런 구획 ...i만 썼던거랑 다르게 i,v다쓴다.
		// fmt.Fprintln(writer, i, v)
		v, _ := strconv.Atoi(v);
		if v > max {
			max = v
		}  // 여기서 else if를 쓰면, -99999보다 작고 999999보다 큰 경우를 찾는다;; 그럼 안되지.
		if v < min {
			min = v
		}
		실제입력갯수 = i+1
	}

	if 실제입력갯수 == T {
		fmt.Fprintln(writer, min, max)
	} else {
		fmt.Fprintln(writer, "err")
		// goto PANIC
	}

  // PANIC: //goto를 사용하지 않는게 낫다.(spaghetti code 60s 갬성이다:=) https://stackoverflow.com/questions/11064981/why-does-go-have-a-goto-statement
	// 특히 else등 조건문과 goto를 섞어쓰면 백빵에러난다.
  // 	fmt.Fprintln(writer, "err")
}

