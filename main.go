package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kavenegar/kavenegar-go"
	"os"
)

type (
	Input struct {
		Alerts []Alert `json:"alerts"`
	}

	Alert struct {
		Status string `json:"status"`
		Label Label `json:"labels"`
	}

	Label struct {
		AlertName string `json:"alertname"`
	}
)
func main() {
	r := gin.Default()
	api := kavenegar.New(os.Getenv("KAVE_API_KEY"))
	sender := ""
	receptor := []string{os.Getenv("RECIPIENT")}

	r.POST("/alert/:region/:severity", func(c *gin.Context) {

		input := &Input{}
		c.BindJSON(input)
		message := fmt.Sprintf("Alert in %s [%s]\n", c.Param("region"), c.Param("severity"))
		for _, alert := range input.Alerts{
			message += fmt.Sprintf("Status: %s\nLabel: %s\n", alert.Status, alert.Label.AlertName)
		}

		if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
			switch err := err.(type) {
			case *kavenegar.APIError:
				fmt.Println(err.Error())
			case *kavenegar.HTTPError:
				fmt.Println(err.Error())
			default:
				fmt.Println(err.Error())
			}
			c.String(500, "ERR")
		} else {
			for _, r := range res {
				fmt.Println("MessageID 	= ", r.MessageID)
				fmt.Println("Status    	= ", r.Status)
			}
			c.String(200, "OK")
		}

		c.String(200, "OK")
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
