package database

const createSchema = `
CREATE TABLE IF NOT EXISTS posts
(
	id SERIAL PRIMARY KEY,
	title TEXT,
	content TEXT,
	author TEXT
)
`

var insertPostSchema = `
INSERT INTO posts(title, content, author, id) VALUES($1,$2,$3, $4) RETURNING id
`
var deletePostSchema = `
DELETE FROM posts WHERE title = $1 and content = $2 and author = $3
`

var updatePostSchema = `
UPDATE posts SET TITLE = $1, content = $2 WHERE author = $3
`
