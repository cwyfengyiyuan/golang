package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"testing"
)

func TestParseCityList(t *testing.T) {
	const a = `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://album.zhenai.com/u/1727435860", nil)
	if err != nil {
		panic(err)
	}
	//Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36
	//Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	all, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s", all)
	re := regexp.MustCompile(a)
	match := re.FindAllSubmatch(all,-1)
	if match != nil {
		for _, m := range match {
			fmt.Printf("%s ", m[1])
		}
	}
}
