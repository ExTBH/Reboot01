package database

// retrieves the count of posts
func GetPostsCount() (int, error) {
	// Prepare the SQL statement to get the count of posts
	stmt, err := db.Prepare(`SELECT COUNT(*) 
                             FROM Post
							 WHERE parent_id IS NULL`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow().Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// retrieves the count of User posts
func GetPostsCountByUser(userId string) (int, error) {
	// Prepare the SQL statement to get the count of posts
	stmt, err := db.Prepare(`SELECT COUNT(*) 
                             FROM Post
							 WHERE user_id = ?
							 AND parent_id IS NULL`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(userId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// retrieves the count of posts associated with a given category name by joining the post_category with the category tables, based on the provided category name.
func GetPostsCountByCategory(categoryName string) (int, error) {
	// Prepare the SQL statement to get the count of posts for the given category name
	stmt, err := db.Prepare(`SELECT COUNT(*) 
                             FROM PostCategory 
                             INNER JOIN Category ON PostCategory.category_id = Category.id
                             WHERE Category.name = ?`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(categoryName).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// retrieves the count of post comments
func GetCommentsCountForPost(parentID string) (int, error) {
	// Prepare the SQL statement to get the count of posts comments
	stmt, err := db.Prepare(`SELECT COUNT(*) 
                             FROM Post 
                             WHERE parent_id = ?`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(parentID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
