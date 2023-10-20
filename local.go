package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

var code string
var USERNAME string

func open(url string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/sendUser", func(w http.ResponseWriter, r *http.Request) {
		USERNAME = r.FormValue("username")
		openbrowser(url)
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		code = r.FormValue("code")
		mov := getMovies(USERNAME)
		fmt.Fprintln(w, "Фільми на які ви можете піти:")
		for _, s := range mov {
			fmt.Fprintln(w, s)
		}
	})

	http.ListenAndServe(":8888", nil)
	openbrowser(url)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
