package fedex

import "fmt"

type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex Send : %s\n", parcel)
}
