package repository

import (
	"context"

	"../model"
)

/*
# repository
## about
- 一般的には倉庫。データを貯めておく場所
- SoftWare的にはデータをためておくもの(データプール)全般を指す
	- ex. RDB(MySQL), in-memoli kvs(Redis), ローカルのデータなど

Go的には実際のデータプールの実装を知らなくともApplicationからデータをCRUDできるように、
インターフェースだけを公開する。そのインターフェース郡をrepositoryと呼ぶ。

実際の実装はinfrastructure層に書く？

## note
- contextを第一引数で渡す
	- なんで？

[?]Repositoryとは？
[?] why context
*/
type User interface {
	All(ctx context.Context) ([]*model.User, error)
	FindById(ctx context.Context, ID int) (*model.User, error)
	FindByToken(ctx context.Context, sid string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	// Where
	// Create
	// Update
	// Destroy
}
