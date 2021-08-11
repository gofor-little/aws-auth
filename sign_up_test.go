package auth_test

import (
	"context"
	"fmt"
	"testing"

	auth "github.com/gofor-little/aws-auth"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	setup(t)
	defer teardown(t)

	testCases := []struct {
		emailAddress         string
		password             string
		passwordConfirmation string
	}{
		{"john@example.com", "test-Password1234!!", "test-Password1234!!"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestSignUp_%d", i), func(t *testing.T) {
			_, err := auth.SignUp(context.Background(), tc.emailAddress, tc.password, tc.passwordConfirmation)
			require.NoError(t, err)
		})
	}
}
