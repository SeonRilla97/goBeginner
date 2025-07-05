package koreaPost

import "fmt"

type PostSender struct {
}

func (f *PostSender) Send(parcel string) {
	fmt.Printf("우체국 택배 : %s\n", parcel)
}
