## setup

### 1. if you have dep

```bash
$ dep ensure
```

### 2. else if you don't have dep

```bash
$ go get golang.org/x/crypto/bcrypt
$ go get github.com/rs/xid
$ go get github.com/pkg/errors
$ go get golang.org/x/net
```

## run
```bash
$ cd go_web
$ go run main.go
```

## deploy
NOTE: This is not nice way. 
```bash
# change file name of main_gea.go to main.go
$ mv main.go ../ && mv main_gae.go main.go
$ gcloud app deploy
$ mv main.go main_gae.go && mv ../main.go . 

# concat
$ mv main.go ../ && mv main_gae.go main.go & gcloud app deploy & mv main.go main_gae.go && mv ../main.go .
```