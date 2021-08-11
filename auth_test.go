package auth_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/gofor-little/env"
	"github.com/stretchr/testify/require"

	auth "github.com/gofor-little/aws-auth"
)

func setup(t *testing.T) {
	if err := env.Load(".env"); err != nil {
		t.Logf("failed to load .env file, ignore if running in CI/CD: %v", err)
	}

	require.NoError(t, auth.Initialize(
		context.Background(),
		env.Get("AWS_PROFILE", ""),
		env.Get("AWS_REGION", ""),
		// Normally the user pool id and user pool client id would be passed here when the package
		// is initialized. Because they are created and destroyed for each test they're set below.
		"",
		"",
	))

	// Create the user pool and set the id for the auth package.
	userPoolOutput, err := auth.CognitoClient.CreateUserPool(context.Background(), &cognitoidentityprovider.CreateUserPoolInput{
		PoolName: aws.String(fmt.Sprintf("waste-scanner-test-user-pool_%d", time.Now().UnixNano())),
	})
	require.NoError(t, err)
	auth.CognitoUserPoolID = *userPoolOutput.UserPool.Id

	// Create the user pool client and set the id for the auth package.
	userPoolClientOutput, err := auth.CognitoClient.CreateUserPoolClient(context.Background(), &cognitoidentityprovider.CreateUserPoolClientInput{
		ClientName: aws.String(fmt.Sprintf("waste-scanner-test-user-pool-client_%d", time.Now().UnixNano())),
		UserPoolId: aws.String(auth.CognitoUserPoolID),
		ExplicitAuthFlows: []types.ExplicitAuthFlowsType{
			types.ExplicitAuthFlowsTypeAllowUserPasswordAuth,
			types.ExplicitAuthFlowsTypeAllowRefreshTokenAuth,
		},
	})
	require.NoError(t, err)
	auth.CognitoUserPoolClientID = *userPoolClientOutput.UserPoolClient.ClientId
}

func teardown(t *testing.T) {
	_, err := auth.CognitoClient.DeleteUserPool(context.Background(), &cognitoidentityprovider.DeleteUserPoolInput{
		UserPoolId: aws.String(auth.CognitoUserPoolID),
	})
	require.NoError(t, err)
}
