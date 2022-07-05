package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var data []int;
var w int = 0;

func push() {
	data = append(data, w+1);
	fmt.Fprintln(writer, "+") // 푸시 출력
	
}

func pop() {
	fmt.Fprintln(writer, "-")
	data = data[0:len(data)-1]
}

func top() int {
	return data[len(data)-1]
}


func main() {
	writer = bufio.NewWriterSize(os.Stdout, 10000000)
    defer writer.Flush()
	
	var T int;
	var num int;
	var check bool = true;

	fmt.Fscanln(reader, &T)

	
	for i:=0;i<T;i++ {
		fmt.Fscanln(reader, &num) 

		for num > w {
			push()
			w += 1
		}

		if num == top() { 
			pop()
		} else {
			check = false ;
			writer.Reset(writer)
		}
}

if check == false {
	fmt.Println("NO")
}
}
