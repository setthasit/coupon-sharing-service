package controllers

import "github.com/gin-gonic/gin"

func responsMessageHttp(c *gin.Context, status int, message string) {
	type baseResponse struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	}

	c.JSON(status, &baseResponse{
		StatusCode: status,
		Message:    message,
	})
}

func responseItemHttp(c *gin.Context, status int, item interface{}) {
	type baseResponse struct {
		StatusCode int         `json:"status_code"`
		Item       interface{} `json:"item"`
	}

	c.JSON(status, &baseResponse{
		StatusCode: status,
		Item:       item,
	})
}

func responseListHttp(c *gin.Context, status int, items []interface{}, length int) {
	type baseResponse struct {
		StatusCode int           `json:"status_code"`
		Items      []interface{} `json:"items"`
		ItemLength int           `json:"item_length"`
	}

	c.JSON(status, &baseResponse{
		StatusCode: status,
		Items:      items,
		ItemLength: length,
	})
}
