package fedex

import "fmt"

type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("FEdex sends %s parcel\n", parcel)
}
