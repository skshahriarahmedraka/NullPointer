<script lang="ts">
	import { onMount } from 'svelte';
	import Emptybox from '$lib/errors/emptybox.svelte';

	// let QuestionList: QuestionDataType[] = [] as QuestionDataType[];
	// let QuestionList:QuesArrWithMetadataType = {} as QuesArrWithMetadataType;

	// let fetchData = async () => {

	// }
	// fetchData()
	// console.log('🚀 ~ file: index.svelte:126 ~ QuestionList:', QuestionList);

	import ZtoA from '$lib/PublicQues/svgs/ZtoA.svelte';
	// import PageNum from '$lib/PageNum/pageNum.svelte';
	import SearchedQues from './searchedQues.svelte';
	import SearchedHash from './SearchedHash.svelte';
	import SearchedBlog from './SearchedBlog.svelte';
	import type { QuesArrWithMetadataType } from '$lib/store/types';
	// import {  fetchQuesArrWithMetadata } from '$lib/store/fetch';
	// import type { QuesArrWithMetadataType } from '$lib/store/types';

	// import { goto } from "$app/navigation";
	let filterState: boolean = true;
	// GetEditedByData(QuestionData.QuesEditedBy);
	let filterType: string = 'time';
	// let filterIndex: number = 0;
	// let pageNumStart: number = 0;
	// let pageNumNow: number = 0;
	// let pageNumEnd: number ;
	// let contentPerPage:number=5
	let Loading:boolean=false

	// export let SearchedString: string 
	// console.log("🚀 ~ file: search.svelte:35 ~ SearchedString:", SearchedString)

	onMount(async () => {
		// QuestionList = await fetchQuesArrWithMetadata(
		// 	filterType,
		// 	pageNumNow*contentPerPage,
		// 	pageNumNow*contentPerPage + contentPerPage,
		// 	filterState ? -1 : 1
		// 	);
		// 	console.log("🚀 ~ file: PublicQues.svelte:32 ~ onMount ~ QuestionList:", QuestionList)
		// CalculatePages();
		// pageNumStart = 0;
		// pageNumNow = pageNumNow*contentPerPage;
		// pageNumEnd = Math.ceil(QuestionList.Metadata.Length / contentPerPage);
		Loading=true
	});
	async function ApplyFilter(filterType1: string, filterState1: boolean) {
		filterType=filterType1
		filterState=filterState1
		// time , views , unanswered , votes
		// if (filterState) {
		// 	QuestionList = await fetchQuesArrWithMetadata(
		// 		filterType,
		// 		pageNumNow*contentPerPage,
		// 		pageNumNow*contentPerPage + contentPerPage,
		// 		-1
		// 	);
		// } else {
		// 	QuestionList = await fetchQuesArrWithMetadata(filterType, pageNumNow*contentPerPage, pageNumNow*contentPerPage + contentPerPage, 1);
		// }
		// console.log("🚀 ~ file: PublicQues.svelte:59 ~ ApplyFilter ~ QuestionList:", QuestionList)
		
		// pageNumStart = 0;
	}

	


	// $: pageNumNow , ApplyFilter(filterType, filterState)

	let activeField = 0;
	function SelectActiveField(num: number) {
		activeField = num;
	}
</script>

{#if Loading}

<div class="flex h-16 w-full flex-row">
	<!-- <div class="ml-5 self-center text-lg text-[#e7e9eb]">{QuestionList.length} Question</div> -->
	<div class=" ml-2 flex h-14 flex-row items-center gap-2">
		<button
			on:click={() => {
				SelectActiveField(0);
			}}
			class=" h-8 w-fit rounded-xl {  activeField===0 ? "bg-blue-500 bg-opacity-20" :"" } px-2 py-1 text-blue-500"
			>Question</button
		>
		<button
			on:click={() => {
				SelectActiveField(1);
			}}
			class=" h-8 w-fit rounded-xl {  activeField===1 ? "bg-blue-500 bg-opacity-20" :"" } bg-[#272727] px-2 py-1 text-blue-500">Hash</button
		>
		<button
			on:click={() => {
				SelectActiveField(2);
			}}
			class=" h-8 w-fit rounded-xl {  activeField===2 ? "bg-blue-500 bg-opacity-20" :"" } bg-[#272727] px-2 py-1 text-blue-500">Blog</button
		>
	</div>

	<div class="grow" />
	<div class=" flex h-9 flex-row justify-around self-center rounded-md border-2 border-[#7d858d]">
		<div
			on:click={() => {filterType="votes"}}
			on:keypress={() => {}}
			class="flex h-full items-center justify-center border-r-2 border-[#7d858d] px-2 text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] {filterType ===
			'votes'
				? 'bg-[#404245]'
				: ''}"
		>
			Rated
		</div>
		<div
			on:click={() => { filterType="views"}}
			on:keypress={() => {}}
			class="flex h-full items-center justify-center border-r-2 border-[#7d858d] px-2 text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] {filterType ===
			'views'
				? 'bg-[#404245]'
				: ''}"
		>
			Popular
		</div>
		<div
			on:click={() => {filterType="time"}}
			on:keypress={() => {}}
			class="flex h-full items-center justify-center border-r-2 border-[#7d858d] px-2 text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] {filterType ===
			'time'
				? 'bg-[#404245]'
				: ''}  "
		>
			Time
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
			// ApplyFilter(filterType, filterState, filterIndex);
			// CalculatePages();
		}}
		on:keydown={() => {}}
		class=" ml-4 flex h-9 w-24 flex-row items-center justify-center self-center rounded-md border-2 border-[#7d858d] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245]"
	>
		<ZtoA class=" {filterState ? ' rotate-180' : ''}" />
		<div class="ml-1 font-semibold text-[#9cc1db]">order</div>
	</div>
</div>
<div class="h-full w-full">
	{#if activeField === 0}
		<SearchedQues     bind:filterState={filterState} bind:filterType={filterType}   />
	{:else if activeField === 1}
		<SearchedHash  bind:filterType={filterType}  bind:filterState={filterState}   />
	{:else if activeField === 2}
		<SearchedBlog  bind:filterType={filterType}  bind:filterState={filterState}  />
	{:else}
		<div class=" flex h-full w-full items-center justify-center text-xl text-white">
			<div class="flex flex-col items-center justify-center">
				<Emptybox class="h-20 " />
				<p class=" font-raleway">Nothing Found</p>
			</div>
		</div>
	{/if}

</div>

{:else}
	 <!-- else content here -->
	 <div class=" flex h-full w-full items-center justify-center text-xl text-white">
		<div class="flex flex-col items-center justify-center">
			<Emptybox class="h-20 " />
			<p class=" font-raleway">Loading</p>
		</div>
	</div>
{/if}