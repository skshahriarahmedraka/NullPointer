<script lang="ts">
	// import { SearchedString } from '$lib/store/store';
	import { goto } from '$app/navigation';
	import { fetchPublicQuestionDataArr, fetchQuesArrWithMetadata, fetchSearchQuesArrWithMetadata } from '$lib/store/fetch';
	import type { QuesArrWithMetadataType, QuestionDataType } from '$lib/store/types';
	import { fetchUserFlairData } from '$lib/store/fetch';
	import { onMount } from 'svelte';
	import Emptybox from '$lib/errors/emptybox.svelte';
	let QuestionList: QuesArrWithMetadataType = {} as QuesArrWithMetadataType;
	// let fetchData = async () => {

	// }
	// fetchData()
	console.log('🚀 ~ file: index.svelte:126 ~ QuestionList:', QuestionList);

	import ZtoA from '$lib/PublicQues/svgs/ZtoA.svelte';
	import PageNum from '$lib/PageNum/pageNum.svelte';
	import { StoredSearchedString } from '$lib/store/store';

	// import { goto } from "$app/navigation";
	export let filterState: boolean ;
	// GetEditedByData(QuestionData.QuesEditedBy);
	export let filterType: string;
	// export let SearchedString: string 
	// console.log("🚀 ~ file: searchedQues.svelte:24 ~ SearchedString:", SearchedString)
	let SearchedString= ''

	

	// let filterIndex: number = 0;
	 let pageNumStart: number = 0;
	let pageNumNow: number = 0;
 let pageNumEnd: number=0;
	$: console.log("pageNumNow:", pageNumNow)
	$: console.log("pageNumStart:", pageNumStart)
	$: console.log("pageNumEnd:", pageNumEnd)
	let contentPerPage: number = 5;
	let Loading: boolean = false;
	// let filterIndex: number = pageNumNow *20;
	onMount(async () => {
		filterType= 'time'
		filterState = true
		StoredSearchedString.subscribe((value) => {
			console.log("🚀 ~ file: searchedQues.svelte:44 ~ StoredSearchedString.subscribe ~ value:", value)
			
			SearchedString = value;
		});
		QuestionList = await fetchSearchQuesArrWithMetadata(
			SearchedString,
			filterType,
			pageNumNow * contentPerPage,
			pageNumNow * contentPerPage + contentPerPage,
			filterState ? -1 : 1
			);
			console.log("🚀 ~ file: searchedQues.svelte:48 ~ onMount ~ QuestionList:", QuestionList)
		pageNumStart = 0;
		// pageNumNow = pageNumNow * contentPerPage;
		pageNumEnd = Math.floor(QuestionList.Metadata.Length / contentPerPage);
		Loading = true;
	});
	async function ApplyFilter(filterType: string, filterState: boolean) {
		// time , views , unanswered , votes
		// let SearchedString= ''
		StoredSearchedString.subscribe((value) => {
			console.log("🚀 ~ file: searchedQues.svelte:65 ~ StoredSearchedString.subscribe ~ value:", value)
			SearchedString = value;
		});
		
		if (filterState ) {
			QuestionList = await fetchSearchQuesArrWithMetadata(
				SearchedString,
				filterType,
				pageNumNow * contentPerPage,
				pageNumNow * contentPerPage + contentPerPage,
				-1
			);
		} else {
			QuestionList = await fetchSearchQuesArrWithMetadata(
				SearchedString,
				filterType,
				pageNumNow * contentPerPage,
				pageNumNow * contentPerPage + contentPerPage,
				1
			);
		}

		pageNumStart = 0;
		// pageNumNow = pageNumNow * contentPerPage;
		pageNumEnd = Math.floor(QuestionList.Metadata.Length / contentPerPage);
	}

	// function CalculatePages() {
	// 	pageNumStart = 0;
	// 	pageNumNow = pageNumNow * contentPerPage;
	// 	pageNumEnd = Math.ceil(QuestionList.Metadata.Length / contentPerPage);
	// }
	$: pageNumNow, ApplyFilter(filterType, filterState);

	// let tempstr = ''
	// StoredSearchedString.subscribe((value) => {
	// 	console.log("🚀 ~ file: searchedQues.svelte:30 ~ StoredSearchedString.subscribe ~ value:", value)
	// 	// console.log('🚀 ~ file: PublicQues.svelte:32 ~ onMount ~ value:', value);
	// 	tempstr = value;
	// });
	// if (tempstr != SearchedString) {
	// 	ApplyFilter(filterType, filterState)
	// }
	$: $StoredSearchedString, ApplyFilter(filterType, filterState)
</script>

{#if Loading}
<div class="flex h-16 w-full flex-row">
	<div class="ml-5 self-center text-lg text-[#e7e9eb]">{QuestionList.Metadata.Length} Question</div>
	<div class="grow" />
</div>
	<div class="h-full w-full">
		{#if QuestionList.QuesList.length != 0}
			{#each QuestionList.QuesList as i}
				<div
					on:click={() => {
						goto(`/q/${i.ID}`);
					}}
					on:keydown={() => {}}
					class=" min-h-32 mx-2 flex max-h-40 w-full cursor-pointer flex-row border-t-2 border-[#404245] py-2 hover:bg-slate-800 hover:bg-opacity-50"
				>
					<!-- quesVote -->
					<div class="my-2 flex h-full w-60 flex-col items-center justify-center gap-1">
						<!-- VoteNumber -->
						<div
							class="  {i.QuesUpvote - i.QuesDownvote > 0 ? 'text-[#e7e9eb]' : 'text-[#959ba0]'} "
						>
							{i.QuesUpvote - i.QuesDownvote} Votes
						</div>
						<!-- Answer Number -->
						{#if i.Answers.length != 0}
							<!-- acceped ans -->
							<div class="rounded-md border-2 border-green-500 px-2 text-[#e7e9eb]">
								{i.Answers.length} Answers
							</div>
						{:else if i.QuesAnsAccepted != ''}
							<!-- Many ans but No accepted -->
							<div class=" rounded-md bg-green-500 px-2 text-white">{i.Answers.length} Answers</div>
						{:else}
							<!-- No answer -->
							<div class="text-[#959ba0]">{i.Answers.length} Answers</div>
						{/if}
						<!-- View Number -->
						<div class="{i.QuesViewed > 0 ? 'text-[#e7e9eb]' : 'text-[#959ba0]'} ">
							{i.QuesViewed} Views
						</div>
					</div>
					<!-- Qtitle &detail -->
					<div class=" flex h-full w-[1200px] flex-col">
						<!-- Title -->
						<div
							class=" line-clamp-2 text-xl text-sky-600 hover:cursor-pointer hover:text-blue-600"
						>
							{i.QuesTitle}
						</div>
						<!-- description -->
						<div class=" text-md line-clamp-2 text-[#e7e9eb]">{i.QuesDescription}</div>
						<!-- tag ans user -->
						<div class="flex flex-row items-center justify-center">
							<div class="  min-h-10 mt-2 flex max-h-20 w-full flex-row flex-wrap">
								{#if i.QuesTags.length != 0}
									{#each i.QuesTags as tag}
										<div
											class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
										>
											{tag}
										</div>
									{/each}
								{:else}
									<div
										class="m-1 rounded-md bg-[#3d4951] bg-opacity-70 px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
									>
										{'No Tag'}
									</div>
								{/if}
							</div>
							<div class="" />
							{#await fetchUserFlairData(i.QuesAskedBy)}
								<p class=" text-white">...waiting</p>
							{:then flairUserData}
								<!-- <p>The number is {number}</p> -->
								<p
									class="mr-1 mt-2 line-clamp-1 h-6 w-36 text-center text-sky-600 hover:cursor-pointer hover:text-blue-600"
								>
									{flairUserData.UserName}
								</p>
								<p class="mx-2 mt-2 font-bold text-[#e7e9eb]">{flairUserData.Badges.Reputation}</p>
							{:catch error}
								<p style="color: red">{error.message}</p>
							{/await}
							<p class=" mt-2 w-80 text-[#959ba0]">{new Date(i.QuesAskedTime).toLocaleString()}</p>
						</div>
					</div>
				</div>
			{/each}
		{:else}
			<!-- else content here -->
			<div class=" flex h-full w-full items-center justify-center text-xl text-white ">
				<div class="flex flex-col items-center justify-center">
					<Emptybox class="h-20 " />
					<p class=" font-raleway">Nothing Found</p>
				</div>
			</div>
		{/if}
	</div>
	<PageNum bind:pageNumStart={pageNumStart} bind:pageNumNow={pageNumNow} bind:pageNumEnd={pageNumEnd} />
{:else}
	<!-- else content here -->
	<div class=" flex h-full w-full items-center justify-center text-xl text-white">
		<div class="flex flex-col items-center justify-center">
			<Emptybox class="h-20 " />
			<p class=" font-raleway">error</p>
		</div>
	</div>
{/if}

<style>
	/* your styles go here */
</style>
