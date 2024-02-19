package database_test

import (
	"path"
	"sandbox/internal/database"
	"sandbox/internal/structs"
	"testing"
)

const PROJECT_PATH = "C:\\Users\\Ruqay"

var DB_PATH = path.Join(PROJECT_PATH, "db.sqlite")

func TestConnect(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCreateUser(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}
	u := structs.User{
		Type:           1,
		Username:       "john",
		Email:          "john@gmail.com",
		HashedPassword: "12345678912345",
	}
	err = database.CreateUser(u)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCheckExistance(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.CheckExistance("User", "username", "john")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetUserByUsername(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetUserByUsername("john")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetCategories(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetCategories()

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetPasswordHashForUserByEmail(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetPasswordHashForUserByEmail("john@gmail.com")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetPasswordHashForUserByUsername(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetPasswordHashForUserByUsername("john")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetUserIdByUsername(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetUserIdByUsername("john")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetUsernameByUserId(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetUsernameByUserId(1)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddSession(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.AddSession(structs.Session{1, "ghjk", 1, 44567})

	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddCategory(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.AddCategory(structs.Category{})

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetSession(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetSession("ghjk")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetPostsCountByCategory(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}
	_, err = database.GetPostsCountByCategory("Sport")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetPostsCountByUser(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}
	_, err = database.GetPostsCountByUser("1")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetPostsByUser(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}
	_, err = database.GetPostsByUser(1, 20, 30, true)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetPostsCount(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetPostsCount()

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetPostsByCategory(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetPostsByCategory("Sport", 10, 20, "time")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetPosts(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetPosts(10, 20, "time")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetCommentsCountForPost(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetCommentsCountForPost("1")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetCommentsForPost(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetCommentsForPost(1, 10, 20)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestUploadImage(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.UploadImage([]byte(""))

	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddPost(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.AddPost(structs.Post{})

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetReactionUsers(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetReactionUsers(1, 1)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetUserReactions(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetUserReactions(1)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddReactionToPost(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.AddReactionToPost(structs.PostReaction{1, 1, 1, 1})

	if err != nil {
		t.Error(err.Error())
	}
}

func TestRemoveReactionFromPost(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.RemoveReactionFromPost(2)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetPost(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetPost(1)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetImage(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetImage(1)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddReport(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.AddReport(structs.Report{})

	if err != nil {
		t.Error(err.Error())
	}
}

func TestRemovePost(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.RemovePost(1)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetReport(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetReport(1)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetReports(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetReports(true)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateReportStatus(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.UpdateReportStatus(1)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetUsersPermissions(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.GetUsersPermissions()

	if err != nil {
		t.Error(err.Error())
	}
}

func TestRemoveImage(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.RemoveImage(1)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddPromoteRequest(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.AddPromoteRequest(structs.PromoteRequest{})

	if err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateRequestStatus(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.UpdateRequestStatus(1)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestSearchContent(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.SearchContent("")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestUpdatePostInfo(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.UpdatePostInfo(&structs.Post{})

	if err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateUserInfo(t *testing.T) {
	err := database.Connect(DB_PATH)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = database.UpdateUserInfo(&structs.User{})

	if err != nil {
		t.Error(err.Error())
	}
}
