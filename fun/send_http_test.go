package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

//미완~
func TestManitto(t *testing.T) {
	nameList := []string{"삐삐", "미미", "티티", "피피", "디디", "지지"}
	exit := false
	var intList []int
	var randResult []int
	var result []string
	var resultNum []int
	max := 0

	for !exit {
		for _, _ = range nameList {
			max++
			intList = append(intList, max)
		}
		//fmt.Println(max, intList)

		for i := 0; i < max; i++ {
			index := len(nameList) - i
			rn := randomNumber(0, index)
			if intContains(randResult, intList[rn]) {
				break
			}
			randResult = append(randResult, intList[rn])
			intList[rn] = index
			//fmt.Println(i, index, rn, intList[rn], intList)
		}
		if len(randResult) != max {
			intList = []int{}
			randResult = []int{}
			max = 0
			continue
		}

		r := []string{}
		for i, v := range randResult {
			if nameList[i] == nameList[v-1] {
				intList = []int{}
				randResult = []int{}
				max = 0
				break
			} else {
				r = append(r, nameList[v-1])
			}
		}
		//fmt.Println("]", len(r) != max)
		if len(r) != max {
			intList = []int{}
			randResult = []int{}
			max = 0
			continue
		} else {
			if len(r) != len(nameList) {
				continue
			}
			result = r
			resultNum = randResult
			break
		}

	}
	fmt.Println("리절트 : ", randResult)
	fmt.Println("목록 : ", nameList)
	fmt.Println("인트리스트 : ", intList)
	//fmt.Printf("[")
	if len(result) != len(nameList) {
		fmt.Println("왜다름;;", len(result), len(nameList))
		return
	}
	fmt.Println("결과 :", result)
	token := "AwRIStuvkTVSitZyvyCacMujm9bCLb"
	roomidList := map[string]string{
		"티티": "wrqwer1234",
		"미미": "qwrqwrqwrwqfasfdafff",
		"삐삐": "afdafa",
		"피피": "assfadf",
		"디디": "asdfasdfs",
		"지지": "afasfasfvzvzx",
	}

	for i, v := range nameList {
		sendWetalk("번호:"+strconv.Itoa(resultNum[i])+" "+result[i], token, roomidList[v])

	}
}
func sendWetalk(content string, token string, roomid string) {
	body := url.Values{}
	for k, v := range map[string]string{
		"id_of_roooooooom": roomid,
		"type":             "1",
		"content_type":     "0",
		"content":          content,
		"c":                "9",
	} {
		body.Set(k, v)
	}

	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Bearer " + token,
	}
	req, err := http.NewRequest("POST", "http://10.70/c/t/send", strings.NewReader(body.Encode()))
	if err != nil {
		fmt.Println("에러", err)
	}
	for key, val := range header {
		req.Header.Add(key, val)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("에러", err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		fmt.Println("에러", str, err)
	}

	res := HttpResult{}
	if err := json.Unmarshal(respBody, &res); err != nil {
		fmt.Println("에러", err)
	}
}
func intContains(l []int, p int) bool {
	for _, v := range l {
		if v == p {
			return true
		}
	}
	return false
}
func isUnique(s []int) bool {
	for _, val := range s {
		if intContains(s, val) {
			return false
		}
	}
	return true
}

func randomNumber(min, max int) int { // cno
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

type HttpResult struct {
	ResultCode interface{} `json:"resultCode"`
	ResultMsg  string      `json:"resultMsg"`
	ResultData interface{} `json:"resultData,omitempty"`
}
