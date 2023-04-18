//함수형 선언하기
//func 함수형(매개변수) (반환타입 또는 반환값) {
// ...
//}
//매개변수가 없을 때는 (매개변수) 부분을 빈 괄호(())로 작성하고, 반환 값이 없을 때는 (반환타입 또는 반환값) 부분을 생략할 수 있다.
//func myFunc() {
//    ...
//}

//func의 매개변수
// 같은 타입인 매개변수가 여러 개일 때는 매개변수 이름을 콤마로 구분하여 작성하고 타입은 마지막에 표기한다.
//아래에서 i, j, k는 모두 int형임
//func myFunc (b bool, s string, i, j, k int) {
//	...
//}

//가변인자
//마지막 매개변수 타입 앞에 생략 부호(…)를 표기하면 여러 개의 값을 배열로 받고 이를 가변인자라고 부름
// 매개변수의 개수가 정해져 있지 않고 유동적으로 변할 때 사용

func myFunc (s string, integers ...int) {
	for i:=0; i <len(integers); i++ {
		fmt.Println(integers[i])
	}
}

//반환값
// Go에서는 값을 하나 이상 반환할 수 있음
// 반환 값이 하나일 때는 괄호를 생략하고, 두 개 이상일 때는 괄호로 묶어준다.

//반환 값이 하나일 때
package main

import "fmt"

func area(w, h int) int {
	return w * h
}
func main() {
	v :=area(3,4)  //return 3 *4값이 v에 저장됨 
	fmt.Println(v)
}

//반환 값이 두개  이상일 때

package main

import "fmt"

func multiply(w, h int) (int, int) {
	return w *2 , h *2
}

func main() {
	
}