package handler

import (
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (H *DatabaseCollections)FavoriteHash(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Content-Type", "application/json")

	x:= model.FavoriteHash{
		{"go","https://github.com/","https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Google_%22G%22_Logo.svg/2048px-Google_%22G%22_Logo.svg.png"},
		{"rust","https://github.com/","https://cdn-icons-png.flaticon.com/512/732/732221.png"},
		{"fiber","https://github.com/","https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSXyTXn1zirGyXZd15ukEGTupWvTasppbf7N5dW3AT8bc7iUDwFxgI3aRyaGYS2NYQUO9c&usqp=CAU"},
		{"gio","https://github.com/",""},
		{"C++","https://github.com/","https://i.pinimg.com/736x/83/fe/22/83fe2289e31398d96e8e6ed9d603c72f.jpg"},
		{"C#","https://github.com/","https://upload.wikimedia.org/wikipedia/commons/thumb/a/ad/Microsoft_Lists_%282020-present%29.svg/2048px-Microsoft_Lists_%282020-present%29.svg.png"},
		{"dart","https://github.com/","https://seeklogo.com/images/M/microsoft-edge-new-2020-logo-2AD376EBAA-seeklogo.com.png"},
		{"flutter","https://github.com/","https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSmZ-U_nmUBZ6SdJBuOjMOajQplOkMqYmYIVbPU14Co9i7jHKqo4kPpQEOkg_0lCE7_yXg&usqp=CAU"},
		{"linux","https://github.com/","https://upload.wikimedia.org/wikipedia/commons/7/70/Firefox_Focus_2021_Icon.png"},
}
	c.JSON(http.StatusOK,x)
}