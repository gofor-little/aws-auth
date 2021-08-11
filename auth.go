package auth

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gofor-little/xerror"
)

var (
	CognitoClient           *cognitoidentityprovider.Client
	CognitoUserPoolID       string
	CognitoUserPoolClientID string
)

func Initialize(ctx context.Context, profile string, region string, cognitoUserPoolID string, cognitoClientID string) error {
	var cfg aws.Config
	var err error

	if profile != "" && region != "" {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile), config.WithRegion(region))
	} else {
		cfg, err = config.LoadDefaultConfig(ctx)
	}
	if err != nil {
		return fmt.Errorf("failed to load default config: %w", err)
	}

	CognitoClient = cognitoidentityprovider.NewFromConfig(cfg)
	CognitoUserPoolID = cognitoUserPoolID
	CognitoUserPoolClientID = cognitoClientID

	return nil
}

func checkPackage() error {
	if CognitoClient == nil {
		return xerror.New("db.CognitoClient is nil, have you called auth.Initialize()?")
	}

	if CognitoUserPoolClientID == "" {
		return xerror.New("db.CognitoClientID is empty, did you call auth.Initialize()?")
	}

	return nil
}
