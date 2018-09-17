package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/vsanna/go_web/config"
	"github.com/vsanna/go_web/domain/model"
	"github.com/vsanna/go_web/registory"
)

var repo registory.Repository
var JstTZ *time.Location = time.FixedZone("Asia/Tokyo", 9*60*60)

func init() {
	cnf := config.NewConfig()
	repo = registory.NewRepository(cnf)
}

/*=================================
* Logginng
=================================*/
type ResponseWriterWithStatus struct {
	http.ResponseWriter
	statusCode int
}

func (rww *ResponseWriterWithStatus) WriteHeader(code int) {
	rww.statusCode = code
	rww.ResponseWriter.WriteHeader(code)
}

func log(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		fmt.Printf("Started %s \"%s\" for %s at %s \n", r.Method, r.URL.Path, r.URL.Host, s.In(JstTZ).Format("2006-01-02 15:04:05 +09:00"))
		rww := &ResponseWriterWithStatus{ResponseWriter: w}
		h(rww, r)

		duration := time.Since(s)
		fmt.Printf("Completed %d %s in %dmsec \n", rww.statusCode, http.StatusText(rww.statusCode), duration.Nanoseconds())
	}
}

/*=================================
* Authentication
=================================*/
type ctxKey string

const currentUserKey ctxKey = "currentUser"

func authenticate(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := detectUserFromRequest(r)
		if err != nil {
			h(w, r)
		} else {
			ctx := r.Context()
			ctx = context.WithValue(ctx, currentUserKey, user)
			h(w, r.WithContext(ctx))
		}
	}
}

/*=================================
* Authorize
=================================*/
func authorize(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if UserSignedIn(r.Context()) {
			h(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "not authorized")
		}
	}
}

/*=================================
* Helpers
=================================*/
func UserSignedIn(ctx context.Context) bool {
	return CurrentUser(ctx) != nil
}

func CurrentUser(ctx context.Context) *model.User {
	val := ctx.Value(currentUserKey)
	if val != nil {
		return val.(*model.User)
	} else {
		return nil
	}
}

func detectUserFromRequest(r *http.Request) (*model.User, error) {
	sid, err := r.Cookie("_sid")
	if err != nil {
		return nil, errors.Wrap(err, "has no _sid")
	}

	urepo := repo.NewUserRepo()
	user, err := urepo.FindByToken(r.Context(), sid.Value)
	if err != nil {
		return nil, errors.Wrap(err, "has no user")
	}

	return user, nil
}

/*=================================
* Rendering
=================================*/
type TemplateOption struct {
	Layout    string
	HasHeader bool
	HasFooter bool
}

func NewTemplateOption() TemplateOption {
	return TemplateOption{
		Layout:    "layouts/application",
		HasHeader: true,
		HasFooter: false,
	}
}

// TODO 本当はrequestに依存したくないけど、しょうがない
func renderHTML(w http.ResponseWriter, r *http.Request, data interface{}, option TemplateOption, viewPaths ...string) {
	viewPaths = append(viewPaths, option.Layout)
	if option.HasHeader {
		viewPaths = append(viewPaths, "shared/header")
	}
	// NOTE こんな感じで拡張できる
	// if option.HasFooter {
	// 	viewPaths = append(viewPaths, "shared/footer")
	// }

	var paths []string
	for _, p := range viewPaths {
		paths = append(paths, fmt.Sprintf("views/%s.html", p))
	}
	t := template.Must(template.ParseFiles(paths...))

	mergedData := genHTMLData(r, data)

	if err := t.ExecuteTemplate(w, "layout", mergedData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 全体で使いたい変数をここで渡す
// TODO これでええんか？
// TODO Funcもmergeするのか...
func genHTMLData(r *http.Request, data interface{}) interface{} {
	var uname string
	user := CurrentUser(r.Context())
	if user == nil {
		uname = ""
	} else {
		uname = user.Name
	}

	return struct {
		Cmn interface{}
		Add interface{}
	}{
		Cmn: struct {
			UserSignedIn    bool
			CurrentUserName string
		}{
			UserSignedIn:    UserSignedIn(r.Context()),
			CurrentUserName: uname,
		},
		Add: data,
	}
}

func renderJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
