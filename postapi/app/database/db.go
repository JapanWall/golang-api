package database

import (
	"fmt"
	"log"
	"postapi/app/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostDB interface {
	Open() error
	Close() error
	CreatePost(p *models.Post) error
	DeletePost(p *models.Post) error
	UpdatePut(p *models.Post) error
	GetPosts() ([]*models.Post, error)
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	fmt.Println("running")
	pg, err := sqlx.Connect("postgres", pgConnStr)
	if err != nil {
		return err
	}

	log.Println("connected to DB!")
	pg.MustExec(createSchema)
	d.db = pg
	return nil
}

func (d *DB) Close() error {
	d.db.Close()
	return nil
}
