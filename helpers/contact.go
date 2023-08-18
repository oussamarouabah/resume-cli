package helpers

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
)

var (
	email    = "rouabaho@gmail.com"
	mobile   = "+213799839476"
	linkedIn = "https://www.linkedin.com/in/oussamarouabah/"
)

func Contact() {
	contactInfo := `

You can contact me at:
Email: %s
Mobile: %s,
LinkedIn: %s

`
	fmt.Printf(contactInfo, yellow(email), yellow(mobile), yellow(linkedIn))
}
