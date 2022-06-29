/*
https://www.acmicpc.net/problem/2445
*********
 *******
  *****
   ***
    *
   ***
  *****
 *******
*********

=> 5.. 1번째 줄 공백 0개, 별갯수는 9개 ==:> 공백은 L-1, 별갯수는 2T-2L+1개( 2*(T-L)+1 ) ==> 바로 반복문 뜌레스홀드로 넣는다ㅇㅇ!

*/ 

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func star(T, L int) { // 주어진 정수값(T), 그리고 라인번쨰수(L)
	// 앞에 찍고  /  별찍고  / 뒤에는 공백없음 그냥 개행문자로 처리
	for i:=0; i<L-1; i++ {
		writer.WriteString(" ")
	}
	for j:=0; j<2*T-2*L+1; j++ {
		writer.WriteString("*")
	}
	// for i:=0; i<L; i++ {
	// 	writer.WriteString("*")
	// }
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

