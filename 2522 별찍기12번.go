// https://www.acmicpc.net/problem/2522
/*

  *
 **
***
 **
  *

 // => 첫번째 줄에는 별이 1개, 즉..n번째 줄에는 별 L개, 공백은 T-L개
*/

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)


func star(T, L int) {
	for j:=0; j<T-L; j++{
		writer.WriteString(" ") //" "를 싹다 쑤셔박어
	}
	for k:=0; k<L; k++{
		writer.WriteString("*") //*를 싹다 쑤셔박어
	}
	writer.WriteString("\n") // 포룹하나 끝나면 개행문자첨부해서 추력해.

}


func main() {

    defer writer.Flush()

	var T int;
	fmt.Fscanln(reader, &T);

	for i:=1; i<T+1; i++ { // 총 T개 찍는다 ==> 범위표에서 뺀 연산 (T+1)  - (1)
		star(T, i)
	}

	for i:=1; i<T; i++ { //총 T-i개 찍는다 ==> (T) - (i)
		star(T, T-i)
	}
}
