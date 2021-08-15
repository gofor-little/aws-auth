package auth

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gofor-little/xerror"
)

// ChangePassword changes a user's password.
//
// - Use auth.ForgotPassword if the user doesn't know their password.
//
// - Use auth.SetPassword if the user has a requirement for their password to be reset.
func ChangePassword(ctx context.Context, accessToken string, oldPassword string, newPassword string) error {
	if err := checkPackage(); err != nil {
		return xerror.Wrap("checkPackage call failed", err)
	}

	if _, err := CognitoClient.ChangePassword(ctx, &cognitoidentityprovider.ChangePasswordInput{
		AccessToken:      aws.String(accessToken),
		PreviousPassword: aws.String(oldPassword),
		ProposedPassword: aws.String(newPassword),
	}); err != nil {
		return xerror.Wrap("failed to change password", err)
	}

	return nil
}
