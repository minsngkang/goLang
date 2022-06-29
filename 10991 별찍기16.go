/*
https://www.acmicpc.net/problem/10991

T= 5

0000*
000* *
00* * *
0* * * *
* * * * *
==> 별은 1번째 줄에는 1개, 2번째 줄에는 2개... n번째 줄에는 n개가 있다
==>> 공백은 1번째 줄에는 4개, 2번째 줄에는 3개...n번째 줄에는 T-n개가 있다 => 이건 별 의미가 없어보인다..
====> L번째 i요소 각각 홀짝구분 ==> 홀짝/짝홀/홀홀/짝짝 이런식 구상..
======> 그냥 별을 한 번 찍었으면, 다음은 " "공백인거야
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
		//근데 홀수번째줄 홀수자리 수는 " "이고, 홀수번째줄 짝수번째 자리는 *을 채운다.
		// 그리고, 짝수번째줄 홀수자리는 " "을 채우고, 짝수번째줄 짝수자리는 *을 채운다.
		if L%2 == 1 && i%2 == 0{
			writer.WriteString("*")
		} 
		if L%2 == 0 && i%2 == 1 {
			writer.WriteString(" ")
		} 
		if L%2 == 1 && i%2 == 1 {
			writer.WriteString(" ")
		} 
		if L%2 == 0 && i%2 == 0 {
			writer.WriteString("*")
		} 
	} 
	writer.WriteString("\n") // 요소 순회 끝나면 개행문자ㅇㅇ!
}

func main() {

    defer writer.Flush()
	var T int;
	fmt.Fscanln(reader, &T);

	for i:=1;i<T+1;i++ {
		star(T, i)
	}
}

