package utils

import "net/http"
import "io/ioutil"
import "fmt"
import "strings"




type Error struct {
	Message string
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
	
	switch(header.(type)) {
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


func CheckError(err error) {
		
  if err != nil {
		fmt.Println(err)
	}
}

func S(bytes []byte) string {
	return string(bytes)
}
