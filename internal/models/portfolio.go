package models

type Portfolio struct {
	Stocks         []Product
	AmericanStocks []Product
	RealStateFund  []Product
	Cryptos        []Product
	Cash           []Product
}
