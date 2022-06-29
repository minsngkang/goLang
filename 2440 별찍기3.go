/*

https://www.acmicpc.net/problem/2440
*****
****
***
**
*
==> 얘는 끝에 공백이가 없다. ㅇㅇ!! 공백 문자열이 안들어있다..!!

*/

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)


func main() {

    defer writer.Flush() //풀러시버튼이 마지막으로 눌리면서 라이터 버퍼에 담겼던 *들이 싹다 출력되는것. 이부분 없으면 출력도 없음.

	/* local variable definition */
	// var i, j int;

	var T int;
	fmt.Fscanln(reader, &T);


	for i:=0; i<T; i++ {

		// j(공백갯수)+k(별의 갯수) = T(주어진 정수값)

		for j:=0; j<T-i; j++{
			writer.WriteString("*") //"*"를 싹다 쑤셔박어
		}
		// for k:=0; k<i; k++{
		// 	writer.WriteString(" ") // 를 싹다 쑤셔박어
		// }
		writer.WriteString("\n") // 포룹하나 끝나면 개행문자첨부해서 추력해.

	}
	// fmt.Print(buf.String())
}

// 챔고: 출력 형식이 잘못되었읍니다: "출력 결과는 정답과 유사하나, 공백, 빈 줄과 같은 문제로 인해서 출력 결과가 일치하지 않은 경우 입니다."
