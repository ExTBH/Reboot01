package database

import (
	"database/sql"
	"sandbox/internal/structs"
)

// Remove reaction from Post by taking ReactionPost ID
func RemoveReactionFromPost(ReactionPostID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Retrieve the reaction details before removing the reaction
	var postReaction structs.PostReaction

	err = db.QueryRow("SELECT post_id, reaction_id FROM PostReaction WHERE id = ?", ReactionPostID).Scan(
		&postReaction.PostId,
		&postReaction.ReactionId)

	if err == sql.ErrNoRows { // iam not sure if u want to return nil if no reaction is found to be deleted
		return nil
	}
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete the reaction from the PostReaction table
	_, err = tx.Exec("DELETE FROM PostReaction WHERE id = ?", ReactionPostID)
	if err != nil {
		tx.Rollback()
		return err
	}

	columnName, err := GetReactionType(postReaction.ReactionId)
	columnName += "_count"

	// Update the specified reaction count column on the Post table
	updateQuery := "UPDATE Post SET " + columnName + " = " + columnName + " - 1 WHERE id = ?"
	updateStmt, err := tx.Prepare(updateQuery)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer updateStmt.Close()

	_, err = updateStmt.Exec(postReaction.PostId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// removes a post, along with its associated reactions and categories, from the database.
func RemovePost(postID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Delete the post reactions associated with the post from the PostReaction table
	_, err = tx.Exec("DELETE FROM PostReaction WHERE post_id = ?", postID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete the post categories associated with the post from the PostCategory table
	_, err = tx.Exec("DELETE FROM PostCategory WHERE post_id = ?", postID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete post and the child posts of the post from the Post table
	_, err = tx.Exec("DELETE FROM Post WHERE id = ? OR parent_id = ?", postID, postID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// delete session from UserSession table, by using session token
func RemoveSession(sessionID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Prepare the SQL statement to delete a session
	stmt, err := tx.Prepare("DELETE FROM UserSession WHERE id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement to delete the session
	_, err = stmt.Exec(sessionID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// RemoveImage removes an image from the UploadedImage table by its ID.
func RemoveImage(imageID int) error {
	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Prepare the SQL statement within the transaction
	stmt, err := tx.Prepare("DELETE FROM UploadedImage WHERE id = ?")
	if err != nil {
		// Rollback the transaction in case of an error
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	// Execute the delete statement within the transaction
	_, err = stmt.Exec(imageID)
	if err != nil {
		// Rollback the transaction in case of an error
		tx.Rollback()
		return err
	}

	// Commit the transaction if everything is successful
	err = tx.Commit()
	if err != nil {
		// Rollback the transaction in case of an error
		tx.Rollback()
		return err
	}

	return nil
}
