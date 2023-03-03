<script lang="ts">
    import {SpaceList ,FavoriteHash} from "../store/store"
    // import {} from "../store/store"

    // export let FavoriteHash:any 
    // export let SpaceList:any

    let focusid=$SpaceList[0][0]
    function SetFocusID(i:any){
        focusid=i
    }
    let HashFocusid: any
    function SetHashFocusID(i:any){
        HashFocusid=i
    }
</script>

<style>
    /* your styles go here */
</style>



<div class=" w-[165px]  flex flex-col mt-16  ">
    <!-- <div class=" w-full h-60 bg-slate-500 m-2 "></div> -->
    <div class="h-10 w-full   rounded-l-lg text-center pt-2 text-sky-600 border-2 border-blue-500 font-semibold mb-2">Collective</div>
    {#each $SpaceList as i}
        <div on:click="{()=>SetFocusID(i[0])}" class="  h-10 w-full hover:cursor-pointer flex flex-row place-content-left bg-[#2d2d2d] hover:bg-[#3d3d3d] rounded-l-xl  pt-2  text-[#e7e9eb] text-sm my-1  {i[0]===focusid ? " border-r-4 border-blue-600":""}">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 4a1 1 0 011-1h4a1 1 0 010 2H6.414l2.293 2.293a1 1 0 11-1.414 1.414L5 6.414V8a1 1 0 01-2 0V4zm9 1a1 1 0 010-2h4a1 1 0 011 1v4a1 1 0 01-2 0V6.414l-2.293 2.293a1 1 0 11-1.414-1.414L13.586 5H12zm-9 7a1 1 0 012 0v1.586l2.293-2.293a1 1 0 111.414 1.414L6.414 15H8a1 1 0 010 2H4a1 1 0 01-1-1v-4zm13-1a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 010-2h1.586l-2.293-2.293a1 1 0 111.414-1.414L15 13.586V12a1 1 0 011-1z" clip-rule="evenodd"></path></svg>
            <div class="line-clamp-1">
                {i[0]}
            </div>
        </div>
    {/each}
        <!-- My Favorite hash Tags -->
    <div class="h-10 w-full rounded-l-lg text-center pt-2 text-sky-600 border-2 border-blue-500 font-semibold mb-2 mt-4">Favorite</div>
    {#each $FavoriteHash as i}
        <div on:click="{()=>SetHashFocusID(i[0])}" class=" hover:cursor-pointer h-10 w-full  flex flex-row place-content-left  hover:bg-[#3d3d3d]  rounded-l-xl  pt-2  text-[#e7e9eb] text-sm my-1  {i[0]===HashFocusid ? " border-r-4 border-r-blue-600 border-y-2 border-l-2 border-y-[#2d2d2d] border-l-[#2d2d2d] ":"border-2 border-[#2d2d2d]"}">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14"></path></svg>
            <!-- <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 4a1 1 0 011-1h4a1 1 0 010 2H6.414l2.293 2.293a1 1 0 11-1.414 1.414L5 6.414V8a1 1 0 01-2 0V4zm9 1a1 1 0 010-2h4a1 1 0 011 1v4a1 1 0 01-2 0V6.414l-2.293 2.293a1 1 0 11-1.414-1.414L13.586 5H12zm-9 7a1 1 0 012 0v1.586l2.293-2.293a1 1 0 111.414 1.414L6.414 15H8a1 1 0 010 2H4a1 1 0 01-1-1v-4zm13-1a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 010-2h1.586l-2.293-2.293a1 1 0 111.414-1.414L15 13.586V12a1 1 0 011-1z" clip-rule="evenodd"></path></svg> -->
            <div class="line-clamp-1">
                {i[0]}
            </div>
        </div>
    {/each}
    
</div>