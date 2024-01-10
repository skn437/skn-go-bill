package components

type ItemType struct {
	Name  string
	Price float32
	Count uint8
}

type BillType struct {
	Title string
	Items []ItemType
	Tip   float32
}
