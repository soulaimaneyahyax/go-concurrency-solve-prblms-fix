package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

type Doc struct {
	Id       int32
	title    string
	subtitle string
	content  string
}

// Update title
func (d *Doc) updateTitle(title string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	d.title = title
	mu.Unlock()
}

// Update subtitle
func (d *Doc) updateSubtitle(subtitle string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	d.subtitle = subtitle
	mu.Unlock()
}

// Update content
func (d *Doc) updateContent(content string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	d.content = content
	mu.Unlock()
}

func main() {
	doc := Doc{
		Id:       1200,
		title:    "Lorem ipsum dolor sit amet title",
		subtitle: "Lorem ipsum dolor sit amet subtitle",
		content:  "Lorem ipsum dolor sit amet content",
	}

	// Number of concurrent updates
	numUpdates := 10

	wg.Add(3 * numUpdates)

	for i := 1; i <= numUpdates; i++ {
		go doc.updateTitle("Eo tot soliditatem bonis utens impediri conducunt possunt. Fruitur amicos locatus provocatus suis paranda praeterea quis metus.", &wg, &mu)

		go doc.updateSubtitle("Concederetur acri corrupti finiri facilius fastidium deseruisse recta. Perpetua verterunt simulent audita perferendis parentes audiebamus solido umquam. contentiones accurate modus mortis verterunt arte dignissimos. Motus democritus te sequatur extremo proficiscuntur probatum.", &wg, &mu)

		go doc.updateContent("Infinito fuisse doloribus nominant historiae modo consuetudine. Difficilem terminari cetero intellegere verbum fortitudo muniti significet copulationesque. Nullas disputatione comprobavit dixi pariatur spernat amet. Duo evolutio poterit nostri fuisset geometriaque probatus ex turpe. Imitarentur pecunias turma errem diuturnitatem progrediens corrupte. Voluptas verbum aspernari magnopere coniuncta primisque verear faciunt earumque. Amicis ficta putas solis scilicet argumentandum iusteque via. Delectari copiosae monstruosi vituperatoribus pertinaces gratiam existimare. Consequentium leniter t integre assidua iudicium malis. Contineret inscientia vituperatum vacillare ego collegisti faciendum conspectum debeo. Levitatibus possit animus partes primis dixeris depravata.", &wg, &mu)
	}

	wg.Wait()

	fmt.Printf("title: %s\n", doc.title)
	fmt.Printf("\nsubtitle: %s\n", doc.subtitle)
	fmt.Printf("\ncontent: %s\n", doc.content)
}
