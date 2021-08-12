# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v0.2.0 - 2021-08-12
### Added
* Added ```ChangePassword```, ```ForgotPassword``` and ```ForgotPasswordConfirmation``` functions.

### Changed
* **BREAKING**: Changed the return type of the ```SignIn``` function from ```*types.AuthenticationResultType``` to ```*cognitoidentityprovider.InitiateAuthOutput```.
* **BREAKING**: Renamed ```CognitoUserPoolClientID``` to ```UserPoolClient```.
* **BREAKING**: Removed ```passwordConfirmation``` parameter from ```SignUp```

## v0.1.0 - 2021-08-11
### Added
* Added a changelog.
* Added a code of conduct.
* Added a pull request template.
* Added a Semantic Commits configuration file.
* Added a Dependabot configuration file.
* Added a CI GitHub action.
* Added a Stale GitHub action.
* Added a LICENSE.
* Added ```SignUp```, ```SignUpConfirm``` and ```SignIn``` functions.
