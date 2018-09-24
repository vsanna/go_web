package usecase

import (
	"context"

	"github.com/vsanna/go_web/domain/repository"
	"github.com/vsanna/go_web/usecase/input"
)

/*
## usecase
- handler(や、コマンドラインツールであればcommand)から呼び出される、実際の処理をまとめたもの
	- handler, command -> リクエスト内容の解釈やパラメーターのvalidation, レスポンスの成形が仕事
	- usecase -> リクエスト内容をもとに処理を行う箇所
- usecaseは名前空間がかぶりやすそう。なので、structのメソッド経由で実行するようにすると良さそう
	- Newでhandler, commandからは必要に応じてrepo, configなど使い回すものを受け取る
	- その他はメソッドの引数として受け取る ex. context
		- 複雑になればinput structにする
*/

type UpdateUserUsecase struct {
	repo repository.User
}

func NewUpdateUserUsecase(repo repository.User) *UpdateUserUsecase {
	return &UpdateUserUsecase{
		repo: repo,
	}
}

func (u *UpdateUserUsecase) Update(ctx context.Context, input *input.UpdateUser) error {
	user := input.User

	name := user.Name
	if n := input.Name; n != "" {
		name = n
	}

	email := user.Email
	if e := input.Email; e != "" {
		email = e
	}

	password := input.Password
	if password != "" {
		user.SetEncryptedPassword(password)
	}

	user.Name = name
	user.Email = email

	return u.repo.Update(ctx, user)
}
