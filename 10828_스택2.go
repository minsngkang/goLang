// 일차원 배열 하나로 스택을 구현하는 방법 // https://www.acmicpc.net/source/17637535

package main

import (
"fmt"
"bufio"
"os"
// "container/list"
"strings"
"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

// 글로발 배열스택 ㅇㅇ;;
var data [10000]int; // 스택은 일차원배열 하나로 구현할 수 있다 ㅇㅇ!!
var size int = 0;

func Push(num int) {
	// 	need type assertion ::num interface{} :: int대신 interface{}를 넣으면 타입에러가 뜬다
	data[size] = num;
	size += 1
}

func Top() {
	if Empty() == true{
		fmt.Fprintln(writer,-1)
	} else {
	fmt.Fprintln(writer, data[size-1] ) // 컴터는 숫자를 0부터세기 때문시.
	}
}

func Empty() bool {
	if size == 0 {
		return true;
	} else {
		return false;
	}
}

func Pop() {
	if Empty() == true {
		fmt.Fprintln(writer,-1)
	} else {
		fmt.Fprintln(writer,data[size-1])
		size -= 1
	}
}

func Size() {
	fmt.Fprintln(writer,size)
}


func main() {
	defer writer.Flush()
	
	var T int;
	fmt.Fscanln(reader, &T);

	for i:=0; i<T; i++ {
		input, _ := reader.ReadString('\n') // 리드스트링으로 받는데 이때 델리미터는 \n
		input = strings.TrimSuffix(input, "\n")
		cmd := strings.Split(input, " ") // 델리미터는 공백 " "

		if cmd[0] == "push" && len(cmd) == 2 { // lena가 2가 아닐땐 패스?시키는 로직..ㅇㅇ
			num, _ := strconv.Atoi( cmd[1])
			Push( num  )
		}

		if cmd[0] == "pop"{
			Pop()
		}
		if cmd[0] == "top"{
			Top()
		}
		if cmd[0] == "size"{
			Size()
		}
		if cmd[0] == "empty"{ 
			if Empty() == true {
				fmt.Fprintln(writer,1)
			} 
			if Empty() == false {
				fmt.Fprintln(writer,0)
			}
		}
}
}
