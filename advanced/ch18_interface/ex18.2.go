package main

import (
	"basic_go/advanced/ch18_interface/fedex"
	"basic_go/advanced/ch18_interface/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

// Send 함수만 제공하면 어떤 구조체든지 상관하지 않는다.
func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)

	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
}
