package secretsmanager

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

var secretCache *secretcache.Cache

func init() {
	secretCache, _ = secretcache.New()
}

type SecretsManager struct {
	ctx    context.Context
	client *secretsmanager.Client
}

func New(ctx context.Context) *SecretsManager {
	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
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
		SecretId:  aws.String(secretName),
		VersionId: aws.String("AWSCURRENT"),
	}
	output, err := s.client.GetSecretValue(s.ctx, input)
	if err != nil {
		return "", err
	}
	return *output.SecretString, nil
}

func (s *SecretsManager) GetSecretValueWithCache(secretName string) (string, error) {
	secret, err := secretCache.GetSecretString(secretName)
	if err != nil {
		return "", err
	}
	return secret, nil
}
