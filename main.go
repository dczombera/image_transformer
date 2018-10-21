package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dczombera/image-transformer/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	ic := controllers.NewImageController()
	r.GET("/:articleId/:fragmentId/:modCount/:version", ic.GetOptimizedImage)
	log.Fatalln(http.ListenAndServe(os.Getenv("ADDR"), r))
}
