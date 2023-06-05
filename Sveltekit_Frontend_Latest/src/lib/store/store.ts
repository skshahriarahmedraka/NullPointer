import { writable } from "svelte/store";
import type {UserDataType} from './types'

// const  Loading=writable(false)

// let UserData = writable({})
// let QuestionData = writable({})
// let RelatedQuestionList = writable([])
// let SpaceList = writable([])
// let FavoriteHash = writable([])
// let BlogList = writable([])
// let HotQues= writable([])
// let PublicQuesData = writable([])


//     async function FetchData() {
//             // USERDATA
//         let UserdataLoading =true
//         let QuestiondataLoading =false
//         let RelatedQuestionListLoading=false
//         let SpaceListLoading=false 
//         let FavoriteHashLoading=false 
//         let BlogListLoading=false
//         let HotQuesLoading=false
//         let PublicQuesDataLoading=false

//         // // UserData
//         // let response = await fetch("http://localhost:8080/raka",{mode:"cors"} )
//         // let j={}
//         // if (response.ok) {
//         //     j = await response.json();
//         //     // console.log("🚀 ~ file: index.svelte ~ line 25 ~ name ~ Userdata : ", j)
//         //     UserData.update((n)=>n=j)
//         //     UserdataLoading=true 
            
//         // } else {
//         //     console.log("HTTP-Error: " + response.status);
//         // }
//         // QUESTIONDATA
//         // let response = await fetch("http://localhost:8080/q/raka",{mode:"cors"} )
//         // let j={}
//         // if (response.ok) { 
//         //     j = await response.json();
//         //     // console.log("🚀 ~ file: index.svelte ~ line 25 ~ name ~ questionData : ", j)
//         //     QuestionData.update((n)=>n=j)
//         //     QuestiondataLoading=true
//         // } else {
//         //     console.log("HTTP-Error: " + response.status);
//         // }

//         // related Question list
//     //     response = await fetch("http://localhost:8080/rq/raka",{mode:"cors"} )
//     //     let k
//     // if (response.ok) { 
//     //     k = await response.json();
//     //     // console.log("🚀 ~ file: store.ts ~ line 45 ~ FetchData ~ RelatedQuestionList : ", k)
//     //     RelatedQuestionList.update((n)=>n=k)
//     //     RelatedQuestionListLoading=true
//     // } else {
//     //     console.log("HTTP-Error: " + response.status);
//     // }

//     // // SPACE LIST
//     // response = await fetch("http://localhost:8080/raka/s",{mode:"cors"} )
//     //     let l
//     //     if (response.ok) { 
//     //         l = await response.json();
//     //         // console.log("🚀 ~ file: store.ts ~ line 45 ~ FetchData ~ SpaceList : ", l)
//     //         SpaceList.update((n)=>n=l)
//     //         SpaceListLoading=true
//     //     } else {
//     //         console.log("HTTP-Error: " + response.status);
//     //     }

//     //     // FAVOURITE HASH
//     // response = await fetch("http://localhost:8080/raka/fh",{mode:"cors"} )
//     // let m
//     // if (response.ok) { 
//     //     m = await response.json();
//     //     // console.log("🚀 ~ file: store.ts ~ line 45 ~ FetchData ~ FavoriteHash : ", m)
//     //     FavoriteHash.update((n)=>n=m)
//     //     FavoriteHashLoading=true
//     // } else {
//     //     console.log("HTTP-Error: " + response.status);
//     // }

//         // // FAVOURITE HASH
//         // response = await fetch("http://localhost:8080/q/blogs",{mode:"cors"} )
//         // let n1
//         // if (response.ok) { 
//         //     n1 = await response.json();
//         //     // console.log("🚀 ~ file: store.ts ~ line 45 ~ FetchData ~ Blogs : ", n1)
//         //     BlogList.update((n)=>n=n1)
//         //     BlogListLoading=true
//         // } else {
//         //     console.log("HTTP-Error: " + response.status);
//         // }
        
//         // // Hot Ques
//         // response = await fetch("http://localhost:8080/q/hot",{mode:"cors"} )
//         // let n2
//         // if (response.ok) { 
//         //     n2 = await response.json();
//         //     // console.log("🚀 ~ file: store.ts ~ line 45 ~ FetchData ~ Hot Questions : ", n2)
//         //     HotQues.update((n)=>n=n2)
//         //     HotQuesLoading=true
//         // } else {
//         //     console.log("HTTP-Error: " + response.status);
//         // }
//          // PublicQuesData
//          response = await fetch("http://localhost:8080/q",{mode:"cors"} )
//          let n3
//          if (response.ok) { 
//              n3 = await response.json();
//             //  console.log("🚀 ~ file: store.ts ~ line 45 ~ FetchData ~ Hot Questions : ", n3)363
//              PublicQuesData.update((n)=>n=n3)
//              PublicQuesDataLoading=true
//          } else {
//              console.log("HTTP-Error: " + response.status);
//          }

//         // CHECK IF EVERYTHING IS OK
//         if (QuestiondataLoading && UserdataLoading && RelatedQuestionListLoading && SpaceListLoading && FavoriteHashLoading && BlogListLoading && HotQuesLoading && PublicQuesDataLoading){
//             Loading.update((n)=>n=true)
//         }

        

//     }
//     FetchData()




// // import { storeFunction } from 'svelte/store'

// // create a store with a default value
// // const dataStore = storeFunction(x)

// // make the store available
// // outside of this file
// // export default dataStore

// // USER DATA
// const SampleUserData :UserDataType =  {
//     ID : "132143",
//     UserID: 'skraka',
//     UserName: 'Sk Shahriar Ahmed Raka',
//     Email: 'skshahra@gmail.com',
//     Password: '123456',
//     UserTitle : 'Full-Stack Web Developer',
//     UserImage:
//         'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
//     Badges: {
//         Reputation: 13452543,
//         Gold: 999,
//         Silver: 888,
//         Bronze: 777
//     },
//     Follower: ['RKA', 'SHAHRIAR	'],
//     Following: ['RKA', 'SHAHRIAR'],
//     // Badges map[string]int
//     Location: 'Dhaka, Bangladesh',
//     MembershipTime: '3 year 5 Month',
//     LastSeen: 'This Week',
//     Aboutme:
//         'A Curious Learner, Full-Stack Web Developer, Security Researcher\nHere are my skills and strengths:\n✓ Expert in Golang\n ✓ Expert in Fiber framework (using Golang) \n ✓ Expert in WebAssembly (using Golang)  Expert in Golang     ✓ Expert in Fiber framework (using Golang)    ✓ Expert in WebAssembly (using Golang) ✓ Expert in database design, development, optimization, and migration    (PostgreSQL, MySQL, MongoDB , Redis) ✓ Expert in ( Grpc, protocol buffer ) ✓ Experienced in ( WebSocket, Socket.io, WebRTC ) for real-time client and server applications ✓ Experienced in ( Svelte.js ) and some knowledge in ( TypeScript ) ✓ Good understanding of ( Docker, Bash, PowerShell, Git,    Nginx, Kubernetes )        Github : github.com/skshahriarahmedraka    Upwork : upwork.com/o/profiles/users/~0107ef3405bffbe57e    Linkedin : linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/  Telegram : t.me/shahriarraka ',
//     Mysite: 'www.shahriarraka.me',
//     Github: 'www.github.com/skshahriarahmedraka',
//     Twitter: 'www.twitter.com/shahriarraka7',
//     Linkedin: 'www.linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/',
//     // "TopTags"    : ["Go","Rust","Python","Svelte","PostgreSQL"],
//     TopTags: [
//         { Name: 'Go', Score: 12, NumberOfPost: 4 },
//         { Name: 'Rust', Score: 10, NumberOfPost: 6 },
//         { Name: 'Python', Score: 7, NumberOfPost: 4 },
//         { Name: 'Svelte', Score: 12, NumberOfPost: 4 },
//         { Name: 'PostgreSQL', Score: 12, NumberOfPost: 4 }
//     ],
//     SelectedPanel: 'Profile',
//     AccountType: 'regular'
// }
// SampleUserData
const UserData = writable({} as UserDataType)
// UserData.update(()=>SampleUserData)
  


// // async function FetchUserData() {
// //     let response = await fetch("http://localhost:8080/raka",{mode:"cors"} )
// //     if (response.ok) { // if HTTP-status is 200-299
// //         x = await response.json();
// //         // console.log("🚀 ~ file: index.svelte ~ line 25 ~ name ~ json : ", x)
// //         UserData.update((n)=>n=x)
// //     } else {
// //         console.log("HTTP-Error: " + response.status);
// //     }
// // }
// // FetchUserData()



// // QUESTION DATA

// const SampleQuestionData:QuestionDataType ={
//     "ID": "143",
//     "QuesTitle": "How to efficiently concatenate strings in go How to efficiently concatenate strings in go How to efficiently concatenate strings in goHow to efficiently concatenate strings in goHow to efficiently concatenate strings in go goHow to efficiently concatenate strings in goHow to",
   
//     "QuesEditedTime": new Date().toISOString(),
//     "QuesViewed": 577000,

//     "QuesUpvote": 4,
//     "QuesDownvote": 32,
//     "QuesBookmark": 23,

//     "QuesTags": [
//         "go",
//         "string",
//         "Concatination",
//         "go"
//     ],
//     "QuesGroup": ["golang", "webdev" ],
//     "QuesAskedBy":  "UserID",
//     "QuesAnsAccepted" : "answerIDNumber",
//     "QuesAskedTime" : new Date().toISOString(),
//     "QuesEditedBy": "154123",
    
//     "QuesDescription": "In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.    In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.    In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.   ",
//     "QuesComment": [
//         "82567384","887874","72643"
//     ],
  
// }
// const QuestionData=writable({} as QuestionDataType )
// QuestionData.update(()=>SampleQuestionData)
    // {
    //     "ID": 1,
    //     "Title": "How to efficiently concatenate strings in go How to efficiently concatenate strings in go How to efficiently concatenate strings in goHow to efficiently concatenate strings in goHow to efficiently concatenate strings in go goHow to efficiently concatenate strings in goHow to",
    //     "AskedTime": " Aug 7, 2011 at 2:14 ",
    //     "ModifiedTime": " Aug 7, 2011 at 2:14 ",
    //     "Viewed": 577000,
    
    //     "Upvote": 4,
    //     "Downvote": 32,
    //     "Bookmark": 23,
    
    //     "Tags": [
    //         "go",
    //         "string",
    //         "Concatination",
    //         "go"
    //     ],
    //     "AskedBy": { "UserID": 154123 },
    //     "EditedBy": { "UserID": 154123 },
    
    //     "Description": "In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.    In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.    In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.   ",
    //     "Comment": [
    //         {
    //             "ID": 1,
    //             "Upvote": 4,
    //             "Downvote": 1,
    //             "UserID": 154123,
    //             "UserName": "sk shahriar ahmed raka",
    //             "Time": " Aug 7, 2011 at 2:14",
    //             "Comment": "no Comment bossNote: This question and most answers seem to have been written before append() came into the language, which is a good solution for this."
    //         }
    //     ],
    //     "Answers": [
    //         {
    //             "ID": 1,
    //             "AnsweredTime": " Aug 7, 2011 at 2:14 ",
    //             "Modified": " Aug 7, 2011 at 2:14 ",
    
    //             "Upvote": 4,
    //             "Downvote": 32,
    //             "Bookmark": 23,
    //             "Accepted": true,
    
    //             "AnsweredBy": { "UserID": 154123 },
    //             "EditedBy": { "UserID": 154123 },
    
    
    //             "Description": "In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.    In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.    In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.   ",
    //             "Comment": [
    //                 {
    //                     "ID": 1,
    //                     "Upvote": 4,
    //                     "Downvote": 1,
    //                     "UserID": 154123,
    //                     "UserName": "sk shahriar ahmed raka",
    //                     "Time": " Aug 7, 2011 at 2:14",
    //                     "Comment": "no Comment bossNote: This question and most answers seem to have been written before append() came into the language, which is a good solution for this."
    //                 }
    //             ]
    //         }
    //     ]
    // }
// let QuestionData = writable({
//     "QuestionTitle":"How to efficiently concatenate strings in go How to efficiently concatenate strings in go How to efficiently concatenate strings in goHow to efficiently concatenate strings in goHow to efficiently concatenate strings in go goHow to efficiently concatenate strings in goHow to",
//     "QuestionAskedTime":"12 years, 5 months ago",
//     "QuestionModified" : "6 months ago",
//     "QuestionViewed": 577000,
//     "QuestionUpvote":4,
//     "QuestionDownvote":32,
//     "QuestionBookmark":23,
//     "QuestionTags":["go","string","Concatination","go","string","Concatination","go","string","Concatination","go","string","go","string","Concatination","Concatination"],
//     "QuestionAskedBy":"yujiitadori",
//     "QuestionAskedByName":"yuji itadori prince vageta son goku",
//     "QuestionAskedByBadge":[999,888,777],
//     "QuestionAskedTimeExact":"Apr 1, 2019 at 21:38",
//     "QuestionAskedByImage":"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg",
    
//     "QuestionEditedBy":"KingofRandom",
//     "QuestionEditedByName":"King of Random prince of all sayian",
//     "QuestionEditedByBadge":[666,555,444],
//     "QuestionEditedByImage":"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg",
//     "QuestionEditedTimeExact":"Jan 12, 2010 at 16:18",
//     "QuestionDescription":`
//     In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.
//     In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.
//     In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.
//     ` ,
//     "QuestionComment":[
//     ["Note: This question and most answers seem to have been written before append() came into the language, which is a good solution for this. It will perform fast like copy() but will grow the slice firs","thomasrutter","Aug 10, 2017 at 1:29"],
//     ["It doesn't just seem very inefficient it has a specific problem that every new non-CS hire we have ever gotten runs into in the first few weeks on the job. It's quadratic - O(n*n). Think of the number sequence: 1 + 2 + 3 + 4 + .... It's n*(n+1)/2","rob lucci","Nov 16, 2017 at 15:22"]
// ],
//     "Answers":[
//         // ["ans(string)","Upvote(int)","downvote(int)","accepted(bool)","Answer(string)","modifiedBy(string)","time(string)","time(string)","EditorImgaelink(string)","answeredImgaelink(string) "]
//         [`994New Way: From Go 1.10 there is a strings.Builder type, please take a look at this answer for more detail. Old Way: Use the bytes package. It has a Buffer type which implements io.Writer. package main import ( "bytes" "fmt" ) func main() { var buffer bytes.Buffer for i := 0; i < 1000; i++ { buffer.WriteString("a") } fmt.Println(buffer.String()) } This does it in O(n) time.`,432,123,true,"inancramsay","Garp","Dec 12, 2019 at 7:30","Jul 9, 2019 at 10:42",[12,99,999],[99,23,11],"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"],
//         [`n Go 1.10+ there is strings.Builder, here.A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use.`,4,43,false,"luffy","icza", "Dec 4, 2015 at 8:01","Oct 9, 2019 at 10:42",[11,9,33],[42,535,22],"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"],
//         [`n Go 1.10+ there is strings.Builder, here.A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use.`,6,3,false,"Dragon","icza", "Dec 4, 2015 at 8:01","Jan 9, 2019 at 10:42",[12,24,53],[99,23,355],"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"],
//         [`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,98,231,false,"monkey D. luffy","annya","May 9, 2019 at 10:42","",[13,354,666],[11,32,535],"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"],
//         [`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,98,1,true ,"Sabo Cheif of stuff","Yur Fozer","Aug 9, 2019 at 10:42","Dec 21, 2022 at 10:42",[66,2,5],[99,11,0],"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"],
//         [`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,9,1,false,"Garp Hero of Nevy","annya","Jun 9, 2019 at 10:42","",[23,53,64],[99,23,24],"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"],
//         [`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,8,31,true,"Akayinu Sakazuki","Spy","Aug 9, 2019 at 10:42","Feb 11, 2020 at 10:42",[76,37,34],[12,23,43],"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"],
//         [`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,9,2,false,"Muzan kibutsuji","liverpool","Apr 9, 2019 at 10:42","",[17,44,53],[12,23,53],"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"],
//         [`In addition to The Go Programming Language Specification, you should read Effective Go. In the section on maps, they say, amongst other things: An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool.f multiple assignment.`,9,31,false,"Fleet Admiral","rengoku","Mar 9, 2019 at 10:42","",[14,35,640],[63,78,9],"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg","https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png"],
//     ]
// })



// // async function FetchQuestionData() {
// //     let response = await fetch("http://localhost:8080/q/raka",{mode:"cors"} )
// //     if (response.ok) { // if HTTP-status is 200-299
// //         x2 = await response.json();
// //         console.log("🚀 ~ file: index.svelte ~ line 25 ~ name ~ json : ", x2)
// //         QuestionData.update((n)=>n=x2)
// //     } else {
// //         console.log("HTTP-Error: " + response.status);
// //     }
// // }
// // FetchQuestionData()

// // QuestionData.update(
// //     n=> n=x
// // )

// const RelatedQuestionList= writable()

// // // RelatedQuestionList.update(
// // //     n=>n=y
// // // )

// const SpaceList = writable([
//     ["Go programming","https://stackoverflow.com/"],
//     ["GTK","https://stackoverflow.com/"],
//     ["QT framework","https://stackoverflow.com/"],
//     ["Gui in Rust","https://stackoverflow.com/"],
//     ["Gopher","https://stackoverflow.com/"],
//     ["Rusty Boys","https://stackoverflow.com/"],
//     ["Google Cloud","https://stackoverflow.com/"],
//     ["Azure services","https://stackoverflow.com/"],
//     ["Backend Engineering","https://stackoverflow.com/"],
//     ["Web DEvelopment","https://stackoverflow.com/"],
// ])


// const FavoriteHash = writable([
//     ["go","https://github.com/","https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Google_%22G%22_Logo.svg/2048px-Google_%22G%22_Logo.svg.png"],
//     ["rust","https://github.com/","https://cdn-icons-png.flaticon.com/512/732/732221.png"],
//     ["fiber","https://github.com/","https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSXyTXn1zirGyXZd15ukEGTupWvTasppbf7N5dW3AT8bc7iUDwFxgI3aRyaGYS2NYQUO9c&usqp=CAU"],
//     ["gio","https://github.com/",""],
//     ["C++","https://github.com/","https://i.pinimg.com/736x/83/fe/22/83fe2289e31398d96e8e6ed9d603c72f.jpg"],
//     ["C#","https://github.com/","https://upload.wikimedia.org/wikipedia/commons/thumb/a/ad/Microsoft_Lists_%282020-present%29.svg/2048px-Microsoft_Lists_%282020-present%29.svg.png"],
//     ["dart","https://github.com/","https://seeklogo.com/images/M/microsoft-edge-new-2020-logo-2AD376EBAA-seeklogo.com.png"],
//     ["flutter","https://github.com/","https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSmZ-U_nmUBZ6SdJBuOjMOajQplOkMqYmYIVbPU14Co9i7jHKqo4kPpQEOkg_0lCE7_yXg&usqp=CAU"],
//     ["linux","https://github.com/","https://upload.wikimedia.org/wikipedia/commons/7/70/Firefox_Focus_2021_Icon.png"],
// ])

// const HotQues= writable([
//     // [QuestionTitle ,SvgIcon,QuesLink ]
//     ["How to efficiently concatenate strings in go" ,`<svg xmlns="http://www.w3.org/2000/svg"  xml:space="preserve" y="0" x="0" id="Layer_1" version="1.1" viewBox="-30.80817 -19.183575 267.00414 115.10145"><style id="style2490" type="text/css">.st4{fill:#00acd7}</style><g transform="translate(-24.7 -77.8657)" id="g2529"><g id="g2498"><g id="g2496"><g id="g2494"><path id="path2492" d="M40.2 101.1c-.4 0-.5-.2-.3-.5l2.1-2.7c.2-.3.7-.5 1.1-.5h35.7c.4 0 .5.3.3.6l-1.7 2.6c-.2.3-.7.6-1 .6z" class="st4"/></g></g></g><g id="g2506"><g id="g2504"><g id="g2502"><path id="path2500" d="M25.1 110.3c-.4 0-.5-.2-.3-.5l2.1-2.7c.2-.3.7-.5 1.1-.5h45.6c.4 0 .6.3.5.6l-.8 2.4c-.1.4-.5.6-.9.6z" class="st4"/></g></g></g><g id="g2514"><g id="g2512"><g id="g2510"><path id="path2508" d="M49.3 119.5c-.4 0-.5-.3-.3-.6l1.4-2.5c.2-.3.6-.6 1-.6h20c.4 0 .6.3.6.7l-.2 2.4c0 .4-.4.7-.7.7z" class="st4"/></g></g></g><g id="g2527"><g id="CXHf1q_3_"><g id="g2524"><g id="g2518"><path id="path2516" d="M153.1 99.3c-6.3 1.6-10.6 2.8-16.8 4.4-1.5.4-1.6.5-2.9-1-1.5-1.7-2.6-2.8-4.7-3.8-6.3-3.1-12.4-2.2-18.1 1.5-6.8 4.4-10.3 10.9-10.2 19 .1 8 5.6 14.6 13.5 15.7 6.8.9 12.5-1.5 17-6.6.9-1.1 1.7-2.3 2.7-3.7h-19.3c-2.1 0-2.6-1.3-1.9-3 1.3-3.1 3.7-8.3 5.1-10.9.3-.6 1-1.6 2.5-1.6h36.4c-.2 2.7-.2 5.4-.6 8.1-1.1 7.2-3.8 13.8-8.2 19.6-7.2 9.5-16.6 15.4-28.5 17-9.8 1.3-18.9-.6-26.9-6.6-7.4-5.6-11.6-13-12.7-22.2-1.3-10.9 1.9-20.7 8.5-29.3 7.1-9.3 16.5-15.2 28-17.3 9.4-1.7 18.4-.6 26.5 4.9 5.3 3.5 9.1 8.3 11.6 14.1.6.9.2 1.4-1 1.7z" class="st4"/></g><g id="g2522"><path id="path2520" d="M186.2 154.6c-9.1-.2-17.4-2.8-24.4-8.8-5.9-5.1-9.6-11.6-10.8-19.3-1.8-11.3 1.3-21.3 8.1-30.2 7.3-9.6 16.1-14.6 28-16.7 10.2-1.8 19.8-.8 28.5 5.1 7.9 5.4 12.8 12.7 14.1 22.3 1.7 13.5-2.2 24.5-11.5 33.9-6.6 6.7-14.7 10.9-24 12.8-2.7.5-5.4.6-8 .9zm23.8-40.4c-.1-1.3-.1-2.3-.3-3.3-1.8-9.9-10.9-15.5-20.4-13.3-9.3 2.1-15.3 8-17.5 17.4-1.8 7.8 2 15.7 9.2 18.9 5.5 2.4 11 2.1 16.3-.6 7.9-4.1 12.2-10.5 12.7-19.1z" class="st4"/></g></g></g></g></g></svg>`,"https://stackoverflow.com/"],
//     // ["How to print struct variables in console?" ,`<svg version="1.1"  xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"> <g id="logo" transform="translate(53, 53)"> <path id="r" transform="translate(0.5, 0.5)" stroke="black" stroke-width="1" stroke-linejoin="round" d="     M -9,-15 H 4 C 12,-15 12,-7 4,-7 H -9 Z     M -40,22 H 0 V 11 H -9 V 3 H 1 C 12,3 6,22 15,22 H 40     V 3 H 34 V 5 C 34,13 25,12 24,7 C 23,2 19,-2 18,-2 C 33,-10 24,-26 12,-26 H -35     V -15 H -25 V 11 H -40 Z"/> <g id="gear" mask="url(#holes)"> <circle r="43" fill="none" stroke="black" stroke-width="9"/> <g id="cogs"> <polygon id="cog" stroke="black" stroke-width="3" stroke-linejoin="round" points="46,3 51,0 46,-3"/> <use xlink:href="#cog" transform="rotate(11.25)"/> <use xlink:href="#cog" transform="rotate(22.50)"/> <use xlink:href="#cog" transform="rotate(33.75)"/> <use xlink:href="#cog" transform="rotate(45.00)"/> <use xlink:href="#cog" transform="rotate(56.25)"/> <use xlink:href="#cog" transform="rotate(67.50)"/> <use xlink:href="#cog" transform="rotate(78.75)"/> <use xlink:href="#cog" transform="rotate(90.00)"/> <use xlink:href="#cog" transform="rotate(101.25)"/> <use xlink:href="#cog" transform="rotate(112.50)"/> <use xlink:href="#cog" transform="rotate(123.75)"/> <use xlink:href="#cog" transform="rotate(135.00)"/> <use xlink:href="#cog" transform="rotate(146.25)"/> <use xlink:href="#cog" transform="rotate(157.50)"/> <use xlink:href="#cog" transform="rotate(168.75)"/> <use xlink:href="#cog" transform="rotate(180.00)"/> <use xlink:href="#cog" transform="rotate(191.25)"/> <use xlink:href="#cog" transform="rotate(202.50)"/> <use xlink:href="#cog" transform="rotate(213.75)"/> <use xlink:href="#cog" transform="rotate(225.00)"/> <use xlink:href="#cog" transform="rotate(236.25)"/> <use xlink:href="#cog" transform="rotate(247.50)"/> <use xlink:href="#cog" transform="rotate(258.75)"/> <use xlink:href="#cog" transform="rotate(270.00)"/> <use xlink:href="#cog" transform="rotate(281.25)"/> <use xlink:href="#cog" transform="rotate(292.50)"/> <use xlink:href="#cog" transform="rotate(303.75)"/> <use xlink:href="#cog" transform="rotate(315.00)"/> <use xlink:href="#cog" transform="rotate(326.25)"/> <use xlink:href="#cog" transform="rotate(337.50)"/> <use xlink:href="#cog" transform="rotate(348.75)"/> </g> <g id="mounts"> <polygon id="mount" stroke="black" stroke-width="6" stroke-linejoin="round" points="-7,-42 0,-35 7,-42"/> <use xlink:href="#mount" transform="rotate(72)"/> <use xlink:href="#mount" transform="rotate(144)"/> <use xlink:href="#mount" transform="rotate(216)"/> <use xlink:href="#mount" transform="rotate(288)"/> </g> </g> <mask id="holes"> <rect x="-60" y="-60" width="120" height="120" fill="white"/> <circle id="hole" cy="-40" r="3"/> <use xlink:href="#hole" transform="rotate(72)"/> <use xlink:href="#hole" transform="rotate(144)"/> <use xlink:href="#hole" transform="rotate(216)"/> <use xlink:href="#hole" transform="rotate(288)"/> </mask> </g> </svg>`,"https://stackoverflow.com/"],
//     // [" What should be the values of GOPATH and GOROOT? " ,`<svg version="1.1"  xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"> <g id="logo" transform="translate(53, 53)"> <path id="r" transform="translate(0.5, 0.5)" stroke="black" stroke-width="1" stroke-linejoin="round" d="     M -9,-15 H 4 C 12,-15 12,-7 4,-7 H -9 Z     M -40,22 H 0 V 11 H -9 V 3 H 1 C 12,3 6,22 15,22 H 40     V 3 H 34 V 5 C 34,13 25,12 24,7 C 23,2 19,-2 18,-2 C 33,-10 24,-26 12,-26 H -35     V -15 H -25 V 11 H -40 Z"/> <g id="gear" mask="url(#holes)"> <circle r="43" fill="none" stroke="black" stroke-width="9"/> <g id="cogs"> <polygon id="cog" stroke="black" stroke-width="3" stroke-linejoin="round" points="46,3 51,0 46,-3"/> <use xlink:href="#cog" transform="rotate(11.25)"/> <use xlink:href="#cog" transform="rotate(22.50)"/> <use xlink:href="#cog" transform="rotate(33.75)"/> <use xlink:href="#cog" transform="rotate(45.00)"/> <use xlink:href="#cog" transform="rotate(56.25)"/> <use xlink:href="#cog" transform="rotate(67.50)"/> <use xlink:href="#cog" transform="rotate(78.75)"/> <use xlink:href="#cog" transform="rotate(90.00)"/> <use xlink:href="#cog" transform="rotate(101.25)"/> <use xlink:href="#cog" transform="rotate(112.50)"/> <use xlink:href="#cog" transform="rotate(123.75)"/> <use xlink:href="#cog" transform="rotate(135.00)"/> <use xlink:href="#cog" transform="rotate(146.25)"/> <use xlink:href="#cog" transform="rotate(157.50)"/> <use xlink:href="#cog" transform="rotate(168.75)"/> <use xlink:href="#cog" transform="rotate(180.00)"/> <use xlink:href="#cog" transform="rotate(191.25)"/> <use xlink:href="#cog" transform="rotate(202.50)"/> <use xlink:href="#cog" transform="rotate(213.75)"/> <use xlink:href="#cog" transform="rotate(225.00)"/> <use xlink:href="#cog" transform="rotate(236.25)"/> <use xlink:href="#cog" transform="rotate(247.50)"/> <use xlink:href="#cog" transform="rotate(258.75)"/> <use xlink:href="#cog" transform="rotate(270.00)"/> <use xlink:href="#cog" transform="rotate(281.25)"/> <use xlink:href="#cog" transform="rotate(292.50)"/> <use xlink:href="#cog" transform="rotate(303.75)"/> <use xlink:href="#cog" transform="rotate(315.00)"/> <use xlink:href="#cog" transform="rotate(326.25)"/> <use xlink:href="#cog" transform="rotate(337.50)"/> <use xlink:href="#cog" transform="rotate(348.75)"/> </g> <g id="mounts"> <polygon id="mount" stroke="black" stroke-width="6" stroke-linejoin="round" points="-7,-42 0,-35 7,-42"/> <use xlink:href="#mount" transform="rotate(72)"/> <use xlink:href="#mount" transform="rotate(144)"/> <use xlink:href="#mount" transform="rotate(216)"/> <use xlink:href="#mount" transform="rotate(288)"/> </g> </g> <mask id="holes"> <rect x="-60" y="-60" width="120" height="120" fill="white"/> <circle id="hole" cy="-40" r="3"/> <use xlink:href="#hole" transform="rotate(72)"/> <use xlink:href="#hole" transform="rotate(144)"/> <use xlink:href="#hole" transform="rotate(216)"/> <use xlink:href="#hole" transform="rotate(288)"/> </mask> </g> </svg>`,"https://stackoverflow.com/"],
//     // ["How do I SET the GOPATH environment variable on Ubuntu? What file must I edit?" ,`<svg xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:cc="http://web.resource.org/cc/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:svg="http://www.w3.org/2000/svg" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd" xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape" width="110.4211" height="109.8461" id="svg2169" sodipodi:version="0.32" inkscape:version="0.45.1" version="1.0" sodipodi:docbase="/home/bene/Desktop" sodipodi:docname="dessin-1.svg" inkscape:output_extension="org.inkscape.output.svg.inkscape"> <defs id="defs2171"> <linearGradient id="linearGradient11301" inkscape:collect="always"> <stop id="stop11303" offset="0" style="stop-color:#ffe052;stop-opacity:1" /> <stop id="stop11305" offset="1" style="stop-color:#ffc331;stop-opacity:1" /> </linearGradient> <linearGradient gradientUnits="userSpaceOnUse" y2="168.1012" x2="147.77737" y1="111.92053" x1="89.136749" id="linearGradient11307" xlink:href="#linearGradient11301" inkscape:collect="always" /> <linearGradient id="linearGradient9515" inkscape:collect="always"> <stop id="stop9517" offset="0" style="stop-color:#387eb8;stop-opacity:1" /> <stop id="stop9519" offset="1" style="stop-color:#366994;stop-opacity:1" /> </linearGradient> <linearGradient gradientUnits="userSpaceOnUse" y2="131.85291" x2="110.14919" y1="77.070274" x1="55.549179" id="linearGradient9521" xlink:href="#linearGradient9515" inkscape:collect="always" /> </defs> <sodipodi:namedview id="base" pagecolor="#ffffff" bordercolor="#666666" borderopacity="1.0" inkscape:pageopacity="0.0" inkscape:pageshadow="2" inkscape:zoom="0.24748737" inkscape:cx="-260.46312" inkscape:cy="316.02744" inkscape:document-units="px" inkscape:current-layer="layer1" width="131.10236px" height="184.25197px" inkscape:window-width="872" inkscape:window-height="624" inkscape:window-x="5" inkscape:window-y="48" /> <metadata id="metadata2174"> <rdf:RDF> <cc:Work rdf:about=""> <dc:format>image/svg+xml</dc:format> <dc:type rdf:resource="http://purl.org/dc/dcmitype/StillImage" /> </cc:Work> </rdf:RDF> </metadata> <g inkscape:label="Calque 1" inkscape:groupmode="layer" id="layer1" transform="translate(-473.36088,-251.72485)"> <g id="g1894" transform="translate(428.42338,184.2561)"> <path style="opacity:1;color:#000000;fill:url(#linearGradient9521);fill-opacity:1;fill-rule:nonzero;stroke:none;stroke-width:1;stroke-linecap:butt;stroke-linejoin:miter;marker:none;marker-start:none;marker-mid:none;marker-end:none;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1;visibility:visible;display:inline;overflow:visible" d="M 99.75,67.46875 C 71.718268,67.468752 73.46875,79.625 73.46875,79.625 L 73.5,92.21875 L 100.25,92.21875 L 100.25,96 L 62.875,96 C 62.875,96 44.9375,93.965724 44.9375,122.25 C 44.937498,150.53427 60.59375,149.53125 60.59375,149.53125 L 69.9375,149.53125 L 69.9375,136.40625 C 69.9375,136.40625 69.433848,120.75 85.34375,120.75 C 101.25365,120.75 111.875,120.75 111.875,120.75 C 111.875,120.75 126.78125,120.99096 126.78125,106.34375 C 126.78125,91.696544 126.78125,82.125 126.78125,82.125 C 126.78125,82.124998 129.04443,67.46875 99.75,67.46875 z M 85,75.9375 C 87.661429,75.937498 89.8125,78.088571 89.8125,80.75 C 89.812502,83.411429 87.661429,85.5625 85,85.5625 C 82.338571,85.562502 80.1875,83.411429 80.1875,80.75 C 80.187498,78.088571 82.338571,75.9375 85,75.9375 z " id="path8615" /> <path id="path8620" d="M 100.5461,177.31485 C 128.57784,177.31485 126.82735,165.1586 126.82735,165.1586 L 126.7961,152.56485 L 100.0461,152.56485 L 100.0461,148.7836 L 137.4211,148.7836 C 137.4211,148.7836 155.3586,150.81787 155.3586,122.53359 C 155.35861,94.249323 139.70235,95.252343 139.70235,95.252343 L 130.3586,95.252343 L 130.3586,108.37734 C 130.3586,108.37734 130.86226,124.03359 114.95235,124.03359 C 99.042448,124.03359 88.421098,124.03359 88.421098,124.03359 C 88.421098,124.03359 73.514848,123.79263 73.514848,138.43985 C 73.514848,153.08705 73.514848,162.6586 73.514848,162.6586 C 73.514848,162.6586 71.251668,177.31485 100.5461,177.31485 z M 115.2961,168.8461 C 112.63467,168.8461 110.4836,166.69503 110.4836,164.0336 C 110.4836,161.37217 112.63467,159.2211 115.2961,159.2211 C 117.95753,159.2211 120.1086,161.37217 120.1086,164.0336 C 120.10861,166.69503 117.95753,168.8461 115.2961,168.8461 z " style="opacity:1;color:#000000;fill:url(#linearGradient11307);fill-opacity:1;fill-rule:nonzero;stroke:none;stroke-width:1;stroke-linecap:butt;stroke-linejoin:miter;marker:none;marker-start:none;marker-mid:none;marker-end:none;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1;visibility:visible;display:inline;overflow:visible" /> </g> </g> </svg>`,"https://stackoverflow.com/"],
//     [" How can I convert a zero-terminated byte array to string? " ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" ><path d="M15.9.476a2.14 2.14 0 0 0-.823.218L3.932 6.01c-.582.277-1.005.804-1.15 1.432L.054 19.373c-.13.56-.025 1.147.3 1.627q.057.087.12.168l7.7 9.574c.407.5 1.018.787 1.662.784h12.35c.646.001 1.258-.3 1.664-.793l7.696-9.576c.404-.5.555-1.16.4-1.786L29.2 7.43c-.145-.628-.57-1.155-1.15-1.432L16.923.695A2.14 2.14 0 0 0 15.89.476z" fill="#326ce5"/><path d="M16.002 4.542c-.384.027-.675.356-.655.74v.188c.018.213.05.424.092.633a6.22 6.22 0 0 1 .066 1.21c-.038.133-.114.253-.218.345l-.015.282c-.405.034-.807.096-1.203.186-1.666.376-3.183 1.24-4.354 2.485l-.24-.17c-.132.04-.274.025-.395-.04a6.22 6.22 0 0 1-.897-.81 5.55 5.55 0 0 0-.437-.465l-.148-.118c-.132-.106-.294-.167-.463-.175a.64.64 0 0 0-.531.236c-.226.317-.152.756.164.983l.138.11a5.55 5.55 0 0 0 .552.323c.354.197.688.428.998.7a.74.74 0 0 1 .133.384l.218.2c-1.177 1.766-1.66 3.905-1.358 6.006l-.28.08c-.073.116-.17.215-.286.288a6.22 6.22 0 0 1-1.194.197 5.57 5.57 0 0 0-.64.05l-.177.04h-.02a.67.67 0 0 0-.387 1.132.67.67 0 0 0 .684.165h.013l.18-.02c.203-.06.403-.134.598-.218.375-.15.764-.265 1.162-.34.138.008.27.055.382.135l.3-.05c.65 2.017 2.016 3.726 3.84 4.803l-.122.255c.056.117.077.247.06.376-.165.382-.367.748-.603 1.092a5.58 5.58 0 0 0-.358.533l-.085.18a.67.67 0 0 0 .65 1.001.67.67 0 0 0 .553-.432l.083-.17c.076-.2.14-.404.192-.61.177-.437.273-.906.515-1.196a.54.54 0 0 1 .286-.14l.15-.273a8.62 8.62 0 0 0 6.146.015l.133.255c.136.02.258.095.34.205.188.358.34.733.456 1.12a5.57 5.57 0 0 0 .194.611l.083.17a.67.67 0 0 0 1.187.131.67.67 0 0 0 .016-.701l-.087-.18a5.55 5.55 0 0 0-.358-.531c-.23-.332-.428-.686-.6-1.057a.52.52 0 0 1 .068-.4 2.29 2.29 0 0 1-.111-.269c1.82-1.085 3.18-2.8 3.823-4.82l.284.05c.102-.093.236-.142.373-.138.397.076.786.2 1.162.34.195.09.395.166.598.23.048.013.118.024.172.037h.013a.67.67 0 0 0 .841-.851.67.67 0 0 0-.544-.446l-.194-.046a5.57 5.57 0 0 0-.64-.05c-.404-.026-.804-.092-1.194-.197-.12-.067-.22-.167-.288-.288l-.27-.08a8.65 8.65 0 0 0-1.386-5.993l.236-.218c-.01-.137.035-.273.124-.378.307-.264.64-.497.99-.696a5.57 5.57 0 0 0 .552-.323l.146-.118a.67.67 0 0 0-.133-1.202.67.67 0 0 0-.696.161l-.148.118a5.57 5.57 0 0 0-.437.465c-.264.302-.556.577-.873.823a.74.74 0 0 1-.404.044l-.253.18c-1.46-1.53-3.427-2.48-5.535-2.67 0-.1-.013-.25-.015-.297-.113-.078-.192-.197-.218-.332a6.23 6.23 0 0 1 .076-1.207c.043-.21.073-.42.092-.633v-.2c.02-.384-.27-.713-.655-.74zm-.834 5.166l-.2 3.493h-.015c-.01.216-.137.4-.332.504s-.426.073-.6-.054l-2.865-2.03a6.86 6.86 0 0 1 3.303-1.799c.234-.05.47-.088.707-.114zm1.668 0c1.505.187 2.906.863 3.99 1.924l-2.838 2.017c-.175.14-.415.168-.618.072s-.333-.3-.336-.524zm-6.72 3.227l2.62 2.338v.015c.163.142.234.363.186.574s-.21.378-.417.435v.01l-3.362.967a6.86 6.86 0 0 1 .974-4.34zm11.753 0c.796 1.295 1.148 2.814 1.002 4.327l-3.367-.97v-.013c-.21-.057-.37-.224-.417-.435s.023-.43.186-.574l2.6-2.327zm-6.404 2.52h1.072l.655.832-.238 1.04-.963.463-.965-.463-.227-1.04zm3.434 2.838c.045-.005.1-.005.135 0l3.467.585c-.5 1.44-1.487 2.67-2.775 3.493l-1.34-3.244a.59.59 0 0 1 .509-.819zm-5.823.015c.196.003.377.104.484.268s.124.37.047.55v.013l-1.332 3.218C11 21.54 10.032 20.325 9.517 18.9l3.437-.583c.038-.004.077-.004.116 0zm2.904 1.4a.59.59 0 0 1 .537.308h.013l1.694 3.057-.677.2c-1.246.285-2.547.218-3.758-.194l1.7-3.057c.103-.18.293-.29.5-.295z" fill="#fff" stroke="#fff" stroke-width=".055"/></svg>`,"https://stackoverflow.com/"],
//     [" go get results in 'terminal prompts disabled' error for github private repo " ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_vue</title><path d="M24.4,3.925H30L16,28.075,2,3.925H12.71L16,9.525l3.22-5.6Z" style="fill:#41b883"/><path d="M2,3.925l14,24.15L30,3.925H24.4L16,18.415,7.53,3.925Z" style="fill:#41b883"/><path d="M7.53,3.925,16,18.485l8.4-14.56H19.22L16,9.525l-3.29-5.6Z" style="fill:#35495e"/></svg>`,"https://stackoverflow.com/"],
//     ["s it possible to capture a Ctrl+C signal (SIGINT) and run a cleanup function, in a \"defer\" fashion?" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_svelte</title><path d="M26.47,5.7A8.973,8.973,0,0,0,14.677,3.246L7.96,7.4a7.461,7.461,0,0,0-3.481,5.009,7.686,7.686,0,0,0,.8,5.058,7.358,7.358,0,0,0-1.151,2.8,7.789,7.789,0,0,0,1.4,6.028,8.977,8.977,0,0,0,11.794,2.458L24.04,24.6a7.468,7.468,0,0,0,3.481-5.009,7.673,7.673,0,0,0-.8-5.062,7.348,7.348,0,0,0,1.152-2.8A7.785,7.785,0,0,0,26.47,5.7" style="fill:#ff3e00"/><path d="M14.022,26.64A5.413,5.413,0,0,1,8.3,24.581a4.678,4.678,0,0,1-.848-3.625,4.307,4.307,0,0,1,.159-.61l.127-.375.344.238a8.76,8.76,0,0,0,2.628,1.274l.245.073-.025.237a1.441,1.441,0,0,0,.271.968,1.63,1.63,0,0,0,1.743.636,1.512,1.512,0,0,0,.411-.175l6.7-4.154a1.366,1.366,0,0,0,.633-.909,1.407,1.407,0,0,0-.244-1.091,1.634,1.634,0,0,0-1.726-.622,1.509,1.509,0,0,0-.413.176l-2.572,1.584a4.934,4.934,0,0,1-1.364.582,5.415,5.415,0,0,1-5.727-2.06A4.678,4.678,0,0,1,7.811,13.1,4.507,4.507,0,0,1,9.9,10.09l6.708-4.154a4.932,4.932,0,0,1,1.364-.581A5.413,5.413,0,0,1,23.7,7.414a4.679,4.679,0,0,1,.848,3.625,4.272,4.272,0,0,1-.159.61l-.127.375-.344-.237a8.713,8.713,0,0,0-2.628-1.274l-.245-.074.025-.237a1.438,1.438,0,0,0-.272-.968,1.629,1.629,0,0,0-1.725-.622,1.484,1.484,0,0,0-.411.176l-6.722,4.14a1.353,1.353,0,0,0-.631.908,1.394,1.394,0,0,0,.244,1.092,1.634,1.634,0,0,0,1.726.621,1.538,1.538,0,0,0,.413-.175l2.562-1.585a4.9,4.9,0,0,1,1.364-.581,5.417,5.417,0,0,1,5.728,2.059,4.681,4.681,0,0,1,.843,3.625A4.5,4.5,0,0,1,22.1,21.905l-6.707,4.154a4.9,4.9,0,0,1-1.364.581" style="fill:#fff"/></svg>`,"https://stackoverflow.com/"],
//     ["What is the shortest way to simply sort an array of structs by (arbitrary) field names?" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_nginx</title><path d="M15.948,2h.065a10.418,10.418,0,0,1,.972.528Q22.414,5.65,27.843,8.774a.792.792,0,0,1,.414.788c-.008,4.389,0,8.777-.005,13.164a.813.813,0,0,1-.356.507q-5.773,3.324-11.547,6.644a.587.587,0,0,1-.657.037Q9.912,26.6,4.143,23.274a.7.7,0,0,1-.4-.666q0-6.582,0-13.163a.693.693,0,0,1,.387-.67Q9.552,5.657,14.974,2.535c.322-.184.638-.379.974-.535" style="fill:#019639"/><path d="M8.767,10.538q0,5.429,0,10.859a1.509,1.509,0,0,0,.427,1.087,1.647,1.647,0,0,0,2.06.206,1.564,1.564,0,0,0,.685-1.293c0-2.62-.005-5.24,0-7.86q3.583,4.29,7.181,8.568a2.833,2.833,0,0,0,2.6.782,1.561,1.561,0,0,0,1.251-1.371q.008-5.541,0-11.081a1.582,1.582,0,0,0-3.152,0c0,2.662-.016,5.321,0,7.982-2.346-2.766-4.663-5.556-7-8.332A2.817,2.817,0,0,0,10.17,9.033,1.579,1.579,0,0,0,8.767,10.538Z" style="fill:#fff"/></svg>`,"https://stackoverflow.com/"],
//     ["exec: \"gcc\": executable file not found in %PATH% when trying go build" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_html</title><polygon points="5.902 27.201 3.655 2 28.345 2 26.095 27.197 15.985 30 5.902 27.201" style="fill:#e44f26"/><polygon points="16 27.858 24.17 25.593 26.092 4.061 16 4.061 16 27.858" style="fill:#f1662a"/><polygon points="16 13.407 11.91 13.407 11.628 10.242 16 10.242 16 7.151 15.989 7.151 8.25 7.151 8.324 7.981 9.083 16.498 16 16.498 16 13.407" style="fill:#ebebeb"/><polygon points="16 21.434 15.986 21.438 12.544 20.509 12.324 18.044 10.651 18.044 9.221 18.044 9.654 22.896 15.986 24.654 16 24.65 16 21.434" style="fill:#ebebeb"/><polygon points="15.989 13.407 15.989 16.498 19.795 16.498 19.437 20.507 15.989 21.437 15.989 24.653 22.326 22.896 22.372 22.374 23.098 14.237 23.174 13.407 22.341 13.407 15.989 13.407" style="fill:#fff"/><polygon points="15.989 7.151 15.989 9.071 15.989 10.235 15.989 10.242 23.445 10.242 23.445 10.242 23.455 10.242 23.517 9.548 23.658 7.981 23.732 7.151 15.989 7.151" style="fill:#fff"/></svg>`,"https://stackoverflow.com/"],
//     ["Go install fails with error: no install location for directory xxx outside GOPATH" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_flutter</title><polyline points="15.383 18.316 18.744 15.042 27.093 15.042 19.697 22.438 15.383 18.316 15.383 18.316 15.383 18.316 15.383 18.316 15.383 18.316" style="fill:#40d0fd"/><polygon points="4.907 16.125 9.106 20.424 27.093 2.287 18.744 2.287 4.907 16.125" style="fill:#41d0fd;isolation:isolate"/><polygon points="11.176 22.479 15.435 26.675 19.697 22.438 15.383 18.316 11.176 22.479" style="fill:#1fbcfd"/><polygon points="15.435 26.675 19.697 22.438 26.989 29.813 18.593 29.813 15.435 26.675" style="fill:#095a9d"/><polygon points="15.435 26.675 19.406 25.354 18.068 24.057 15.435 26.675" style="fill:#0e5199"/></svg>`,"https://stackoverflow.com/"],
//     ["Function declaration syntax: things in parenthesis before function name" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_docker</title><path d="M18.191,13.071H20.7v2.566H21.97a5.5,5.5,0,0,0,1.744-.292,4.462,4.462,0,0,0,.848-.383,3.149,3.149,0,0,1-.589-1.623,3.427,3.427,0,0,1,.616-2.416l.264-.305.314.253a4,4,0,0,1,1.575,2.538,3.837,3.837,0,0,1,2.913.271l.345.2-.181.354a3.629,3.629,0,0,1-3.648,1.74c-2.173,5.413-6.9,7.976-12.642,7.976A7.958,7.958,0,0,1,6.3,20.211l-.025-.043-.226-.459a7.28,7.28,0,0,1-.579-3.693l.035-.38H7.648V13.071h2.51v-2.51h5.02V8.051h3.012v5.02Z" style="fill:#3a4e55"/><path d="M26.324,14.021A3.311,3.311,0,0,0,24.906,11.2a3.072,3.072,0,0,0,.289,3.821,5.279,5.279,0,0,1-3.225,1.037H5.883a6.779,6.779,0,0,0,.667,3.737l.183.335a6.2,6.2,0,0,0,.379.569h0q.992.064,1.829.045h0a8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.4,8.4,0,0,1-1.845.3c.044,0-.046.007-.046.007l-.082.007c-.291.016-.6.02-.925.02-.351,0-.7-.007-1.083-.026l-.01.007a7.882,7.882,0,0,0,6.063,2.41c5.56,0,10.276-2.465,12.365-8,1.482.152,2.906-.226,3.553-1.49a3.5,3.5,0,0,0-3.122-.022" style="fill:#00aada"/><path d="M26.324,14.021A3.311,3.311,0,0,0,24.906,11.2a3.072,3.072,0,0,0,.289,3.821,5.279,5.279,0,0,1-3.225,1.037H6.836a5.223,5.223,0,0,0,2.106,4.686h0a8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.83,8.83,0,0,1-1.894.314L9.543,21.1c1.892.971,4.636.967,7.782-.241a21.868,21.868,0,0,0,9.1-6.889l-.1.048" style="fill:#27b9ec"/><path d="M5.913,17.732a6.431,6.431,0,0,0,.637,2.061l.183.335a6.2,6.2,0,0,0,.379.569q.992.064,1.829.045a8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.826,8.826,0,0,1-1.891.307l-.1,0c-.291.016-.6.026-.922.026-.351,0-.709-.007-1.1-.026a7.913,7.913,0,0,0,6.076,2.413c4.76,0,8.9-1.807,11.3-5.8Z" style="fill:#088cb9"/><path d="M6.98,17.732a4.832,4.832,0,0,0,1.961,3.01,8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.959,8.959,0,0,1-1.9.307c1.892.971,4.628.957,7.773-.252a20.545,20.545,0,0,0,5.377-3.13Z" style="fill:#039cc7"/><path d="M9.889,13.671h.172v1.813H9.889V13.671Zm-.33,0h.179v1.813H9.559V13.671Zm-.33,0h.179v1.813H9.23V13.671Zm-.33,0h.179v1.813H8.9V13.671Zm-.33,0h.179v1.813H8.57V13.671Zm-.323,0h.172v1.813H8.248V13.671Zm-.181-.181h2.175v2.176H8.066V13.49Z" style="fill:#00acd3"/><path d="M12.4,11.161h.172v1.813H12.4V11.161Zm-.33,0h.179v1.813H12.07V11.161Zm-.33,0h.179v1.813H11.74V11.161Zm-.33,0h.179v1.813H11.41V11.161Zm-.33,0h.178v1.813h-.178V11.161Zm-.323,0h.172v1.813h-.172V11.161Zm-.181-.181h2.176v2.176H10.577V10.979Z" style="fill:#00acd3"/><path d="M12.4,13.671h.172v1.813H12.4V13.671Zm-.33,0h.179v1.813H12.07V13.671Zm-.33,0h.179v1.813H11.74V13.671Zm-.33,0h.179v1.813H11.41V13.671Zm-.33,0h.178v1.813h-.178V13.671Zm-.323,0h.172v1.813h-.172V13.671Zm-.181-.181h2.176v2.176H10.577V13.49Z" style="fill:#26c2ee"/><path d="M14.909,13.671h.172v1.813h-.172V13.671Zm-.33,0h.179v1.813H14.58V13.671Zm-.33,0h.179v1.813H14.25V13.671Zm-.33,0H14.1v1.813h-.179V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.323,0h.172v1.813h-.172V13.671Zm-.181-.181h2.176v2.176H13.087V13.49Z" style="fill:#00acd3"/><path d="M14.909,11.161h.172v1.813h-.172V11.161Zm-.33,0h.179v1.813H14.58V11.161Zm-.33,0h.179v1.813H14.25V11.161Zm-.33,0H14.1v1.813h-.179V11.161Zm-.33,0h.179v1.813h-.179V11.161Zm-.323,0h.172v1.813h-.172V11.161Zm-.181-.181h2.176v2.176H13.087V10.979Z" style="fill:#26c2ee"/><path d="M17.42,13.671h.172v1.813H17.42V13.671Zm-.33,0h.179v1.813H17.09V13.671Zm-.33,0h.179v1.813H16.76V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.33,0h.179v1.813H16.1V13.671Zm-.323,0h.172v1.813h-.172V13.671ZM15.6,13.49h2.176v2.176H15.6V13.49Z" style="fill:#26c2ee"/><path d="M17.42,11.161h.172v1.813H17.42V11.161Zm-.33,0h.179v1.813H17.09V11.161Zm-.33,0h.179v1.813H16.76V11.161Zm-.33,0h.179v1.813h-.179V11.161Zm-.33,0h.179v1.813H16.1V11.161Zm-.323,0h.172v1.813h-.172V11.161Zm-.181-.181h2.176v2.176H15.6V10.979Z" style="fill:#00acd3"/><path d="M17.42,8.65h.172v1.813H17.42V8.65Zm-.33,0h.179v1.813H17.09V8.65Zm-.33,0h.179v1.813H16.76V8.65Zm-.33,0h.179v1.813h-.179V8.65Zm-.33,0h.179v1.813H16.1V8.65Zm-.323,0h.172v1.813h-.172V8.65ZM15.6,8.469h2.176v2.176H15.6V8.469Z" style="fill:#26c2ee"/><path d="M19.93,13.671H20.1v1.813H19.93V13.671Zm-.33,0h.178v1.813H19.6V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.323,0h.172v1.813h-.172V13.671Zm-.181-.181h2.176v2.176H18.107V13.49Z" style="fill:#00acd3"/><path d="M12.616,19.193a.6.6,0,1,1-.6.6.6.6,0,0,1,.6-.6" style="fill:#d5eef2"/><path d="M12.616,19.363a.431.431,0,0,1,.156.029.175.175,0,1,0,.241.236.43.43,0,1,1-.4-.265" style="fill:#3a4e55"/><path d="M2,17.949H29.92c-.608-.154-1.923-.362-1.707-1.159-1.105,1.279-3.771.9-4.444.267-.749,1.087-5.111.674-5.415-.173-.939,1.1-3.85,1.1-4.789,0-.3.847-4.666,1.26-5.415.173-.673.631-3.338,1.012-4.444-.267.217.8-1.1,1.005-1.707,1.159" style="fill:#3a4e55"/><path d="M14.211,23.518a5.287,5.287,0,0,1-2.756-2.711,9.2,9.2,0,0,1-1.987.3q-.436.024-.917.025-.554,0-1.168-.033a7.942,7.942,0,0,0,6.145,2.43q.344,0,.683-.013" style="fill:#c0dbe1"/><path d="M12.007,21.773a5.206,5.206,0,0,1-.552-.966,9.2,9.2,0,0,1-1.987.3,6.325,6.325,0,0,0,2.539.664" style="fill:#d5eef2"/></svg>`,"https://stackoverflow.com/"],
// ])


// const BlogList = writable([
//     // ["blogTite","link"]
//     ["Stack under attack: what we learned about handling DDoS attacks","https://heroicons.dev"],
//     ["An unfiltered look back at April Fools’ 2022","https://heroicons.dev"],
//     ["Solidity Tutorial - How to Build and Deploy an NFT Minting dApp with Solidity and React 🛠","https://heroicons.dev"],
//     ["Understanding XA Transactions With Practical Examples in Go ","https://heroicons.dev"],
//     ["Why you should use a Go backend in Flutter ","https://heroicons.dev"],
//     ["Is anyone interested in contributing to a new OSS built with Go? Please join to develop a NO-CODE workflow engine! ","https://heroicons.dev"],
// ])


const SearchedString = writable("" as string)



export {UserData,SearchedString}
// export { SpaceList,FavoriteHash}

