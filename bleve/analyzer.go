package bleve

import (
	"errors"

	"github.com/bleve-search/bleve/analysis"
	"github.com/bleve-search/bleve/registry"
)

type JiebaAnalyzer struct {
}

func analyzerConstructor(config map[string]interface{}, cache *registry.Cache) (*analysis.Analyzer, error) {
	tokenizerName, ok := config["tokenizer"].(string)
	if !ok {
		return nil, errors.New("must specify tokenizer")
	}
	tokenizer, err := cache.TokenizerNamed(tokenizerName)
	if err != nil {
		return nil, err
	}
	alz := &analysis.Analyzer{
		Tokenizer: tokenizer,
	}
	return alz, nil
}

func init() {
	registry.RegisterAnalyzer("sego", analyzerConstructor)
}
