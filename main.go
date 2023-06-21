package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var kFlag = flag.String("k", "こんにちわ", "hiragana string value that is converted to romaji.")
var version string = "vX.X.X"

// https://www.ezairyu.mofa.go.jp/passport/hebon.html
var kanaToRomaji = map[rune]string{
	'あ': "a", 'い': "i", 'う': "u", 'え': "e", 'お': "o",
	'か': "ka", 'き': "ki", 'く': "ku", 'け': "ke", 'こ': "ko",
	'さ': "sa", 'し': "shi", 'す': "su", 'せ': "se", 'そ': "so",
	'た': "ta", 'ち': "chi", 'つ': "tsu", 'て': "te", 'と': "to",
	'な': "na", 'に': "ni", 'ぬ': "nu", 'ね': "ne", 'の': "no",
	'は': "ha", 'ひ': "hi", 'ふ': "fu", 'へ': "he", 'ほ': "ho",
	'ま': "ma", 'み': "mi", 'む': "mu", 'め': "me", 'も': "mo",
	'や': "ya", 'ゆ': "yu", 'よ': "yo",
	'ら': "ra", 'り': "ri", 'る': "ru", 'れ': "re", 'ろ': "ro",
	'わ': "wa", 'を': "wo",
	'ん': "n",
	'が': "ga", 'ぎ': "gi", 'ぐ': "gu", 'げ': "ge", 'ご': "go",
	'ざ': "za", 'じ': "ji", 'ず': "zu", 'ぜ': "ze", 'ぞ': "zo",
	'だ': "da", 'ぢ': "ji", 'づ': "zu", 'で': "de", 'ど': "do",
	'ば': "ba", 'び': "bi", 'ぶ': "bu", 'べ': "be", 'ぼ': "bo",
	'ぱ': "pa", 'ぴ': "pi", 'ぷ': "pu", 'ぺ': "pe", 'ぽ': "po",
	'っ': "x", 'ー': "",
}
var kanaToRomajiCombo = map[string]string{
	"きゃ": "kya", "きゅ": "kyu", "きょ": "kyo",
	"しゃ": "sha", "しゅ": "shu", "しょ": "sho",
	"ちゃ": "cha", "ちゅ": "chu", "ちょ": "cho",
	"にゃ": "nya", "にゅ": "nyu", "にょ": "nyo",
	"ひゃ": "hya", "ひゅ": "hyu", "ひょ": "hyo",
	"みゃ": "mya", "みゅ": "myu", "みょ": "myo",
	"りゃ": "rya", "りゅ": "ryu", "りょ": "ryo",
	"ぎゃ": "gya", "ぎゅ": "gyu", "ぎょ": "gyo",
	"じゃ": "ja", "じゅ": "ju", "じょ": "jo",
	"びゃ": "bya", "びゅ": "byu", "びょ": "byo",
	"ぴゃ": "pya", "ぴゅ": "pyu", "ぴょ": "pyo",

	"しぇ": "she", "ちぇ": "che", "てぃ": "ti",
	"ふぁ": "fua", "ふぃ": "fui", "ふぇ": "fue",
	"でぃ": "dei", "でゅ": "deyu", "うぃ": "ui",
	"ゔぁ": "ba", "ゔぃ": "bi", "ゔ": "bu", "ゔぇ": "be", "ゔぉ": "bo",
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Name:     %s\nVersion:  %s\nOPTIONS:\n",
			os.Args[0], version,
		)
		flag.PrintDefaults()
	}
	flag.Parse()
	x, err := convertKanaToRomaji(*kFlag)
	if err != nil {
		panic(err)
	}
	fmt.Println(x)

}

func convertKanaToRomaji(ks string) (rs string, err error) {
	kr := []rune(ks)
	for i := 0; i < len(kr); i++ {
		if v, ok := kanaToRomaji[kr[i]]; ok {
			if i+1 == len(kr) {
				rs += v
				continue
			}
			if v, ok := kanaToRomajiCombo[string(kr[i:i+2])]; ok {
				rs += v
				i++
				continue
			}
			rs += v
			continue
		}
		return rs, errors.New("invalid character: " + string(kr[i]))
	}

	rs = rule1(rs)
	rs = rule2(rs)
	rs = rule3(rs)
	return rs, err
}

func rule1(rs string) string {
	rr := []rune(rs)
	for i := 0; i < len(rr)-1; i++ {
		x := string(rr[i : i+2])
		if x == "nb" || x == "nm" || x == "np" {
			rr[i] = 'm'
		}
	}
	return string(rr)
}

func rule2(rs string) string {
	x := kanaToRomaji['っ']
	rr := []rune(rs)
	for i := len(rr) - 1; i > 0; i-- {
		if string(rr[i]) == x {
			if i+3 < len(rr) {
				y := string(rr[i+1 : i+4])
				if y == "chi" || y == "cha" || y == "chu" || y == "cho" {
					rr[i] = 't'
					continue
				}
			}
			rr[i] = rr[i+1]
		}
	}
	return string(rr)
}

func rule3(rs string) string {
	rr := []rune(rs)
	for i := 0; i < len(rr)-1; i++ {
		x := string(rr[i : i+2])
		if x == "uu" || x == "ou" {
			rr = append(rr[:i+1], rr[i+2:]...)
		}
		if x == "oo" && i != len(rr)-2 {
			rr = append(rr[:i+1], rr[i+2:]...)
		}
		if x == "ou" && i == len(rr)-2 {
			rr = rr[:i+1]
		}
	}
	return string(rr)
}
