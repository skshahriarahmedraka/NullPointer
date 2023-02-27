<script lang="ts">
    import {UserData} from "../store/store"
    // export let UserData:any
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->
<div class="h-full w-3/12 flex flex-col ">
    <!-- stats -->
    <!-- Stats Title -->
    <div class=" text-[#e7e9eb] text-2xl ml-4 ">
        Stats
    </div>
    <div class=" w-auto h-48  m-4 rounded-xl border-[1px]  border-gray-500  ">
        <div class="w-full h-32 flex flex-col ">
            <div class="  w-full h-1/2 flex justify-center items-center">
                <div class=" flex flex-col w-1/2 justify-center items-center ">
                    <p class="text-[#e7e9eb] text-xl font-semibold">{755}</p>
                    
                    <p class=" text-gray-400 text-sm">
                        reputation
                    </p>
                </div>
                <div class="flex flex-col w-1/2 justify-center items-center">
                    <p class="text-[#e7e9eb] text-xl font-semibold">{"65k"}</p>
                    
                    <p class="text-gray-400 text-sm">
                        reached
                    </p>
                </div>
            </div>
            <div class="w-full h-1/2 flex justify-center items-center">
                <div class=" flex flex-col w-1/2 justify-center items-center ">
                    <p class="text-[#e7e9eb] text-xl font-semibold">{843}</p>
                    
                    <p class=" text-gray-400 text-sm">
                        Answers
                    </p>
                </div>
                <div class="flex flex-col w-1/2 justify-center items-center">
                    <p class="text-[#e7e9eb] text-xl font-semibold">{53}</p>
                    
                    <p class="text-gray-400 text-sm">
                        Questions
                    </p>
                </div>
            </div>
        </div>
        <div class="w-full h-16 flex flex-col justify-center items-center  ">
            <svg class="w-6 h-6 text-[#e7e9eb]" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"></path></svg> 
            <div class="text-sm text-gray-400">top 6% this year</div>
        </div>
    </div>
    <!-- Tags -->
    <div class=" text-[#e7e9eb] text-2xl ml-4 ">
        Top Tags
    </div>
    <div class="w-auto h-auto  m-4 rounded-xl border-[1px]  border-gray-500 p-2 flex flex-col gap-1  ">
        {#each $UserData["TopTags"] as tag,j }
        <div class="flex flex-row">
            <div class="text-[#9bc0da] w-fit hover:text-teal-200 hover:bg-slate-600 bg-[#3d4951] m-1 px-2 py-1 rounded-md hover:cursor-pointer">
                {tag}
            </div>
            <div class="grow"></div>
            <div class="mt-1 mr-2 text-lg text-[#e7e9eb]">
                {$UserData["TopTagsPercent"][j]}%
            </div>
        </div>
             
        {/each}
        
        
    </div>
</div>