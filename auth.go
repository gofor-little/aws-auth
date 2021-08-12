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
	// CognitoClient is the client that interacts with Cognito.
	CognitoClient           *cognitoidentityprovider.Client
	// CognitoUserPoolID is the id of the user pool in Cognito.
	CognitoUserPoolID       string
	// CognitoClientID is the id of the user pool client in Cognito.
	CognitoClientID string
)

// Initialize will initialize the auth package. Both profile and region parameters are option if authentication can be
// achieved via another method. For example, environment variables or IAM roles.
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
	CognitoClientID = cognitoClientID

	return nil
}

func checkPackage() error {
	if CognitoClient == nil {
		return xerror.New("db.CognitoClient is nil, have you called auth.Initialize()?")
	}

	if CognitoClientID == "" {
		return xerror.New("db.CognitoClientID is empty, did you call auth.Initialize()?")
	}

	return nil
}
