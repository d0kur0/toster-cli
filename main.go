package main

import (
	"fmt"
	"log"
	"os"

	"github.com/d0kur0/toster-liker/httpClient"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli/v2"
)

var SID string
var target string

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "sid",
				Value:       "",
				Usage:       "Session identify for you account",
				Destination: &SID,
			},
			&cli.StringFlag{
				Name:        "target",
				Value:       "",
				Usage:       "User name for set likes",
				Destination: &target,
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "setLikes",
				Usage:  "Each all answers and set likes",
				Action: setLikes,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setLikes(c *cli.Context) error {
	var page = 1
	var counter = 0

	httpClient.SetSID(SID)

	for page > 0 {
		res, err := httpClient.GetRequest(fmt.Sprintf("https://qna.habr.com/user/%s/answers?page=%d", target, page))
		defer res.Body.Close()
		if err != nil {
			return err
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".btn_like:not(.btn_active)").Each(func(i int, s *goquery.Selection) {
			answerId, _ := s.Attr("data-answer-id")

			_, reqErr := httpClient.PostRequest("https://qna.habr.com/answer/like", fmt.Sprintf("answer_id=%s", answerId))

			if reqErr != nil {
				log.Println("Error then request set like", reqErr)
				return
			}

			fmt.Printf("Like for answer %s set\n", answerId)
			counter += 1
		})

		nextPage := doc.Find(".paginator__item.next > .paginator__item-link")
		if len(nextPage.Nodes) == 0 {
			log.Println("All answers set like, total:", counter)
			page = -1
		}

		page += 1
	}

	return nil
}
