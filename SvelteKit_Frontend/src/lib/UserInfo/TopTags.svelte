<script lang="ts">

import {UserData} from "../store/store"
    
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->
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