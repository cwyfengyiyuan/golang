package parser

import (
	"learn/crawler/engine"
	"learn/crawler_distributed/config"
	"regexp"
)
var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	genderUrlRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
)
func ParseCity(contents []byte, _ string) engine.ParseResult {
	var str []string
	result := engine.ParseResult{}
	matches := genderUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		str = append(str, string(m[1]))
	}
	matches = profileRe.FindAllSubmatch(contents, -1)
	flag := 0
	for _, m := range matches {
		gender := str[flag]
		flag++
		if flag >= len(str) {
			flag = 0
		}
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			Parser: NewProfileParser(string(m[2]), gender),
		})
	}
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}