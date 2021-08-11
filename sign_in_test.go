package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/stretchr/testify/require"

	auth "github.com/gofor-little/aws-auth"
)

func TestSignIn(t *testing.T) {
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
			_, err := auth.SignUp(context.Background(), tc.emailAddress, tc.password, tc.password)
			require.NoError(t, err)

			// Confirm the user without the use of a confirmation code so we can test the sign in.
			_, err = auth.CognitoClient.AdminConfirmSignUp(context.Background(), &cognitoidentityprovider.AdminConfirmSignUpInput{
				UserPoolId: aws.String(auth.CognitoUserPoolID),
				Username:   aws.String(tc.emailAddress),
			})
			require.NoError(t, err)

			_, err = auth.SignIn(context.Background(), tc.emailAddress, tc.password)
			require.NoError(t, err)
		})
	}
}
