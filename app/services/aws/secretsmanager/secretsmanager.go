package secretsmanager

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	"log"
)

func init() {
	log.Println("SecretsManager init")
}

type SecretsManager struct {
	ctx         context.Context
	client      *secretsmanager.Client
	secretCache *secretcache.Cache
}

func New(ctx context.Context) *SecretsManager {
	secretCache, _ := secretcache.New()
	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Println("error loading aws config")
		panic(err)
	}
	client := secretsmanager.NewFromConfig(awsConfig)
	return &SecretsManager{
		ctx:         ctx,
		client:      client,
		secretCache: secretCache,
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

func (s *SecretsManager) GetSecretValueWithCache(secretName string) (string, error) {
	secret, err := s.secretCache.GetSecretString(secretName)
	if err != nil {
		return "", err
	}
	return secret, nil
}
