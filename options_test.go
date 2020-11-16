package deepl_test

import (
	"net/url"
	"testing"

	"github.com/bounoable/deepl"
	"github.com/stretchr/testify/assert"
)

func TestSourceLang(t *testing.T) {
	vals := make(url.Values)
	assert.Equal(t, "", vals.Get("source_lang"))
	deepl.SourceLang(deepl.German)(vals)
	assert.Equal(t, string(deepl.German), vals.Get("source_lang"))
}

func TestSplitSentences(t *testing.T) {
	splits := []deepl.SplitSentence{
		deepl.SplitNone,
		deepl.SplitDefault,
		deepl.SplitNoNewlines,
	}

	for _, split := range splits {
		t.Run(split.String(), func(t *testing.T) {
			vals := make(url.Values)
			deepl.SplitSentences(split)(vals)
			assert.Equal(t, vals.Get("split_sentences"), split.Value())
		})
	}
}

func TestPreserveFormatting(t *testing.T) {
	vals := make(url.Values)
	assert.Equal(t, "", vals.Get("preserve_formatting"))
	deepl.PreserveFormatting(true)(vals)
	assert.Equal(t, "1", vals.Get("preserve_formatting"))
	deepl.PreserveFormatting(false)(vals)
	assert.Equal(t, "0", vals.Get("preserve_formatting"))
	deepl.PreserveFormatting(true)(vals)
	assert.Equal(t, "1", vals.Get("preserve_formatting"))
}

func TestFormality(t *testing.T) {
	formalities := []deepl.Formal{
		deepl.DefaultFormal,
		deepl.LessFormal,
		deepl.MoreFormal,
	}

	for _, f := range formalities {
		t.Run(f.String(), func(t *testing.T) {
			vals := make(url.Values)
			deepl.Formality(f)(vals)
			assert.Equal(t, f.Value(), vals.Get("formality"))
		})
	}
}
