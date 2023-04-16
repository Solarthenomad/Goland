//if
//불타입 조건문만 정상적인 조건문으로 인식한다.
package main

//func main() {
//	i:=1
//	if i {

//	}// 잘못된 코드임. 조건문에는 반드시 불타입이 들어가야함!
//}

//if condition {
//     return x
//} return y
//형식으로 코드 작성해야함

func main() {
	i:=0
	if i >0 {
       //

	}

}

//올바른 작성법임

//조건문 if에서 초기화 구문을 사용하는것 역시 가능함 
// if 초기화구문; 조건식 {

//}

//v := Compute() 
//    if v<0 {
//		fmt.Println(v, "는 음수입니다.")

//}

if v:=Compute(); if v<0{
	fmt.Println(v, "는 음수입니다.")
}

//switch 구문
// switch 값 {
//case 조건1:
//    …
//case 조건2:
//    …
//default:
//    …
//}

switch i {
case -1, -2 :
	fmt.Println( i,"는 음수입니다.");
case 1, 2 :
	fmt.Prinln(i, "는 양수입니다.");
}

// 일치하는 case 조건을 만나면 바로 switch 문을 빠져나오므로 switch 문에는 break를 쓰지 않아도 된다
//. 일치하는 조건을 만났지만 switch 문을 빠져나오지 않고 다음 case로 넘어가려면 fallthrough를 표기해야 한다.



