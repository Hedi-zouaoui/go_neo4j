package utils




type AddUser_response struct {
	 NewId int64  `json:"to"`
     Name string  `json:"name"`
	 Err string `json:"err"`
}

type TransferUser_response struct {
	  To int64  `json:"to"`
	  From int64  `json:"from"`
     Methode string  `json:"methode"`

}

type DeleteUser_response struct {
	Id int64  	`json:"id"`
	Err string     `json:"err"`
}




