/*
백준 9093단어뒤집기 https://www.acmicpc.net/problem/9093
일차원배열을 사용한 풀이

일차원배열 data[10000]을 만들면 [0 0 0 0  ... 0 ]이렇다.
1만개 자리가 생긴다. 그러나 이것을 구분해주는 것이 필요하다.
push하고 pop하려면 그 지점을 나태내는 무언가가 필요하다. 그것이 여기선 size변수다.

푸시 = 몇번 번지수에 a값을 할당하는 것, 값을 넣고 번지수 +1
팝 = 번지수에서 a값을 빼는 것, 빼고 번지수를 -1

*/


package main

import (
"fmt"
"bufio"
"os"
"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

// 문자열 배열은 golang에서 byte배열
var data[10000] byte;
var size int = 0;


func Pop() {
	fmt.Fprintf(writer, "%c" ,data[size-1])
	size -= 1
}

func Push(v byte) {
	data[size] = v
	size += 1
}

func main() {
  defer writer.Flush()
	var T int;
	fmt.Fscanln(reader, &T);
  
	for i:=0; i<T; i++ {
		input,_ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		word_list := strings.Split(input, " ")

		for i, v:= range word_list {
			
			for j:=0;j<len(v);j++{
				Push(v[j])
			}
			for k:=0; k<len(v);k++{
				Pop()
			}
			if i < len(word_list) {
			fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprint(writer, "\n")
		}
	}
