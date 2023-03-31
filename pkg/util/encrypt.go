package util

import (
	"regexp"
	"strconv"
	"strings"
)

// EmojiEncode Emoji表情转码
func EmojiDecode(s string) string {
	//emoji表情的数据表达式
	var re *regexp.Regexp
	var reg *regexp.Regexp
	var src []string
	re = regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]")
	//提取emoji数据表达式
	reg = regexp.MustCompile("\\[\\\\u|]")
	src = re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		var e = reg.ReplaceAllString(src[i], "")
		var p, err = strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}

//表情转换
func EmojiCode(s string) string {
	var ret string
	var rs []rune
	rs = []rune(s)

	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			ret += `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
		} else {
			ret += string(rs[i])
		}
	}
	return ret
}
