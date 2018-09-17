package registory

import (
	"../config"
	"../domain/repository"
	"../infra/dao"
)

/*
# registory

## about
- [?] これはなに？
- 実装だけ見ると、daoのDB操作をrepositoryのinterfaceに隠す場所？
	- repository.User.Findでdao.User.Findを隠す


repo = registory.NewRepository(cnf)
// repoは
userRepo = repo.NewUserRepo()


## note
- Repository, repoはここにあるべき?

userrepo := registry.NewUserRepo()
userrepo.Create()

などのほうがわかりやすくないか？
repository.NewUserRepo() だともっと良さそう


*/

// このinterfaceいる？
type Repository interface {
	NewUserRepo() repository.User
}

type repo struct {
	Cnf *config.Config
}

func NewRepository(cnf *config.Config) Repository {
	return &repo{
		Cnf: cnf,
	}
}

//
func (r *repo) NewUserRepo() repository.User {
	return &dao.User{
		Cnf: r.Cnf,
	}
}
