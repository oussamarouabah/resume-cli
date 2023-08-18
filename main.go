package main

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/oussamarouabah/resume-cli/helpers"
)

// the questions to ask
var qs = &survey.Question{

	Name: "choice",
	Prompt: &survey.Select{
		Message: "What would you like to know?",
		Options: []string{"About", "Experience", "Skills", "Contact", "Exit"},
		Default: "About",
	},
}

func welcome() {
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Println("")
	fmt.Println("")
	fmt.Printf(
		"Hey there! I'm %s, a full stack web3 developer and currently learning new technologies.\n\n",
		yellow("Oussama ROUABAH"),
	)
}

func main() {
	welcome()

	answer := struct {
		Choice string `survey:"choice"`
	}{}

	for {
		err := survey.Ask([]*survey.Question{qs}, &answer)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		switch answer.Choice {
		case "About":
			helpers.About()
		case "Experience":
			err := helpers.ShowExperience("./data/experience/experience.json")
			if err != nil {
				log.Fatal(err)
			}
		case "Skills":
			err := helpers.ShowSkills("./data/skills/skills.json")
			if err != nil {
				log.Fatal(err)
			}
		case "Contact":
			helpers.Contact()
		case "Exit":
			fmt.Println("Bye! Have a great day!")
			return
		}
	}

}
