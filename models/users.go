package httpauth

import (
	"errors"
	"fmt"
	"github.com/asdine/storm"
	"os"
)

type BoltdbAuthBackend struct {
	filePath string
	users    map[string]UserData
}

// NewBoltdbAuthBackend initializes a new backend by loading a map of users
// from a file.
// If the file doesn't exist, returns an error.
func NewBoltdbAuthBackend(filepath string) (b BoltdbAuthBackend, e error) {
	b.filePath = filepath
	if _, err := os.Stat(b.filePath); err == nil {
		db, _ := storm.Open(b.filePath)
		defer db.Close()
		db.Get("httpauth", "userdata", &b.users)
		if b.users == nil {
			b.users = make(map[string]UserData)
		}
	}
	return b, nil
}

// User returns the user with the given username. Error is set to
// ErrMissingUser if user is not found.
func (b BoltdbAuthBackend) User(username string) (user UserData, e error) {
	if user, ok := b.users[username]; ok {
		return user, nil
	}
	return user, ErrMissingUser
}

// Users returns a slice of all users.
func (b BoltdbAuthBackend) Users() (us []UserData, e error) {
	for _, user := range b.users {
		us = append(us, user)
	}
	return
}

// SaveUser adds a new user, replacing one with the same username, and flushes
// to the db.
func (b BoltdbAuthBackend) SaveUser(user UserData) error {
	b.users[user.Username] = user
	err := b.save()
	return err
}

func (b BoltdbAuthBackend) save() error {
	db, err := storm.Open(b.filePath)
	defer db.Close()
	if err != nil {
		return errors.New("boltdbauthbackend: failed to edit auth file")
	}
	db.Set("httpauth", "userdata", b.users)
	return nil
}

// DeleteUser removes a user, raising ErrDeleteNull if that user was missing.
func (b BoltdbAuthBackend) DeleteUser(username string) error {
	_, err := b.User(username)
	if err == ErrMissingUser {
		return ErrDeleteNull
	} else if err != nil {
		return fmt.Errorf("boltdbauthbackend: %v", err)
	}
	delete(b.users, username)
	return b.save()
}

// Close cleans up the backend. Currently a no-op for gobfiles.
func (b BoltdbAuthBackend) Close() {

}
