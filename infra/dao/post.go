package dao

import (
	"context"
	"sort"

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
- なぜ？ => daoの実装を気にせずに済ませたいから

*/

type Post struct {
	Cnf *config.Config
}

func (post *Post) Find(id int) (*model.Post, error) {
	return post.Cnf.DB["posts"][0].(*model.Post), nil
}

func (klass *Post) Create(ctx context.Context, record *model.Post) (*model.Post, error) {
	mux.Lock()
	defer mux.Unlock()

	klassName := "posts"

	var data []*model.Post
	for _, d := range klass.Cnf.DB[klassName] {
		data = append(data, d.(*model.Post))
	}

	// 1. idのセット
	sort.Slice(data, func(i, j int) bool { return data[i].ID < data[j].ID })
	record.ID = data[0].ID + 1

	// 2. insert
	klass.Cnf.DB[klassName][record.ID] = record

	return nil, nil
}
