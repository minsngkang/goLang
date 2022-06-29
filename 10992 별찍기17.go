/*
https://www.acmicpc.net/problem/10992

T= 4
   *
  * *
 *   *
*******

==> 공백찍기 + 별이랑공백조합군 찍기
-==>> 최초 공백은 T-L개 
======> 조합군들에서, i=0; i=마지막수 일때 *을찍고 else는 " "공백을 찍는다. 조합군들 갯수는 2*L -1 개 찍는다.
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
 // 일단 첫번째 줄에는 한개, 두번째 줄에는 두개... L번째 줄에는 L개 
 // 공백 + 별과 공백조합 , 이렇게 나눠서 찍을수가 있다.
 // 최초 등장하는 공백 갯수는 T-L개다
	for i:=0; i<T-L; i++{ // 공백을 찍는 곳
		writer.WriteString(" ")
	}
 	//여기서부터 별과 공백의 조합을 찍는다. 이렇게 찍히는 애덜의 렝쓰는? 1,3,5,7,9...이런식으로 증가한다
 	//일단, 2L-1개 찍는것으로 확인
	for i:=0; i<2*L-1; i++{
		if i==0 || i==2*L-2 {
			writer.WriteString("*")
		} else {
			writer.WriteString(" ")
		}

	} 
	writer.WriteString("\n")
}

func bottomLine(T int) {
	for i:=0; i<2*T-1; i++{
		writer.WriteString("*")
	}
	writer.WriteString("\n")
}

func main() {

    defer writer.Flush()
	var T int;
	fmt.Fscanln(reader, &T);

	for i:=1;i<T;i++ { // 총 T-1개 줄을 찍고
		star(T, i)
	}
	bottomLine(T) // 마지막 번째 줄은 완전 도배시키는 애로 찍는다.
}

