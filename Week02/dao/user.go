package dao

import (
	"context"
	"database/sql"
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
	err := db.QueryRowContext(ctx, "SELECT name FROM user WHERE id=?", id).Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 转换为自定义错误
			err = UserNotFind
		} else {
			err = errors.Wrap(err, "failed to find user")
		}
	}
	return name, err
}
