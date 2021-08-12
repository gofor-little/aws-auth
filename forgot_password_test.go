package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/stretchr/testify/require"

	auth "github.com/gofor-little/aws-auth"
)

func TestForgotPassword(t *testing.T) {
	setup(t)
	defer teardown(t)

	testCases := []struct {
		emailAddress string
		password     string
	}{
		{"john@example.com", "test-Password1234!!"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestForgotPassword%d", i), func(t *testing.T) {
			_, err := auth.SignUp(context.Background(), tc.emailAddress, tc.password)
			require.NoError(t, err)

			// Confirm the user without the use of a confirmation code so we can test the sign in.
			_, err = auth.CognitoClient.AdminConfirmSignUp(context.Background(), &cognitoidentityprovider.AdminConfirmSignUpInput{
				UserPoolId: aws.String(auth.CognitoUserPoolID),
				Username:   aws.String(tc.emailAddress),
			})
			require.NoError(t, err)

			// Set the email verified to true so we can use it in the forgot password request.
			_, err = auth.CognitoClient.AdminUpdateUserAttributes(context.Background(), &cognitoidentityprovider.AdminUpdateUserAttributesInput{
				UserAttributes: []types.AttributeType{
					{
						Name:  aws.String("email_verified"),
						Value: aws.String("true"),
					},
				},
				UserPoolId: aws.String(auth.CognitoUserPoolID),
				Username:   aws.String(tc.emailAddress),
			})
			require.NoError(t, err)

			_, err = auth.ForgotPassword(context.Background(), tc.emailAddress)
			require.NoError(t, err)
		})
	}
}
