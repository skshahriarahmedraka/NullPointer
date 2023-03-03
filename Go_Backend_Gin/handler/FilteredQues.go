package handler

import (
	"app/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)







func (H *DatabaseCollections)FiltedQuestion(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Content-Type", "application/json")
	tab:=c.Query("tab")
    fmt.Println("🚀 ~ file: FilteredQues.go ~ line 20 ~ func ~ tab : ", tab)
	if tab=="" {
		
	}else if tab=="newest"{
    fmt.Println("🚀 ~ file: FilteredQues.go ~ line 25 ~ func ~ newest : ")

	}else if tab=="unanswered"{
    fmt.Println("🚀 ~ file: FilteredQues.go ~ line 28 ~ func ~ unanswered : ")

	}else if tab=="bountied"{
    fmt.Println("🚀 ~ file: FilteredQues.go ~ line 31 ~ func ~ bountied : ")

	} else if tab=="voted" {
    fmt.Println("🚀 ~ file: FilteredQues.go ~ line 34 ~ func ~ voted : ")

	}
	x:=model.PublicQues{
		{235,999,1626842,true,[]interface{}{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},[]interface{}{"go","rust","backend","linux"},"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
        {235,9,1626842,true,[]interface{}{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},[]interface{}{"go","rust","backend","linux"},"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
        {0,0,0,false,[]interface{}{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},[]interface{}{"go","rust","backend","linux"},"How to check if a map contains a key in Go? How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec 185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
        {786,999,1626842,false,[]interface{}{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},[]interface{}{"go","rust","backend","linux"},"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
        {235,999,1626842,true,[]interface{}{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},[]interface{}{"go","rust","backend","linux"},"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
        {235,0,1626842,false,[]interface{}{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},[]interface{}{"go","rust","backend","linux"},"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"}, 
	}
	c.JSON(http.StatusOK,x)	

}







        // // {voteNumber(int),AnsNum(int),Viewnum(num),AnsAccepted(bool),askedby(array),tag(Array),QuesTitle(string),QuesDes(string) }
        // {235,999,1626842,true,{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},{"go","rust","backend","linux"},"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
        // {235,9,1626842,true,{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},{"go","rust","backend","linux"},"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
        // {0,0,0,false,{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},{"go","rust","backend","linux"},"How to check if a map contains a key in Go? How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec 185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
        // {786,999,1626842,false,{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},{"go","rust","backend","linux"},"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
        // {235,999,1626842,true,{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},{"go","rust","backend","linux"},"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
        // {235,0,1626842,false,{"yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"},{"go","rust","backend","linux"},"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"},
    
