package database

import (
	"postapi/app/models"
)

func (d *DB) CreatePost(p *models.Post) error {
	res, err := d.db.Exec(insertPostSchema, p.Title, p.Content, p.Author, p.ID)

	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) DeletePost(p *models.Post) error {
	_, err := d.db.Exec(deletePostSchema, p.Title, p.Content, p.Author)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) UpdatePut(p *models.Post) error {
	_, err := d.db.Exec(updatePostSchema, p.Title, p.Content, p.Author)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetPosts() ([]*models.Post, error) {
	var posts []*models.Post
	err := d.db.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		return posts, err
	}

	return posts, nil
}
