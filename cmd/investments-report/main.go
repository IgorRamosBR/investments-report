package main

import (
	"investments-report/internal/handlers"
	"investments-report/internal/io/reader"
	"investments-report/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	xlsxReader := reader.NewXlsxReader("")
	reportService := services.NewReportService(xlsxReader)
	reportHandler := handlers.NewReportHandler(reportService)

	r := gin.Default()
	r.GET("/report", reportHandler.GenerateReport)

	r.Run(":5000")
}
