package main

import (
"fmt"
)

// 슬라이스 선언
var sli []byte;

func main() {

	// 슬라이스에 어팬두(푸시) // 슬라이스 스택 푸시
	sli = append(sli, 'c')
	fmt.Println(sli[len(sli)-1])
  
	sli = append(sli, 'd')
  // 슬라이스 탑(맨 윗 데이타 보기)
	fmt.Println(sli[len(sli)-1])

	// 슬라이스에서 길이 하나 빼서 줄이기. 팝.  // 슬라이스 스택 팝
	sli = sli[0:len(sli)-1]
	fmt.Println(sli)

}
