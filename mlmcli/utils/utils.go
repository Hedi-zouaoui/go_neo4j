package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	
	"net/http"


)


func AddUser(task AddUser_request) AddUser_response {
	url := "http://localhost:8089/v1/Adduser"
	fmt.Println("task in add user function " , task) 
	taskJson, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}
fmt.Println(taskJson)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(taskJson)))
	fmt.Println("req est : " , req )
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)

	AddUser_response  := AddUser_response{}
	_ = json.Unmarshal(responseData, &AddUser_response )

	return 	AddUser_response 
}



func TransferUser(task TransferUser_request) TransferUser_response {
	url := "http://localhost:8088/v1/transfer_user"
	fmt.Println("task in add user function " , task) 
	taskJson, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}
fmt.Println(taskJson)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(taskJson)))
	fmt.Println("req est : " , req )
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)

	TransferUser_response := TransferUser_response{}
	_ = json.Unmarshal(responseData, &TransferUser_response )

	return 	TransferUser_response 
}





func DeleteUser(task DeleteUser_request) DeleteUser_response {
	url := "http://localhost:8087/v1/Deleteuser"
	fmt.Println("task in add user function " , task) 
	taskJson, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}
fmt.Println(taskJson)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(taskJson)))
	fmt.Println("req est : " , req )
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)

	DeleteUser_response := DeleteUser_response{}
	_ = json.Unmarshal(responseData, &DeleteUser_response )

	return 	DeleteUser_response 
}
