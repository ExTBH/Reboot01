# Structs

## UserRequest
The request for signing up / loging in with a user.
For signup all fields are required except `image`.
For Log in only `username` and `password`.
```go
type UserRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Image     []byte `json:"image"`
}
```