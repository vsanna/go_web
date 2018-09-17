package config

// NOTE 本当は依存しない
import (
	"fmt"

	"../domain/model"
)

func NewConfig() *Config {
	userTable := Table{}
	userNames := []string{"anna", "bob", "carl", "danny", "ethan", "Franky"}
	for i := 0; i < len(userNames); i++ {
		u, _ := model.NewUser(userNames[i], "user"+fmt.Sprint(i+1)+"@example.com", "pass"+fmt.Sprint(i+1))
		u.ID = i + 1
		userTable[i+1] = u
	}

	p1 := model.NewPost("title1", "body1")
	p1.ID = 1

	postTable := Table{
		1: p1,
	}

	return &Config{
		Host: "127.0.0.1",
		Port: "3567",
		User: "Admin",
		Pass: "none",
		DB: map[string]Table{
			"users": userTable,
			"posts": postTable,
		},
	}
}

type Config struct {
	Host string
	Port string
	User string
	Pass string
	DB   map[string]Table
}

type Table map[int]interface{}
