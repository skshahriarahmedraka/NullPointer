

<script lang="ts">
    
    // import Navbar from "$lib/Navbar/index.svelte"
    // import Collectives from "$lib/Collectives/index.svelte"
    // import RelatedQues from "$lib/RelatedQues/index.svelte"
    import Ques from "$lib/Ques/index.svelte"
    import Navbar from '$lib/Navbar/index.svelte';
	import Collectives from "$lib/Collectives/index.svelte"
	import Footer from '$lib/Footer/footer.svelte';

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
import type { PageData } from './$types';
	import type { CookieInfo1Type, QuestionDataType, UserDataType } from "$lib/store/types";
	import { afterUpdate, onMount } from "svelte";
	import { getCookieValue } from "$lib/store/utils";
	import { fetchUserData } from "$lib/store/fetch";
	import { UserData } from "$lib/store/store";

export let data: PageData;
  let QuestionData:QuestionDataType 
  QuestionData =  data.QuestionData
  console.log("🚀 ~ file: +page.svelte:79 ~ QuestionData:", QuestionData)

//   let loadingState: boolean = false;
// 	afterUpdate(() => {
// 		loadingState = true;
// 	});


    let loadingState: boolean = false;
	onMount(async () => {
		let CookieValueInfo1: string = getCookieValue('Info1');
		console.log('🚀 ~ file: +page.svelte:30 ~ onMount ~ CookieValueInfo1:', CookieValueInfo1);
		let InfoCookieData = JSON.parse(atob(CookieValueInfo1)) as CookieInfo1Type;

		// let InfoCookieData = JSON.parse(Buffer.from(CookieValueInfo1, 'base64').toString('utf-8')) as CookieInfo1Type;
		console.log('🚀 ~ file: +page.ts:24 ~ load ~ InfoCookieData:', InfoCookieData);
		let UserDataValue = {} as UserDataType;
		UserData.subscribe((value) => {
			UserDataValue = value;
		});
		if (UserDataValue.ID != InfoCookieData.UUID) {
			const GetUserData = await fetchUserData(InfoCookieData.UUID);
			console.log('🚀 ~ file: +page.ts:24 ~ InitializeData ~ GetUserData:', GetUserData);
			UserData.update(() => GetUserData);
		}
		loadingState = true;
	});
const HeadLogo = new URL("../../../../lib/icons/favicon.png", import.meta.url).href;

</script>

<style>
    /* your styles go here */
</style>
<svelte:head>
	<title>{QuestionData.QuesTitle}</title>
	<link rel="icon"  href={HeadLogo} />
</svelte:head>



{#if loadingState}

	<!-- ////////////////////////////// -->

    
    <div class="   flex   w-full flex-col  justify-center overflow-x-hidden overflow-y-hidden bg-[#181818] ">
        <Navbar />
        <div class="flex w-full flex-row justify-center   ">
            
            <!-- COLLECTIVE -->
            <Collectives />
            <!-- QUESTION LIST -->
    
            <Ques  QuestionData={QuestionData} />
        </div>
        <Footer />
    </div>
{:else}
	<LoadingSVG/>
{/if}












