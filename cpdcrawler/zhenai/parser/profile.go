package parser

import (
	"learn/crawler/engine"
	"learn/crawler/model"
	"regexp"
)

const all = `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`
func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.User = append(profile.User, name)
	re := regexp.MustCompile(all)
	match := re.FindAllSubmatch(contents,-1)
	if match != nil {
		for _, m := range match {
			profile.User = append(profile.User, string(m[1]))
		}
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}