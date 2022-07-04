/*
https://www.acmicpc.net/problem/9012
문제를 푸는 요령
((()))
oooxxx
o:푸시, size+1   x:팝, size-1
x가 더 많이 나온경우 => bool :false
ox 아다리가 전부 맞으면 size=0, 뒤틀림이 없으면 bool=true이 된다.
*/

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)


var word_sac [10000]byte;
var size int = 0;
var check bool = true;


func push(a byte) { // a byte;
	word_sac[size] = a
	size += 1
	// fmt.Println("푸시합니다")
}

func pop(){
	if size > 0 {
	size -= 1;
	} else {
		check = false
	}
}

func main() {

    defer writer.Flush()
	var T int;
	fmt.Fscanln(reader, &T);


	for i:=0;i<T;i++{
		input,_ := reader.ReadString('\n')

		// '(' 이면 푸시, ')' 이면 팝
		for j:=0; j<len(input); j++{

			if input[j] == '(' {
				push(input[j])
			}

			if input[j] == ')' {
				pop()
			}
		}
		if check == true && size == 0 {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
    // 첫번재 줄이끝나면 스택 초기화
		check = true;
		size = 0; 
	}	
}
