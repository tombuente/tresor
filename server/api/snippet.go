package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/tombuente/tresor/server/service/snippet"
	"github.com/tombuente/tresor/spec/snippetspec"
)

type snippetHandler struct {
	service snippet.Service
}

func snippetResToSpec(snippet snippet.Snippet) snippetspec.SnippetRes {
	return snippetspec.SnippetRes{
		Key:      snippet.Key,
		Content:  snippet.Content,
		Language: snippet.Language.Name,
	}
}

func NewSnippetRouter(service snippet.Service) chi.Router {
	handler := snippetHandler{
		service: service,
	}

	router := chi.NewRouter()
	router.Post("/", handler.postSnippet)
	router.Route("/{key}", func(router chi.Router) {
		router.Get("/", handler.getSnippet)
	})

	return router
}

func (handler snippetHandler) getSnippet(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	snippet, err := handler.service.GetSnippet(r.Context(), key)
	if err != nil {
		render.Render(w, r, NewDerivedErrorRenderer(err, http.StatusNotFound))
		return
	}

	render.JSON(w, r, snippetResToSpec(snippet))
}

func (handler snippetHandler) postSnippet(w http.ResponseWriter, r *http.Request) {
	var body snippetspec.SnippetRes
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		render.Render(w, r, NewTextErrorRenderer("cannot parse body", http.StatusBadRequest))
		return
	}

	language, err := handler.service.GetLanguageByName(r.Context(), body.Language)
	if err != nil {
		render.Render(w, r, NewDerivedErrorRenderer(err, http.StatusNotFound))
		return
	}

	snippet, err := handler.service.AddSnippet(r.Context(), body.Content, language)
	if err != nil {
		render.Render(w, r, NewDerivedErrorRenderer(err, http.StatusNotFound))
		return
	}

	render.JSON(w, r, snippetResToSpec(snippet))
}
