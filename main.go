package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Comment struct {
	ID            int        `json:"id"`
	Text          string     `json:"text"`
	ParentComment *Comment   `json:"-"`
	ChildComments []*Comment `json:"child_comments,omitempty"`
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/recursive_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
			return
		}

		// Retrieve the comments from the database
		rows, err := db.Query("SELECT comment_id, text, parent_comment_id FROM comments")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Map to store comment objects by ID
		commentMap := make(map[int]*Comment)

		// Iterate through the rows and create Comment objects
		for rows.Next() {
			var commentID, parentCommentID sql.NullInt64
			var text sql.NullString

			err = rows.Scan(&commentID, &text, &parentCommentID)
			if err != nil {
				log.Fatal(err)
			}

			// Create a new Comment object and add it to the map
			comment := &Comment{
				ID:   int(commentID.Int64),
				Text: text.String,
			}
			commentMap[comment.ID] = comment

			// Check if this comment has a parent comment
			if parentCommentID.Valid {
				parentComment, ok := commentMap[int(parentCommentID.Int64)]
				if !ok {
					// Parent comment hasn't been created yet, create a new one
					parentComment = &Comment{ID: int(parentCommentID.Int64), ChildComments: []*Comment{}}
					commentMap[parentComment.ID] = parentComment
				}

				// Add the comment as a child of its parent
				parentComment.ChildComments = append(parentComment.ChildComments, comment)
				comment.ParentComment = parentComment
			}
		}

		// Check for any errors during iteration
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		// Convert the map of comments to a slice
		comments := []*Comment{}
		for _, comment := range commentMap {
			// Only add top-level comments to the slice
			if comment.ParentComment == nil {
				comments = append(comments, comment)
			}
		}

		// Convert the comments slice to JSON
		jsonBytes, err := json.MarshalIndent(comments, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		// Set the content type and write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	})

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
