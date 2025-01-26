package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

type Doc struct {
	Id       int32
	Title    string
	Subtitle string
	Content  string
}

// Update Title
func (d *Doc) updateTitle(newTitle string, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	d.Title = newTitle
	mu.Unlock()

	wg.Done()
}

// Update Subtitle
func (d *Doc) updateSubtitle(newSubtitle string, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	d.Subtitle = newSubtitle
	mu.Unlock()

	wg.Done()
}

// Update Content
func (d *Doc) updateContent(newContent string, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	d.Content = newContent
	mu.Unlock()

	wg.Done()
}

func main() {
	doc := Doc{
		Id:       1200,
		Title:    "Lorem ipsum dolor sit amet Title",
		Subtitle: "Lorem ipsum dolor sit amet Subtitle",
		Content:  "Lorem ipsum dolor sit amet Content",
	}

	// Number of concurrent updates
	numUpdates := 10

	wg.Add(3 * numUpdates)

	for i := 1; i <= numUpdates; i++ {
		go doc.updateTitle("Eo tot soliditatem bonis utens impediri conducunt possunt. Fruitur amicos locatus provocatus suis paranda praeterea quis metus.", &wg, &mu)

		go doc.updateSubtitle("Concederetur acri corrupti finiri facilius fastidium deseruisse recta. Perpetua verterunt simulent audita perferendis parentes audiebamus solido umquam. Contentiones accurate modus mortis verterunt arte dignissimos. Motus democritus te sequatur extremo proficiscuntur probatum.", &wg, &mu)

		go doc.updateContent("Infinito fuisse doloribus nominant historiae modo consuetudine. Difficilem terminari cetero intellegere verbum fortitudo muniti significet copulationesque. Nullas disputatione comprobavit dixi pariatur spernat amet. Duo evolutio poterit nostri fuisset geometriaque probatus ex turpe. Imitarentur pecunias turma errem diuturnitatem progrediens corrupte. Voluptas verbum aspernari magnopere coniuncta primisque verear faciunt earumque. Amicis ficta putas solis scilicet argumentandum iusteque via. Delectari copiosae monstruosi vituperatoribus pertinaces gratiam existimare. Consequentium leniter t integre assidua iudicium malis. Contineret inscientia vituperatum vacillare ego collegisti faciendum conspectum debeo. Levitatibus possit animus partes primis dixeris depravata.", &wg, &mu)
	}

	wg.Wait()

	fmt.Printf("Title: %s\n", doc.Title)
	fmt.Printf("\nSubtitle: %s\n", doc.Subtitle)
	fmt.Printf("\nContent: %s\n", doc.Content)
}
