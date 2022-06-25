//https://www.acmicpc.net/problem/10951
// 입력은 여러 개의 테스트 케이스로 이루어져 있다. 각 테스트 케이스는 한 줄로 이루어져 있으며, 각 줄에 A와 B가 주어진다. (0 < A, B < 10)

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout) 

func main() {

  defer writer.Flush()

	var A, B int;
  
	n, _:= fmt.Fscan(reader, &A, &B) // 첫번째 n은 렝쓰, 두번째 변수는 _인데 이것은 err리턴값을 안쓰겠다는 것이다.
	       // Fscan시키는 동작인데, 그것의 리턴값을 쓸 수 있다니.. 신기할 뿐
  
  if n == 2 { //만약 가져온 값이 2개면, 다음을 진행시켜
	fmt.Fprintln(writer, A+B)
	main() // 다시 재실행한다. 즉 조건이 맞음녀 EOF 뜰 때까지 무한루프 한다.
	}
}

// C: scanf의 리턴값은 성공적으로 입력받은 변수의 갯수다. ㅇㅇ!
