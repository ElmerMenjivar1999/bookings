package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w,r)
	})
}

//NoSurf adds CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler{
	csrfHandlers := nosurf.New(next)

	csrfHandlers.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandlers

}

//SessionLoad loads and saves hte session on every request
func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}