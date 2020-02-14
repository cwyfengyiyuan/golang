package parser

import (
	"learn/crawler/engine"
	"learn/crawler/model"
	"regexp"
)

const all = `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`
var idUrl = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)
func ParseProfile(contents []byte, name string, gender string, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender
	re := regexp.MustCompile(all)
	match := re.FindAllSubmatch(contents,-1)
	if match != nil {
		for _, m := range match {
			profile.User = append(profile.User, string(m[1]))
		}
	}
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:url,
				Type:"zhenai",
				Id:extractString([]byte(url), idUrl),
				Payload:profile,
			},
		},
	}
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string  {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}else {
		return ""
	}
}
func ProfileParser(name string, gender string) engine.ParserFunc  {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, name, gender, url )
	}
}