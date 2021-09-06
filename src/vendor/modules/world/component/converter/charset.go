package converter

import (
	"errors"
	"strings"

	"github.com/jarlyyn/ansi"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

//ToUTF8 : convert from CJK encoding to UTF-8
func ToUTF8(from string, s []byte) ([]byte, error) {
	var decoder *encoding.Decoder
	switch strings.ToLower(from) {
	case "gbk", "cp936", "windows-936":
		decoder = simplifiedchinese.GBK.NewDecoder()
	case "gb18030":
		decoder = simplifiedchinese.GB18030.NewDecoder()
	case "gb2312":
		decoder = simplifiedchinese.HZGB2312.NewDecoder()
	case "big5", "big-5", "cp950":
		decoder = traditionalchinese.Big5.NewDecoder()
	case "euc-kr", "euckr", "cp949":
		decoder = korean.EUCKR.NewDecoder()
	case "euc-jp", "eucjp":
		decoder = japanese.EUCJP.NewDecoder()
	case "shift-jis":
		decoder = japanese.ShiftJIS.NewDecoder()
	case "iso-2022-jp", "cp932", "windows-31j":
		decoder = japanese.ISO2022JP.NewDecoder()
	case "utf8":
		return s, nil
	default:
		return s, errors.New("Unsupported encoding " + from)
	}
	d, e := decoder.Bytes(s)
	decoder.Reset()
	if e != nil {
		return nil, e
	}
	return d, nil
}

// FromUTF8 convert from UTF-8 encoding to CJK encoding
func FromUTF8(to string, s []byte) ([]byte, error) {
	var encoder *encoding.Encoder
	switch strings.ToLower(to) {
	case "gbk", "cp936", "windows-936":
		encoder = simplifiedchinese.GBK.NewEncoder()
	case "gb18030":
		encoder = simplifiedchinese.GB18030.NewEncoder()
	case "gb2312":
		encoder = simplifiedchinese.HZGB2312.NewEncoder()
	case "big5", "big-5", "cp950":
		encoder = traditionalchinese.Big5.NewEncoder()
	case "euc-kr", "euckr", "cp949":
		encoder = korean.EUCKR.NewEncoder()
	case "euc-jp", "eucjp":
		encoder = japanese.EUCJP.NewEncoder()
	case "shift-jis":
		encoder = japanese.ShiftJIS.NewEncoder()
	case "iso-2022-jp", "cp932", "windows-31j":
		encoder = japanese.ISO2022JP.NewEncoder()
	case "utf8":
		return s, nil
	default:
		return s, errors.New("Unsupported encoding " + to)
	}
	d, e := encoder.Bytes(s)
	encoder.Reset()
	if e != nil {
		return nil, e
	}
	return d, nil
}

func init() {
	ansi.FlagIgnoreC1 = true
}
