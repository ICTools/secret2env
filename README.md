# Secret2Env

Secret2Env is a Go-based tool that retrieves secrets from AWS Secrets Manager and saves them into a .env file. This script is designed to facilitate secret management in environments that use .env files for configuration.


### Prerequisites

- AWS Account: You need an AWS account with secrets stored in AWS Secrets Manager.
- EC2 Instance with IAM Role: If running this script on an EC2 instance, the IAM role must have secretsmanager:GetSecretValue permission.

## Quick Start

Using the Pre-built Binary:

```bash
wget https://github.com/ICTools/secret2env/releases/download/v1.1.0/secret2env
chmod +x secret2env
```

Then, you can run the binary with options:

```bash
./secret2env --secretName "ictools" --region "eu-west-3"
```

## Running from Source Code (Go Required)

If you prefer to run Secret2Env from the source code, make sure Go (recommended version 1.23+) is installed and follow these steps.

### Installation

#### Clone the project:

```bash
git clone https://github.com/ICTools/secret2env.git
cd secret2env
```

#### Install dependencies:
```bash
go mod tidy
```

#### IAM Configuration

Ensure that the EC2 instance or AWS user running this script has the necessary permissions to access the secrets. A minimal IAM policy might look like this:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Statement1",
            "Effect": "Allow",
            "Action": [
                "secretsmanager:GetSecretValue"
            ],
            "Resource": [
                "arn:aws:secretsmanager:eu-west-3:account-id:secret:ictools"
            ]
        }
    ]
}
```

Replace account-id and ictools with the values specific to your account and secret. You can find the arn of your secrets manager directly in the secrets manager (via the AWS console).

#### Full Go command:

```bash
go run main.go --secretName "ictools" --region "eu-west-3" --outputDir "/path/to/directory" --fileName ".env.local" --versionStage "AWSCURRENT"
```

#### Using shortcuts:

```bash
go run main.go -s "ictools" -r "eu-west-3" -o "/path/to/directory" -f ".env.local" -v "AWSCURRENT"
```

#### Parameters

- `--secretName` (-s): The name of the secret in AWS Secrets Manager.

- `--region` (-r): The AWS region where the secret is stored (e.g., eu-west-3).

- `--outputDir` (-r) (OPTIONAL) : The directory where the file will be saved. Current directory by default.

- `--fileName` (-f) (OPTIONAL) : The name of the output file. ".env" by default.

- `--versionStage` (-v) (OPTIONAL) : The version stage of AWS secret. AWSCURRENT by default. 