package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

const (
	googleOauthClientID = "488806591869-skelvuocd7ffcehn04fopj352gaesp2h.apps.googleusercontent.com"
	githubOauthClientID = "a4a6a1cb8b8c42b404e6"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	if err := t.templ.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "",
		MaxAge: -1,
	})
	w.Header()["Location"] = []string{"/chat"}
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()

	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey("FzPf8jNqxnmscTMvNUB816BLqeJu8n")

	googleSecurityKey := os.Getenv("GOOGLE_OAUTH_SECRET_KEY")
	githubSecurityKey := os.Getenv("GITHUB_OAUTH_SECRET_KEY")
	gomniauth.WithProviders(
		google.New(googleOauthClientID, googleSecurityKey, "http://localhost:8080/auth/callback/google"),
		github.New(githubOauthClientID, githubSecurityKey, "http://localhost:8080/auth/callback/github"),
	)

	r := newRoom()
	// r.tracer = trace.New(os.Stdout)

	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/room", r)

	// チャットルーム開始
	go r.run()

	log.Println("Webサーバーを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}
