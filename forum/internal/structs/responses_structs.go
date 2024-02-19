package structs

// MARK: Session
// Might not be needed
// type SessionResponse struct {
// 	Token  string
// 	Expiry int
// }

// MARK: Category
type CategoryResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	IconURL     string `json:"icon_url"`
}

// MARK: Reaction
type PostReactionResponse struct {
	Reaction   string   `json:"reaction"`
	Count      int      `json:"count"`
	DidReact   bool     `json:"did_react"`
	WhoReacted []string `json:"who_reacted"`
}

// MARK: Post
type PostResponse struct {
	Id         int                    `json:"id"`
	ParentId   int                    `json:"parent_id"`
	Title      string                 `json:"title"`
	Message    string                 `json:"message"`
	ImageURL   string                 `json:"image_url"`
	Categories []CategoryResponse     `json:"categories"`
	Reactions  []PostReactionResponse `json:"reactions"`
}

// MARK: User
type UserTypeResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type UserResponse struct {
	Username  string           `json:"username"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	ImageURL  string           `json:"image_url"`
	Type      UserTypeResponse `json:"type"`
}

// Badges: github, twitter ...
type UserBadgeResponse struct {
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
	Link    string `json:"link"`
}

type UserProfileResponse struct {
	User   UserResponse        `json:"user"`
	Badges []UserBadgeResponse `json:"badges"`
}
