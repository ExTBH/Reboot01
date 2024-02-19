package database

import (
	"database/sql"
	"fmt"
	"sandbox/internal/structs"
	"time"
)

// retrieves user information from the database based on the provided username and returns a structs.User and error
func GetUserByUsername(username string) (*structs.User, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id, type_id, username, first_name, last_name, 
		date_of_birth, email, hashed_password, image_id, banned_until,
		github_name, linkedin_name, twitter_name FROM User WHERE username = ?`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var bannedUntil sql.NullTime

	// Execute the SQL statement and retrieve the user information
	var u structs.User
	err = stmt.QueryRow(username).Scan(
		&u.Id,
		&u.Type,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.DateOfBirth,
		&u.Email,
		&u.HashedPassword,
		&u.ImageId,
		bannedUntil,
		&u.GithubName,
		&u.LinkedinName,
		&u.TwitterName)

	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist, return nil with no error
	}

	if err != nil && bannedUntil.Valid {
		return nil, err
	}

	// Assign the value from sql.NullTime to u.BannedUntil
	if bannedUntil.Valid {
		u.BannedUntil = bannedUntil.Time
	} else {
		u.BannedUntil = time.Time{} // Set a default value for u.BannedUntil (e.g., time.Time{})
	}

	return &u, nil
}

// retrieves all the values from the category table and returns them as a slice of Category structs
func GetCategories() ([]structs.Category, error) {
	// Prepare the SQL statement to select all values from the category table
	stmt, err := db.Prepare("SELECT id, name, description, color FROM category")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the category data
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []structs.Category
	for rows.Next() {
		var category structs.Category
		err := rows.Scan(&category.Id, &category.Name, &category.Description, &category.Color)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func getPassword(username, ptype string) (string, error) {
	stmt, err := db.Prepare(fmt.Sprintf("SELECT hashed_password FROM User WHERE %s = ?", ptype))
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var password string
	err = stmt.QueryRow(username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

// retrieves the password hash for a given username from the User table and returns it as a string
func GetPasswordHashForUserByUsername(username string) (string, error) {
	return getPassword(username, "username")
}

// retrieves the password hash for a given email from the User table and returns it as a string
func GetPasswordHashForUserByEmail(email string) (string, error) {
	return getPassword(email, "email")
}

// retrieves the usrId for a given username from the User table and returns it as a int
func GetUserIdByUsername(username string) (int, error) {
	// Prepare the SQL statement to select the password hash for the given username
	stmt, err := db.Prepare("SELECT id FROM User WHERE username = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var UserId int

	err = stmt.QueryRow(username).Scan(&UserId)
	if err != nil {
		return 0, err
	}

	return UserId, nil
}

// retrieves the username for a given userId from the User table and returns it as a string
func GetUsernameByUserId(userId int) (string, error) {
	var username string

	stmt, err := db.Prepare("SELECT username FROM User WHERE id = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	err = stmt.QueryRow(userId).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}

// retrieves a session from the UserSession table based on the provided session token and returns it as a Session struct
// before u use this function you should check the existence of the session first, it will create a problem if u use it before the validation
func GetSession(sessionToken string) (structs.Session, error) {
	// Prepare the SQL statement to select the session based on the session token
	stmt, err := db.Prepare("SELECT id, token, user_id, creation_time FROM UserSession WHERE token = ?")
	if err != nil {
		return structs.Session{}, err
	}
	defer stmt.Close()

	var session structs.Session
	err = stmt.QueryRow(sessionToken).Scan(&session.Id, &session.Token, &session.UserId, &session.CreationTime)

	if err != nil {
		return structs.Session{}, err
	}

	return session, nil
}

// retrieves the category IDs associated with a given post ID from the PostCategory table.
// It takes a post ID as a parameter and returns a slice of integers representing the category IDs.
func getPostCategories(postID int) ([]int, error) {
	// Create a slice to store the category IDs
	var categoryIDs []int

	// Execute the SQL query to retrieve the category IDs associated with the post ID
	rows, err := db.Query("SELECT category_id FROM PostCategory WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the retrieved rows and extract the category IDs
	for rows.Next() {
		var categoryID int
		err := rows.Scan(&categoryID)
		if err != nil {
			return nil, err
		}
		categoryIDs = append(categoryIDs, categoryID)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Return the retrieved category IDs
	return categoryIDs, nil
}

func getPostsHelper(rows *sql.Rows) ([]structs.Post, error) {
	var posts []structs.Post
	for rows.Next() {
		var post structs.Post
		err := rows.Scan(
			&post.Id,
			&post.UserId,
			&post.ParentId,
			&post.Title,
			&post.Message,
			&post.ImageId,
			&post.Time,
			&post.LikeCount,
			&post.DislikeCount,
			&post.LoveCount,
			&post.HahaCount,
			&post.SkullCount)
		if err != nil {
			return nil, err
		}

		post.CategoriesIDs, err = getPostCategories(post.Id)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

// retrieve the posts from the DB which can be sorted based on time(most liked, most disliked, oldest, latest) -- use 'DESC' for latest sorting
func GetPosts(count, offset int, sortType string) ([]structs.Post, error) {
	// Prepare the SQL statement to get posts with specified count, offset and sort type
	stmt, err := db.Prepare(`SELECT id, user_id, parent_id, title, message, image_id, time, like_count, 
							 dislike_count, love_count, haha_count, skull_count
							 FROM Post 
							 WHERE parent_id IS NULL
							 ORDER BY ?
							 LIMIT ? 
							 OFFSET ?`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sortType, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getPostsHelper(rows)
}

// retrieve the User posts from the DB, if you want to retrieve comments set the bool value to isComment = true
func GetPostsByUser(userId, count, offset int, isComment bool) ([]structs.Post, error) {
	comment := ""
	if isComment {
		comment = "NOT"
	}
	// Prepare the SQL statement to get User posts with specified userId, count and offset
	stmt, err := db.Prepare(fmt.Sprintf(`SELECT id, user_id, parent_id, title, message, image_id, time, like_count, 
							 dislike_count, love_count, haha_count, skull_count
							 FROM Post 
							 WHERE user_id = ? 
							 AND parent_id IS %s NULL
							 ORDER BY time DESC
							 LIMIT ? 
							 OFFSET ?`, comment))

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getPostsHelper(rows)
}

// retrieve the posts from the DB based on category, which can be sorted based on time(most liked, most disliked, oldest, latest) -- use 'DESC' for latest sorting
func GetPostsByCategory(categoryName string, count, offset int, sortType string) ([]structs.Post, error) {
	// Prepare the SQL statement to get posts by category with specified count, offset and sort type
	stmt, err := db.Prepare(`SELECT Post.id, Post.user_id, Post.parent_id, Post.title, Post.message, Post.image_id,
							 Post.time, Post.like_count, Post.dislike_count, Post.love_count, Post.haha_count, Post.skull_count
                             FROM Post 
                             INNER JOIN PostCategory ON Post.id 				 = PostCategory.post_id
                             INNER JOIN Category 	 ON PostCategory.category_id = Category.id
                             WHERE Category.name = ?
							 AND Post.parent_id IS NULL
							 ORDER BY ?
							 LIMIT ? 
							 OFFSET ?`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(categoryName, sortType, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getPostsHelper(rows)
}

// retrieves post comments by taken parent ID -post ID-
func GetCommentsForPost(parentID, count, offset int) ([]structs.Post, error) {
	// Prepare the SQL statement to get comments for post with specified count and offset
	stmt, err := db.Prepare(`SELECT id, user_id, parent_id, title, message, image_id, time, like_count, 
							 dislike_count, love_count, haha_count, skull_count
                             FROM Post 
							 WHERE parent_id = ?
							 LIMIT ? 
							 OFFSET ?`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(parentID, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getPostsHelper(rows)
}

// retrieve users ID for a rxn on post by taking the postId and rxnId
func GetReactionUsers(postId, reactionId int) ([]int, error) {
	var userIds []int

	stmt, err := db.Prepare("SELECT user_id FROM PostReaction WHERE post_id	= ? AND reaction_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(postId, reactionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userId int
		if err := rows.Scan(&userId); err != nil {
			return nil, err
		}
		userIds = append(userIds, userId)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userIds, nil
}

// retrieve posts that the user reacted to (liked, disliked) by taking the user ID
func GetUserReactions(userId int) ([]structs.Post, error) {
	stmt, err := db.Prepare(`
		SELECT Post.id, Post.user_id, Post.parent_id, Post.title, Post.message, Post.image_id,
		Post.time, Post.like_count, Post.dislike_count, Post.love_count, Post.haha_count, Post.skull_count
		FROM Post
		INNER JOIN PostReaction ON Post.id = PostReaction.post_id
		WHERE PostReaction.user_id = ?`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getPostsHelper(rows)
}

// Helper function to get the corresponding reaction name in the Post struct
func GetReactionType(reactionId int) (string, error) {

	stmt, err := db.Prepare("SELECT type FROM ReactionType WHERE id = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var reactionType string
	err = stmt.QueryRow(reactionId).Scan(&reactionType)
	if err != nil {
		return "", err
	}

	return reactionType, nil
}

func GetPost(postID int) (*structs.Post, error) {
	// Prepare the query statement
	stmt, err := db.Prepare(`SELECT id, user_id, parent_id, title, message, image_id, time, like_count, 
							dislike_count, love_count, haha_count, skull_count
							FROM Post 
							WHERE id = ?`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// Execute the query with the provided postID as a parameter
	row := stmt.QueryRow(postID)

	// Initialize a new Post structure
	post := &structs.Post{}

	// Scan the retrieved values into the Post structure
	err = row.Scan(
		&post.Id,
		&post.UserId,
		&post.ParentId,
		&post.Title,
		&post.Message,
		&post.ImageId,
		&post.Time,
		&post.LikeCount,
		&post.DislikeCount,
		&post.LoveCount,
		&post.HahaCount,
		&post.SkullCount)

	if err != nil {
		if err == sql.ErrNoRows {
			// Post with the given ID not found
			return nil, nil
		}
		return nil, err
	}

	post.CategoriesIDs, err = getPostCategories(post.Id)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func GetImage(imageID int) ([]byte, error) {
	// Prepare the query statement
	stmt, err := db.Prepare("SELECT data FROM UploadedImage WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the query with the provided imageID as a parameter
	row := stmt.QueryRow(imageID)

	// Initialize a byte slice to store the image data
	var imageData []byte

	// Scan the retrieved image data into the byte slice
	err = row.Scan(&imageData)
	if err != nil {
		if err == sql.ErrNoRows {
			// Image with the given ID not found
			return nil, nil
		}
		return nil, err
	}

	return imageData, nil
}

// retrieves a report from the database based on the provided report Id and returns the corresponding Report-struct
func GetReport(reportID int) (structs.Report, error) {
	// Prepare the SELECT statement
	stmt, err := db.Prepare(`SELECT id, reporter_user_id, reported_user_id, report_message, 
							reported_post_id, time, is_post_report, is_pending, report_response
							FROM Report 
							WHERE id = ?`)
	if err != nil {
		return structs.Report{}, err
	}
	defer stmt.Close()

	// Execute the SELECT statement with the report ID as a parameter
	row := stmt.QueryRow(reportID)

	// Initialize a Report struct to store the retrieved values
	var report structs.Report

	// Scan the retrieved values into the Report struct
	err = row.Scan(
		&report.Id,
		&report.ReporterId,
		&report.ReportedId,
		&report.Reason,
		&report.PostId,
		&report.Time,
		&report.IsPostReport,
		&report.IsPending,
		&report.ReportResponse)

	if err != nil {
		if err == sql.ErrNoRows {
			// Report with the given ID not found
			return structs.Report{}, nil
		}
		return structs.Report{}, err
	}

	return report, nil
}

// retrieves either the history of resolved reports (isPending set to false) or the active reports awaiting
// resolution (isPending set to true) from the database, based on the provided isPending parameter.
func GetReports(isPendding bool) ([]structs.Report, error) {
	// Prepare the SELECT statement
	stmt, err := db.Prepare(`SELECT id, reporter_user_id, reported_user_id, report_message, reported_post_id, 
							time, is_post_report, is_pending, report_response
							FROM Report 
							WHERE is_pending = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the SELECT statement with is_pending set to true
	rows, err := stmt.Query(isPendding)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the retrieved reports
	var reports []structs.Report

	// Iterate over the result set and scan the values into the Report structs
	for rows.Next() {
		var report structs.Report
		err = rows.Scan(
			&report.Id,
			&report.ReporterId,
			&report.ReportedId,
			&report.Reason,
			&report.PostId,
			&report.Time,
			&report.IsPostReport,
			&report.IsPending,
			&report.ReportResponse)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

// retrieves the permissions for all users from the UserRole table.
// It returns a slice of structs.UserType objects, each containing the user's permissions.
func GetUsersPermissions() ([]structs.UserType, error) {
	var userPerms []structs.UserType

	rows, err := db.Query(`SELECT id, role_name, description, can_post, can_react, can_manage_category, can_delete, can_report, can_promote FROM UserRole`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set and scan each row into a UserType object
	for rows.Next() {
		var userPerm structs.UserType
		var canPost, canReact, canManageCategory, canDelete, canReport, canPromote bool
		// Initialize the permissions map
		var permissions map[string]bool
		err = rows.Scan(
			&userPerm.Id,
			&userPerm.Name,
			&userPerm.Description,
			&canPost,
			&canReact,
			&canManageCategory,
			&canDelete,
			&canReport,
			&canPromote)

		if err != nil {
			return nil, err
		}

		// Assign the permissions to the map
		permissions = map[string]bool{
			"can_post":          canPost,
			"CanReact":          canReact,
			"CanManageCategory": canManageCategory,
			"CanDelete":         canDelete,
			"CanReport":         canReport,
			"CanPromote":        canPromote,
		}

		// Assign the permissions map to the UserType object
		userPerm.Perms = permissions

		userPerms = append(userPerms, userPerm)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userPerms, nil
}

// retrieves a Request from the PromoteRequest table based on the given requestId.
func GetRequest(requestID int) (structs.PromoteRequest, error) {
	var request structs.PromoteRequest

	// Prepare the SQL statement
	stmt, err := db.Prepare("SELECT id, user_id, description, time, is_pending FROM PromoteRequest WHERE id = ?")
	if err != nil {
		return request, err
	}
	defer stmt.Close()

	// Execute the SQL statement to retrieve the request
	err = stmt.QueryRow(requestID).Scan(
		&request.Id,
		&request.UserId,
		&request.Reason,
		&request.Time,
		&request.IsPending)

	if err != nil {
		return request, err
	}

	return request, nil
}

// retrieves all Requests from the PromoteRequest table based on the given isPending status.
func GetRequests(isPending bool) ([]structs.PromoteRequest, error) {
	var requests []structs.PromoteRequest

	// Prepare the SQL statement
	stmt, err := db.Prepare("SELECT id, user_id, description, time, is_pending FROM PromoteRequest WHERE is_pending = ?")
	if err != nil {
		return requests, err
	}
	defer stmt.Close()

	// Execute the SQL statement to retrieve the requests
	rows, err := stmt.Query(isPending)
	if err != nil {
		return requests, err
	}
	defer rows.Close()

	// Iterate over the result set and populate the requests slice
	for rows.Next() {
		var request structs.PromoteRequest
		err := rows.Scan(
			&request.Id, 
			&request.UserId, 
			&request.Reason, 
			&request.Time, 
			&request.IsPending)

		if err != nil {
			return requests, err
		}
		requests = append(requests, request)
	}

	return requests, nil
}
