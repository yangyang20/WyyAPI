package common

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Options struct {
	Crypto string
	Cookie string
	Proxy  string
	RealIP string
	Ua     string
	Url    string
}

func Request(method string, urlS string, params interface{}, options Options) *http.Response {
	jar, _ := cookiejar.New(nil)
	trans := &http.Transport{
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: trans,
		Jar:       jar,
	}

	var data map[string]string
	switch options.Crypto {
	case "weapi":
		re3, _ := regexp.Compile("\\w*api")
		urlS = re3.ReplaceAllString(urlS, "weapi")
		data = WeapiCrypto(params)
	case "eapi":
		re3, _ := regexp.Compile("\\w*api")
		urlS = re3.ReplaceAllString(urlS, "eapi")
		data = EapiCryptp(options.Url, params)
	default:
		panic("meiyou jiami leixing")
	}

	var paramsS string
	for k, v := range data {
		paramsS += k + "=" + url.QueryEscape(v) + "&"
	}
	paramsS = strings.TrimRight(paramsS, "&")
	//fmt.Println("str:", params)

	req, err := http.NewRequest(method, urlS, strings.NewReader(paramsS))
	if err != nil {
		fmt.Println("requestErr:", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if options.Cookie != "" {
		req.Header.Set("Cookie", options.Cookie)
	}
	if Global_Cookie != nil {
		for _, cookie := range Global_Cookie {
			req.AddCookie(cookie)
		}
	}

	req.Header.Set("Referer", "https://music.163.com")

	if options.Crypto == "eapi" {

		cookie, err := req.Cookie("osver")
		if err == nil {
			//系统版本
			req.Header.Set("osver", cookie.Value)
		}
		cookie, err = req.Cookie("deviceId")
		if err == nil {
			//encrypt.base64.encode(imei + '\t02:00:00:00:00:00\t5106025eb79a5247\t70ffbaac7')
			req.Header.Set("deviceId", cookie.Value)
		}

		cookie, err = req.Cookie("appver")
		if err == nil {
			req.Header.Set("appver", cookie.Value) // app版本
		} else {
			req.Header.Set("appver", "8.7.01") // app版本
		}

		cookie, err = req.Cookie("mobilename")
		if err == nil {
			//设备model
			req.Header.Set("mobilename", cookie.Value)
		}

		cookie, err = req.Cookie("versioncode")
		if err == nil {
			//versioncode: cookie.versioncode || '140', //版本号
			req.Header.Set("versioncode", cookie.Value)
		} else {
			req.Header.Set("versioncode", "140")
		}

		cookie, err = req.Cookie("resolution")
		if err == nil {
			// //设备分辨率
			req.Header.Set("resolution", cookie.Value)
		} else {
			req.Header.Set("resolution", "1920x1080")
		}

		cookie, err = req.Cookie("buildver")
		if err == nil {
			req.Header.Set("buildver", cookie.Value)
		} else {
			req.Header.Set("buildver", strconv.FormatInt(time.Now().Unix(), 10))
		}

		cookie, err = req.Cookie("buildver")
		if err == nil {
			req.Header.Set("__csrf", cookie.Value)
		}

		cookie, err = req.Cookie("os")
		if err == nil {
			req.Header.Set("os", cookie.Value)
		} else {
			req.Header.Set("os", "android")
		}
		cookie, err = req.Cookie("channel")
		if err == nil {
			req.Header.Set("channel", cookie.Value)
		}

		//req.Header.Set("requestId", "")
	}

	resp, err := client.Do(req)

	//defer resp.Body.Close()

	//resp.Cookies()

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	// handle error
	//}
	//fmt.Println("resp:", resp)
	return resp
}
