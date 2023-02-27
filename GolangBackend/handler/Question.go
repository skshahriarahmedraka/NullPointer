package handler

import (
	"app/model"
	"encoding/json"
	// "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)







func (H *DatabaseCollections)Question(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	x:= model.QuestionData{
		QuestionTitle: "How to efficiently concatenate strings in go " ,
		QuesitonAskedTime :"12 years, 5 months ago",
		QuestionModified : "6 months ago",
		QuestionViewed: 577000,
		QuestionUpvote:4,
		QuestionDownvote:32,
		QuestionBookmark:23,
		QuestionTags:[]string{"go","string","Concatination","go","string","Concatination","go","string","Concatination","go","string","go","string","Concatination","Concatination"},
		QuestionAskedBy:"yujiitadori",
		QuestionAskedByName:"yuji itadori prince vageta son goku",
		QuestionAskedByBadge:[]int{999,888,777},
		QuestionAskedTimeExact:"Apr 1, 2019 at 21:38",
		QuestionAskedByImage:"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg",
		QuestionEditedBy:"KingofRandom",
		QuestionEditedByName:"King of Random prince of all sayian",
		QuestionEditedByBadge:[]int{666,555,444},
		QuestionEditedByImage:"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg",
		QuestionEditedTimeExact:"Jan 12, 2010 at 16:18",
		QuestionDescription:`
		In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.
		In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.
		In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.
		` ,
		QuestionComment:[][]string{
				{"Note: This question and most answers seem to have been written before append() came into the language, which is a good solution for this. It will perform fast like copy() but will grow the slice firs","thomasrutter","Aug 10, 2017 at 1:29"},
				{"It doesn't just seem very inefficient it has a specific problem that every new non-CS hire we have ever gotten runs into in the first few weeks on the job. It's quadratic - O(n*n). Think of the number sequence: 1 + 2 + 3 + 4 + .... It's n*(n+1)/2","rob lucci","Nov 16, 2017 at 15:22"}},
		Answers:[][]interface{}{
			// ["ans(string)","Upvote(int)","downvote(int)","accepted(bool)","Answer(string)","modifiedBy(string)","time(string)","time(string)","EditorImgaelink(string)","answeredImgaelink(string) "]
			{`994New Way: From Go 1.10 there is a strings.Builder type, please take a look at this answer for more detail. Old Way: Use the bytes package. It has a Buffer type which implements io.Writer. package main import ( "bytes" "fmt" ) func main() { var buffer bytes.Buffer for i := 0; i < 1000; i++ { buffer.WriteString("a") } fmt.Println(buffer.String()) } This does it in O(n) time.`,432,123,true,"inancramsay","Garp","Dec 12, 2019 at 7:30","Jul 9, 2019 at 10:42",[]int{12,99,999},[]int{99,23,11},"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"},
			{`n Go 1.10+ there is strings.Builder, here.A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use.`,4,43,false,"luffy","icza", "Dec 4, 2015 at 8:01","Oct 9, 2019 at 10:42",[]int{11,9,33},[]int{42,535,22},"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"},
			{`n Go 1.10+ there is strings.Builder, here.A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use.`,6,3,false,"Dragon","icza", "Dec 4, 2015 at 8:01","Jan 9, 2019 at 10:42",[]int{12,24,53},[]int{99,23,355},"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"},
			{`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,98,231,false,"monkey D. luffy","annya","May 9, 2019 at 10:42","",[]int{13,354,666},[]int{11,32,535},"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"},
			{`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,98,1,true ,"Sabo Cheif of stuff","Yur Fozer","Aug 9, 2019 at 10:42","Dec 21, 2022 at 10:42",[]int{66,2,5},[]int{99,11,0},"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"},
			{`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,9,1,false,"Garp Hero of Nevy","annya","Jun 9, 2019 at 10:42","",[]int{23,53,64},[]int{99,23,24},"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"},
			{`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,8,31,true,"Akayinu Sakazuki","Spy","Aug 9, 2019 at 10:42","Feb 11, 2020 at 10:42",[]int{76,37,34},[]int{12,23,43},"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"},
			{`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,9,2,false,"Muzan kibutsuji","liverpool","Apr 9, 2019 at 10:42","",[]int{17,44,53},[]int{12,23,53},"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"},
			{`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,9,31,false,"Fleet Admiral","rengoku","Mar 9, 2019 at 10:42","",[]int{14,35,640},[]int{63,78,9},"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"}}}
	
			z,err := json.Marshal(x)
			if err !=nil {
				log.Fatalln("🔥❌",err)
			}
			// fmt.Println("🚀 ~ file: Question.go ~ line 58 ~ func ~ z : ", z)
			var y model.QuestionData
			json.Unmarshal(z,&y)
			// fmt.Println("🚀 ~ file: Question.go ~ line 71 ~ func ~ y : ", y)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.JSON(http.StatusOK,x)	

}


