package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/highhi/slspolicygen"
	"github.com/manifoldco/promptui"
)

func main() {
	templates := &promptui.PromptTemplates{
		Prompt:  promptui.IconInitial + " {{ . | cyan | bold }} ",
		Valid:   promptui.IconInitial + " {{ . | cyan | bold }} ",
		Invalid: promptui.IconInitial + " {{ . | cyan | bold }} ",
		Success: promptui.IconInitial + " {{ . | cyan | bold }} ",
	}

	serviceNamePrompt := promptui.Prompt{
		Label:     "Your Serverless service name",
		Templates: templates,
		Validate: func(input string) error {
			if input == "" {
				return errors.New("service name is required")
			}
			return nil
		},
	}

	accountIdPrompt := promptui.Prompt{
		Label:     "Your AWS account id",
		Templates: templates,
		Validate: func(input string) error {
			if input == "" {
				return errors.New("account id is required")
			}
			return nil
		},
	}

	dynamoDBPrompt := promptui.Prompt{
		Label:     "Does your service rely on DynamoDB?[y/n]",
		Templates: templates,
		Validate: func(input string) error {
			if input != "y" && input != "n" {
				return errors.New("please `y` or `n`")
			}
			return nil
		},
	}

	s3Prompt := promptui.Prompt{
		Label:     "Does your service rely on S3?[y/n]",
		Templates: templates,
		Validate: func(input string) error {
			if input != "y" && input != "n" {
				return errors.New("please `y` or `n`")
			}
			return nil
		},
	}

	serviceName, _ := serviceNamePrompt.Run()
	accountID, _ := accountIdPrompt.Run()
	dynamoDB, _ := dynamoDBPrompt.Run()
	s3, _ := s3Prompt.Run()

	file, err := os.Create("./sls-policy.json")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = slspolicygen.Gen(file, map[string]interface{}{
		"account":          accountID,
		"servicename":      serviceName,
		"dynamoDBRequired": dynamoDB == "y",
		"s3Required":       s3 == "y",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(promptui.IconGood + " Generated sls-policy.json!")
}
