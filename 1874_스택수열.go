/* 스택수열
https://www.acmicpc.net/problem/1874
문제를 푸는 방법

8( 입력 정수 갯수)

4
3
6


(테스트 케이스1)

+(푸시 => '내 스택'에 1 쌓임) : num4 size1, w1->w2
+(푸시 => 내 스택에 2 쌓임) :num4 size2, w2->w3
+(푸시 => 내 스택에 3 쌓임) : num4 size 3, w3->w4
+(푸시 => 내 스택에 4 쌓임) : num4 size 4, w4->w5
-(푸시 => 내 스택에서 4 뺌) => 4 출력됨 : num4 size 4, w5
-(푸시 => 내 스택에서 3 뺌) => 3 출력됨 : num3 size 4, w5
+(푸시 => 내 스택에 5 쌓임) num6 => size5, w5->w6
+(푸시 => 내 스택에 6 쌓임) num6 = > size6, w6->w7
-(푸시 => 내 스택에서 6 뺌) => 6 출력됨, w7
...
===> pop할때 출력된다.
data[size-1] 값이 num과 같으면 팝한다.

(테스트 케이스2)
5

1
2
5
3
4

+(푸시) 스택에 1추가 num1 size1
-팝 num1 size1 
+푸시 스택에 2추가 num2 size2 w2
-팝 num2 size2
num5 size2 
+푸시 w3
+푸시 w4
+푸시 w5
num5 size5 w5
-팝
num5 size5 w4
num3 size5 w4


이런식으로 운용하겠다.
저게 가능하면 저렇게 찍고, 아니면 버퍼를 비워고 NO를 찍는다.

*/

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)


var data [100000]int; //내 스택수열, 팝해야 출력됨
var size int = 0; //스택 커서Cursor!
var w int = 0; //대기수열

func push() {
	data[size] = w+1;
	size += 1;
	fmt.Fprintln(writer, "+") // 푸시 출력
}

func pop() {
	fmt.Fprintln(writer, "-") // -를 출력한다
	size -= 1;
}


func main() {
	
  writer = bufio.NewWriterSize(os.Stdout, 10000000) // 버퍼 고치니까 출력오류가 안뜬다. 버퍼가 꽉차면 비우는 것.. GO의 버퍼는 4096바이트. 넘 작다. 키운다.
  // https://golangbyexample.com/write-to-a-file-go/ <-버퍼 사이즈 커스텀하는 방법.
  
  defer writer.Flush()
	
	
	var T int;
	var num int;
  var check bool = false; // 체크변수는 팝했을때 값이 num이랑 다른 경우를 체크하는 변수다.(팝대기값 != num이면 true);
	// var bf bytes.Buffer // 별도의 버퍼를 잡는방법도 있다.

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ { // 첫번째 줄, 두번째 줄, 세번째 줄 ...
		fmt.Fscanln(reader, &num) // 그때 사용자 입력값을 num으로 받고,
		for num > w { // 입력값이 5인데, 현재 웨잇리스트(스택으로 이동할 수 있는 대기열)는 0인경우 => 5-0 5-1 5-2 5-3 5-4 5-5
			push()
			w += 1
		}

		if num == data[size-1] {  // data[size]가 아닌 data[size-1]이다. 번지 숫자를 잘 세야한다.
			pop()
		} else {
			check = true;
			// if found == false {
      writer.Reset(writer) // 버퍼 비우기 :: https://socketloop.com/tutorials/golang-reset-buffer-example
			// found = true ; // 또찍지 않겠다 예) 1 5 2 3 4
			// fmt.Println("NO")
			// bf.Reset()
      // os.Exit(0) 이런게 있다고?.. 참고
			// }
		}
    }

	if check == true {

		// bf.Reset()
		// bf.WriteString("NO")
		// writer.WriteString(bf.String())
		fmt.Println("NO")
	}
}
