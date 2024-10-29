package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func main() {
	secretName := flag.String("secretName", "", "The name of the secret in AWS Secrets Manager")
	region := flag.String("region", "", "The AWS region where the secret is stored")
	outputDir := flag.String("outputDir", ".", "The directory where the .env file will be saved. Current folder by default.")
	fileName := flag.String("fileName", ".env", "The output filename. '.env' by default.")
	versionStage := flag.String("versionStage", "AWSCURRENT", "Version stage of AWS secret. AWSCURRENT by default.")

	flag.StringVar(secretName, "s", "", "Shortcut for secretName")
	flag.StringVar(region, "r", "", "Shortcut for region")
	flag.StringVar(outputDir, "o", ".", "Shortcut for outputDir")
	flag.StringVar(fileName, "f", ".", "Shortcut for fileName")
	flag.StringVar(versionStage, "v", ".", "Shortcut for versionStage")
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
		VersionStage: aws.String(*versionStage),
	}

	result, err := svc.GetSecretValue(ctx, input)
	if err != nil {
		log.Fatalf("Unable to retrieve secret: %v", err)
	}

	var secretData map[string]string
	if err := json.Unmarshal([]byte(*result.SecretString), &secretData); err != nil {
		log.Fatalf("Unable to parse secret string: %v", err)
	}

	envFilePath := filepath.Join(*outputDir, *fileName)

	if _, err := os.Stat(envFilePath); err == nil {
		log.Printf("File already exists at %s. It will be overwritten.\n", envFilePath)
	}

	file, err := os.Create(envFilePath)
	if err != nil {
		log.Fatalf("Unable to create %s: %v", envFilePath, err)
	}
	defer file.Close()

	for key, value := range secretData {
		_, err := file.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		if err != nil {
			log.Fatalf("Unable to write to %s: %v", envFilePath, err)
		}
	}

	log.Printf("%s created successfully\n", envFilePath)
}
