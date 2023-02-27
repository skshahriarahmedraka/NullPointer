<script lang="ts">
// import { toggle_class } from "svelte/internal";
// import {PublicQuesData} from "../store/store"

// export let PublicQuesData:any


    let PublicQuesData =[
        // [voteNumber(int),AnsNum(int),Viewnum(num),AnsAccepted(bool),askedby(array),tag(Array),QuesTitle(string),QuesDes(string) ]
        [235,999,1626842,true,["yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"],["go","rust","backend","linux"],"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"],
        [235,9,1626842,true,["yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"],["go","rust","backend","linux"],"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"],
        [0,0,0,false,["yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"],["go","rust","backend","linux"],"How to check if a map contains a key in Go? How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec 185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"],
        [786,999,1626842,false,["yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"],["go","rust","backend","linux"],"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"],
        [235,999,1626842,true,["yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"],["go","rust","backend","linux"],"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"],
        [235,0,1626842,false,["yuji","Yuji Itadori",999,"Jan 12, 2010 at 16:18"],["go","rust","backend","linux"],"How to check if a map contains a key in Go?","185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"],
    ]
    
</script>

<style>
    /* your styles go here */
</style>

<div class="w-full h-full">
   {#each PublicQuesData as i}
        <div class=" min-h-32 max-h-40 w-full border-t-2 mx-2 py-2 border-[#404245] flex flex-row">
            <!-- quesVote -->
            <div class="h-full w-60 my-2 flex flex-col justify-center items-center gap-1 ">
                <!-- VoteNumber -->
                <div class="  {i[0]>0 ? "text-[#e7e9eb]" : "text-[#959ba0]" } ">{i[0]} Votes</div>
                <!-- Answer Number -->
                {#if i[3]}
                     <!-- acceped ans -->
                     <div class=" bg-green-500 text-white px-2 rounded-md">{i[1]} Answers</div>
                {:else if  i[1]!=0}
                     <!-- Many ans but No accepted -->
                     <div class="text-[#e7e9eb] border-2 border-green-500 rounded-md px-2">{i[1]} Answers</div>
                {:else}
                     <!-- No answer -->
                     <div class="text-[#959ba0]">{i[1]} Answers</div>
                {/if}
                <!-- View Number -->
                <div class="{i[0]>0 ? "text-[#e7e9eb]" : "text-[#959ba0]" } ">{i[2]} Views</div>
            </div>
            <!-- Qtitle &detail -->
            <div class=" w-[1200px] h-full flex flex-col">
                <!-- Title -->
                <div class=" text-xl text-sky-600 hover:text-blue-600 hover:cursor-pointer line-clamp-2">{i[6]}</div>
                <!-- description -->
                <div class=" text-md text-[#e7e9eb] line-clamp-2">{i[7]}</div>
                <!-- tag ans user -->
                <div class="flex flex-row  justify-center items-center ">
                    <div class="  min-h-10 max-h-20  w-full mt-2 flex flex-row flex-wrap">
                        <!-- {#each i[5] as tag}
                            <div class="text-[#9bc0da] hover:text-teal-200 hover:bg-slate-600 bg-[#3d4951] m-1 px-2 py-1 rounded-md hover:cursor-pointer">
                                {tag}
                            </div>
                        {/each}    -->
                    </div>
                    <div class=""></div>
                    
                    <!-- <p class="h-6 w-36 text-center mt-2 mr-1 text-sky-600 hover:text-blue-600 hover:cursor-pointer line-clamp-1">{i[4][1]}</p> -->
                    <!-- <p class="text-[#e7e9eb] mt-2 mr-1">{i[4][2]}</p> -->
                    <!-- <p class=" w-80 text-[#959ba0] mt-2 ">{i[4][3]}</p> -->
                </div>
            </div>
        </div>
   {/each}

</div>