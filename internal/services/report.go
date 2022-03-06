package services

import (
	"investments-report/internal/models"
)

type ReportService interface {
	GenerateReport() (models.Report, error)
}

type ReportServiceImpl struct {
	PortfolioService PortfolioService
}

func NewReportService(portfolioService PortfolioServiceImpl) ReportService {
	return ReportServiceImpl{
		PortfolioService: portfolioService,
	}
}

func (s ReportServiceImpl) GenerateReport() (models.Report, error) {
	portfolio, err := s.PortfolioService.GetPortfolio()
	if err != nil {
		return models.Report{}, err
	}

	report := s.createReport(portfolio)

	return report, nil
}

func (s ReportServiceImpl) createReport(portfolio models.Portfolio) models.Report {
	report := models.Report{
		Stocks:         s.calculateItem(portfolio.Stocks),
		AmericanStocks: s.calculateItem(portfolio.AmericanStocks),
		RealStateFund:  s.calculateItem(portfolio.RealStateFund),
		Cryptos:        s.calculateItem(portfolio.Cryptos),
		Cash:           s.calculateItem(portfolio.Cash),
	}
	report.CalculateQuota()
	return report
}

func (s ReportServiceImpl) calculateItem(products []models.Product) models.Item {
	var amountValue, amountBalance float64

	for _, product := range products {
		amountValue += product.Value
		amountBalance += product.Balance
	}

	return models.Item{
		Products:      products,
		Amount:        amountBalance,
		Profitability: (amountBalance - amountValue) / amountValue * 100,
	}
}
