package models

type Report struct {
	Stocks         Item
	AmericanStocks Item
	RealStateFund  Item
	Cryptos        Item
	Cash           Item
}

func (r Report) CalculateQuota() {

}
