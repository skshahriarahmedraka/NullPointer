

<script lang="ts">
    
    // import Navbar from "$lib/Navbar/index.svelte"
    // import Collectives from "$lib/Collectives/index.svelte"
    // import RelatedQues from "$lib/RelatedQues/index.svelte"
    import Ques from "$lib/Ques/index.svelte"
    // import Ans from "$lib/Ans/index.svelte"
    // import {UserData,QuestionData,Loading} from "$lib/store/store"
    import LoadingSVG from "$lib/Loading/index.svelte"

    // export let SpaceList:any
    // export let FavoriteHash:any 
    // export let UserData:any
    // export let QuestionData:any
    // export let RelatedQuestionList:any
    // export let RelatedQuestionListLoading:boolean
    // export let QuestiondataLoading:boolean   
    // export let QuestionData:any

// import { onMount } from 'svelte';

    // let loading=false 
    // onMount(async()=>{
    //     let x2={}   
    //     let response = await fetch("http://localhost:8080/q/raka",{mode:"cors"} )
    //     if (response.ok) { // if HTTP-status is 200-299
    //         x2 = await response.json();
    //         // console.log("🚀 ~ file: index.svelte ~ line 25 ~ name ~ QuestionData : ", x2)
    //         QuestionData.update((n)=>n=x2)
    //         loading=true 
    //     } else {
    //         console.log("HTTP-Error: " + response.status);
    //     }
    // })
    // console.log("🚀 ~ file: store.ts ~ line 100 ~ FetchQuestionData ~ QuestionData : ",  typeof($QuestionData))


// let QuestionData= {
//     "QuestionTitle":"How to efficiently concatenate strings in go How to efficiently concatenate strings in go How to efficiently concatenate strings in goHow to efficiently concatenate strings in goHow to efficiently concatenate strings in go goHow to efficiently concatenate strings in goHow to",
//     "QuestionAskedTime":"12 years, 5 months ago",
//     "QuestionModified" : "6 months ago",
//     "QuestionViewed": 577000,
//     "QuestionUpvote":4,
//     "QuestionDownvote":32,
//     "QuestionBookmark":23,
//     "QuestionTags":["go","string","Concatination","go","string","Concatination","go","string","Concatination","go","string","go","string","Concatination","Concatination"],
//     "QuestionAskedBy":"yujiitadori",
//     "QuestionEditedBy":"KingofRandom",
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
//         [`994New Way: From Go 1.10 there is a strings.Builder type, please take a look at this answer for more detail. Old Way: Use the bytes package. It has a Buffer type which implements io.Writer. package main import ( "bytes" "fmt" ) func main() { var buffer bytes.Buffer for i := 0; i < 1000; i++ { buffer.WriteString("a") } fmt.Println(buffer.String()) } This does it in O(n) time.`,123,true,"inancramsay","Dec 12, 2019 at 7:30"],
//         [`n Go 1.10+ there is strings.Builder, here.A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use.`,43,false,"luffy","icza", "Dec 4, 2015 at 8:01"] 
    
//         // ["ans(string)","vote(int)","accepted(bool)","user(string)","time(string)"]
//     ]


// }


  
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<!-- 
{#if $Loading}
    
    <div class="bg-[#181818]   w-full flex flex-row justify-center   ">
        
        <Navbar {UserData}/>
        <!-- COLLECTIVE -->
        <!-- <Collectives {FavoriteHash} {SpaceList} />  -->
        <!-- QUESTION LIST -->
        <!-- <Ques/> -->
        

        <!-- RELATED QUESTION -->
        <!-- <RelatedQues/> -->
<!-- 
    </div>
{:else}
{/if} -->

{#if true}
<!-- content here -->
<Ques  />
{:else}
<!-- else content here -->
<LoadingSVG/>
{/if}









