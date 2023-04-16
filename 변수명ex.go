package main
import "fmt"

const (
	Running = 1  <<iota 
	Waiting  //1
	Send //1
	Receive //1
)  //상수 네개 : Running , Waiting , Send, Receive를 그룹으로 선언함 
//Running에만 1을 넣어주고 나머지는 넣어주지 않음 -> 나머지 변수에도 1이 자동할당된다는 의미

//Running << iota : 1

func main() {
	start := Running | Send
	display(start)
}

func display(start int) {
	if start&Running == Running {
		ftm.Println("Running")
	}
	if start&Waiting == Waiting {
		fmt.Println("Waiting")
	}
	if start&Send == Send {
		fmt.Println("Send")
	}
	if start&Receive == Receive {
		fmt.Println("Receive")
	}
}