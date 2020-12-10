package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// 用户信息
type User struct {
	ID   int
	Name string
}

func NewUser() User {
	return User{}
}

var (
	ctx         context.Context
	db          *sql.DB
	UserNotFind = errors.New("dao: user not find")
)

func (u User) FindUserNameByID(id int) (string, error) {
	var name string
	query := fmt.Sprintf("SELECT name FROM user WHERE id=%d", id)
	err := db.QueryRowContext(ctx, query).Scan(&name)
	if err != nil {
		err = errors.Wrapf(UserNotFind, fmt.Sprintf("sql: %s error: %v", query, err))
	}
	return name, err
}
