//var 키워드로 선언한 뒤에 변수 이름과 타입을 표기한다.
//변수를 선언할 때 초기값을 지정하지 않으면 제로값으로 초기화해준다. 
//정수 타입의 제로값 : 0
//실수 타입의 제로 값 : 0.0
//문자열 타입의 제로 값 : ""
package main

var a int
var b string

//변수 여러 개를 한 번에 선언해주기 

var name, id, address string

//타입이 다른 변수 여러 개를 한꺼번에 선언해주기
var (
	name string
	age int
	weight float
)


// 변수 선언과 동시에 값을 할당할 때는 타입을 생략할 수 있다.

var c = true

//이때 타입을 생략하면 변수의 타입은 초깃값의 타입으로 정해지게 된다. 

//위와 같은 선언 방식말고도 엄청 컴팩트한 선언 방식이 있는데, 이를 짧은 선언이라하고 var을 생략한다.
start :=1 
//이미 선언된 변수에 값을 할당할 때는 :=를 사용하는 것이 불가능
//:=는 변수 선언과 동시에 초깃값을 할당할 때만 사용 

if v:= getValue() {
	fmt.Println(v)
}

//go는 대소문자 구분 : NUMBER, number , Number 세개는 모두 다 다른 변수가 됨 

//특정 객체를 반환하는(return) 함수나 메서드의 이름은 명사형으로 짓고 + get을 붙이지 않는다.

fun Connection() *Conn {
	//GetConnection()으로 쓰지 않는다.

	//..어쩌구
	return conn
}

//특정 객체를 변형하거나 설정하는 함수의 이름 앞에는 set을 붙인다.
func SetId(i int) {...}

//상수 : const , 불, 숫자, 문자열 타입으로만 선언 가능

const limit = 64 //상수는 타입을 선언하지 않아도 됨 

const max uint64 = 1024 
//이거 맞긴한데 굳이 쓸필요는 없음 

const max = 1024 * 1024
const c = getNumber() //함수는 상수형 변수의 값이 되는 것이 불가능하다.

const (
	RED= 0
	ORANGE = 1
	YELLOW =2
)

// 열거형 
//차례로 1씩 증가하는 상수의 묶음 
//일일이 만들 때
const (
	Sunday =0
	Monday = 1
	Tuesday = 2
	Thursday = 3
	Friday = 4
	Saturday = 5
)

//iota를 쓰면 위와같은 개고생을 할필요가 없음
const (
	Sunday = iota
	Monday
	Tuesday
	Thursday
	Friday
	Saturday
)

type Color int 

const (
	RED Color = iota  //0 : 새로운 const Group인 Color이 선언되었으므로 0부터 시작하게 됨 
	ORANGE  //1
	YELLOW  //2
	GREEN  //3
)

//정수뿐만 아니라 부동소수점 타입에서 사용하는 것이 가능하다.

type ByteSize int64  //ByteSize 

const (
	_=iota
	KB
	MB
	GB 
	TB
	PB
	EB


)

const (
	DEFAULT_RATE = 5+0.3 * iota //5
	GREEN_RATE     //5.3
	SILVER_RATE    //5.6
	GOLD_RATE   //5.9
)