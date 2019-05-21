package main

import "fmt"

func TwentyNine() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Received message: ", msg)

	default:
		fmt.Println("No message received")
	}
	msg := "hi"
	select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
	}
	
	select {
	case messages <- msg:
		fmt.Println("Recived Message", msg)
	case sig := <- signals:
		fmt.Println("received Signal: ", sig)
	default:
		fmt.Println("No activity detected")
	}

}
