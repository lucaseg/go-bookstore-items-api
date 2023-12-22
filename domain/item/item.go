package item

type Item struct {
	Id                string
	Seller            string
	Title             string
	Description       Description
	Pictures          []Picture
	Video             string
	Price             float32
	AvailableQuantity int
	SoldQuantity      int
	Status            string
}

type Description struct {
	PlainText string
	Html      string
}

type Picture struct {
	Id  int64
	Url string
}
