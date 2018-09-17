package model

import (
	"../../lib"
)

/*
## model
- いわゆるmodelのインスタンスメソッドに書くようなビジネスロジックを記述する
- classメソッドはclassメソッド週にあたる存在を書く or packageに切り出す or model.NewUserのようにmodel pkgに直接関数として記述する
*/

type User struct {
	ID                int
	Name              string
	Email             string
	EncryptedPassword string
	NormalizedEmail   string
	AccessToken       string
}

func NewUser(name, email, password string) (*User, error) {
	user := &User{
		Name:  name,
		Email: email,
	}
	user.setNormalizedEmail()
	user.SetAccessToken()
	user.SetEncryptedPassword(password)
	return user, nil
}

/* private */

func (u *User) setNormalizedEmail() error {
	u.NormalizedEmail = normalizeEmail(u.Email)
	// if err != nil {
	// 	return errors.Wrap(err, "cannot set normalized email")
	// }
	return nil
}

func normalizeEmail(email string) string {
	// TODO 正規化する
	return email
}

func (u *User) SetAccessToken() error {
	token, err := lib.SecureRandom()
	if err != nil {
		return err
	}
	// NOTE TODO tokenの重複チェックが必要
	// やっぱりmodelとDBは密になるしか無い? or この処理はdaoに移すべき?
	// でも分けると気持ち悪い...
	u.AccessToken = token
	return nil
}

func (u *User) SetEncryptedPassword(password string) error {
	u.EncryptedPassword = lib.Encrypt(password)
	return nil
}
