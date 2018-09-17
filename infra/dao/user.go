package dao

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"sync"

	"github.com/vsanna/go_web/config"
	"github.com/vsanna/go_web/domain/model"
)

/*
# infra/dao

## about
- 実際にDBからデータを取得し、modelのstructとして返す
- dao.Userはrepository.Userのinterfaceを実装する。
- applicationはrepository.Userの実装に依存する。



## note
- なんで？？

*/

type User struct {
	Cnf *config.Config
}

var mux sync.Mutex

func (user *User) FindById(ctx context.Context, id int) (*model.User, error) {
	for _, d := range user.Cnf.DB["users"] {
		u := d.(*model.User)
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("not found")
}

func (user *User) FindByToken(ctx context.Context, sid string) (*model.User, error) {
	for _, d := range user.Cnf.DB["users"] {
		u := d.(*model.User)
		if u.AccessToken == sid {
			return u, nil
		}
	}
	return nil, errors.New("not found")
}

func (user *User) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	for _, d := range user.Cnf.DB["users"] {
		u := d.(*model.User)
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("not found")
}

// TODO この型変換どうするべき？
func (user *User) All(ctx context.Context) ([]*model.User, error) {
	var res []*model.User
	data := user.Cnf.DB["users"]

	for _, u := range data {
		res = append(res, u.(*model.User))
	}

	return res, nil
}

// NOTE: ココらへんはgormに本来は任せる
func (klass *User) Create(ctx context.Context, record *model.User) error {
	mux.Lock()
	defer mux.Unlock()
	klassName := "users"
	var data []*model.User
	for _, d := range klass.Cnf.DB[klassName] {
		data = append(data, d.(*model.User))
	}

	// 1. idのセット
	sort.Slice(data, func(i, j int) bool { return data[i].ID > data[j].ID })
	fmt.Println("debug: ", data, data[0])
	record.ID = data[0].ID + 1

	// 2. insert
	klass.Cnf.DB[klassName][record.ID] = record

	return nil
}

func (klass *User) Update(ctx context.Context, record *model.User) error {
	mux.Lock()
	defer mux.Unlock()
	klassName := "users"
	klass.Cnf.DB[klassName][record.ID] = record
	return nil
}
