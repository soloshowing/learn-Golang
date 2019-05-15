package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([未婚|离异|丧偶])</div>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(contents, ageRe))

	if err != nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := marriageRe.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
