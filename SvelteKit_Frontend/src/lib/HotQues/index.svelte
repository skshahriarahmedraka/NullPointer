<script lang="ts">
    import { HotQues } from "../store/store";
    // export let HotQues:any 
</script>

<style>
    /* your styles go here */
</style>

<div class="w-full h-full flex flex-col">
    <div class=" text-lg text-indigo-500 font-semibold ml-4">Hot Questions</div>
    
    {#each $HotQues as i}
         <div class="flex flex-row ml-3">
             <!-- icon -->
             <div class="h-10 w-10 mt-3">
                 {@html i[1]}
             </div>
             <!-- Ques -->
             <div class="text-[#32a1f6] hover:text-sky-600 hover:cursor-pointer mt-2 mx-2 ">
                {i[0]}
            </div>
         </div>
    {/each}
</div>