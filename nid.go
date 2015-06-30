package nid

import (
	"regexp"
	"strings"
)

var (
	validPattern  = regexp.MustCompile(`\A[0-9a-zåäö-]*\z`)
	squishPattern = regexp.MustCompile(`(\s+)`)
	stripPattern  = regexp.MustCompile(`([^0-9a-zåäö-])`)

	dashSpace = strings.NewReplacer("-", " ")
	spaceDash = strings.NewReplacer(" ", "-", "_", "-")
)

// Case returns a nid based on the input text
func Case(text string) string {
	if text == "" {
		return ""
	}

	return cleanup(transliterate(squish(prepare(text))))
}

// Possible checks if a candidate string is a possible nid
func Possible(candidate string) bool {
	return validPattern.MatchString(candidate)
}

func cleanup(s string) string {
	return stripPattern.ReplaceAllString(spaceDash.Replace(s), "")
}

func transliterate(s string) string {
	return transliterations.Replace(s)
}

func squish(s string) string {
	return squishPattern.ReplaceAllString(strings.TrimSpace(s), " ")
}

func prepare(s string) string {
	return dashSpace.Replace(strings.ToLower(s))
}

var transliterations = strings.NewReplacer(
	"À", "A",
	"Á", "A",
	"Â", "A",
	"Ã", "A",
	"Ä", "Ä",
	"Å", "Å",
	"Æ", "Ä",
	"Ç", "C",
	"È", "E",
	"É", "E",
	"Ê", "E",
	"Ë", "E",
	"Ì", "I",
	"Í", "I",
	"Î", "I",
	"Ï", "I",
	"Ð", "D",
	"Ñ", "N",
	"Ò", "O",
	"Ó", "O",
	"Ô", "O",
	"Õ", "O",
	"Ö", "Ö",
	"×", "x",
	"Ø", "Ö",
	"Ù", "U",
	"Ú", "U",
	"Û", "U",
	"Ü", "U",
	"Ý", "Y",
	"Þ", "Th",
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
	"Ā", "A",
	"ā", "a",
	"Ă", "A",
	"ă", "a",
	"Ą", "A",
	"ą", "a",
	"Ć", "C",
	"ć", "c",
	"Ĉ", "C",
	"ĉ", "c",
	"Ċ", "C",
	"ċ", "c",
	"Č", "C",
	"č", "c",
	"Ď", "D",
	"ď", "d",
	"Đ", "D",
	"đ", "d",
	"Ē", "E",
	"ē", "e",
	"Ĕ", "E",
	"ĕ", "e",
	"Ė", "E",
	"ė", "e",
	"Ę", "E",
	"ę", "e",
	"Ě", "E",
	"ě", "e",
	"Ĝ", "G",
	"ĝ", "g",
	"Ğ", "G",
	"ğ", "g",
	"Ġ", "G",
	"ġ", "g",
	"Ģ", "G",
	"ģ", "g",
	"Ĥ", "H",
	"ĥ", "h",
	"Ħ", "H",
	"ħ", "h",
	"Ĩ", "I",
	"ĩ", "i",
	"Ī", "I",
	"ī", "i",
	"Ĭ", "I",
	"ĭ", "i",
	"Į", "I",
	"į", "i",
	"İ", "I",
	"ı", "i",
	"Ĳ", "IJ",
	"ĳ", "ij",
	"Ĵ", "J",
	"ĵ", "j",
	"Ķ", "K",
	"ķ", "k",
	"ĸ", "k",
	"Ĺ", "L",
	"ĺ", "l",
	"Ļ", "L",
	"ļ", "l",
	"Ľ", "L",
	"ľ", "l",
	"Ŀ", "L",
	"ŀ", "l",
	"Ł", "L",
	"ł", "l",
	"Ń", "N",
	"ń", "n",
	"Ņ", "N",
	"ņ", "n",
	"Ň", "N",
	"ň", "n",
	"ŉ", "'n",
	"Ŋ", "NG",
	"ŋ", "ng",
	"Ō", "O",
	"ō", "o",
	"Ŏ", "O",
	"ŏ", "o",
	"Ő", "O",
	"ő", "o",
	"Œ", "OE",
	"œ", "oe",
	"Ŕ", "R",
	"ŕ", "r",
	"Ŗ", "R",
	"ŗ", "r",
	"Ř", "R",
	"ř", "r",
	"Ś", "S",
	"ś", "s",
	"Ŝ", "S",
	"ŝ", "s",
	"Ş", "S",
	"ş", "s",
	"Š", "S",
	"š", "s",
	"Ţ", "T",
	"ţ", "t",
	"Ť", "T",
	"ť", "t",
	"Ŧ", "T",
	"ŧ", "t",
	"Ũ", "U",
	"ũ", "u",
	"Ū", "U",
	"ū", "u",
	"Ŭ", "U",
	"ŭ", "u",
	"Ů", "U",
	"ů", "u",
	"Ű", "U",
	"ű", "u",
	"Ų", "U",
	"ų", "u",
	"Ŵ", "W",
	"ŵ", "w",
	"Ŷ", "Y",
	"ŷ", "y",
	"Ÿ", "Y",
	"Ź", "Z",
	"ź", "z",
	"Ż", "Z",
	"ż", "z",
	"Ž", "Z",
	"ž", "z",
)
