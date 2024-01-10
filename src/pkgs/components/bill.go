package components

import (
	"bufio"
	"fmt"
	"os"
	"skn-go-bill/src/pkgs/utils"
	"strconv"
	"strings"
)

func getPrompt(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	return input, err
}

func (bill *BillType) Format() string {
	var formattedString string = "Bill Info \n"

	formattedString += fmt.Sprintf("Title: %s \n", bill.Title)

	var total float32 = 0.0

	for _, value := range bill.Items {
		var subTotal float32 = value.Price * float32(value.Count)
		formattedString += fmt.Sprintf("%-25s $%f \n", value.Name+":", subTotal)
		total += subTotal
	}

	total += bill.Tip

	formattedString += fmt.Sprintf("%-25s $%f \n", "Tip:", bill.Tip)
	formattedString += fmt.Sprintf("%-25s $%f \n", "Total:", total)

	return formattedString
}

func processBill(reader *bufio.Reader, bill *BillType) {
	var option string = `
		Choose an option:
		(a) Press 'a' to add items
		(b) Press 't' to add tips
		(c) Press 's' to save the bill
		Your Selection: `

	input, _ := getPrompt(reader, option)

	switch input {
	case "a":
		name, _ := getPrompt(reader, "Enter Item Name: ")
		count, _ := getPrompt(reader, "Enter Item Count: ")
		price, _ := getPrompt(reader, "Enter Item Price: ")

		count64, _ := strconv.ParseUint(count, 10, 8)
		var count8 uint8 = uint8(count64)

		price64, _ := strconv.ParseFloat(price, 32)
		var price32 float32 = float32(price64)

		var item *ItemType = &ItemType{
			Name:  name,
			Count: count8,
			Price: price32,
		}

		bill.Items = append(bill.Items, *item)

		processBill(reader, bill)
	case "t":
		tip, _ := getPrompt(reader, "Enter Tip Amount: ")

		tip64, _ := strconv.ParseFloat(tip, 32)
		var tip32 float32 = float32(tip64)

		bill.Tip = tip32

		processBill(reader, bill)
	case "s":
		var data []byte = []byte(bill.Format())

		var err error = os.WriteFile("logs/"+bill.Title+".log", data, 0644)

		if err != nil {
			panic(err)
		}

		fmt.Println("Bill Created Succesfully!")
	default:
		fmt.Println("Invalid Selection!")

		processBill(reader, bill)
	}
}

func CreateBill() *BillType {

	var reader *bufio.Reader = utils.GetReader()

	title, _ := getPrompt(reader, "Enter The Title For The Bill: ")

	var items []ItemType = []ItemType{}

	var bill *BillType = &BillType{
		Title: title,
		Items: items,
		Tip:   0,
	}

	processBill(reader, bill)

	return bill
}
