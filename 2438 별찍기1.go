/* https://www.acmicpc.net/problem/2438
첫째 줄에는 별 1개, 둘째 줄에는 별 2개, N번째 줄에는 별 N개를 찍는 문제
입력
첫째 줄에 N(1 ≤ N ≤ 100)이 주어진다.
*/

package main

import (
"fmt"
"bufio"
"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) 
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {

    defer writer.Flush() //풀러시버튼이 마지막으로 눌리면서 라이터 버퍼에 담겼던 *들이 싹다 출력되는것. 이부분 없으면 출력도 없음.

	/* local variable definition */
	// var i, j int;

	var T int;
	fmt.Fscanln(reader, &T);

	// 최초 생각은 var 별창고 string; 이거였으나, <== 버퍼를 별창고처럼 쓰는것이 가능하다.

	for i:=0; i<T; i++ {
		for j:=0; j<i+1; j++{
			writer.WriteString("*") //*를 싹다 쑤셔박어
			// fmt.Fprint(writer, "*")
		}
		writer.WriteString("\n") // 포룹하나 끝나면 개행문자첨부해서 추력해.
	}
}

/*
C언어로 제작한 똑같은 코드.
==똑같이 생겼다.


#include <stdio.h>

int main(){
    int i;
    int j;
    int n;
    scanf("%d", &n);
    
    for (i=0;i<n;i++){
//        printf("\n");
        for (j=0;j<=i;j++){
            printf("*");
        }printf("\n");
        
    }

return 0;

}
*/
