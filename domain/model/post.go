package model

/*
## model
- いわゆるmodelのインスタンスメソッドに書くようなビジネスロジックを記述する
- classメソッドはclassメソッド週にあたる存在を書く or packageに切り出す or model.NewUserのようにmodel pkgに直接関数として記述する
*/

type Post struct {
	ID    int
	Title string
	Body  string
}

func NewPost(title, body string) *Post {
	return &Post{
		Title: title,
		Body:  body,
	}
}
