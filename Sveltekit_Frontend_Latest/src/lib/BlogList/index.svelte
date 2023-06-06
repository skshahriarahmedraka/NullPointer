<script lang="ts">
	import { goto } from "$app/navigation";
	import Trash from "$lib/icons/trash.svelte";
	import { fetchBlogList } from "$lib/store/fetch";
	import type { BlogDataType } from "$lib/store/types";
	import { onMount } from "svelte";


    // import {BlogList} from "../store/store"
//     const BlogList=[
//     // ["blogTite","link"]
//     ["Stack under attack: what we learned about handling DDoS attacks","https://heroicons.dev"],
//     ["An unfiltered look back at April Fools’ 2022","https://heroicons.dev"],
//     ["Solidity Tutorial - How to Build and Deploy an NFT Minting dApp with Solidity and React 🛠","https://heroicons.dev"],
//     ["Understanding XA Transactions With Practical Examples in Go ","https://heroicons.dev"],
//     ["Why you should use a Go backend in Flutter ","https://heroicons.dev"],
//     ["Is anyone interested in contributing to a new OSS built with Go? Please join to develop a NO-CODE workflow engine! ","https://heroicons.dev"],
// ]

let loadingState: boolean = false;
let BlogList: BlogDataType[] = [] as BlogDataType[]
onMount(async () => {
    BlogList =await  fetchBlogList("time", 0, 10, 1);
    console.log("🚀 ~ file: index.svelte:24 ~ onMount ~ BlogList:", BlogList)
    loadingState = true;
});
	const StarLoading = new URL('$lib/icons/starLoading.gif', import.meta.url).href;

</script>

<style>
    /* your styles go here */
</style>



{#if loadingState}
     
     <div class=" m-4 p-2 pr-2  h-full  flex flex-col border-2 border-[#393939] rounded-md text-[#e7e9eb]">
         <div on:click={()=>{goto("/b")}} on:keypress={()=>{}} class=" text-lg text-indigo-500 font-semibold hover:text-indigo-400 hover:cursor-pointer ">Blogs</div>
     
         {#if BlogList.length != 0}
           
              {#each BlogList as blog}
              <div on:click={()=> {goto(`/b/${blog.ID}`)}} on:keypress={()=>{}} class=" w-full   mt-3 flex flex-row justify-center items-center hover:text-zinc-400 cursor-pointer ">
                  <svg class=" w-5 h-5 mt-0 mx-1  " fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"></path></svg>
                  <div class=" w-fit text-sm line-clamp-2  mr-1">{blog.Title}</div>
              </div>
              
              {/each}
         {:else}
              <div class="text-xl text-white h-full w-full flex flex-col justify-center aspect-1 items-center ">
                 <!-- No Blog -->
                 <!-- <Trash class="h-10 " /> -->
                 <img src="{StarLoading}" alt="Star Loading" class="">
                 <p class=" font-raleway ">No Blogs</p>
              </div>
         {/if}
     </div>
{:else}
     <!-- else content here -->
     <div class=" m-4 p-2 pr-2  h-full  flex flex-col justify-center items-center border-2 border-[#393939] rounded-md text-[#e7e9eb]">
        <img src="{StarLoading}" alt="Star Loading" class=" h-10 w-10">
    </div>

{/if}