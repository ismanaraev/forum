package repository

import (
	"database/sql"
	"forumv2/internal/models"
)

type pictureStorage struct {
	db *sql.DB
}

func newPictureStorage(db *sql.DB) *pictureStorage {
	return &pictureStorage{db: db}
}

func (p *pictureStorage) AddPictureToPost(id models.PostID, pic models.Picture) error {
	stmt, err := p.db.Prepare(`INSERT INTO picture (postID, value, type, size) VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, string(pic.Value), pic.Type, pic.Size)
	if err != nil {
		return err
	}
	return nil
}

func (p *pictureStorage) GetPicturesByPostID(id models.PostID) ([]models.Picture, error) {
	stmt, err := p.db.Prepare(`SELECT value, type, size FROM picture WHERE postID = $1`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	var res []models.Picture
	for rows.Next() {
		var temp models.Picture
		var tmpStr string
		var imageType string
		var size int
		err = rows.Scan(&tmpStr, &imageType, &size)
		if err != nil {
			return nil, err
		}
		temp.Value = tmpStr
		temp.Type, err = models.StringToImageType(imageType)
		if err != nil {
			return nil, err
		}
		temp.Size = size
		res = append(res, temp)
	}
	return res, nil
}
