package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	auth "github.com/gofor-little/aws-auth"
	"github.com/stretchr/testify/require"
)

func TestSetPassword(t *testing.T) {
	setup(t)
	defer teardown(t)

	testCases := []struct {
		emailAddress string
		password     string
	}{
		{"john@example.com", "test-Password1234!!"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestSignIn_%d", i), func(t *testing.T) {
			_, err := auth.CognitoClient.AdminCreateUser(context.Background(), &cognitoidentityprovider.AdminCreateUserInput{
				UserPoolId:        aws.String(auth.CognitoUserPoolID),
				Username:          aws.String(tc.emailAddress),
				TemporaryPassword: aws.String(tc.password),
			})
			require.NoError(t, err)

			// Sign in so we can get a session to set a new password.
			output, err := auth.SignIn(context.Background(), tc.emailAddress, tc.password)
			require.NoError(t, err)

			_, err = auth.SetPassword(context.Background(), *output.Session, tc.emailAddress, tc.password)
			require.NoError(t, err)
		})
	}
}
