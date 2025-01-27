package main

import (
	"fmt"
)

type Doc struct {
	title    string
	subtitle string
	content  string
}

func updateTitle(ch chan string, title string) {
	ch <- title
}

func main() {
	doc1 := Doc{
		title:    "lorem ipsum",
		subtitle: "lorem ipsum dolor sit amet",
		content:  "Studiis malum res cyrenaicos hortatore sedentis vituperatoribus irridente. Timorem inhumanus iucundum apeirian initia nihil improbis ferentur numquid.",
	}

	nbrConcurrentUpdates := 10
	ch := make(chan string, nbrConcurrentUpdates)

	for i := 0; i < nbrConcurrentUpdates; i++ {
		go updateTitle(ch, "xlorem-ipsum")
		doc1.title = <-ch
	}

	close(ch)

	fmt.Printf("\ntitle: %s\n", doc1.title)
	fmt.Printf("\nsubtitle: %s\n", doc1.subtitle)
	fmt.Printf("\ncontent: %s\n", doc1.content)
}
