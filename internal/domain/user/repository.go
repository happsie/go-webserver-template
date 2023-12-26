package user

import (
	"errors"
	"github.com/google/uuid"
	"github.com/happsie/go-webserver-template/internal/architecture"
)

type Repository struct {
	Container architecture.Container
}

func (r Repository) Create(user User) error {
	res, err := r.Container.DB.NamedExec(`INSERT INTO users (id, email, display_name, created_at, updated_at, version) 
									VALUES (:user.id, :user.displayName, :user.createdAt, :user.updatedAt, 1)`, user)
	if err != nil {
		return nil
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if affected == 0 {
		return errors.New("minumum affected rows not reached")
	}
	return nil
}

func (r Repository) Read(ID uuid.UUID) (User, error) {
	user := User{}
	err := r.Container.DB.Get(&user, "SELECT * FROM users WHERE id = ?", ID)
	if err != nil {
		return User{}, nil
	}
	return user, nil
}

func (r Repository) Update(user User) error {
	res, err := r.Container.DB.NamedExec(`UPDATE users
										SET id = :user.id, display_name = :user.displayName, created_at = :user.createdAt, updated_at = :user.updatedAt, version = :version + 1
										WHERE id = :user.id AND version = user.version`, user)
	if err != nil {
		return nil
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if affected == 0 {
		return errors.New("minumum affected rows not reached")
	}
	return nil
}

func (r Repository) Delete(ID uuid.UUID) error {
	res, err := r.Container.DB.NamedExec(`DELETE * FROM users where id = :ID`, ID)
	if err != nil {
		return nil
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if affected == 0 {
		return errors.New("minumum affected rows not reached")
	}
	return nil
}
