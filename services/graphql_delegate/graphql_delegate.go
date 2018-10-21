package graphql_delegate

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

type Response struct {
	data struct {
		article Article
	}
}

type Article struct {
	teasers []Teaser
}

type Teaser struct {
	image Image
}

type Image struct {
	id       string
	binaryID string `json:"binaryId`
}

var client *graphql.Client
var articleTeaserReq *graphql.Request

func init() {
	client = graphql.NewClient(os.Getenv("GRAPHQL_ENDPOINT"))
	articleTeaserReq = graphql.NewRequest(`
		query fetchArticleTeaser($input: ArticleInput) {
			article(input: $input) {
				teasers {
					id
					binaryId
				}
			}
		}
	`)
}

func FetchArticleTeaser(articleId, fragmentId string) (string, error) {
	articleTeaserReq.Var("input", articleId)
	ctx := context.Background()
	var resp Response

	if err := client.Run(ctx, articleTeaserReq, &resp); err != nil {
		log.Println(err)
		return "", err
	}

	binaryID, err := findBinaryId(fragmentId, resp.data.article.teasers)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return binaryID, err
}

func findBinaryId(fragmendID string, teasers []Teaser) (string, error) {
	for _, t := range teasers {
		if t.image.id == fragmendID {
			return t.image.binaryID, nil
		}
	}
	return "", errors.New("No matching binary id")
}
