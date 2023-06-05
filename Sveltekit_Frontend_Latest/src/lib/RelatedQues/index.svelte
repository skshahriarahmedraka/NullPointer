<script lang="ts">
    // import {RelatedQuestionList} from "../store/store"
    // export let RelatedQuestionList:any
    // your script goes here
    // let x= $RelatedQuestionList
    // console.log(typeof x)
    function My(a:any){
        return String(a[0])
    }
    let RelatedQuesList= [
    
    ["How to differentiate empty string and nothing in a map",342,true,"https://www.youtube.com/"],
    
    ["go - checking if key of one map present in another",2,false,"https://www.youtube.com/"],
    ["How to tell if the value of a map is undefined in Go?",54,true,"https://www.youtube.com/"],
    
    ["Checking if key exist in map which return interface type in go",777,true,"https://www.youtube.com/"],
    
    
    ["What is \"_\" (underscore comma) in a Go declaration?",7,false,"https://www.youtube.com/"],
    
    ["What is the Big O performance of maps in golang?",99,true,"https://www.youtube.com/"],
    ["How to check if a slice has a given index in Go?",999,false,"https://www.youtube.com/"],
    
    ["How can I create an array that contains unique strings?",87,true,"https://www.youtube.com/"],
    
    
    ["Check if key exists in multiple maps in one condition",70,true,"https://www.youtube.com/"],
    
    ["What is benefit of not raising error (like Python) when key is not in map?",9,true,"https://www.youtube.com/"],
    ]

</script>

<style>
    /* your styles go here */
</style>

<div class="w-[300px]  flex flex-col  ">
    <div class=" text-sky-500 text-center text-xl font-semibold  my-3">Related Questions</div>
    {#each RelatedQuesList as i}
        <div class="flex flex-row hover:cursor-pointer">
            <div class=" w-12 h-6 rounded-md mt-1   { i[2] ? "bg-green-500" : "bg-gray-600" }  text-center  ">{i[1]}</div>
            <div  class=" w-full  text-sky-600 hover:text-sky-500  ml-3 mb-3 pr-2 text-sm">{i[0]}</div>
        </div> 
    {/each}
    
</div>