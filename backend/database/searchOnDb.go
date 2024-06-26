package database

import (
	"database/sql"
	"real-time-backend/backend/modals"
)

func GetUserTable() ([]modals.UserRegistration, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var user modals.UserRegistration
		err := rows.Scan(&user.Id, &user.Nickname, &user.Age, &user.Gender, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Date, &user.Img)
		return user, err
	}
	results, err := FetchDb("SELECT * FROM Users", executor)
	if err != nil {
		return nil, err
	}
	return ConvertResults[modals.UserRegistration](results)
}

func GetCategoriesTable() ([]modals.Categories, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var categorie modals.Categories
		err := rows.Scan(&categorie.Id, &categorie.Name, &categorie.Description)
		return categorie, err
	}
	results, err := FetchDb("SELECT * FROM Categories", executor)
	if err != nil {
		return nil, err
	}
	return ConvertResults[modals.Categories](results)
}

// Home JS
func GetPostTable() ([]modals.Post, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var post modals.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Username, &post.Creation, &post.Title, &post.Description)
		return post, err
	}
	results, err := FetchDb("SELECT * FROM Posts ORDER BY created_at DESC", executor)
	if err != nil {
		return nil, err
	}
	return ConvertResults[modals.Post](results)
}

func GetCommentsTable() ([]modals.Comments, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var comments modals.Comments
		err := rows.Scan(&comments.Id, &comments.PostId, &comments.Userid, &comments.Date, &comments.Content)
		return comments, err
	}
	results, err := FetchDb("SELECT * FROM Comments", executor)
	if err != nil {
		return nil, err
	}
	return ConvertResults[modals.Comments](results)
}

func GetLikesCommentsTable() ([]modals.LikesComments, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var likesComments modals.LikesComments
		err := rows.Scan(&likesComments.Id, &likesComments.Userid, &likesComments.CommentsId, &likesComments.Date, &likesComments.Sentiments)
		return likesComments, err
	}
	results, err := FetchDb("SELECT * FROM LikesComments", executor)
	if err != nil {
		return nil, err
	}
	return ConvertResults[modals.LikesComments](results)
}

func GetPostCategoriesTable() ([]modals.Postscategories, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var PostCategories modals.Postscategories
		err := rows.Scan(&PostCategories.Postid, &PostCategories.Categoryid)
		return PostCategories, err
	}
	results, err := FetchDb("SELECT * FROM PostCategories", executor)
	if err != nil {
		return nil, err
	}
	return ConvertResults[modals.Postscategories](results)
}

func GetPostLikesTable() ([]modals.PostsLikes, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var PostLikes modals.PostsLikes
		err := rows.Scan(&PostLikes.Id, &PostLikes.Userid, &PostLikes.Postid, &PostLikes.Date, &PostLikes.Sentiment)
		return PostLikes, err
	}
	results, err := FetchDb("SELECT * FROM PostsLike", executor)
	if err != nil {
		return nil, err
	}
	return ConvertResults[modals.PostsLikes](results)
}
