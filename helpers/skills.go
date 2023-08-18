package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

type Skills struct {
	WebTech   string `json:"web-technologies"`
	Languages string `json:"web-languages"`
	AIML      string `json:"ai-ml"`
	DevOps    string `json:"dev-ops"`
}

func ShowSkills(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var sk Skills
	err = json.Unmarshal(content, &sk)
	if err != nil {
		return err
	}

	fmt.Printf("\nLanguages: %s\n", green(sk.Languages))
	fmt.Printf("\nWeb Technologies: %s\n", green(sk.WebTech))
	fmt.Printf("\nAI/ML: %s\n", green(sk.AIML))
	fmt.Printf("\nDevOps: %s\n", green(sk.DevOps))

	return nil
}
