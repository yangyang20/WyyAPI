package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	urlS := "https://music.163.com/weapi/login/cellphone"
	paramsS := "params=xj9pIo4XIyHZhCCwc6T6CU%2BchyUs0CVUl%2Br8F8%2B3PeR5U2whnkHMdHFpWspiWBtIeKBkzQCALWUDZjtPWejxO43EInmEvWtWbLuXAAg8W3qRsJeSXT06XP6zZuGQVK%2FR85RR2ms1VLZjiDNlD1HTwUafdOrmBVRWBEVZ9r6cZyQvjea8h0GxXcbgmw5F%2Ff5FvZoXG5HYN1nUTmL5nueILlneda03bycprIGgZlRl6yU%3D&encSecKey=6f4b4ea7acb3f8e3d54d4bc3ac99af2771250e200bd91773bab461284367186c385b1b0525aa591a27b28fd148fa3c2880b600af6ffa6fd5a2cb6709b3656870e592de678737076b2dd26f73854f4bcd7f6bd7e4920df63c0c6aad3f08596ef5c78a0f6b6e6f1d3beb005f76bba3dd54e80d5defe13154add82276782ad4d99b                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          "
	req, err := http.NewRequest("POST", urlS, strings.NewReader(paramsS))
	if err != nil {
		fmt.Println("requestErr:", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://music.163.com")
	req.Header.Set("Cookie", "os=pc; appver=2.9.7")

	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("res:", string(body))
}
