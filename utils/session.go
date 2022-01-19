package utils

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionMessage struct {
	Success string `json:"success"`
	Error   string `json:"error"`
}

// GetSuccessErrorSession returns first parameter as error message second as success message
func GetSuccessErrorSession(c *gin.Context) SessionMessage {
	var errorItem string
	var success string
	session := sessions.Default(c)
	val := session.Flashes("error")
	if val != nil {
		errorItem = val[0].(string)
	}
	val = session.Flashes("success")
	if val != nil {
		success = val[0].(string)
	}
	err := session.Save()
	if err != nil {
		fmt.Println(err.Error())
	}

	return SessionMessage{
		Success: success,
		Error:   errorItem,
	}

}
