package twittergraphql

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)



var (
	UsernameMinLength = 2
	PasswordMinLength = 6
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")


type AuthService interface{
	Register(ctx context.Context, input RegisterInput)


}

type RegisterInput struct{
	Email string
	Username string
	Password string
	ConfirmPassword string
} 


// Sanitize cleans up the input fields for registration by trimming whitespace from the Email and Username,
// and converting the Email to lowercase for consistency.
func (in RegisterInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)
	in.Username = strings.TrimSpace(in.Username)
}


// Validate checks the validity of the RegisterInput fields.
func (in RegisterInput) Validate() error {
	// Check if the username is at least the minimum required length.
	if len(in.Username) < UsernameMinLength {
		return fmt.Errorf("%w: username not long enough, (%d) characters at least", ErrValidation, UsernameMinLength)
	}

	// Check if the email format is valid using a regular expression.
	if !emailRegexp.MatchString(in.Email) {
		return fmt.Errorf("%w: email not valid", ErrValidation)
	}

	// Check if the password is at least the minimum required length.
	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf("%w: password not long enough, (%d) characters at least", ErrValidation, PasswordMinLength)
	}

	// Ensure that the password and confirm password match.
	if in.Password != in.ConfirmPassword {
		return fmt.Errorf("%w: confirm password must match the password", ErrValidation)
	}

	// If all checks pass, return nil indicating no errors.
	return nil
}

