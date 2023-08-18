package helpers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()

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

func About() {
	fmt.Println("")
	fmt.Printf("I am a proficient and driven %s with extensive experience in emerging technologies.\n", yellow("full-stack developer"))
	fmt.Printf("I possess diverse technical skills, including proficiency in programming languages such as %s.\n", yellow("C++, Python, JavaScript, Solidity, Rust, and more"))
	fmt.Printf("With practical knowledge in %s I have the capacity to tackle complex projects in the tech industry.\n\n", yellow("blockchain, web3, and machine learning."))
	fmt.Printf("I have honed my technical skills through %s at multiple companies, and %s and am excited to continue building my skills and contributing to %s in the tech industry.\n", yellow("internships"), yellow("online courses"), yellow("innovative projects"))
	fmt.Printf("Along with my technical expertise, I possess strong soft skills such as %s. I am a team player and enjoy collaborating with others to achieve common goals.\n", yellow("communication, critical thinking, and problem-solving"))
	fmt.Printf("Furthermore, I have demonstrated natural leadership skills through various group projects.\n\n\n")
}

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
