package v1

import (
	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	e.POST("/storage_report", storageReportHandler)
	e.POST("/storage_recommendation", storageRecommendationHandler)
}
