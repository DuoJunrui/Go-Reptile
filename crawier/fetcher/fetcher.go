package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Fetch(urlStr string) ([]byte, error) {
	/*请求容易出现403 Forbidden*/
	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}

	// 用户名密码, 若已添加白名单则不需要添加
	username := "t14908251189716"
	password := "yilcxggg"

	// 隧道服务器
	proxy_raw := "tps714.kdlapi.com:15818"
	proxy_str := fmt.Sprintf("http://%s:%s@%s", username, password, proxy_raw)
	proxy, err := url.Parse(proxy_str)

	newUrl := strings.Replace(urlStr, "http://", "https://", 1)
	cookie, err := ioutil.ReadFile("E:\\GolandProjects\\Study\\crawier\\fetcher\\cookie.txt")
	if err != nil {
		panic(err)
	}

	//  请求目标网页
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	req, _ := http.NewRequest("GET", newUrl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Add("Cookie", string(cookie))
	resp, err := client.Do(req)
	if err != nil {
		// 请求发生异常
		fmt.Println(err.Error())
	}
	defer resp.Body.Close() //保证最后关闭Body

	if resp.StatusCode != http.StatusOK {
		if http.StatusAccepted == resp.StatusCode {
			log.Printf("http 202, fetch this url again...")
			Fetch(newUrl)
		}
		return nil, fmt.Errorf("wrong status code： %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
