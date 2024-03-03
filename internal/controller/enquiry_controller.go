package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EnquiryController struct {
}

func NewEnquiryController() *EnquiryController {
	return &EnquiryController{}
}

// just define a simple endpoint to directly return save enquiry
func (ec *EnquiryController) SaveEnquiry(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Enquiry saved"})
}
