package structs

type UserRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Password     string `json:"password"`
	Image        []byte `json:"image"`
	DateOfBirth  int64  `json:"date_of_birth"`
	GithubName   string `json:"github_name"`
	LinkedinName string `json:"linkedin_name"`
	TwitterName  string `json:"twitter_name"`
}

type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"name"`
}

type PostReportRequest struct {
	PostID int    `json:"post_id"`
	Reason string `json:"reason"`
}

type UserReportRequest struct {
	Username int    `json:"username"`
	Reason   string `json:"reason"`
}

type PromoteUserRequest struct {
	Reason string `json:"reason"`
}
