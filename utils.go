package utils

import "net/http"
import "io/ioutil"
import "fmt"





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


func CheckError(err error) {
		
  if err != nil {
		fmt.Println(err)
	}
}

func S(bytes []byte) string {
	return string(bytes)
}
