package controllers

import (
	"net/http"

	"github.com/emikohmann/arq-software/ej-auth/domain"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var cred domain.Credentials
	if err := c.BindJSON(&cred); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := service.Login(cred)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		//abort with startus
	}
	//unmarshall body
	//call to service

	c.JSON(http.StatusOK, token)
}
