package utils

import "net/http"
import "io/ioutil"
import "fmt"
import "strings"
import "io"
import "bytes"
import (
	"time"
)


type Error struct {
	Message string
	Status string
}

func DecodeRequestBody(req *http.Request) []byte {
		defer req.Body.Close()		
		body, err   := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		}
		return body
}

func DecodeResponseBody(res *http.Response) []byte {
	defer res.Body.Close()			
	body, err   := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func checkBearerToken(header interface{}) bool {
	
	var authHead string
	auth := map[string]interface{}
	
	switch v :=  header.(type)  {
		case http.Request:
		auth["request"] = header.(http.Request).Header.Get("Authorization")
		case *http.Response:
		auth["response"] = header.(*http.Response).Header.Get("Authorization")
	}
	if auth["request"].(string) != "" {
		authHead = auth["request"].(string)
	} else if auth["response"].(string) != "" {
		authHead = auth["response"].(string)
	}
	
	
	tokenType := strings.Split(authHead, " ")[0]
	token := strings.Split(authHead, " ")[1]
	
	if tokenType != "Bearer" {
		return false
	}
	if token == "" {
		return false
	}
	return true
}


func makeRequest (method string, url string, body io.Reader) *http.Response {
		 req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	         CheckError(err, "Wrong " + method + " request")
		
		 client := http.Client{
			Timeout: 30 * time.Second,
		}
		res, err := client.Do(req)
	        CheckError(err, method + " Response error: ")
}		return res


func CheckError(err error, message string) {
		
  if err != nil {
	  fmt.Println(message + err.Error())
	}
}

func S(bytes []byte) string {
	return string(bytes)
}
