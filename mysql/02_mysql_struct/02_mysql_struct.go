package main

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// User struct
type User struct {
	ID          int64       `json:"id" db:"user_id"`
	FirstName   string      `json:"firstName" db:"first_name"`
	LastName    string      `json:"lastName" db:"last_name"`
	BirthDate   string      `json:"birthDate" db:"birth_date"`
	AccountType AccountType `json:"accountType" db:"account_type"`
	CreatedAt   time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time   `json:"updatedAt" db:"updated_at"`
}

// AccountType represents the type of each account, which is int8 underneath.
type AccountType int8

// Account type const
const (
	Free         AccountType = 1
	Trial        AccountType = 2
	Subscribe    AccountType = 3
	Professional AccountType = 4
	Admin        AccountType = 20
)

// Scan implements driver.Scanner interface
// and turns a tinyint field from MySQL to the AccountType
func (r *AccountType) Scan(value interface{}) error {
	// If value is nil, false
	if value == nil {
		// Set the value of the pointer r to AccountType(0)
		*r = AccountType(2)
		return nil
	}

	if v, ok := value.(int64); ok {
		*r = AccountType(int8(v))
		return nil
	}

	return errors.New("Type assertion of AccountType failed")
}

// Value implements driver.Valuer interface
// and turns the AccountType into a tinyint field (TINYINT) for MySQL storage.
func (r AccountType) Value() (driver.Value, error) {
	return int64(r), nil
}

func main() {
	db, err := sql.Open("mysql", "root:ultraman7@/4xbuddy_db?parseTime=true&charset=utf8")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	query := `
		select
			user_id,
			birth_date,
			first_name,
			last_name,
			photo_url,
			avatar_url,
			member_id,
			account_type,
			preferred_lang,
			is_banned,
			is_subscribed,
			receive_all_updates,
			confirmed,
			deleted,
			created_at,
			updated_at
		from users
		where user_id = ?
	`
	u := &User{}
	var rows *sql.Rows
	rows, err = db.Query(query, 1)
	if err != nil && err == sql.ErrNoRows {
		fmt.Println(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.BirthDate,
			&u.AccountType,
			&u.CreatedAt,
			&u.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(u)
}
