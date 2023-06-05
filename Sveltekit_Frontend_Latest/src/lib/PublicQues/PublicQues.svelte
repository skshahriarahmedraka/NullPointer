<script lang="ts">
	import { goto } from '$app/navigation';
	import { fetchPublicQuestionDataArr, fetchQuesArrWithMetadata } from '$lib/store/fetch';
	import type { QuesArrWithMetadataType, QuestionDataType } from '$lib/store/types';
	import { fetchUserFlairData } from '$lib/store/fetch';
	import { onMount } from 'svelte';
	import Emptybox from '$lib/errors/emptybox.svelte';

	

	let QuestionList:QuesArrWithMetadataType = {} as QuesArrWithMetadataType;
	// let fetchData = async () => {

	// }
	// fetchData()
	console.log('🚀 ~ file: index.svelte:126 ~ QuestionList:', QuestionList);

	import ZtoA from '$lib/PublicQues/svgs/ZtoA.svelte';
	import PageNum from '$lib/PageNum/pageNum.svelte';

	// import { goto } from "$app/navigation";
	let filterState: boolean = true;
	// GetEditedByData(QuestionData.QuesEditedBy);
	let filterType: string = 'time';
	// let filterIndex: number = 0;
	let pageNumStart: number=0;
	let pageNumNow: number =0;
	let pageNumEnd: number;
	let contentPerPage:number=5
	let Loading:boolean=false
	onMount(async () => {
		QuestionList = await fetchQuesArrWithMetadata(
			filterType,
			pageNumNow*contentPerPage,
			pageNumNow*contentPerPage + contentPerPage,
			filterState ? -1 : 1
			);
			console.log("🚀 ~ file: PublicQues.svelte:32 ~ onMount ~ QuestionList:", QuestionList)
		// CalculatePages();
		pageNumStart = 0;
		pageNumNow = pageNumNow*contentPerPage;
		pageNumEnd = Math.ceil(QuestionList.Metadata.Length / contentPerPage);
		Loading=true
	});
	async function ApplyFilter(filterType: string, filterState: boolean) {
		// time , views , unanswered , votes
		if (filterState) {
			QuestionList = await fetchQuesArrWithMetadata(
				filterType,
				pageNumNow*contentPerPage,
				pageNumNow*contentPerPage + contentPerPage,
				-1
			);
		} else {
			QuestionList = await fetchQuesArrWithMetadata(filterType, pageNumNow*contentPerPage, pageNumNow*contentPerPage + contentPerPage, 1);
		}
		
		pageNumStart = 0;
		pageNumNow = pageNumNow*contentPerPage;
		pageNumEnd = Math.ceil(QuestionList.Metadata.Length / contentPerPage);
	}

	function CalculatePages() {
		pageNumStart = 0;
		pageNumNow = pageNumNow*contentPerPage;
		pageNumEnd = Math.ceil(QuestionList.Metadata.Length / contentPerPage);
	}
	$: pageNumNow , ApplyFilter(filterType, filterState)
</script>

{#if Loading}
	 <!-- content here -->

<div class="flex h-16 w-full flex-row">
	<div class="ml-5 self-center text-lg text-[#e7e9eb]">{QuestionList.Metadata.Length} Question</div>
	<div class="grow" />
	<div class=" flex h-9 flex-row justify-around self-center rounded-md border-2 border-[#7d858d]">
		<div
			on:click={() => {
				ApplyFilter('time', filterState);
				CalculatePages();
			}}
			on:keypress={() => {}}
			class="flex h-full items-center justify-center border-r-2 border-[#7d858d] px-2 text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] {filterType ===
			'time'
				? 'bg-[#404245]'
				: ''}  "
		>
			Time
		</div>
		<div
			on:click={() => {
				ApplyFilter('views', filterState);
				CalculatePages();
			}}
			on:keypress={() => {}}
			class="flex h-full items-center justify-center border-r-2 border-[#7d858d] px-2 text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] {filterType ===
			'views'
				? 'bg-[#404245]'
				: ''}"
		>
			Popular
		</div>
		<div
			on:click={() => {
				ApplyFilter('votes', filterState);
				CalculatePages();
			}}
			on:keypress={() => {}}
			class="flex h-full items-center justify-center border-r-2 border-[#7d858d] px-2 text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] {filterType ===
			'votes'
				? 'bg-[#404245]'
				: ''}"
		>
			Rated
		</div>
		<div
			on:click={() => {
				ApplyFilter('unanswered', filterState);
				CalculatePages();
			}}
			on:keypress={() => {}}
			class="flex h-full items-center justify-center border-r-0 border-[#7d858d] px-2 text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] {filterType ===
			'unanswered'
				? 'bg-[#404245]'
				: ''}"
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
			ApplyFilter(filterType, filterState);
			CalculatePages();
		}}
		on:keydown={() => {}}
		class=" ml-4 flex h-9 w-24 flex-row items-center justify-center self-center rounded-md border-2 border-[#7d858d] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245]"
	>
		<ZtoA class=" {filterState ? ' rotate-180' : ''}" />
		<div class="ml-1 font-semibold text-[#9cc1db]">order</div>
	</div>
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
					<div class="  {i.QuesUpvote - i.QuesDownvote > 0 ? 'text-[#e7e9eb]' : 'text-[#959ba0]'} ">
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
					<div class=" line-clamp-2 text-xl text-sky-600 hover:cursor-pointer hover:text-blue-600">
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
		<div class=" flex h-full w-full items-center justify-center text-xl text-white">
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
			<p class=" font-raleway">Loading</p>
		</div>
	</div>
{/if}
<style>
	/* your styles go here */
</style>
