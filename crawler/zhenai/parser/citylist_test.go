package parser

import (
	"fmt"
	"regexp"
	"testing"
)
//Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36
//Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)
var u = "http://album.zhenai.com/u/1135238424"
func TestParseCityList(t *testing.T) {
	//const a = `<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`
	//
	//client := &http.Client{}
	//req, err := http.NewRequest("GET", "http://www.zhenai.com/zhenghun/nanjing", nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//all, _ := ioutil.ReadAll(resp.Body)
	////log.Printf("%s", all)
	//re := regexp.MustCompile(a)
	//match := re.FindAllSubmatch(all,-1)
	//if match != nil {
	//	for _, m := range match {
	//		fmt.Printf("%s ", m[1])
	//	}
	//}
	s := extractString([]byte(u), idUrlRe)
	fmt.Println(s)
}
