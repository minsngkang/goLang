/*
https://www.acmicpc.net/problem/9093 단어뒤집기,
문제를 푸는 아이디어
i am happy
oxooxooooox ...[통째로 스트링으로 읽어서 i번째 탐색, 그때]
o:인풋, 스택에 쌓기
x:아웃풋, 스택에 있는 모든 애들 팝
*/

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var word_sac []byte;

func reverseWord() {
	for i:=0; i<len(word_sac); i++ {
		fmt.Fprintf(writer, "%c", word_sac[len(word_sac)-i-1]) // 맨끝요소가 안잡히는 것은, 컴터가 숫자세는 방식이 달라서.. 그래서 -1
	}
  word_sac = nil // 클리어런스. word_sac :=[] 이런식은 다시 선언하는 방식이라 왜 다시 사용안했냐고 물어본다. nil이 최고다. => https://yourbasic.org/golang/clear-slice/
}

func main() {
  defer writer.Flush()
	var T int;
	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		input,_ := reader.ReadString('\n')
		// input = strings.TrimSuffix(input, "\n")
		// word_list = strings.Split(input, " ")

		for j:=0;j<len(input);j++{
			if input[j] == ' ' || input[j] == '\n' {
				// fmt.Println("공백")
				reverseWord()
				fmt.Fprint(writer, " ")

			} else {
				word_sac = append(word_sac, input[j])
				// fmt.Println("인풋", input[j])
			}
			
		}
		fmt.Fprintln(writer, "")
	}
}



/* 참고 
	// var charArray [4]uint8; // byte == uint8
	// // charArray.append('s')
	// // 배열은 길이가 고정돼 있기 때문시 어팬두를 쓸 수가 없다. 왜냐묜 길이가 늘어나기 때뮨시다.
	// // 배열은 할당만 가능하다. 스트링인데 어팬두 처럼 추가하기 기능을 쓰려면 슬라잇스를 쓰는게 낫겠다.

	// charArray[0] = 'l'
	// charArray[1] = 'o'
	// charArray[2] = 'v'
	// charArray[3] = 'e'


	// a := append(charArray, 'l')
	// 문자를 저장하는데, 그 저장값은 문자가 아니라 (u)int8이다. 
	// https://levelup.gitconnected.com/demystifying-bytes-runes-and-strings-in-go-1f94df215615

	// c := make([]byte, 0) // 이게 파이썬스러운 그런 리스트
	// c = append(c, 1)
	// c = append(c, 1)
	// c = append(c, 2)
	// c = c[0:len(c)]


	// d := make([]string, 0)
	// // fmt.Println(&d)
	// d = append(d, "sexy" )
	// d = append(d, "lady" )
	// // fmt.Println(&d)

	// for _, v:= range d {
	// 	fmt.Println(v)
	// }

*/
