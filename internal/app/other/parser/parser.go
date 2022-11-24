package parser

import (
	"container/list"
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"factorio-calculator/internal/app/other/models"
	"factorio-calculator/internal/app/other/token"
)

var regexps []*regexp.Regexp
var seps []rune

func init() {
	seps = make([]rune, token.sepEnd)

	seps[models.CurvedO] = '{'
	seps[models.CurvedC] = '}'
	seps[models.RollO] = '('
	seps[models.RollC] = ')'
	seps[models.Assign] = '='

	regexps = make([]*regexp.Regexp, token.notSepEnd)

	regexps[models.CurvedO] = regexp.MustCompile("\\{")
	regexps[models.CurvedC] = regexp.MustCompile("}")
	regexps[models.RollO] = regexp.MustCompile("\\(")
	regexps[models.RollC] = regexp.MustCompile("\\)")
	regexps[models.Assign] = regexp.MustCompile("=")
	regexps[models.NumFloat] = regexp.MustCompile("\\d+\\.\\d+")
	regexps[models.NumInt] = regexp.MustCompile("\\d+")
	regexps[models.String] = regexp.MustCompile("\"[a-z\\d-]+\"")
	regexps[models.Field] = regexp.MustCompile("[a-z\\d-]+")
	regexps[models.BeginData] = regexp.MustCompile("\\w+:\\w+")
	regexps[models.Bool] = regexp.MustCompile("true|false")
}

type WordParser struct {
	IsDebug bool
}

func (p WordParser) Parse(str string) []models.Token {
	words := p.makeWords(str)
	tokens := make([]models.Token, 0, len(words))

	for _, word := range words {
		for tType, rgx := range regexps {
			if rgx.MatchString(word) {
				tokens = append(tokens, models.Token{
					Type:  models.Type(tType),
					value: word,
				})

				break
			} else if tType == len(regexps)-1 {
				tokens = append(tokens, models.Token{
					Type:  models.Unknown,
					value: word,
				})
			}
		}
	}

	return tokens
}

func (p WordParser) makeWords(str string) []string {
	fields := strings.FieldsFunc(str, func(r rune) bool {
		return unicode.IsSpace(r) || r == ','
	}) // remove formatting symbols

	words := make([]string, 0, int(token.sepEnd)*len(fields))

	for _, raw := range fields {
		subRes := list.New()
		subRes.PushBack(raw)

		for i := token.sepBeg; i < token.sepEnd; i++ {
			for it := subRes.Front(); it != nil; it = it.Next() {
				it = p.splitWord(it, subRes, seps[i])
			}
		}

		for it := subRes.Front(); it != nil; it = it.Next() {
			words = append(words, it.Value.(string))
		}
	}

	return words
}

func (p WordParser) splitWord(it *list.Element, l *list.List, sep rune) *list.Element {
	prevID := 0
	str := it.Value.(string)
	size := utf8.RuneCountInString(str)
	newIt := it

	search := func() int {
		if prevID == size {
			return -1
		}
		if tmp := strings.IndexRune(str[prevID:], sep); tmp >= 0 {
			return prevID + tmp
		}

		return -1
	}

	for ID := search(); ID >= 0 && prevID < size; ID = search() {
		if ID > prevID {
			newIt = l.InsertAfter(str[prevID:ID], newIt)
		}

		newIt = l.InsertAfter(str[ID:ID+1], newIt)

		prevID = ID + 1
	}

	if prevID != 0 && prevID < size {
		newIt = l.InsertAfter(str[prevID:], newIt)
	}

	if prevID != 0 {
		tmp := it.Next()
		l.Remove(it)
		it = tmp
	}

	if p.IsDebug {
		fmt.Print(str, " splited by ", string(sep), " to: ")
		for prntIt := it; it != newIt.Next() && prntIt != nil; prntIt = prntIt.Next() {
			fmt.Print(prntIt.Value.(string), "|")
		}
		fmt.Println()
	}

	return newIt
}
