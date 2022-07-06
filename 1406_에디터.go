
/*
백준 1406번  https://www.acmicpc.net/problem/1406
// 첫째 줄에는 초기에 편집기에 입력되어 있는 문자열이 주어진다. 이
// 문자열은 길이가 N이고, 영어 소문자로만 이루어져 있으며, 길이는 100,000을 넘지 않는다. 
// 둘째 줄에는 입력할 명령어의 개수를 나타내는 정수 M(1 ≤ M ≤ 500,000)이 주어진다. 
// 셋째 줄부터 M개의 줄에 걸쳐 입력할 명령어가 순서대로 주어진다. 명령어는 위의 네 가지 중 하나의 형태로만 주어진다.

문제를 푸는 방법
abcd
3
P x
L
P y

abcd
LEFT: []
RIGHT: []
====
P x :: abcdx (left에 x추가됨)
L :: right에 abcd right에 x이렇게 됨
P y :: abcdy , x
----
왼쪽출력, 오른쪽 출력
=> 결과물 abcdyx

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
// var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	reader = bufio.NewReaderSize(os.Stdin, 10000000) // 필요한 경우 이렇게 리더 버퍼 사이즈도 키운다. 입력이 60만개
	writer = bufio.NewWriterSize(os.Stdout, 10000000) // 버퍼가 커지면 계산 시간이약간은줄어드는 효과가 있긴한데미세하다. 늘어나는 메모리는 짝당 7메가 정도

    defer writer.Flush()

	var T int;
	var left []byte;
	var right []byte;

	input, _ := reader.ReadString('\n') 
	input = strings.TrimSuffix(input, "\n") // 맨끝에 개행문자가 들어간 것을 제거한다. 바이트스택에 불필요하게 쌓인다
	fmt.Fscanln(reader, &T);

	i:=0 // 왼쪽 스택에 일단 다 넣어라
	for i<len(input) {
		left = append(left, input[i])
		i++
	}

	for i:=0;i<T;i++ { 
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSuffix(cmd, "\n") // 개행문자 제거에 목숨걸어라.
		cmd2 := strings.Split(cmd, " ") 
    // (주의!!) 짜르면 리스트가 되는것 => 스트링인 오브젝에다가 []스트링을 재할당? 당연 안되지. 분리를 위해 cmd2를 썼다.
		// ./1406.go:87:10: cannot use strings.Split(cmd, " ") (value of type []string) as type string in assignment
    
		
		if cmd2[0] == "P" && len(cmd2) == 2 {
			s := []byte(cmd2[1])
			left = append(left, s[0])
		} 
      // https://stackoverflow.com/questions/8032170/how-to-assign-string-to-bytes-array
			// []byte("Here is a string....")
			// 스트링 => 바잇으로 형변환 한것은, 바잇트 어레이가 되기 때문시. [0]을 붙여줘야 한글자 짜리 바잇이 완성된다.

		if cmd2[0] == "L" && len(left)>0 { //커서를 왼쪽으로 옮기기 ==> 오른쪽에 문자열이 추가됨 :: 근데 커서가 맨앞이면 무시됨 ㅇㅇ; 숫자세는것 빡시다 ㅇㅇ;; >0
			right = append(right, left[len(left)-1])
			left = left[0:len(left)-1]
		}
		if cmd2[0] == "D" && len(right)>0 { 
			left = append(left, right[len(right)-1])
			right = right[0:len(right)-1]
		}
		if cmd2[0] == "B" && len(left)>0{
			left = left[0:len(left)-1]
			right = right //오른쪽은 그대로임
		}
	}

	// 수택 이터레이션..
	// for _, v := range left {
	// 	fmt.Fprintf(writer, "%c", v)
	// }
	for i, _ := range left {
		fmt.Fprintf(writer, "%c", left[i])
	}

   // 오른쪽 스택의 경우 출력하면 거꾸로 나오기 때문시, 거꾸로 출력해준다.
	for i, _ := range right {
		// fmt.Fprintf(writer, "%c", v)
		fmt.Fprintf(writer, "%c", right[len(right)-i-1])
	}
	fmt.Fprintln(writer, "")
}



/* 비교
		scanner.Scan() // 한줄 통째로 스트링으로 읽어와서 '\n' 개행문자도 제거해준다. 그걸 스캐너라 하노라.
		cmd := scanner.Text()
// fmt.Fscan(reader, &cmd) // 이게 내가 필요한 것이였나.. <<딱 한 토큰만 읽는 것>>..ㅇㅇ;; 왜냐.. " "하고 '\n'모두 딜리미터가 되기 때문시다.


if t == 'P' {
	var tp byte
	fmt.Fscanf(reader, " %c", &tp)
	a = append(a, tp)
 ===> 이런식으로 Fscanf " %c"를 쓴 사람도 있었다.
*/
