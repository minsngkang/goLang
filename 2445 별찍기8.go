/*
https://www.acmicpc.net/problem/2445
*        *
**      **
***    ***
****  ****
**********
****  ****
***    ***
**      **
*        *

*/

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func star(T, L int) { // 주어진 정수값(T), 그리고 라인번째수(L) ==> 정수값 그리고 몇번째 라인인지 알면 그때의 별이 몇개고 공백이 몇개인지 찍어낼수가 있지.(점화식)
	// 만약 주어진정수값이 5일때의 첫째줄에는 별이 몇개가 들어가나? => 앞에 1개 .... 뒤에 1개.. 즉 L개 * 2 => L + L
	// 그때예 공백갯수는 2T - 2L개
	// 이것을 로직으로.. 앞에 *찍고  /  공백  찍고  / 뒤에 *찍고
	for i:=0; i<L; i++ {
		writer.WriteString("*")
	}
	for j:=0; j<2*T-2*L; j++ {
		writer.WriteString(" ")
	}
	for i:=0; i<L; i++ {
		writer.WriteString("*")
	}
	writer.WriteString("\n")
}



func main() {

    defer writer.Flush() 
	
	var T int;
	fmt.Fscanln(reader, &T);
	// star(T, 1)
	// star(T, 2)

	for i:=1; i<T+1; i++ {
		star(T, i)
	}
	for i:=1; i<T; i++{
		star(T, T-i)
	}

}

