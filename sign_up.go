package auth

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gofor-little/xerror"
)

// SignUp signs a new user up.
func SignUp(ctx context.Context, emailAddress string, password string, passwordConfirmation string) (string, error) {
	if password != passwordConfirmation {
		return "", xerror.New("password and password confirmation missmatch")
	}

	output, err := CognitoClient.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(CognitoUserPoolClientID),
		Password: aws.String(password),
		Username: aws.String(emailAddress),
	})
	if err != nil {
		return "", xerror.Wrap("failed to sign up user", err)
	}

	return *output.UserSub, nil
}

// SignUpConfirm confirms a newly signed up user with the confirmation code they received.
func SignUpConfirm(ctx context.Context, emailAddress string, confirmationCode string) error {
	_, err := CognitoClient.ConfirmSignUp(ctx, &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(CognitoUserPoolClientID),
		ConfirmationCode: aws.String(confirmationCode),
		Username:         aws.String(emailAddress),
	})
	if err != nil {
		return xerror.Wrap("failed to confirm sign up", err)
	}

	return nil
}
