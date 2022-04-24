package main

import (
	"bytes"
	"html/template"
	"log"
	"net/url"
	"runtime"

	"github.com/zserge/lorca"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./frontend/html/*"))
}

func main() {

	exists := lorca.LocateChrome()
	if len(exists) == 0 {
		lorca.PromptDownload()
		return
	}

	var w bytes.Buffer

	tpl.ExecuteTemplate(&w, "main.gohtml", nil)

	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	// Create UI with data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(w.String()), "", 600, 200, args...)
	if err != nil {
		log.Fatal(err)
	}

	defer ui.Close()
	// Create a GoLang function callable from JS
	// ui.Bind("hello", func() string { return "World!" })
   
	// // Call above `hello` function then log to the JS console
	// ui.Eval("hello().then( (x) => { console.log(x) })")
   
	// Wait until UI window is closed
	<-ui.Done()
}