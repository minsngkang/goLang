package main

// https://www.acmicpc.net/problem/1000

import (
"fmt"
"bufio"
"os"
)

// 버퍼입출력
var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {

    defer writer.Flush() // 지연함수, 함수를 종료하기 바로 직전에 버퍼에 들어있는 값을 영구 저장소로 옮긴다
  
    var A, B int;
    fmt.Fscan(reader, &A, &B) // 공백/개행문자 기준으로 입력값을 받는다 Fscan
    fmt.Fprintln(writer, A+B) // 값(A+B)을 그대로 문자열로 만들고, 해당 문자열에 개행문자를 붙여서 저장한다
  
}

