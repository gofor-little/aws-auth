package auth

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/gofor-little/xerror"
)

// SignIn will attempt to sign a user in, returning the result.
func SignIn(ctx context.Context, emailAddress string, password string) (*types.AuthenticationResultType, error) {
	if err := checkPackage(); err != nil {
		return nil, xerror.Wrap("checkPackage call failed", err)
	}

	output, err := CognitoClient.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME": emailAddress,
			"PASSWORD": password,
		},
		ClientId: aws.String(CognitoUserPoolClientID),
	})
	if err != nil {
		return nil, xerror.Wrap("failed to initiate auth", err)
	}

	return output.AuthenticationResult, nil
}
