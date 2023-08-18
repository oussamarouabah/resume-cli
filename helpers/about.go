package helpers

import "fmt"

func About() {
	fmt.Println("")
	fmt.Printf("I am a proficient and driven %s with extensive experience in emerging technologies.\n", yellow("full-stack developer"))
	fmt.Printf("I possess diverse technical skills, including proficiency in programming languages such as %s.\n", yellow("C++, Python, JavaScript, Solidity, Rust, and more"))
	fmt.Printf("With practical knowledge in %s I have the capacity to tackle complex projects in the tech industry.\n\n", yellow("blockchain, web3, and machine learning."))
	fmt.Printf("I have honed my technical skills through %s at multiple companies, and %s and am excited to continue building my skills and contributing to %s in the tech industry.\n", yellow("internships"), yellow("online courses"), yellow("innovative projects"))
	fmt.Printf("Along with my technical expertise, I possess strong soft skills such as %s. I am a team player and enjoy collaborating with others to achieve common goals.\n", yellow("communication, critical thinking, and problem-solving"))
	fmt.Printf("Furthermore, I have demonstrated natural leadership skills through various group projects.\n\n\n")
}
