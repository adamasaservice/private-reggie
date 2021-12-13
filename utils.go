package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//GetEnv - Gets the env variable if set, otherwise resturns the fallback
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type Block struct {
	Try   func()
	Catch func(Exception)
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {

	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func HttpGet(url string) ([]byte, int) {

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("erorr", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("erorr", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode, url)
	}

	return body, resp.StatusCode
}

func GetShasum(url string, os string, arch string) string {
	shasums, status := HttpGet(url)

	log.Println(status)

	temp := strings.Split(string(shasums), "\n")
	shasum := ""

	r, _ := regexp.Compile(os + `_` + arch)
	for _, v := range temp {
		if r.MatchString(v) {
			split := strings.Split(v, " ")
			shasum = split[0]
		}

	}
	return shasum

}

func replaceVars(original string, namespace string, name string, version string, os string, arch string) string {

	newStr := strings.ReplaceAll(original, `{{namespace}}`, namespace)
	newStr = strings.ReplaceAll(newStr, `{{name}}`, name)
	newStr = strings.ReplaceAll(newStr, `{{version}}`, version)
	newStr = strings.ReplaceAll(newStr, `{{os}}`, os)
	newStr = strings.ReplaceAll(newStr, `{{arch}}`, arch)

	return newStr
}