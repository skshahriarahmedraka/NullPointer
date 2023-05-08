<script lang="ts">

	import Groupdecoration from "../svgs/groupdecoration.svelte";
    export let  GroupData: {
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
    // console.log("🚀 ~ file: groupCom.svelte:5 ~ GroupData:", GroupData)

    // your script goes here
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class=" h-16 w-full m-2 ml-3 flex flex-row ">
    <div class=" relative grid place-content-center ">

        <Groupdecoration class=" fill-[#b8dffd]"/>
        <img src="{GroupData.Image}" alt="" class=" h-6 w-6 object-cover absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 rounded-md">
    </div>
    <div class=" h-full w-[150px] flex flex-col ml-2 ">
        <div class=" text-[#e7e9eb] text-base font-bold  break-words line-clamp-2">{GroupData.Name}</div>
        <div class=" text-[#e7e9eb] text-sm">{GroupData.NumberOfMember}</div>
    </div>
</div>