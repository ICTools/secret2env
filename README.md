# Secret2Env

Secret2Env is a Go-based tool that retrieves secrets from AWS Secrets Manager and saves them into a .env file. This script is designed to facilitate secret management in environments that use .env files for configuration.

## Prerequisites

- Go: Make sure Go is installed (recommended version 1.18+).
- AWS Account: You need an AWS account with secrets stored in AWS Secrets Manager.
- EC2 Instance with IAM Role: If running this script on an EC2 instance, the IAM role must have secretsmanager:GetSecretValue permission.

## Installation

### Clone the project:

```bash
git clone https://github.com/your-username/secret2env.git
cd secret2env
```


### Initialize the Go module (if not already done):

```bash
go mod init secret2env
```

### Install dependencies:
```bash
go mod tidy
```

### IAM Configuration

Ensure that the EC2 instance or AWS user running this script has the necessary permissions to access the secrets. A minimal IAM policy might look like this:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "secretsmanager:GetSecretValue",
      "Resource": "arn:aws:secretsmanager:eu-west-3:account-id:secret:ictools"
    }
  ]
}
```

Replace account-id and ictools with the values specific to your account and secret.

## Usage

This script supports two primary options: --secretName (or -s for shorthand) and --region (or -r for shorthand).

### Full command:

```bash
go run main.go --secretName "ictools" --region "eu-west-3"
```

### Using shortcuts:

```bash
go run main.go -s "ictools" -r "eu-west-3"
```

### Parameters

- `--secretName` (-s): The name of the secret in AWS Secrets Manager.

- `--region` (-r): The AWS region where the secret is stored (e.g., eu-west-3).