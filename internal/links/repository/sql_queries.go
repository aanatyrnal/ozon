package repository

const (
	createLink = `INSERT INTO links (link, short_link)
					VALUES ($1, $2)
					RETURNING short_link;`

	readLink = `select link
				from links
				WHERE short_link = $1`
)
