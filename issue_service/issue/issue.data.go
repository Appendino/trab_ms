package issue

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/appendino/trab_ms/issue_service/database"
)

func getIssue(issueID int) (*Issue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `select issueid,
	username,
	description
	FROM issues
	WHERE issueid = ?`, issueID)
	issue := &Issue{}
	err := row.Scan(&issue.IssueID, &issue.Username, &issue.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return issue, nil
}

func insertIssue(issue Issue) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO issues
	(username,
		description) VALUES(?,?)`,
		issue.Username,
		issue.Description)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}
