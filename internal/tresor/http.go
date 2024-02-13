package tresor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jinzhu/copier"
	"github.com/tombuente/tresor/rest"
)

const encodingBase = 36

type snippetHandler struct {
	service Service
}

func NewRouter(service Service) chi.Router {
	h := snippetHandler{
		service,
	}

	r := chi.NewRouter()
	r.Route("/snippets", func(r chi.Router) {
		r.Get("/{key}", h.getCodeSnippet)
		r.Post("/", h.postCodeSnippet)
	})

	return r
}

func (h snippetHandler) getCodeSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "key"), encodingBase, 64)
	if err != nil {
		fmt.Println("bad body")
		return
	}

	snippet, err := h.service.GetSnippet(r.Context(), id)
	if err != nil {
		fmt.Println("bad body")
		return
	}

	var snippetResponse rest.SnippetResponse
	copier.Copy(&snippetResponse, snippet)
	snippetResponse.Key = strconv.FormatInt(snippet.ID, 36)
	render.JSON(w, r, snippetResponse)
}

func (h snippetHandler) postCodeSnippet(w http.ResponseWriter, r *http.Request) {
	var body rest.SnippetRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("bad body")
		return
	}

	var snippet Snippet
	copier.Copy(&snippet, body)
	newSnippet, err := h.service.CreateSnippet(r.Context(), snippet)

	var snippetResponse rest.SnippetResponse
	copier.Copy(&snippetResponse, newSnippet)
	snippetResponse.Key = strconv.FormatInt(newSnippet.ID, 36)
	render.JSON(w, r, snippetResponse)
}
