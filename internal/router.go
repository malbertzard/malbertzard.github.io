package internal

import (
	"log"
	handlers "malbertzard_homepage/internal/handler"
	"net/http"
)

var Routes = []struct {
	Path    string
	Handler http.Handler
}{
	// Style elements
	{"/", http.FileServer(http.Dir("../static/"))},

	{"/section/styles/", http.HandlerFunc(handlers.StylesHandler)},

	{"/section/header", http.HandlerFunc(handlers.HeaderHandler)},
	{"/section/main", http.HandlerFunc(handlers.MainHandler)},
	{"/section/footer", http.HandlerFunc(handlers.FooterHandler)},


	{"/section/home", http.HandlerFunc(handlers.HomeHandler)},
	{"/section/blog", http.HandlerFunc(handlers.BlogHandler)},
	{"/section/game", http.HandlerFunc(handlers.GameHandler)},
	{"/section/portfolio", http.HandlerFunc(handlers.PortfolioHandler)},

	{"/page/contact", http.HandlerFunc(handlers.PageContactHandler)},
}

func Route() {
	for _, route := range Routes {
		http.Handle(route.Path, route.Handler)
	}

	log.Println("Listening on :3001...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
