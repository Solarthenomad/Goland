package main
import "fmt"

const (
	Running = 1  <<iota 
	Waiting 
	Send
	Receive
)

func main() {
	start := Running | Send
	display(start)
}

func display(start int) {
	
}