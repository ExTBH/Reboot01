package structs

import "time"

type UserTypeId int
type PostReactionType string

const (
	UserTypeIdGuest     UserTypeId = 0
	UserTypeIdUser      UserTypeId = 1
	UserTypeIdModerator UserTypeId = 2
	UserTypeIdAdmin     UserTypeId = 3
	// Post reactions types
	PostReactionTypeLike    PostReactionType = "like"
	PostReactionTypeDislike PostReactionType = "dislike"
	PostReactionTypeLove    PostReactionType = "love"
	PostReactionTypeHaha    PostReactionType = "haha"
	PostReactionTypeSkull   PostReactionType = "skull"
)

type UserType struct {
	Id          int
	Name        string
	Description string
	Perms       map[string]bool
}

type User struct {
	Id             int
	Type           UserTypeId
	Username       string
	Email          string
	FirstName      string
	LastName       string
	DateOfBirth    time.Time
	HashedPassword string
	ImageId        int
	BannedUntil    time.Time
	GithubName     string
	LinkedinName   string
	TwitterName    string
}

type Post struct {
	Id            int
	UserId        int
	ParentId      *int   //u can change the type to pointer or set the default parent ID as -1, choose what do u want
	Title         string // to do because it causes a problem retrieving the null value as an integer
	Message       string
	ImageId       int
	Time          time.Time
	LikeCount     int //set them as zero not null pls, when u add a post
	DislikeCount  int
	LoveCount     int
	HahaCount     int
	SkullCount    int
	CategoriesIDs []int
}

type Category struct {
	Id          int
	Name        string
	Description string
	Color       string
}

type PostCategory struct {
	Id         int
	PostId     int
	CategoryId int
}

type ReactionType struct {
	Id       int
	Reaction string
}

type PostReaction struct {
	Id         int
	PostId     int
	UserId     int
	ReactionId int
}

type Session struct {
	Id           int
	Token        string
	UserId       int
	CreationTime int64
}

type Image struct {
	Id   int
	Data []byte
}

type Report struct {
	Id             int
	ReporterId     int
	ReportedId     int
	Reason         string
	PostId         int // can be null
	Time           time.Time
	IsPostReport   bool
	IsPending      bool
	ReportResponse string
}

type PromoteRequest struct {
	Id        int
	UserId    int
	Reason    string
	Time      time.Time
	IsPending bool
}
