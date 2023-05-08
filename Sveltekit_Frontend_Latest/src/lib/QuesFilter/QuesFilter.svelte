<script lang="ts">
    // your script goes here

	import ZtoA from "./svgs/ZtoA.svelte";

    import { goto } from "$app/navigation";
    let filterState: boolean = true;

</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class="flex h-16 w-full flex-row">
    <div class="ml-5 self-center text-lg text-[#e7e9eb]">{22587351} Questions</div>
    <div class="grow" />
    <div class=" flex h-9  flex-row self-center justify-around rounded-md border-2 border-[#7d858d]">
        <div
            class="flex h-full px-2 items-center justify-center border-r-2 border-[#7d858d] text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245]"
        >
            Time
        </div>
        <div
            class="flex h-full px-2  items-center justify-center border-r-2 border-[#7d858d] text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245]"
        >
            Popular
        </div>
        <div
            class="flex h-full px-2  items-center justify-center border-r-2 border-[#7d858d] text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245]"
        >
            Rated
        </div>
        <div
            class="flex h-full px-2  items-center justify-center border-r-0 border-[#7d858d] text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245]"
        >
            Unanswered
        </div>
        <!-- <div
            class="flex h-full w-16 items-center justify-center text-[#e7e9eb] hover:cursor-pointer hover:bg-[#404245]"
        >
            More
        </div> -->
    </div>
    <div
        on:click={() => {
            filterState = !filterState;
        }}
        on:keydown={() => {}}
        class=" hover:cursor-pointer ml-4 flex h-9 w-24 flex-row items-center justify-center self-center rounded-md border-2 border-[#7d858d] hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245]"
    >
        <ZtoA class=" {filterState ?" rotate-180" :""}"/>
        <div class="ml-1 font-semibold text-[#9cc1db]">order</div>
    </div>
</div>