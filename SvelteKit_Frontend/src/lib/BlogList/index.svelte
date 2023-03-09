<script lang="ts">
	import { goto } from "$app/navigation";


    // import {BlogList} from "../store/store"
    const BlogList=[
    // ["blogTite","link"]
    ["Stack under attack: what we learned about handling DDoS attacks","https://heroicons.dev"],
    ["An unfiltered look back at April Fools’ 2022","https://heroicons.dev"],
    ["Solidity Tutorial - How to Build and Deploy an NFT Minting dApp with Solidity and React 🛠","https://heroicons.dev"],
    ["Understanding XA Transactions With Practical Examples in Go ","https://heroicons.dev"],
    ["Why you should use a Go backend in Flutter ","https://heroicons.dev"],
    ["Is anyone interested in contributing to a new OSS built with Go? Please join to develop a NO-CODE workflow engine! ","https://heroicons.dev"],
]
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class=" m-4 p-2 pr-2  h-full  flex flex-col border-2 border-[#393939] rounded-md text-[#e7e9eb]">
    <div class=" text-lg text-indigo-500 font-semibold">Blogs</div>
    {#if BlogList.length != 0}
      
         {#each BlogList as blog}
         <div on:click={()=> {goto(`/blog/${blog[1]}`)}} on:keypress={()=>{}} class=" w-full   mt-3 flex flex-row justify-center items-center hover:text-zinc-400 cursor-pointer ">
             <svg class=" w-5 h-5 mt-0 mx-1  " fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"></path></svg>
             <div class=" w-fit text-sm line-clamp-2  mr-1">{blog[0]}</div>
         </div>
         
         {/each}
    {:else}
         <div class="text-xl text-white h-full w-full">
            No Blog
         </div>
    {/if}
</div>