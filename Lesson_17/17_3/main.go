package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	l := log.New(os.Stdout, "", 0)
	logHandler := logMiddleware(l)
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: authHandler(
			logHandler(mux),
		),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}
func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	mLog := fmt.Sprintf("resp:", msg)
	WriteLogFile(mLog)
	fmt.Fprint(res, msg)
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			http.Error(w, "пользователь не авторизован",
				http.StatusUnauthorized)
			r.Header.Add("is-auth", "false")
		} else {
			r.Header.Add("is-auth", "true")
		}
		h.ServeHTTP(w, r)
	})
}

func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mLog := fmt.Sprintf("url:", r.URL)
			WriteLogFile(mLog)
			auth := r.Header.Get("is-auth")
			if auth == "true" {
				h.ServeHTTP(w, r)
			}
		})
	}
}

func WriteLogFile(msg string) {
	go func() {
		file, err := os.OpenFile("./log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, fs.ModePerm)
		if err != nil {
			log.Fatalln("Не удалось открыть файл для логов:", err)
		}
		defer file.Close()

		log.SetOutput(file)
		log.Printf(msg)
	}()
}
