/*

https://www.acmicpc.net/problem/2441
*****
 ****
  ***
   **
    *
==> 왼쪽에 공백포함

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
	/* local variable definition */
	// var i, j int;

	var T int;
	fmt.Fscanln(reader, &T);


	for i:=0; i<T; i++ {

		// j(공백갯수)+k(별의 갯수) = T(주어진 정수값)

		for j:=0; j<i; j++{
			writer.WriteString(" ") //" "를 싹다 쑤셔박어
		}
		for k:=0; k<T-i; k++{
			writer.WriteString("*") //*를 싹다 쑤셔박어
		}
		writer.WriteString("\n") // 포룹하나 끝나면 개행문자첨부해서 추력해.

	}
}
