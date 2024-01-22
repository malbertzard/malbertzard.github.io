package handlers

import (
	"net/http"
)

func StylesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../public/styles.css")
}

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler := NewTemplateHandler("header.html")
	templateHandler.ServeHTTP(w, r)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler := NewTemplateHandler("main.html")
	templateHandler.ServeHTTP(w, r)
}

func FooterHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler := NewTemplateHandler("footer.html")
	templateHandler.ServeHTTP(w, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler := NewTemplateHandler("sections/home.html")
	templateHandler.ServeHTTP(w, r)
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler := NewTemplateHandler("sections/blog.html")
	templateHandler.ServeHTTP(w, r)
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler := NewTemplateHandler("sections/game.html")
	templateHandler.ServeHTTP(w, r)
}

func PortfolioHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler := NewTemplateHandler("sections/portfolio.html")
	templateHandler.ServeHTTP(w, r)
}

func PageContactHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler := NewTemplateHandler("page/contact.html")
	templateHandler.ServeHTTP(w, r)
}
