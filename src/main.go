package main

import (
	"fmt"
	"skn-go-bill/src/pkgs/components"
)

func main() {
	var bill *components.BillType = components.CreateBill()

	fmt.Println(bill.Format())
}
