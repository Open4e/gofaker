package lorem

import (
	"math/rand"
	"strings"
	"time"
)

var words []string = []string{
	"alias", "consequatur", "aut", "perferendis", "sit", "voluptatem",
	"accusantium", "doloremque", "aperiam", "eaque", "ipsa", "quae", "ab",
	"illo", "inventore", "veritatis", "et", "quasi", "architecto",
	"beatae", "vitae", "dicta", "sunt", "explicabo", "aspernatur", "aut",
	"odit", "aut", "fugit", "sed", "quia", "consequuntur", "magni",
	"dolores", "eos", "qui", "ratione", "voluptatem", "sequi", "nesciunt",
	"neque", "dolorem", "ipsum", "quia", "dolor", "sit", "amet",
	"consectetur", "adipisci", "velit", "sed", "quia", "non", "numquam",
	"eius", "modi", "tempora", "incidunt", "ut", "labore", "et", "dolore",
	"magnam", "aliquam", "quaerat", "voluptatem", "ut", "enim", "ad",
	"minima", "veniam", "quis", "nostrum", "exercitationem", "ullam",
	"corporis", "nemo", "enim", "ipsam", "voluptatem", "quia", "voluptas",
	"sit", "suscipit", "laboriosam", "nisi", "ut", "aliquid", "ex", "ea",
	"commodi", "consequatur", "quis", "autem", "vel", "eum", "iure",
	"reprehenderit", "qui", "in", "ea", "voluptate", "velit", "esse",
	"quam", "nihil", "molestiae", "et", "iusto", "odio", "dignissimos",
	"ducimus", "qui", "blanditiis", "praesentium", "laudantium", "totam",
	"rem", "voluptatum", "deleniti", "atque", "corrupti", "quos",
	"dolores", "et", "quas", "molestias", "excepturi", "sint",
	"occaecati", "cupiditate", "non", "provident", "sed", "ut",
	"perspiciatis", "unde", "omnis", "iste", "natus", "error",
	"similique", "sunt", "in", "culpa", "qui", "officia", "deserunt",
	"mollitia", "animi", "id", "est", "laborum", "et", "dolorum", "fuga",
	"et", "harum", "quidem", "rerum", "facilis", "est", "et", "expedita",
	"distinctio", "nam", "libero", "tempore", "cum", "soluta", "nobis",
	"est", "eligendi", "optio", "cumque", "nihil", "impedit", "quo",
	"porro", "quisquam", "est", "qui", "minus", "id", "quod", "maxime",
	"placeat", "facere", "possimus", "omnis", "voluptas", "assumenda",
	"est", "omnis", "dolor", "repellendus", "temporibus", "autem",
	"quibusdam", "et", "aut", "consequatur", "vel", "illum", "qui",
	"dolorem", "eum", "fugiat", "quo", "voluptas", "nulla", "pariatur",
	"at", "vero", "eos", "et", "accusamus", "officiis", "debitis", "aut",
	"rerum", "necessitatibus", "saepe", "eveniet", "ut", "et",
	"voluptates", "repudiandae", "sint", "et", "molestiae", "non",
	"recusandae", "itaque", "earum", "rerum", "hic", "tenetur", "a",
	"sapiente", "delectus", "ut", "aut", "reiciendis", "voluptatibus",
	"maiores", "doloribus", "asperiores", "repellat"}

func Word() string {
	return words[rand.Intn(len(words))]
}

func Words(numWords int) (words string) {
	if numWords < 1 {
		return ""
	}

	words = Word()
	for i := 1; i < numWords; i++ {
		words = words + " " + Word()
	}
	return
}

func Sentence(numWords int) (sentence string) {
	if numWords < 1 {
		return ""
	}
	sentence = strings.Title(Word())
	for i := 1; i < numWords; i++ {
		sentence = sentence + " " + Word()
	}
	sentence = strings.TrimSpace(sentence) + "."
	return
}

func Sentences(numSentences int) (sentences string) {
	if numSentences < 1 {
		return ""
	}
	sentences = Sentence(6)
	for i := 1; i < numSentences; i++ {
		sentences = sentences + " " + Sentence(6)
	}
	return
}

func Paragraph(numSentences int) (paragraph string) {
	if numSentences < 1 {
		return ""
	}
	paragraph = Sentences(numSentences)
	return
}

func Paragraphs(numParagraphs int) (paragraphs string) {
	if numParagraphs < 1 {
		return ""
	}
	paragraphs = Paragraph(6)
	for i := 1; i < numParagraphs; i++ {
		paragraphs = paragraphs + "\n\n" + Paragraph(6)
	}
	return
}

func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func init() {
	Seed()
}
