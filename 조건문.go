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

switch i {
case 0: //i가 0일 때 아무것도 하지 않고 나감
case 1: //i가 1일떄만
    f() //수행한다.
}

switch i {
case 0:
	fallthrough
case 1:
	f()  //결국 f함수는 i가 0,1일 때 모두 수행하는 것임
}



switch level {
case 1:
	doFirst()   //doFirst는 레벨 1일 때 수행함
	fallthrough
case 2:
	doSecond()  //레벨 1,2는 doSecond를 수행함
	fallthrough
case 3:
	doThird() //레벨 1,2,3은 doThird를 수행함
}

//for 문
// 다른 언어와 다르게 go는 while이 없음. 모든 조건문은 다 for문으로 작성해준다. 


//for 초기화구문;조건식;후속작업 {
//	..어쩌구저쩌구
//}

//for {
//
//}
//이렇게 되면 for 문 내부의 코드를 무한반복하게 됨 

//break & Continue
//for 문을 강제로 종료해야 할 때는 break 키워드를 사용한다. 현재 수행하는 반복 작업을 건너뛰고 다음 반복 작업을 수행해야 할 때는 continue 키워드를 사용한다.

for {
	i = i-1 
	...
	if i<0 {
		break
	}
	if i%2 == 0 {
		continue
	}
}
for {
	i = i+1 
	if i >100 {
		break
	} if i % 2 == 0 {
		continue
	}
}

