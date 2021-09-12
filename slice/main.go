package main

import (
	"fmt"
	"sort"
	"time"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 31},
		{"백두산", 52},
		{"류", 42},
		{"켄", 38},
		{"송하나", 18},
	}
	sort.Sort(Students(s))
	fmt.Println(s)
}

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	ID    int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(pdt *Product) *Parcel {
	var p = &Parcel{Pdt: pdt, ShippedTime: time.Now()}
	return p
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}
