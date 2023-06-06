<script lang="ts">
	import { onMount } from 'svelte';

	export let pageNumStart: number;
	export let pageNumNow: number;
	export let pageNumEnd: number;
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
	pageNumStart = pageNumStart;
	let Arr: number[] = [];

	function CalculatePageNum() {
		Arr = [];
		if (pageNumEnd <= 3) {
			for (let i = 0; i <= pageNumEnd; i++) {
				Arr.push(i);
			}

			return;
		} else if (pageNumEnd > 3) {
			if (pageNumNow - 1 >= 0) {
				Arr.push(pageNumNow - 1);
			}
			Arr.push(pageNumNow);
			if (pageNumNow + 1 <= pageNumEnd) {
				Arr.push(pageNumNow + 1);
			}
			if (Arr[0] > 1) {
				Arr.unshift(0, -1);
			} else if (Arr[0] == 1) {
				Arr.unshift(0);
			}
			if (Arr[Arr.length - 1] < pageNumEnd - 1) {
				Arr.push(-1, pageNumEnd);
			} else if (Arr[Arr.length - 1] == pageNumEnd - 1) {
				Arr.push(pageNumEnd);
			}
		}
	}
	onMount(() => {
		CalculatePageNum();
	});
</script>

<!-- markup (zero or more items) goes here -->

<div class=" flex h-32 w-full items-center justify-center gap-2 font-sf-pro">
	<!-- //////////////////////////////////////////////////////////////////// -->
	{#if pageNumNow != 0}
		<!-- content here -->
		<button
			class="h-8 w-fit"
			on:click={() => {
				pageNumNow = pageNumNow - 1;
				// CalculatePageNum();
			}}
			><p
				class=" h-8 w-fit rounded-md border-[1px] border-[#a9a9a9] border-opacity-40 px-2 py-1 text-center text-sm font-bold text-[#a9a9a9]"
			>
				Prev
			</p></button
		>
	{/if}
	{#each Arr as i}
		{#if i != -1}
			<!-- content here -->
			<button
				class="h-8 w-8 rounded-md {i === pageNumNow
					? 'bg-[#f28225]'
					: 'border-[1px] border-[#a9a9a9] border-opacity-40 text-[#a9a9a9]'} "
				on:click={() => {
					pageNumNow = i;
					// CalculatePageNum();
				}}><p class="text-sm font-bold">{i + 1}</p></button
			>
		{:else}
			<!-- else content here -->
			<button class="h-8 w-8 rounded-md" on:click={() => {}}
				><p class="text-sm font-bold text-[#a9a9a9]">...</p></button
			>
		{/if}
	{/each}
	{#if pageNumNow != pageNumEnd}
		<button
			class="h-8 w-fit"
			on:click={() => {
				pageNumNow = pageNumNow + 1;
				// CalculatePageNum();
			}}
			><p
				class=" h-8 w-fit rounded-md border-[1px] border-[#a9a9a9] border-opacity-40 px-2 py-1 text-center text-sm font-bold text-[#a9a9a9]"
			>
				Next
			</p></button
		>
	{/if}
</div>
