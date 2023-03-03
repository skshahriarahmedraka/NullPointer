
<script lang="ts">
    import UserInfo from "$lib/UserInfo/Title.svelte"
    import Stats from "$lib/UserInfo/Stats.svelte"
    import Abouts from "$lib/UserInfo/Abouts.svelte"
    import {UserData} from "$lib/store/store"

    // export let UserData:any
    // console.log("🚀 ~ file: index@root.svelte ~ line 47 ~ UserData", UserData)
    
</script>

<style>
    /* your styles go here */
</style>


    <!-- settings -->
    <div class=" w-[1100px] min-h-screen max-h-full  bg-[#2d2d2d] mt-12 flex flex-col   ">
        <!-- image & name & social media -->
        <UserInfo  />
        <!-- Button : profile, activity, -->
        <div class="  pl-10 w-full flex flex-row justify-start items-center gap-5">
            <div on:click={()=> $UserData["SelectedPanel"]="Profile"} on:keydown={()=>{}} class="ml-3 -mr-3 h-9 w-24 rounded-xl {$UserData["SelectedPanel"]==="Profile" ? "bg-sky-500 hover:bg-blue-600" :"border-2 border-blue-600 " }  hover:cursor-pointer flex justify-center items-center ">
                <p class=" my-auto  font-semibold text-base {$UserData["SelectedPanel"]==="Profile" ? "text-gray-200" :"text-sky-500" }">Profile</p>
            </div>
            <div on:click={()=> $UserData["SelectedPanel"]="Activity"} on:keydown={()=>{}} class="ml-3 -mr-3 h-9 w-24 rounded-xl {$UserData["SelectedPanel"]==="Activity" ? "bg-sky-500 hover:bg-blue-600" :"border-2 border-blue-600 " }   hover:cursor-pointer flex justify-center items-center ">
                <p class=" my-auto  font-semibold text-base {$UserData["SelectedPanel"]==="Activity" ? "text-gray-200" :"text-sky-500" }">Activity</p>
            </div>
            <div on:click={()=> $UserData["SelectedPanel"]="Settings"} on:keydown={()=>{}} class="ml-3 -mr-3 h-9 w-24 rounded-xl {$UserData["SelectedPanel"]==="Settings" ? "bg-sky-500 hover:bg-blue-600" :"border-2 border-blue-600 " }  hover:cursor-pointer flex justify-center items-center ">
                <p class=" my-auto  font-semibold text-base {$UserData["SelectedPanel"]==="Settings" ? "text-gray-200" :"text-sky-500" }">Settings</p>
            </div>
        </div>
        <!-- details about me -->
        {#if $UserData["SelectedPanel"]==="Profile"}
            <div class="h-full w-full flex flex-row mt-5">
                <!-- Stats and Tags -->
                <Stats />
                <!-- About -->
                <Abouts />

            </div>
             <!-- content here -->
        {:else if   $UserData["SelectedPanel"]==="Activity"}
             <div class="">Activity</div>
        {:else}
             <div class="">Settings</div>
             <!-- else content here -->
        {/if}
        
        
    </div>

