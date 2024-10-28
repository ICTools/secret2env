package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func main() {
	secretName := flag.String("secretName", "", "The name of the secret in AWS Secrets Manager")
	region := flag.String("region", "", "The AWS region where the secret is stored")
	flag.StringVar(secretName, "s", "", "Shortcut for secretName")
	flag.StringVar(region, "r", "", "Shortcut for region")
	flag.Parse()

	if *secretName == "" || *region == "" {
		log.Fatal("Both --secretName (or -s) and --region (or -r) parameters are required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(*region))
	if err != nil {
		log.Fatalf("Error loading AWS configuration: %v", err)
	}

	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(*secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(ctx, input)
	if err != nil {
		log.Fatalf("Unable to retrieve secret: %v", err)
	}

	var secretData map[string]string
	if err := json.Unmarshal([]byte(*result.SecretString), &secretData); err != nil {
		log.Fatalf("Unable to parse secret string: %v", err)
	}

	if _, err := os.Stat(".env"); err == nil {
		log.Println(".env file already exists. It will be overwritten.")
	}

	file, err := os.Create(".env")
	if err != nil {
		log.Fatalf("Unable to create .env file: %v", err)
	}
	defer file.Close()

	for key, value := range secretData {
		_, err := file.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		if err != nil {
			log.Fatalf("Unable to write to .env file: %v", err)
		}
	}

	log.Println(".env file created successfully with secret values")
}
