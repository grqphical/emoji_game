package routes

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grqphical/emoji-game/internal/api"
	"github.com/grqphical/emoji-game/templates"
)

func MakeRouter() http.Handler{
	r := chi.NewRouter()

	component := templates.Index("Emoji Game")

	r.Use(middleware.RequestID)
  	r.Use(middleware.RealIP)
  	r.Use(middleware.Logger)
  	r.Use(middleware.Recoverer)


	r.Get("/", templ.Handler(component).ServeHTTP)
	r.Get("/emoji", GetRandomEmojiHandler)
	r.Post("/guess", CheckGuessHandler)
	r.Get("/play-again", PlayAgainHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return r
}

func GetRandomEmojiHandler(w http.ResponseWriter, r *http.Request) {
	emoji, err := api.GetRandomEmoji()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	component := templates.EmojiTemplate(emoji)

	component.Render(r.Context(), w)
	w.Header().Add("Content-Type", "text/html")
}

func CheckGuessHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	w.Header().Add("Accept", "application/x-www-form-urlencoded")
	if err != nil {
		http.Error(w, "Invalid data recieved", http.StatusBadRequest)
		return
	}
	emojis := r.FormValue("emoji")
	guess := r.FormValue("guess")

	w.Header().Add("Content-Type", "text/html")

	validChoices := strings.Split(emojis, "≊")

	for _, choice := range validChoices {
		if strings.ToLower(guess) == strings.TrimSpace(choice) {
			component := templates.Correct()
			component.Render(r.Context(), w)
			return
		}
	}

	component := templates.Wrong(emojis)
	component.Render(r.Context(), w)
	
}

func PlayAgainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Hx-Trigger", "reload")
	w.Write([]byte(" "))
}