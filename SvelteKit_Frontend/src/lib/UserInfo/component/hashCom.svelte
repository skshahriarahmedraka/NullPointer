<script lang="ts">
	// import { Following } from '$lib/Navbar/svgs/following.svelte';

	// import Groupdecoration from "../svgs/groupdecoration.svelte";
    export let  HashData: {
        ID: string;
        Name: string;
    Image: string;
    Description: string;
    CoverImage: string;
    NumberOfMember: number;
} ={} as  {
    ID: string;
    Name: string;
    Image: string;
    Description: string;
    CoverImage: string;
    NumberOfMember: number;
}
    console.log("🚀 ~ file: groupCom.svelte:5 ~ HashData:", HashData)
    

    // your script goes here
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class=" h-16 w-full m-2 ml-3 flex flex-row ">
    <div class=" relative grid place-content-center ">

        <!-- <Groupdecoration class=" fill-[#b8dffd]"/> -->
        <img class="ml-1 h-10 w-10 object-cover " src={HashData.Image} alt="" />
        <!-- <img src="{HashData.Image}" alt="" class=" h-6 w-6 object-cover absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 rounded-md"> -->
    </div>
    <div class=" h-full w-[150px] flex flex-col ml-2 ">
        <div class=" text-[#e7e9eb] text-base font-bold  break-words line-clamp-1 ">{HashData.Name}</div>
        <div class=" text-[#e7e9eb] text-sm">{HashData.NumberOfMember} Follower</div>
    </div>
</div>