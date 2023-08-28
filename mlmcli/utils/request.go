package utils



type AddUser_request struct {
	  To int64  `json:"to"`
     Name string  `json:"name"`
}



type TransferUser_request struct {
	  To int64  `json:"to"`
	  From int64  `json:"from"`
     Methode string  `json:"methode"`

}




type DeleteUser_request struct {
  Id int64  `json:"id"`

}


