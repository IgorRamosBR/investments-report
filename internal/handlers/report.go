package handlers

import (
	"investments-report/internal/services"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	ReportService services.ReportService
}

func NewReportHandler(reportService services.ReportService) ReportHandler {
	return ReportHandler{
		ReportService: reportService,
	}
}

func (h ReportHandler) GenerateReport(c *gin.Context) {
	report, err := h.ReportService.GenerateReport()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, report)
}
