package user

import (
	"database/sql"
	"fmt"

	"github.com/gaba-bouliva/buyit/types"
)



type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User , error) {
	rows, err := s.db.Query("SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	usr := new(types.User)
	for rows.Next() {
		usr, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if usr.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return usr, nil
}

func (s *Store)GetUserByID(id int) (*types.User, error) {

	return nil, nil
}

func (s *Store)CreateUser(types.User) error {

	return nil
}


func scanRowIntoUser(rows *sql.Rows) (*types.User, error){
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}