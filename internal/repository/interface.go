package repository

import (
	"database/sql"
)

type repository struct {
	*userStorage
	*postStorage
	*sessionStorage
	*commentStorage
	*reactionsStorage
	*categoriesStorage
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		newUserSQLite(db),
		newPostSQLite(db),
		newSessionSQLite(db),
		newCommentsSQLite(db),
		newReactionsSQLite(db),
		newCategoriesStorage(db),
	}
}
