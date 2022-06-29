/*

https://www.acmicpc.net/problem/2444
	*
   ***
  *****
 *******
*********
 *******
  *****
   ***
    *
==> 2442번을 찍고 아랫쪽 다시 찍기. 둘을 함수로 분리시키면 편리하군.

*/

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)



func star(T int) { // func 내함수(변수 int) 리턴값의자료형 {중괗로 } <=== 형식..

	for i:=0; i<T; i++ {

		for j:=0; j<T-i-1; j++{
			writer.WriteString(" ") //" "를 싹다 쑤셔박어
		}

		for k:=0; k<i+1; k++{
			writer.WriteString("*") //*를 싹다 쑤셔박어
		}
		// 앵콜
		for l:=0; l<i; l++{
			writer.WriteString("*") //*를 또찍어
		}

		// for k:=0; k<2*i+1; k++{ // k/l 두개식을 하나로 합치는 이 방식도 가능함
		// 	writer.WriteString("*") //*를 싹다 쑤셔박어
		// }

		writer.WriteString("\n") // 포룹하나 끝나면 개행문자첨부해서 추력해.

	}

}

func reverseStar(T int) {

	for i:=0; i<T; i++ {

		for j:=0; j<i+1; j++{
			writer.WriteString(" ") //" "를 싹다 쑤셔박어
		}
		for k:=0; k<T-i; k++{
			writer.WriteString("*") //*를 싹다 쑤셔박어
		}
		// 앵콜
		for l:=0; l<T-i-1; l++{
			writer.WriteString("*") //*를 또찍어
		}
		
		// for l:=0; l<2*(T-i)-1; l++{ // k, l 두개식을 하나로 합친것. 그러나 속도는 펼쳐놓은게 더 빠르다.(8ms => ...4ms
		// 	writer.WriteString("*") //*를 또찍어
		// }

		writer.WriteString("\n") // 포룹하나 끝나면 개행문자첨부해서 추력해.
}
}




func main() {

    defer writer.Flush() //풀러시버튼이 마지막으로 눌리면서 라이터 버퍼에 담겼던 *들이 싹다 출력되는것. 이부분 없으면 출력도 없음.

	var T int; // 테스트케이스는 몇개?
	fmt.Fscanln(reader, &T);
	star(T)
	reverseStar(T-1)
	// fmt.Print(buf.String())
}
