package repository

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"forumv2/internal/models"
)

type pictureStorage struct {
	db *sql.DB
}

func newPictureStorage(db *sql.DB) *pictureStorage {
	return &pictureStorage{db: db}
}

func (p *pictureStorage) AddPictureToPost(id models.PostID, pic models.Picture) error {
	stmt, err := p.db.Prepare(`INSERT INTO picture (postID, value) VALUES ($1,$2)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, string(pic.Value))
	if err != nil {
		return err
	}
	return nil
}

func (p *pictureStorage) GetPictureByPostID(id models.PostID) (models.Picture, error) {
	stmt, err := p.db.Prepare(`SELECT value FROM picture WHERE postID = $1`)
	if err != nil {
		return models.Picture{}, err
	}
	row := stmt.QueryRow(id)
	if err != nil {
		return models.Picture{}, err
	}
	var res models.Picture
	var tmpStr string
	err = row.Scan(&tmpStr)
	if err != nil {
		return models.Picture{}, err
	}
	decoder := base64.NewDecoder(base64.StdEncoding, bytes.NewReader([]byte(tmpStr)))
	_, err = decoder.Read(res.Value)
	if err != nil {
		return models.Picture{}, err
	}
	return res, nil
}
