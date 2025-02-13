package main

import (
	"fmt"
)

type Doc struct {
	title    string
	subtitle string
	content  string
}

// Update title
func updateTitle(ch chan string, title string) {
	ch <- title
}

// Update subtitle
func updateSubtitle(ch chan string, subtitle string) {
	ch <- subtitle
}

// Update content
func updateContent(ch chan string, content string) {
	ch <- content
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

		go updateSubtitle(ch, "xlorem-ipsum-dolor-sit-amet")
		doc1.subtitle = <-ch

		go updateContent(ch, "xStudiis-malum-res-cyrenaicos-hortatore-sedentis-vituperatoribus-irridente.-Timorem-inhumanus-iucundum-apeirian-initia-nihil-improbis-ferentur-numquid.")
		doc1.content = <-ch
	}

	close(ch)

	fmt.Printf("\ntitle: %s\n", doc1.title)
	fmt.Printf("\nsubtitle: %s\n", doc1.subtitle)
	fmt.Printf("\ncontent: %s\n", doc1.content)
}
