/*

Package nid is used to create nids (slugs/tags)

Installation

Just go get the package:

    go get -u github.com/TV4/nid

Usage

A small usage example

    package main

    import (
    	"fmt"

    	"github.com/TV4/nid"
    )

    func main() {
    	fmt.Println(nid.Case("Let's_Dance ")) // lets-dance
    }

*/
package nid

import (
	"regexp"
	"strings"
)

// Default is the default configuration used when calling nid.Case and nid.Possible
var Default = New()

// WithÅÄÖ allows åäö in nids (and replaces ø with ö)
var WithÅÄÖ = New(AllowÅÄÖ)

// Case returns a nid based on the input text
func Case(text string) string {
	return Default.Case(text)
}

// Possible checks if a candidate string is a possible nid
func Possible(candidate string) bool {
	return Default.Possible(candidate)
}

// Nid contains the configuration used to create and validate nids
type Nid struct {
	ValidPattern     *regexp.Regexp
	SquishPattern    *regexp.Regexp
	StripPattern     *regexp.Regexp
	DashPattern      *regexp.Regexp
	DashSpace        *strings.Replacer
	Transliterations *strings.Replacer
}

// New returns a *Nid that implements the Interface
func New(options ...func(*Nid)) *Nid {
	n := &Nid{
		ValidPattern:     validPattern,
		SquishPattern:    squishPattern,
		StripPattern:     stripPattern,
		DashPattern:      dashPattern,
		DashSpace:        dashSpace,
		Transliterations: transliterations,
	}

	for _, option := range options {
		option(n)
	}

	return n
}

// AllowÅÄÖ is a functional option that can be used in New
func AllowÅÄÖ(n *Nid) {
	SetValidPattern(regexp.MustCompile(`\A[0-9a-zåäö-]*\z`))(n) // Note: Allows empty nids
	SetStripPattern(regexp.MustCompile(`[^0-9a-zåäö-]`))(n)
	SetTransliterations(transliterationsWithÅÄÖ)(n)
}

// SetValidPattern is a functional option that can be used in New
func SetValidPattern(r *regexp.Regexp) func(*Nid) {
	return func(n *Nid) {
		if r != nil {
			n.ValidPattern = r
		}
	}
}

// SetStripPattern is a functional option that can be used in New
func SetStripPattern(r *regexp.Regexp) func(*Nid) {
	return func(n *Nid) {
		if r != nil {
			n.StripPattern = r
		}
	}
}

// SetTransliterations is a functional option that can be used in New
func SetTransliterations(r *strings.Replacer) func(*Nid) {
	return func(n *Nid) {
		if r != nil {
			n.Transliterations = r
		}
	}
}

// Case returns a nid based on the input text
func (n *Nid) Case(text string) string {
	if text == "" {
		return ""
	}

	return n.strip(n.transliterate(n.squish(n.prepare(text))))
}

// Possible checks if a candidate string is a possible nid
func (n *Nid) Possible(candidate string) bool {
	if strings.Contains(candidate, "--") {
		return false
	}

	return n.ValidPattern.MatchString(candidate)
}

func (n *Nid) strip(s string) string {
	return n.DashPattern.ReplaceAllString(n.StripPattern.ReplaceAllString(s, ""), "-")
}

func (n *Nid) transliterate(s string) string {
	return n.Transliterations.Replace(s)
}

func (n *Nid) squish(s string) string {
	return n.SquishPattern.ReplaceAllString(s, " ")
}

func (n *Nid) prepare(s string) string {
	return strings.TrimSpace(n.DashSpace.Replace(strings.ToLower(s)))
}

var (
	validPattern     = regexp.MustCompile(`\A[0-9a-z-]{1,}\z`)
	squishPattern    = regexp.MustCompile(`\s+`)
	stripPattern     = regexp.MustCompile(`[^0-9a-z-]`)
	dashPattern      = regexp.MustCompile("-{1,}")
	dashSpace        = strings.NewReplacer("-", " ", "_", " ", "–", " ", "—", " ")
	transliterations = strings.NewReplacer(
		" ", "-",
		"_", "-",
		"×", "x",
		"ß", "ss",
		"à", "a",
		"á", "a",
		"â", "a",
		"ã", "a",
		"ä", "a",
		"å", "a",
		"æ", "a",
		"ç", "c",
		"è", "e",
		"é", "e",
		"ê", "e",
		"ë", "e",
		"ì", "i",
		"í", "i",
		"î", "i",
		"ï", "i",
		"ð", "d",
		"ñ", "n",
		"ò", "o",
		"ó", "o",
		"ô", "o",
		"õ", "o",
		"ö", "o",
		"ø", "o",
		"ù", "u",
		"ú", "u",
		"û", "u",
		"ü", "u",
		"ý", "y",
		"þ", "th",
		"ÿ", "y",
		"ā", "a",
		"ă", "a",
		"ą", "a",
		"ć", "c",
		"ĉ", "c",
		"ċ", "c",
		"č", "c",
		"ď", "d",
		"đ", "d",
		"ē", "e",
		"ĕ", "e",
		"ė", "e",
		"ę", "e",
		"ě", "e",
		"ĝ", "g",
		"ğ", "g",
		"ġ", "g",
		"ģ", "g",
		"ĥ", "h",
		"ħ", "h",
		"ĩ", "i",
		"ī", "i",
		"ĭ", "i",
		"į", "i",
		"ı", "i",
		"ĳ", "ij",
		"ĵ", "j",
		"ķ", "k",
		"ĸ", "k",
		"ĺ", "l",
		"ļ", "l",
		"ľ", "l",
		"ŀ", "l",
		"ł", "l",
		"ń", "n",
		"ņ", "n",
		"ň", "n",
		"ŉ", "'n",
		"ŋ", "ng",
		"ō", "o",
		"ŏ", "o",
		"ő", "o",
		"œ", "oe",
		"ŕ", "r",
		"ŗ", "r",
		"ř", "r",
		"ś", "s",
		"ŝ", "s",
		"ş", "s",
		"š", "s",
		"ţ", "t",
		"ť", "t",
		"ŧ", "t",
		"ũ", "u",
		"ū", "u",
		"ŭ", "u",
		"ů", "u",
		"ű", "u",
		"ų", "u",
		"ŵ", "w",
		"ŷ", "y",
		"ź", "z",
		"ż", "z",
		"ž", "z",
	)

	transliterationsWithÅÄÖ = strings.NewReplacer(
		" ", "-",
		"_", "-",
		"×", "x",
		"ß", "ss",
		"à", "a",
		"á", "a",
		"â", "a",
		"ã", "a",
		"ä", "ä",
		"å", "å",
		"æ", "ä",
		"ç", "c",
		"è", "e",
		"é", "e",
		"ê", "e",
		"ë", "e",
		"ì", "i",
		"í", "i",
		"î", "i",
		"ï", "i",
		"ð", "d",
		"ñ", "n",
		"ò", "o",
		"ó", "o",
		"ô", "o",
		"õ", "o",
		"ö", "ö",
		"ø", "ö",
		"ù", "u",
		"ú", "u",
		"û", "u",
		"ü", "u",
		"ý", "y",
		"þ", "th",
		"ÿ", "y",
		"ā", "a",
		"ă", "a",
		"ą", "a",
		"ć", "c",
		"ĉ", "c",
		"ċ", "c",
		"č", "c",
		"ď", "d",
		"đ", "d",
		"ē", "e",
		"ĕ", "e",
		"ė", "e",
		"ę", "e",
		"ě", "e",
		"ĝ", "g",
		"ğ", "g",
		"ġ", "g",
		"ģ", "g",
		"ĥ", "h",
		"ħ", "h",
		"ĩ", "i",
		"ī", "i",
		"ĭ", "i",
		"į", "i",
		"ı", "i",
		"ĳ", "ij",
		"ĵ", "j",
		"ķ", "k",
		"ĸ", "k",
		"ĺ", "l",
		"ļ", "l",
		"ľ", "l",
		"ŀ", "l",
		"ł", "l",
		"ń", "n",
		"ņ", "n",
		"ň", "n",
		"ŉ", "'n",
		"ŋ", "ng",
		"ō", "o",
		"ŏ", "o",
		"ő", "o",
		"œ", "oe",
		"ŕ", "r",
		"ŗ", "r",
		"ř", "r",
		"ś", "s",
		"ŝ", "s",
		"ş", "s",
		"š", "s",
		"ţ", "t",
		"ť", "t",
		"ŧ", "t",
		"ũ", "u",
		"ū", "u",
		"ŭ", "u",
		"ů", "u",
		"ű", "u",
		"ų", "u",
		"ŵ", "w",
		"ŷ", "y",
		"ź", "z",
		"ż", "z",
		"ž", "z",
	)
)
