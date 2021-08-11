## A package for authentication via AWS Cognito

![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gofor-little/aws-auth?include_prereleases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gofor-little/aws-auth)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/gofor-little/aws-auth/main/LICENSE)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/gofor-little/aws-auth/CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofor-little/aws-auth)](https://goreportcard.com/report/github.com/gofor-little/aws-auth)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gofor-little/aws-auth)](https://pkg.go.dev/github.com/gofor-little/aws-auth)

### Introduction
* Authenticate users via AWS Cognito

### Example
```go
package main

import (
	"context"

	auth "github.com/gofor-little/aws-auth"
)

func main() {
	// Initialize the auth package.
	if err := auth.Initialize(context.Background(), "AWS_PROFILE", "AWS_REGION", "USER_POOL_ID", "CLIENT_ID"); err != nil {
		panic(err)
	}

	// Sign in.
	if _, err := auth.SignIn(context.Background(), "john@example.com", "password1234"); err != nil {
		panic(err)
	}
}
```

### Testing
Ensure the following environment variables are set, usually with a .env file.
* ```AWS_PROFILE``` (an AWS CLI profile name)
* ```AWS_REGION``` (a valid AWS region)

Run ```go test ./...``` at the root of the project.