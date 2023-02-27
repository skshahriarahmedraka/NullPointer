<script lang="ts">
    import {RelatedQuestionList} from "../store/store"
    // export let RelatedQuestionList:any
    // your script goes here
    // let x= $RelatedQuestionList
    // console.log(typeof x)
    function My(a:any){
        return String(a[0])
    }

</script>

<style>
    /* your styles go here */
</style>

<div class="w-[300px]  flex flex-col  ">
    <div class=" text-sky-500 text-center text-xl font-semibold  my-3">Related Questions</div>
    {#each $RelatedQuestionList as i}
        <div class="flex flex-row hover:cursor-pointer">
            <div class=" w-12 h-6 rounded-md mt-1   { i[2] ? "bg-green-500" : "bg-gray-600" }  text-center  ">{i[1]}</div>
            <div  class=" w-full  text-sky-600 hover:text-sky-500  ml-3 mb-3 pr-2 text-sm">{i[0]}</div>
        </div> 
    {/each}
    
</div>