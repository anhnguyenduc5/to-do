package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandlerV1(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	return r
}
