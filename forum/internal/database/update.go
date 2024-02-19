package database

import (
	"fmt"
	"sandbox/internal/structs"
)

func UpdateRepor(reportID int, atype, value string) error {
	// Prepare the UPDATE statement
	stmt, err := db.Prepare(fmt.Sprintf("UPDATE Report SET %s = ? WHERE id = ?", atype))
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the UPDATE statement to mark the report as resolved
	_, err = stmt.Exec(value, reportID)
	if err != nil {
		return err
	}

	return nil
}

// mark the report as resolved -not pendding-
func UpdateReportStatus(reportID int) error {
	return UpdateRepor(reportID, "is_pending", "false")
}

// add response to the report
func UpdateReportResponse(reportID int, response string) error {
	return UpdateRepor(reportID, "report_response", response)
}

// updates the information of a user in the User table.
func UpdateUserInfo(newUserInfo *structs.User) error {
	// Prepare the SQL statement to update the user
	stmt, err := db.Prepare(`UPDATE User SET type_id = ?, username = ?, first_name = ?, last_name = ?, 
	date_of_birth = ?, email = ?, hashed_password = ?, image_id = ?, banned_until = ?, github_name = ?, 
	linkedin_name = ?, twitter_name = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement to update the user
	_, err = stmt.Exec(
		newUserInfo.Type,
		newUserInfo.Username,
		newUserInfo.FirstName,
		newUserInfo.LastName,
		newUserInfo.DateOfBirth,
		newUserInfo.Email,
		newUserInfo.HashedPassword,
		newUserInfo.ImageId,
		newUserInfo.BannedUntil,
		newUserInfo.GithubName,
		newUserInfo.LinkedinName,
		newUserInfo.TwitterName,
		newUserInfo.Id)
	if err != nil {
		return err
	}

	return nil
}

// updates the IsPending status of a request to false in the PromoteRequest table.
func UpdateRequestStatus(requestID int) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare("UPDATE PromoteRequest SET is_pending = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement to update the IsPending status
	_, err = stmt.Exec(false, requestID)
	if err != nil {
		return err
	}

	return nil
}

// updates the information of a post in the Post table.
func UpdatePostInfo(newpost *structs.Post) error {
	// Prepare the SQL statement to update the post
	stmt, err := db.Prepare(`UPDATE Post SET title = ?, message = ?, image_id = ? WHERE id = ?`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement to update the post
	_, err = stmt.Exec(
		newpost.Title,
		newpost.Message,
		newpost.ImageId,
		newpost.Id)
	if err != nil {
		return err
	}

	return nil
}
