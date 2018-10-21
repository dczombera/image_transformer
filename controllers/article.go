package controllers

import (
	"log"
	"net/http"

	"github.com/dczombera/image-transformer/services/graphql_delegate"
	"github.com/julienschmidt/httprouter"
)

type ImageController struct{}

func NewImageController() *ImageController {
	return &ImageController{}
}

func (ic ImageController) GetOptimizedImage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	binaryId, err := graphql_delegate.FetchArticleTeaser(p.ByName("articleId"), p.ByName("fragmentId"))
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	qp := r.URL.Query()
}
