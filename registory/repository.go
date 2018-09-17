package registory

import (
	"github.com/vsanna/go_web/config"
	"github.com/vsanna/go_web/domain/repository"
	"github.com/vsanna/go_web/infra/dao"
)

/*
# registory

## about
- 実実装をinterface値に変換する関数群。
- 実装だけ見ると、daoのDB操作をrepositoryのinterfaceに隠す場所？
	- repository.User.Findでdao.User.Findを隠す


*/

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

/*
# ex. もしmailerがあれば(かつsendgrid)
// registory/mailer.go
package registory
type mail struct {
	Cnf *config.Config
}

type Mailer interface {
	NewSendGrid() mailer.Sendgrid // interface値を返す
}

func NewMailer(cnf *config.Config) Mailer {
	return &mail{
		Cnf: cnf,
	}
}

func (m *mailer) NewUserMailer() mailer.UserMailer {
	return &mailerservice.UserMailer {
		Cnf: m.Cnf
	}
}

// domain/mailer/user_mailer.go
package mailer
interface UserMailer interface {
	Send() error
}

// infra/mailerservice/user_mailer.go
package mailerservice
type UserMailer {
	Cnf: *config.Config
}

func (m *UserMailer) Send() error {
	// 個別外部サービスのAPIをcallしてメールを送る
}
*/
