package auth

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gofor-little/xerror"
)

// ForgotPassword will initiate a forgot password request.
func ForgotPassword(ctx context.Context, emailAddress string) (*cognitoidentityprovider.ForgotPasswordOutput, error) {
	output, err := CognitoClient.ForgotPassword(ctx, &cognitoidentityprovider.ForgotPasswordInput{
		ClientId: aws.String(CognitoClientID),
		Username: aws.String(emailAddress),
	})
	if err != nil {
		return nil, xerror.Wrap("failed to send forgot password request", err)
	}

	return output, nil
}

// ForgotPasswordConfirm will confirm a forgot password request.
func ForgotPasswordConfirm(ctx context.Context, confirmationCode string, emailAddress string, newPassword string) error {
	if _, err := CognitoClient.ConfirmForgotPassword(ctx, &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         aws.String(CognitoClientID),
		ConfirmationCode: aws.String(confirmationCode),
		Password:         aws.String(newPassword),
		Username:         aws.String(emailAddress),
	}); err != nil {
		return xerror.Wrap("failed to send forgot password confirmation request", err)
	}

	return nil
}
