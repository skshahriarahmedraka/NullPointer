package model

type RegisterData struct {
	UserName string `json:"UserName" bson:"UserName" validate:"omitempty,min=4,max=64,UserName"`
	Email    string `json:"Email" bson:"Email" validate:"required,min=4,max=100,Email"`
	Password string `json:"Password"	bson:"Password" validate:"omitempty,min=4,max=100"`
}
