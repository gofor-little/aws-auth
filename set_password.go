package auth

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/gofor-little/xerror"
)

// SetPassword sets a password for a user that has a requirement for their password to be changed. The session parameter
// can be obtained from the output.Session return value of auth.SignIn.
//
// - Use auth.ForgotPassword if the user doesn't know their password.
//
// - Use auth.ChangePassword and auth.ChangePasswordConfirm to update a user's password that doesn't require resetting.
func SetPassword(ctx context.Context, session string, emailAddress string, password string) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error) {
	output, err := CognitoClient.RespondToAuthChallenge(ctx, &cognitoidentityprovider.RespondToAuthChallengeInput{
		ChallengeName: types.ChallengeNameTypeNewPasswordRequired,
		ClientId:      aws.String(CognitoClientID),
		ChallengeResponses: map[string]string{
			"NEW_PASSWORD": password,
			"USERNAME":     emailAddress,
		},
		Session: aws.String(session),
	})
	if err != nil {
		return nil, xerror.Wrap("failed to set password", err)
	}

	return output, nil
}
