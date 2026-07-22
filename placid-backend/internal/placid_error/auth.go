package placiderror

import "errors"

var ErrInsecurePassword error = errors.New("Password cannot be less than 6 characters")
var ErrInvalidEmail error = errors.New("Please provide a valid email")
var ErrInvalidCredentials error = errors.New("Invalid credentials")
var ErrBadRequest error = errors.New("Bad request")
var ErrInternalServerError error = errors.New("Internal Server Error")

var ErrEmailExistsNewsletter error = errors.New("Email already exists in the newsletter")
