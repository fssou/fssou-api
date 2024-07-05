package secretsmanager

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log"
)

type SecretsManager struct {
	ctx    context.Context
	client *secretsmanager.Client
}

func New(ctx context.Context) *SecretsManager {
	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Println("error loading aws config")
		panic(err)
	}
	client := secretsmanager.NewFromConfig(awsConfig)
	return &SecretsManager{
		ctx:    ctx,
		client: client,
	}
}

func (s *SecretsManager) GetSecretValue(secretName string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}
	output, err := s.client.GetSecretValue(s.ctx, input)
	if err != nil {
		return "", err
	}
	return *output.SecretString, nil
}
