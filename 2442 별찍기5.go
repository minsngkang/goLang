/*

https://www.acmicpc.net/problem/2442
    *
   ***
  *****
 *******
*********
==> 오른쪽에 공백없음. i-1번 더 앵콜.

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

    defer writer.Flush() 

	var T int; // 테스트케이스는 몇개?
	fmt.Fscanln(reader, &T);


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
		writer.WriteString("\n") // 포룹하나 끝나면 개행문자첨부해서 추력해.

	}
}
