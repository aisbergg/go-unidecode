package unidecode_test

import (
	"strings"
	"testing"

	"github.com/aisbergg/go-unidecode/pkg/unidecode"
)

func seedCorpus(f *testing.F) {
	f.Add("", uint8(0), "")
	f.Add("abc", uint8(1), "abc")
	f.Add("北京", uint8(2), "?")
	f.Add("abc北京", uint8(3), "0")
	f.Add("ネオアームストロングサイクロンジェットアームストロング砲", uint8(0), "")
	f.Add("30 𝗄𝗆/𝗁", uint8(0), "")
	f.Add("kožušček", uint8(0), "")
	f.Add("ⓐⒶ⑳⒇⒛⓴⓾⓿", uint8(0), "")
	f.Add("Hello, World!", uint8(0), "")
	f.Add(`\n`, uint8(0), "")
	f.Add(`北京abc\n`, uint8(0), "")
	f.Add(`'"\r\n`, uint8(0), "")
	f.Add("ČŽŠčžš", uint8(0), "")
	f.Add("ア", uint8(0), "")
	f.Add("α", uint8(0), "")
	f.Add("a", uint8(0), "")
	f.Add("ch\u00e2teau", uint8(0), "")
	f.Add("vi\u00f1edos", uint8(0), "")
	f.Add("Efﬁcient", uint8(0), "")
	f.Add("příliš žluťoučký kůň pěl ďábelské ódy", uint8(0), "")
	f.Add("PŘÍLIŠ ŽLUŤOUČKÝ KŮŇ PĚL ĎÁBELSKÉ ÓDY", uint8(0), "")
	f.Add("\ua500", uint8(0), "")
	f.Add("\u1eff", uint8(0), "")
	f.Add("\U000fffff", uint8(0), "")
	f.Add("\U0001d5a0", uint8(0), "")
	f.Add("\U0001d5c4\U0001d5c6/\U0001d5c1", uint8(0), "")
	f.Add("\u2124\U0001d552\U0001d55c\U0001d552\U0001d55b \U0001d526\U0001d52a\U0001d51e \U0001d4e4\U0001d4f7\U0001d4f2\U0001d4ec\U0001d4f8\U0001d4ed\U0001d4ee \U0001d4c8\U0001d4c5\u212f\U0001d4b8\U0001d4be\U0001d4bb\U0001d4be\U0001d4c0\U0001d4b6\U0001d4b8\U0001d4be\U0001d4bf\u212f \U0001d59f\U0001d586 \U0001d631\U0001d62a\U0001d634\U0001d622\U0001d637\U0001d626?!", uint8(0), "")
}

func FuzzUnidecode(f *testing.F) {
	seedCorpus(f)
	f.Fuzz(func(t *testing.T, s string, errHandling uint8, replacement string) {
		if errHandling > 3 {
			t.Skip()
		}
		errHnd := unidecode.ErrorHandling(errHandling)
		_, _ = unidecode.Unidecode(s, errHnd, replacement)
	})
}

func FuzzWriter(f *testing.F) {
	seedCorpus(f)
	f.Fuzz(func(t *testing.T, s string, errHandling uint8, replacement string) {
		if errHandling > 3 {
			t.Skip()
		}
		errHnd := unidecode.ErrorHandling(errHandling)

		bld := strings.Builder{}
		bld.Grow(len(s) + len(s)/3)
		w := unidecode.NewWriter(&bld, errHnd, replacement)
		w.WriteString(s)
	})
}
