package output

/*
## output
- usecase層の返り値を担うstruct
- structにしておくことでテストのしやすさ、返り値の内容の明確さを担保する

*/
type User struct {
	ID   int
	Name string
}

type Users []User
