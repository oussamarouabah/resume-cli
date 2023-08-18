package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

type Expariences struct {
	Experiences []Experience `json:"Experiences"`
}

type Experience struct {
	Position    string   `json:"Position"`
	Company     string   `json:"Company"`
	Location    string   `json:"Location"`
	Duration    string   `json:"Duration"`
	Description []string `json:"Description"`
}

func ShowExperience(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var exp Expariences
	err = json.Unmarshal(content, &exp)
	if err != nil {
		return err
	}

	for idx, ex := range exp.Experiences {
		fmt.Printf("\n%s : #%d\n", yellow("Experience"), idx+1)
		fmt.Printf("%s : %s\n", yellow("Postion"), green(ex.Position))
		fmt.Printf("%s : %s\n", yellow("Company"), green(ex.Company))
		fmt.Printf("%s : %s\n", yellow("Location"), green(ex.Location))
		fmt.Printf("%s : %s\n", yellow("Duration"), green(ex.Duration))
		fmt.Printf("%s:\n", yellow("Description"))
		for _, d := range ex.Description {
			fmt.Printf("- %s\n", green(d))
		}
	}

	return nil
}
