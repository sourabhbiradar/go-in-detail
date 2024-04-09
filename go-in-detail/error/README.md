# errors & error handling

error type is of interface type.

Ways to handle errors :
a) implementing error interface
b) using errors.New()
c) using fmt.Errorf()

# errors.Is() & errors.As()

errors.Is() used to check error is of particular type.

errors.As() checks if error matches non-nil pointer either to type that implements error or to any interface type.

They runs 'forloop' through error chain internally.

# wrapping & unwrapping errors

wrapping error : error wrapping another error & adding context to error message.

unwrapping error : seperating original error from wrapped error.


- main.go covers errors & error handling
- isasErr.go covers comparing errors.
- wrapErr.go covers wrapping & unwrapping errors

