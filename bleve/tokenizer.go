package bleve

import (
	"errors"
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
	"github.com/solos/sego"
)

const (
	Name = "solos/sego"
)

type SegoTokenizer struct {
	handle *sego.Segmenter
}

func NewSegoTokenizer(dictpath, stop_words string) *SegoTokenizer {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary(dictpath)
	//x := gojieba.NewJieba(dictpath, hmmpath, userdictpath, idf, stop_words)
	return &SegoTokenizer{&segmenter}
}

func (x *SegoTokenizer) Free() {
	//x.handle.Free()
}

func (x *SegoTokenizer) Tokenize(sentence []byte) analysis.TokenStream {
	result := make(analysis.TokenStream, 0)
	pos := 1
	words := x.handle.Segment(sentence)
	for _, word := range words {
		token := analysis.Token{
			Term:     []byte(word.Token().Text()),
			Start:    word.Start(),
			End:      word.End(),
			Position: pos,
			Type:     analysis.Ideographic,
		}
		result = append(result, &token)
		pos++
	}
	return result
}

func tokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	dictpath, ok := config["dictpath"].(string)
	if !ok {
		return nil, errors.New("config dictpath not found")
	}

	stop_words, ok := config["stop_words"].(string)
	if !ok {
		return nil, errors.New("config stop_words not found")
	}
	return NewSegoTokenizer(dictpath, stop_words), nil
}

func init() {
	registry.RegisterTokenizer(Name, tokenizerConstructor)
}
