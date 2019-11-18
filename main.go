package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/betsegawlemma/webprogmem/entity"
	"github.com/betsegawlemma/webprogmem/menu/service"
)

var tmpl = template.Must(template.ParseGlob("delivery/web/templates/*"))
var categoryCache service.CategoryCache

func index(w http.ResponseWriter, r *http.Request) {

	idRaw := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		id = 1
	}
	cat, err := categoryCache.Category(id)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "index.layout", cat)
}

func init() {
	categoryCache = service.NewCategoryCache()
	breakfast := entity.Category{ID: 1, Name: "Breakfast", Description: "Lorem ipsum", Image: "bkt.png"}
	lunch := entity.Category{ID: 2, Name: "Lunch", Description: "Lorem ipsum", Image: "lnc.png"}
	dinner := entity.Category{ID: 3, Name: "Dinner", Description: "Dinner Cateogry", Image: "dnr.png"}
	snack := entity.Category{ID: 4, Name: "Snack", Description: "Snack Cateogry", Image: "snk.png"}
	categoryCache.StoreCategory(&breakfast)
	categoryCache.StoreCategory(&lunch)
	categoryCache.StoreCategory(&dinner)
	categoryCache.StoreCategory(&snack)
}

func main() {
	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8181", nil)
}
