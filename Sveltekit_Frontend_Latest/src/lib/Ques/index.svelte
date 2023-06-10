<script lang="ts">
	// import {QuestionData} from "../store/store"
	import RelatedQues from '$lib/RelatedQues/index.svelte';
	import { marked } from 'marked';

	import Ans from '$lib/Ans/index.svelte';
	import MarkDownWriter from '$lib/Write/index.svelte';
	import AvatarDefault from '$lib/icons/avatarDefault.svg';
	import { fetchAnsData, fetchPostAnsData, fetchUserFlairData } from '$lib/store/fetch';
	import type { AnswerDataType, QuestionDataType, UserFlairDataType } from '$lib/store/types';
	import { onMount } from 'svelte';
	import Down from './svgs/Down.svelte';
	import Up from './svgs/Up.svelte';
	import Share from './svgs/share.svelte';
	import Edit from './svgs/edit.svelte';
	import Follow from './svgs/follow.svelte';
	import RoundDot from './svgs/RoundDot.svelte';
	import Filter from './svgs/Filter.svelte';
	import Book from '$lib/Loading/book.svelte';
	import { UserData } from '$lib/store/store';
	import WriteAns from '$lib/WriteAns/index.svelte';
	import Emptybox from '$lib/errors/emptybox.svelte';
	import { goto } from '$app/navigation';

	export let QuestionData: QuestionDataType;
	// export let RelatedQuestionList:any
	// export let RelatedQuestionListLoading:boolean
	// export let QuestionData:any
	let DropDownData = {
		Show: false,
		'a-z': true,
		List: ['Newest', 'Recent Activity', 'Highest score', 'Most frequent', 'bounty']
	};
	function DropDownClick() {
		DropDownData['Show'] = !DropDownData['Show'];
	}
	function DropDownBlur() {
		if (DropDownData['Show']) {
			DropDownData['Show'] = false;
		}
	}
	let writeAns = false;
	function RoundNum(x: number) {
		if (x === null) {
			return 0;
		} else if (x >= 1000000000) {
			x /= 1000000000;
			return String(x.toFixed(2)) + 'B';
		} else if (x >= 1000000) {
			x /= 1000000;
			return String(x.toFixed(2)) + 'M';
		} else if (x >= 1000) {
			x /= 1000;
			return String(x.toFixed(2)) + 'K';
		} else {
			return String(x);
		}
	}

	let AskedBy: UserFlairDataType;
	let EditedBy: UserFlairDataType;

	// async function GetAskbyData(ID: string) {
	// }
	// async function GetEditedByData(ID: string) {
	// }

	let Loading = false;
	onMount(async () => {
		if (parseInt(QuestionData.QuesAskedBy) != 0) {
			AskedBy = await fetchUserFlairData(QuestionData.QuesAskedBy);
			// GetAskbyData(QuestionData.QuesAskedBy);
		}
		if (parseInt(QuestionData.QuesEditedBy) != 0) {
			EditedBy = await fetchUserFlairData(QuestionData.QuesEditedBy);
			// GetEditedByData(QuestionData.QuesEditedBy);
		}
		Loading = true;
	});
	//     var isoTimestamp = new Date().toISOString();
	// var localTimestamp = new Date(isoTimestamp).toLocaleString();

	// console.log(localTimestamp);
	let UserWrittenAns: AnswerDataType = {
		ID: '',
		QuesID: QuestionData.ID,
		AnsweredTime: new Date().toISOString(),
		EditedTime: new Date().toISOString(),

		Upvote: 0,
		Downvote: 0,
		Votes: [],
		Bookmark: [],
		Accepted: false,

		AnsweredBy: $UserData.UserID,
		EditedBy: '',

		Description: '',
		Comment: []
	};
	async function PostMyAns() {
		console.log('🚀 ~ file: index.svelte:97 ~ UserWrittenAns:', UserWrittenAns);
		let PostedAns: AnswerDataType = await fetchPostAnsData(QuestionData.ID, UserWrittenAns);
		console.log('🚀 ~ file: index.svelte:100 ~ PostMyAns ~ PostedAns:', PostedAns);
	}
</script>

{#if Loading}
	<div class=" mt-2 max-h-full min-h-screen w-[1100px] bg-[#2d2d2d] pl-5">
		<div class=" ml-3 mt-3 flex h-[150px] w-[1050px] flex-col text-[#e7e9eb]">
			<!-- QUESTION TITLE -->
			<div class=" flex flex-row">
				<p class=" line-clamp-3 basis-10/12 font-raleway text-2xl">{QuestionData.QuesTitle}</p>
				<div
				on:click={()=>{goto("/ask")}}
				on:keypress={()=>{}}
					class="ml-3 flex h-12 basis-2/12 items-center justify-center rounded-md bg-[#0964aa] hover:bg-blue-600 hover:cursor-pointer"
				>
					<p class="  my-auto text-xl font-semibold text-gray-200">Ask Question</p>
				</div>
			</div>
			<div class="ml-3 mt-2 flex flex-row space-x-3">
				<div class="">
					<p class="inline-flex text-[#959ba0]">Asked</p>
					{new Date(QuestionData.QuesAskedTime).toLocaleString()}
				</div>
				<div class="">
					<p class="inline-flex text-[#959ba0]">Modified</p>
					{new Date(QuestionData.QuesEditedTime).toLocaleString()}
				</div>
				<div class="">
					<p class="inline-flex text-[#959ba0]">Viewed</p>
					{QuestionData.QuesViewed}
				</div>
			</div>
		</div>
		<!-- HORIZONTAL BAR -->
		<div class=" ml-3 h-[1px] w-[1050px] overflow-hidden bg-stone-400" />
		<div class=" flex flex-row">
			<!-- QUESTION DESCRIPTION CONTAINER  -->
			<div class="w-[850px]">
				<!-- QUESTION DESCRIPTION -->
				<div class="">
					<div class=" my-3 flex w-full flex-row pt-2">
						<!-- Ratings -->
						<div class=" flex w-14 flex-col">
							<!-- UP -->
							<Up />
							<div
								class=" mx-2 text-center text-xl {QuestionData.QuesUpvote -
									QuestionData.QuesDownvote <
								0
									? ' text-red-500'
									: ' text-green-500'}"
							>
								{QuestionData.QuesUpvote - QuestionData.QuesDownvote}
							</div>
							<!-- DOWN -->
							<Down />
						</div>
						<!-- Question Detail  -->
						<div class="w-full font-sf-pro text-[#e7e9eb]">
							<div
								class=" prose mt-5 min-w-full max-w-full overflow-hidden break-words bg-inherit p-5 text-[#e7e9eb] prose-headings:text-[#e7e9eb] prose-p:text-[#e7e9eb] prose-a:text-blue-500 prose-blockquote:border-sky-400 prose-blockquote:text-[#e7e9eb] prose-figure:text-white prose-figcaption:text-[#e7e9eb] prose-strong:text-[#e7e9eb] prose-em:text-[#e7e9eb] prose-code:text-[#e7e9eb] prose-pre:text-[#e7e9eb] prose-ol:text-[#e7e9eb] prose-ul:text-[#e7e9eb] prose-li:text-[#e7e9eb] prose-li:marker:text-white prose-table:text-[#e7e9eb] prose-thead:text-[#e7e9eb] prose-tr:border-4 prose-tr:border-gray-300 prose-tr:text-[#e7e9eb] prose-th:border-2 prose-th:border-gray-300 prose-th:text-[#e7e9eb] prose-td:border-2 prose-td:border-gray-300 prose-td:text-[#e7e9eb] prose-img:text-[#e7e9eb] prose-video:text-[#e7e9eb] prose-hr:bg-gray-500 prose-hr:text-[#e7e9eb]"
							>
								{@html marked(QuestionData.QuesDescription)}
							</div>

							<!-- Question Tag -->
							<div class="  min-h-10 mt-2 flex max-h-20 w-full flex-row flex-wrap">
								{#each QuestionData.QuesTags as tag}
									<div
										class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
									>
										{tag}
									</div>
								{/each}
							</div>
							<!-- share edit askedBy modifiedBy -->
							<div class="mt-2 flex flex-row">
								<button class="" on:click={() => {}}>
									<Share />
								</button>
								<button class="" on:click={() => {}}>
									<Edit />
								</button>
								<button class="" on:click={() => {}}>
									<Follow />
								</button>
								<!-- Gap -->
								<div class="grow" />
								<!-- MODIFIED BY -->
								{#if parseInt(QuestionData.QuesEditedBy) != 0}
									<div class=" flex h-20 w-60 flex-col overflow-hidden px-2">
										<p class="text-sm text-white">
											Edited : {new Date(QuestionData.QuesEditedTime).toLocaleString()}
										</p>
										<div class="flex h-12 w-full flex-row">
											{#if EditedBy.UserImage != ''}
												<img
													src={EditedBy.UserImage}
													alt="Modified By UserImage"
													class=" aspect-square active:ring-offset-base-50 mt-1 h-10 w-10 cursor-pointer rounded-3xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500 active:rounded-md active:ring active:ring-blue-600"
												/>
											{:else}
												<img
													src={AvatarDefault}
													alt="Modified By UserImage"
													class=" aspect-square active:ring-offset-base-50 mt-1 h-10 w-10 cursor-pointer rounded-full object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500 active:rounded-md active:ring active:ring-blue-600"
												/>
											{/if}
											<div class=" flex flex-col">
												{#if EditedBy != null}
													<!-- content here -->
													<div class=" ml-2 line-clamp-1">{EditedBy.UserName}</div>
												{:else}
													<div class=" ml-2 line-clamp-1">{'Edited Name'}</div>
													<!-- else content here -->
												{/if}
												<!-- <div class=" ml-2 line-clamp-1">{EditedBy.UserName}</div> -->
												<div class=" ml-2 flex flex-row">
													{#if EditedBy.Badges.Reputation != 0}
														<RoundDot class="ml-1 mt-2 h-2 w-2 place-content-center fill-white" />
														<p class="mx-1 text-white">{RoundNum(EditedBy.Badges.Reputation)}</p>
													{/if}
													{#if EditedBy.Badges.Gold != 0}
														<RoundDot
															class="ml-1 mt-2 h-2 w-2 place-content-center fill-[#ffcc01]"
														/>
														<p class="mx-1 text-[#ffcc01]">{RoundNum(EditedBy.Badges.Gold)}</p>
													{/if}
													{#if EditedBy.Badges.Silver != 0}
														<RoundDot
															class="ml-1 mt-2 h-2 w-2 place-content-center fill-[#b4b8bc]"
														/>
														<p class="mx-1 text-[#b4b8bc]">{RoundNum(EditedBy.Badges.Silver)}</p>
													{/if}
													{#if EditedBy.Badges.Bronze != 0}
														<RoundDot
															class="ml-1 mt-2 h-2 w-2 place-content-center fill-[#d1a684]"
														/>
														<p class="mx-1 text-[#d1a684]">{RoundNum(EditedBy.Badges.Bronze)}</p>
													{/if}
												</div>
											</div>
										</div>
									</div>
								{/if}
								<!-- Asked BY -->
								<div class=" ml-2 flex h-20 w-56 flex-col overflow-hidden px-2">
									<p class="text-sm text-white">
										Asked : {new Date(QuestionData.QuesAskedTime).toLocaleString()}
									</p>
									<div class="flex h-12 w-full flex-row">
										{#if AskedBy.UserImage != ''}
											<!-- svelte-ignore a11y-img-redundant-alt -->
											<img
												src={AskedBy.UserImage}
												alt="Answered By Image"
												class=" aspect-square active:ring-offset-base-50 mt-1 h-10 w-10 cursor-pointer rounded-3xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500 active:rounded-md active:ring active:ring-blue-600"
											/>
										{:else}
											<img
												src={AvatarDefault}
												alt="Modified By UserImage"
												class=" aspect-square active:ring-offset-base-50 mt-1 h-10 w-10 cursor-pointer rounded-full object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500 active:rounded-md active:ring active:ring-blue-600"
											/>
										{/if}
										<div class=" flex flex-col">
											{#if AskedBy != null}
												<!-- content here -->
												<div class=" ml-2 line-clamp-1">{AskedBy.UserName}</div>
											{:else}
												<div class=" ml-2 line-clamp-1">{'Asked Name'}</div>
												<!-- else content here -->
											{/if}
											<div class=" flex flex-row">
												{#if AskedBy.Badges.Gold != 0}
													<RoundDot class="ml-1 mt-2 h-2 w-2 place-content-center fill-[#ffcc01]" />

													<p class="mx-1 text-[#ffcc01]">{RoundNum(AskedBy.Badges.Gold)}</p>
												{/if}
												{#if AskedBy.Badges.Silver != 0}
													<RoundDot class="ml-1 mt-2 h-2 w-2 place-content-center fill-[#b4b8bc]" />

													<p class="mx-1 text-[#b4b8bc]">{RoundNum(AskedBy.Badges.Silver)}</p>
												{/if}
												{#if AskedBy.Badges.Bronze != 0}
													<RoundDot class="ml-1 mt-2 h-2 w-2 place-content-center fill-[#d1a684]" />

													<p class="mx-1 text-[#d1a684]">{RoundNum(AskedBy.Badges.Bronze)}</p>
												{/if}
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
					<!-- Number of ANSWER -->
					<div
						class=" m-4 mr-2 flex h-14 flex-row place-content-center place-items-center rounded-lg border-[1px] border-gray-600 text-2xl text-[#e7e9eb]"
					>
						<div class="ml-3">{QuestionData.Answers.length} Answers</div>
						<!-- write answer button -->
						<div
							on:click={() => {
								writeAns = !writeAns;
							}}
							on:keypress={() => {}}
							class="ml-3 cursor-pointer rounded-md border-2 border-gray-600 bg-gray-500 bg-opacity-25 px-2 hover:bg-gray-800"
						>
							Write Answer
						</div>
						<div class="grow" />
						<!-- drop down -->
						<div class="m-2 mr-4">
							<div class=" relative inline-block">
								<button
									on:blur={DropDownBlur}
									on:click={DropDownClick}
									class="inline-flex items-center rounded border-2 border-[#688fac] p-1 pl-2 text-lg font-medium text-[#e7e9eb]"
								>
									<Filter class="mr-2 h-6 w-6 fill-[#688fac]" />
									<span class="mr-1">Most Populer</span>
								</button>
								{#if DropDownData['Show']}
									<ul
										class=" absolute w-full rounded border-2 border-[#688fac] bg-slate-700 text-lg font-medium text-[#e7e9eb]"
									>
										<li class=" p-1 pl-2 hover:bg-slate-500">Most Populer</li>
										<li class="p-1 pl-2 hover:bg-slate-500">Most Shared</li>
										<li class="p-1 pl-2 hover:bg-slate-500">Most commented</li>
									</ul>
								{/if}
							</div>
						</div>
						<!-- <div class="w-14 h-8 bg-slate-200 "></div> -->
					</div>
					<div class=" m-5 w-[95%]">
						{#if writeAns}
							<WriteAns
								bind:markdown={UserWrittenAns.Description}
								on:PostMyAns={PostMyAns}
								on:CancelAnswer={() => {
									UserWrittenAns.Description = '';
									writeAns = !writeAns;
								}}
							/>
						{/if}
					</div>
				</div>
				<!-- answers -->
				{#if QuestionData.Answers.length != 0}
					{#each QuestionData.Answers as ans}
						{#await fetchAnsData(ans.ID)}
							<p>...waiting</p>
						{:then AnsData}
							<Ans ans={AnsData} />
						{:catch error}
							<p style="color: red">{error.message}</p>
						{/await}
					{/each}
				{:else}
					<div class=" flex h-full w-full items-center justify-center text-xl text-white">
						<div class="flex flex-col items-center justify-center">
							<Emptybox class="h-20 " />
							<p class=" font-raleway">No Answer Yet</p>
						</div>
					</div>
				{/if}
			</div>

			<!--List of Answers -->
			<!-- RELATED QUESTION -->
			<!-- {#if true}
				<RelatedQues />
			{/if} -->
		</div>
	</div>
{:else}
	<div
		class=" mt-2 flex max-h-full min-h-screen w-[1100px] items-center justify-center bg-[#2d2d2d] pl-5"
	>
		<Book />
	</div>
{/if}

<style>
</style>
