package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type firstJson struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

type operationExpectedJson struct {
	OperationType string `json:"operation_type" validate:"required,eq=addition|eq=subtraction|eq=multiplication"`
	X             int64  `json:"x"`
	Y             int64  `json:"y"`
}
type operationJsonResponse struct {
	SlackUsername string `json:"slackUsername"`
	OperationType string `json:"operation_type" validate:"required,eq=addition|eq=subtraction|eq=multiplication"`
	Result        int64  `json:"result"`
}

func main() {
	port := "3000"
	r := gin.New()
	//r.GET("/", giveResponse())
	r.POST("/", performOperation())
	r.Run(":" + port)

	//.Println("Server listening on port: " + port)
}

func performOperation() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var operation operationExpectedJson
		operation := &operationExpectedJson{
			OperationType: "addition",
			X:             8,
			Y:             8,
		}
		var operationResp operationJsonResponse
		var validate = validator.New()
		if err := c.BindJSON(&operation); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if validationErr := validate.Struct(operation); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		operationResp.SlackUsername = "Meshach"

		switch operation.OperationType {
		case "addition":
			result := operation.X + operation.Y
			operationResp.Result = result
			operationResp.OperationType = operation.OperationType
		case "subtraction":
			result := operation.X - operation.Y
			operationResp.Result = result
			operationResp.OperationType = operation.OperationType
		case "multiplication":
			result := operation.X * operation.Y
			operationResp.Result = result
			operationResp.OperationType = operation.OperationType
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Operation_type is unknown"})
		}
		c.JSON(http.StatusOK, operationResp)

	}
}

func giveResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		myRes := &firstJson{
			SlackUsername: "Meshach",
			Backend:       true,
			Age:           22,
			Bio:           "A junior backend developer",
		}
		//if err := c.BindJSON(&myRes); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		//	return
		//}
		c.JSON(http.StatusOK, myRes)

	}
}
