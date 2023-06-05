<script lang="ts">
	export let pageNumStart :number;
	export let pageNumNow :number;
	export let pageNumEnd :number ;
	// let PresentArray = Array.from(Array(6).keys()).slice(1);
	// let func = function () {
	// 	if (pageNumNow < 3) {
	// 		PresentArray = Array.from(Array().keys()).slice(1);
	// 	} else if (pageNumNow > pageNumEnd - 2) {
	// 		PresentArray = Array.from(Array(pageNumNow + 5).keys()).slice(pageNumEnd - 5);
	// 	} else {
	// 		PresentArray = Array.from(Array(pageNumNow + 3).keys()).slice(pageNumNow - 2);
	// 	}
	// };
	// $: func();
	function CalculatePageNum (){

	}
</script>

<!-- markup (zero or more items) goes here -->


<div class=" flex h-32 w-full items-center justify-center font-sf-pro ">
	{#if pageNumEnd!=0 }
		<div class="flex items-center justify-center gap-2 ">
			{#if pageNumNow != 0}
				<button
					class="h-8 w-fit   "
					on:click={() => {
						pageNumNow = pageNumNow - 1;
					}}><p class=" h-8 px-2 py-1 w-fit rounded-md text-center text-sm font-bold text-[#a9a9a9] border-[1px] border-[#a9a9a9] border-opacity-40">Prev</p></button
				>
			{/if}
			<button
				class="h-8 w-8  rounded-md  {1 === pageNumNow ? 'bg-[#f28225]' : 'text-[#a9a9a9] border-[1px] border-[#a9a9a9] border-opacity-40'}"
				on:click={() => {
					pageNumNow = 0;
				}}><p class="text-sm font-bold ">1</p></button
			>

			{#if pageNumNow > 3}
				<button
					class="h-8 w-8 rounded-md "
					on:click={() => {
						
					}}><p class="text-sm font-bold text-[#a9a9a9] ">...</p></button
				>
			{/if}
			<div class="flex items-center justify-center gap-2">
				{#if pageNumNow < 4 && pageNumEnd >0}
					
					{#each Array.from(Array(5+1 ).keys()).slice(2) as i}
						<button
							class="h-8 w-8  rounded-md  {i === pageNumNow ? 'bg-[#f28225]' : 'text-[#a9a9a9] border-[1px] border-[#a9a9a9] border-opacity-40'} "
							on:click={() => {
								pageNumNow = i;
							}}><p class="text-sm font-bold ">{i+1}</p></button
						>
					{/each}
				{:else if pageNumNow > pageNumEnd - 4}
					{#each Array.from(Array(pageNumEnd ).keys()).slice(pageNumEnd-4) as i}
						<button
							class="h-8 w-8 rounded-md {i === pageNumNow ? 'bg-[#f28225]' : 'text-[#a9a9a9] border-[1px] border-[#a9a9a9] border-opacity-40'} "
							on:click={() => {
								pageNumNow = i;
							}}><p class="text-sm font-bold ">{i+1}</p></button
						>
					{/each}
				{:else}
					<!-- content here -->
                    {#each Array.from(Array(pageNumNow+3).keys()).slice(pageNumNow-2) as i }
                    <button
                        class="h-8 w-8 rounded-md  { i===pageNumNow ? "bg-[#f28225]" : "text-[#a9a9a9] border-[1px] border-[#a9a9a9] border-opacity-40"} "
                        on:click={() => {
							pageNumNow = i;
                        }}><p class="text-sm font-bold ">{ i+1}</p></button
                    >
                {/each}
				{/if}
			</div>
            {#if pageNumNow < pageNumEnd - 3}
				<button
					class="h-8 w-8 rounded-md"
					on:click={() => {
				
					}}><p class="text-sm font-bold text-[#a9a9a9] ">...</p></button
				>
			{/if}
			{#if pageNumNow>=5}
				 <button
					 class="h-8 w-8 rounded-md {pageNumEnd === pageNumNow ? 'bg-[#f28225]' : 'text-[#a9a9a9] border-[1px] border-[#a9a9a9] border-opacity-40'}"
					 on:click={() => {
						 pageNumNow = pageNumEnd;
					 }}><p class="text-sm font-bold ">{pageNumEnd+1}</p></button
				 >
			{/if}
			{#if pageNumNow != pageNumEnd}
				<button
					class="h-8 w-fit  "
					on:click={() => {
						pageNumNow = pageNumNow + 1
					}}><p class=" h-8 px-2 py-1 w-fit rounded-md text-center text-sm font-bold text-[#a9a9a9] border-[1px] border-[#a9a9a9] border-opacity-40">Next</p></button
				>
			{/if}
		</div>
	{:else}
		<p class="" />
	{/if}
</div>


<style>
	/* your styles go here */
</style>


{#if pageNumEnd!=0 }
	 {#if pageNumNow != 0}
		 <!-- prev  -->
	 {/if}
	 
	 <!-- 1 button -->
	 {#if pageNumEnd<3}
		 {#each 0 as pageNumEnd}
			 <!-- content here -->
		 {/each}
	{:else if  pageNumNow > 3}
			<!-- 1 -->
			<!-- ... -->
			<!-- page Now  -->
			<!-- ... -->
			<!-- Page End -->

	 {/if}

	 
{:else}
	 <!-- else content here -->
{/if}



{#if pageNumNow > 3}
		 <!-- ... -->
	 {/if}
	 {#if condition}
	 <!-- content here -->
	{:else if  otherCondition}
		<!-- else if content here -->
	{:else}
		<!-- else content here -->
	{/if}