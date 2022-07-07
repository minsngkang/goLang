/*
https://www.acmicpc.net/problem/10845

// GO의 채널로 구현한 큐는 프론트/빽 이터레이션 할라면 채널을 닫아야함.
// 채널이 닫히면 다시 만들어야함;;
// 번거로움. 자동화하기는 어려움.
// 채널구현큐는 그냥 팝/푸시만 됨;;

*/

package main

import (
"fmt"
"bufio"
"os"
"strings"
"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var data []int; // 큐 슬라이스. 슬라이스로 구현한 큐 // 얘는 레퍼런스 타입이라 Data하든 data하든 별 상관이 없나보군 ㅇㅇ;;


func Push(a int){
	// fmt.Println("푸시 실행")
	data = append(data, a)
}

func Pop() {
	if len(data) == 1{
		fmt.Fprintln(writer, data[0])
		data = nil // 슬라이스 클리어하기=nil. https://yourbasic.org/golang/clear-slice/
	} else if len(data) == 0{
		fmt.Fprintln(writer, -1)
	} else {
	fmt.Fprintln(writer, data[0]) // 일단 맨앞의 숫자를 찍고
	data = data[1:len(data)] //한칸 비운다 ㅇㅇ! len(data)-1까지하면 )조건이라 -1이 더까이는 듯. 그래서 +1해줌
	}
}

func Size() {
	fmt.Fprintln(writer, len(data))
}

func Empty() {
	if len(data) == 0{
		// fmt.Println("비어있음")
		fmt.Fprintln(writer, 1)
	} else {
		// fmt.Println("owing")
		fmt.Fprintln(writer, 0)
	}
}

func Front() {
	if len(data) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, data[0])
	}
}

func Back() {
	if len(data) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, data[len(data)-1])
	}
}

func main() {

    defer writer.Flush()

	var T int;
	fmt.Fscanln(reader, &T);

	for i:=0;i<T;i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		cmd := strings.Split(input, " ")

		switch cmd[0]{
		case "push":
			num, _ := strconv.Atoi(cmd[1])
			Push(num)
		case "pop":
			Pop()
		case "size":
			Size()
		case "empty":
			Empty()
		case "front":
			Front()
		case "back":
			Back()
		}
	}
	// fmt.Fprintln(writer, data)
}
