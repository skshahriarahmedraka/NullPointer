package model



type LoginData struct{
	Email string `json:"Email"`
	Password string `json:"Password"`
}


type RegisterData struct{
	ID string `json:"ID" bson:"_id"`
	UserName string `json:"UserName" bson:"UserName,omitempty"`
	Email string `json:"Email" bson:"Email,omitempty"`
	Password string `json:"Password" bson:"Password,omitempty"`
}