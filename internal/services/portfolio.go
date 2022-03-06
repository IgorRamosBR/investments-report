package services

import (
	"investments-report/internal/io/reader"
	"investments-report/internal/models"
	"strconv"
	"strings"
)

type PortfolioService interface {
	GetPortfolio() (models.Portfolio, error)
}

type PortfolioServiceImpl struct {
	Reader reader.XlsxReader
}

func NewPortfolioService(reader reader.XlsxReader) PortfolioService {
	return PortfolioServiceImpl{
		Reader: reader,
	}
}

func (s PortfolioServiceImpl) GetPortfolio() (models.Portfolio, error) {
	rows, err := s.Reader.ReadFile()
	if err != nil {
		return models.Portfolio{}, err
	}

	portfolio := models.Portfolio{}
	for i := 1; i < len(rows); i++ {
		product, err := s.createProduct(rows, i)
		if err != nil {
			return models.Portfolio{}, err
		}

		s.addProduct(&portfolio, product)
	}
	return portfolio, nil
}

func (s PortfolioServiceImpl) createProduct(rows [][]string, i int) (models.Product, error) {
	value, err := strconv.ParseFloat(strings.ReplaceAll(rows[i][4], ",", "."), 32)
	if err != nil {
		return models.Product{}, err
	}
	balance, err := strconv.ParseFloat(strings.ReplaceAll(rows[i][5], ",", "."), 32)
	if err != nil {
		return models.Product{}, err
	}
	profitability, err := strconv.ParseFloat(strings.ReplaceAll(rows[i][6], ",", "."), 32)
	if err != nil {
		return models.Product{}, err
	}
	product := models.Product{
		Name:          rows[i][0],
		Class:         rows[i][1],
		Broker:        rows[i][2],
		Value:         value,
		Balance:       balance,
		Profitability: profitability,
	}
	return product, nil
}

func (s PortfolioServiceImpl) addProduct(report *models.Portfolio, product models.Product) {
	if product.Broker == "AVENUE" {
		report.AmericanStocks = append(report.AmericanStocks, product)
		return
	}
	switch product.Class {
	case "Ação":
		report.Stocks = append(report.Stocks, product)
	case "Fundo Imobiliário":
		report.RealStateFund = append(report.RealStateFund, product)
	case "Criptomoeda":
		report.Cryptos = append(report.Cryptos, product)
	default:
		report.Cash = append(report.Cash, product)
	}
}
