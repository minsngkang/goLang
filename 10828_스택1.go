// GO의 리스트 자료형을 사용해서 스택을 구현하는 방식
// fmt.Println을 쓰면 한줄한줄 즉각 피드백이 오지만 결과적으로는 매우 느린 코드가 된다.(380ms)
// fmt.Fprintln(writer, )를 사용했을 경우 8ms수준으로 확 줄어든다.
// 내부에서 제공하는 링크드리스트도 충분히 빠르다.(일차원배열과 동일한 8ms수준)

package main

import (
"fmt"
"bufio"
"os"
"container/list"
"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin) //버퍼 리더의 포인터 생성.(읽기 인스턴스)
// var reader = bufio.NewReader(os.Stdin) 이렇게 쓴 애도 있었는데.. 별차이는 없는듯 https://www.acmicpc.net/source/17637535
var writer *bufio.Writer = bufio.NewWriter(os.Stdout) //버퍼 라이터의 포인터 생성.(쓰기 인스턴스)


type 나의스택 struct{ // (내부)리스트를 상속해서 만든 나의스택. 고는 네이티브 스택이 없다.
	v필드 *list.List // v는 *list.List를 상속한다. ::v는 필드라는 명칭. Stack은 *list.List 필드를 가지고 있다
		// 필드명 없이 메서드만 쓴 경우는 구조체 임베딩이라고 한다. 나예스택은 *list.List라는 v필드를 가지고 있다.
  // 필드명을 대문자로 쓰면 퍼블릭, 소문자로 쓰면 프라이빗이 된다.
}

// GO의 리스트는 현재 변수에 저장된 값과 이전과 다음 값에 접근할 수 있는 포인터 변수를 가지고 있다. 
// type Element struct {
// 	Value interface{}
// 	Next  *Element
// 	Prve  *Element
//   }


func NewStack() *나의스택 { //func 함수명 리턴값자료형 ::형태로 쓰여져있다. // 생성자스럽게쓰고 있다.
	return &나의스택{list.New()} // 구조체인스턴스 '나의스택'를 생성하고 포인터<자료형>를 리턴한다. 찐&는 메모리주소값를 뜻함 
}


// 뭔가 다른 방식으로 써있는데?::=> func키워드, 찐함수명 사이에 있는 (q *나의스택) ==> 이것을 <리시버>라고 한다. q로 정의된 리시버.
// 리시버는 구조체에 메서드를 연결하도록 하는 구멍이다.
// 리시버(q)를 통해서 인스턴스 값에 접근하거나 점(.)연산자를 통해 메서드를 호출할 수 있다.
// 인터페이스란 메서드의 집합이다.
// 인터페이스 안에는 구현이 없는 메서드를 선언한다. 내부동작을 감추기 위해 사용된다.
// 여기 나온건 빈 인터페이스로 인터페이스에 아무 메서드도 정의돼 있지 않기때문에 << 모든 타입을 저장할 수 있는 그런 특징예 인터페이스다.>>:큰그릇
// ==>func (~) Push(a interface{}) :: 푸시함수는 빈 인터페이스를 사용해서 모든 값을 받겠다..ㅇㅇ!!라는 의미다.

func (q *나의스택) Push(a interface{}) { //매개변수로 구조체 포인터를 받음. *나의스택 하면 투입하는 나의스택 값이 수정됨.(요리가능). 
									// 그냥 나의스택 넣으면 카피만 떠서 값만 찍어내는 용도.(원래 들어있던 데이터 유지)
									// func (T int) 기억나는가... 그것이다.
									// 그럼 여기서 인터페이스란 또 뭔가?
	q.v필드.PushBack(a)
}

func (q *나의스택) Pop() interface{} { // 반대로.. 그럼 모든걸 담을수 있는 큰리턴은 인터페이스 리턴인가?? ㅇㅇ 그런것 갑다. 인터페이스를 리턴하는 함수 잘 됨.ㅇㅇ!!
	back:= q.v필드.Back()
	if back == nil {
		fmt.Fprintln(writer,-1) // 만약 들어있는 정수가 없을 경우에는 -1을 출력한다
		return nil
	}
	fmt.Fprintln(writer,back.Value) // 닷밸류를 해줘야 값이 찍혀보인다. 그냥 back하면 메모리주소값만 나옴.
	// Fprintln => 입력 끝났을때 종료하기 직전에 밸류들 싹다 일괄배송(버퍼에 말린것들 한꺼번에)
	// Println => 입력이 끝나든 말든, 그 시점에 그냥 자동적으로 배송(각개 배송)
  // 팝 하면 싸이즈값을 변경시켜줘야한다..? 아니다. 자동으로 바뀐다..;;<?!>:: *list.List 작동
	return q.v필드.Remove(back)
}


// 제일 위에 있는 녀석 출력하기 ==> 매개변수가 필요한가? ㄴㄴ 그냥 찍으면 된다. void()마냥 반환하는애도없고, 매개변수도 필요하지 않...는게 아니라 ==> 리시버가 필요하지
func (q *나의스택) Top() {
	back:= q.v필드.Back()
	if back == nil {
		fmt.Fprintln(writer,-1)
	}
  
	// top
	// nothing
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x1098c3d]
	// SigFault는 버퍼문제가 아니었다. 널포인터 에러
  
	if back != nil {
		fmt.Fprintln(writer,back.Value)
	}
}

func (q *나의스택) Size() {
	size:= q.v필드.Len()
	fmt.Fprintln(writer,size) // 요거는 .Value를 안해도된다;;;
}

func (q *나의스택) Empty() {
	size:= q.v필드.Len()
	if size == 0 {
		fmt.Fprintln(writer,1) //스택이 비어있으면 1 아니면 0출력;; 문제를 잘 읽자 ㅇㅇ;;
	} else {
		fmt.Fprintln(writer,0)
	}
}



// func my_sum (T int) interface{} {
// 	fmt.Fprintln(writer,"5555")
// 	return T+5
// }



func main() {

    defer writer.Flush()
	stack := NewStack() // 스택구조체 포인터를 리턴함. 뉴스택애는 메모리주소를 반환하는 애임. 그러므로, stack과 메모리주소가 맞춰짐.

	var T int;
	fmt.Fscanln(reader, &T); // 생성자 흉내기법. reader는 포인터 자료형. &T는 메모리주소

	for i:=0; i<T; i++ { // 입력은 버퍼로 받지만, 출력은 각각 일괄배송이란다 ㅇㅇ;;
		input, _ := reader.ReadString('\n') // 리드스트링으로 받는데 이때 델리미터는 \n
		input = strings.TrimSuffix(input, "\n")
		cmd := strings.Split(input, " ") // 델리미터는 공백 " "


		if cmd[0] == "push" && len(cmd) == 2 { // lena가 2가 아닐땐 패스?시키는 로직..ㅇㅇ
			// fmt.Fprintln(writer, "푸시감지", cmd[1])

			// if len(cmd) == 2{
			stack.Push(cmd[1])
			// }
			// ./10828.go:71:14: invalid operation: cannot index stack.Push (value of type func(a interface{})) :: push()가 아닌 push[] 썼을때의 에러문재
			
		}

		// 	stack.Push[cmd[1]]
		// }
		if cmd[0] == "pop"{
			stack.Pop()
		}

		if cmd[0] == "top"{
			stack.Top()
		}

		if cmd[0] == "size"{
			stack.Size()
		}

		if cmd[0] == "empty"{ 
			stack.Empty()
		}
	}
}


// 참조: // https://dev-yakuza.posstree.com/ko/golang/data-structure/#%EC%8A%A4%ED%83%9D :: 리스트로 스택구현하기
