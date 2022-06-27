// https://www.acmicpc.net/problem/2439
/*
    *
   **
  ***
 ****
*****
*/

package main

import (
"fmt"
"bufio"
"os"
  //"bytes"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout) // 여기도 마찬가지로 초기화된 공간이다.
// 임의의 버퍼를 셋팅하는 방법도 있긴 하다.
// var buf bytes.Buffer ==>> 윗 방식이 아니더라도 writer대신 버퍼를 쓸 수가 있다. 다만 이렇게 만든 buf에는 플러시버튼이 없어, fmt.Println(buf.String())이렇게 출력해준다.

func main() {

    defer writer.Flush() //풀러시버튼이 마지막으로 눌리면서 라이터 버퍼에 담겼던 *들이 싹다 출력되는것. 이부분 없으면 출력도 없음.

	/* local variable definition */
	// var i, j int;

	var T int;
	fmt.Fscanln(reader, &T);


	for i:=0; i<T; i++ {

		// j(공백갯수)+k(별의 갯수) = T(주어진 정수값)

		for j:=0; j<T-i-1; j++{
			writer.WriteString(" ") //" "를 싹다 쑤셔박어
		}
		for k:=0; k<i+1; k++{
			writer.WriteString("*") //*를 싹다 쑤셔박어
		}
		writer.WriteString("\n") // 포룹하나 끝나면 개행문자첨부해서 추력해.

	}
  // 	fmt.Print(buf.String())
}
