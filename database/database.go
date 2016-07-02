package database

import (
	"github.com/asdine/storm"
)

var Storage DBStorage

type DBStorage struct {
	DB     *storm.DB
	Opened bool
}

func (s *DBStorage) Open(filePath string) error {
	var err error
	s.DB, err = storm.Open(filePath)
	if err == nil {
		s.Opened = true
		defer s.DB.Close()
	}
	return err
}

func (s *DBStorage) Close() error {
	s.Opened = false
	return s.DB.Close()
}
