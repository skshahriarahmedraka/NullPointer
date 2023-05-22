<script lang="ts">
	// import {QuestionData} from "../store/store"
	import { marked } from 'marked';

	import AvatarDefault from '$lib/icons/avatarDefault.svg';
	import type { AnswerDataType, UserFlairDataType } from '$lib/store/types';
	import { onMount } from 'svelte';
	import Down from './svgs/down.svelte';
	import Edit from './svgs/edit.svelte';
	import NetworkSignal from './svgs/networkSignal.svelte';
	import RoundDot from './svgs/roundDot.svelte';
	import Share from './svgs/share.svelte';
	import Up from './svgs/up.svelte';
	import Book from '$lib/Loading/book.svelte';
	import { fetchUserFlairData } from '$lib/store/fetch';

	let ans: AnswerDataType= {
		ID: "78667896",
		QuesID: "78667896",
		AnsweredTime: new Date().toISOString(),
		EditedTime: new Date().toISOString(),

		Upvote: 4,
		Downvote: 32,
		Bookmark: 23,
		Accepted: true,

		AnsweredBy: '46262554',
		EditedBy: '46262554',

		Description:
			"In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.    In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.    In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.   ",
		Comment: [
			{
				ID: '1',
				Upvote: 4,
				Downvote: 1,
				UserID: '154123',
				UserName: 'sk shahriar ahmed raka',
				CommentTime: new Date().toISOString(),
				Comment:
					'no Comment bossNote: This question and most answers seem to have been written before append() came into the language, which is a good solution for this.'
			}
		]
	};
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
	let AnsGivenBy: UserFlairDataType = {
		ID: '56654653',
		UserID: 'skraka',
		UserName: ' sheikh Ahmed Raka',

		UserImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
		Badges: { Reputation: 681385285, Gold: 8893785, Silver: 646234, Bronze: 77455345 },
		Location: '',
		Aboutme: ''
	};

	// let AskedBy: UserFlairDataType;
	let EditedBy: UserFlairDataType = {
		ID: '56654653',
		UserID: 'skraka',
		UserName: ' sheikh Ahmed Raka',

		UserImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
		Badges: { Reputation: 681385285, Gold: 8893785, Silver: 646234, Bronze: 77455345 },
		Location: '',
		Aboutme: ''
	};

	// async function GetAskbyData(ID: string) {
	// }
	// async function GetEditedByData(ID: string) {
	// }

	let Loading = false;
	onMount(async () => {
		if (ans.AnsweredBy != '') {
			AnsGivenBy = await fetchUserFlairData(ans.AnsweredBy);
			// GetAskbyData(QuestionData.QuesAskedBy);
		}
		if (ans.EditedBy != '') {
			EditedBy = await fetchUserFlairData(ans.EditedBy);
			// GetEditedByData(QuestionData.QuesEditedBy);
		}
		Loading = true;
	});
</script>

{#if Loading}
	<!-- content here -->

	<div class=" mb-5 mr-[300px] mt-5 flex w-full flex-row pr-3 text-[#e7e9eb]">
		<!-- Ratings -->
		<div class=" flex w-14 flex-col">
			<!-- UP -->
			<Up />
			<div
				class=" mx-2 text-center text-xl {ans.Upvote - ans.Downvote < 0
					? ' text-red-500'
					: ' text-green-500'}"
			>
				{RoundNum(ans.Upvote - ans.Downvote)}
			</div>
			<!-- DOWN -->
			<Down />
		</div>
		<div class="flex flex-col">
			<!-- ANSWER -->
			<div
				class=" prose mt-5 min-w-full max-w-full overflow-hidden break-words bg-inherit p-5 text-[#e7e9eb] prose-headings:text-[#e7e9eb] prose-p:text-[#e7e9eb] prose-a:text-blue-500 prose-blockquote:border-sky-400 prose-blockquote:text-[#e7e9eb] prose-figure:text-white prose-figcaption:text-[#e7e9eb] prose-strong:text-[#e7e9eb] prose-em:text-[#e7e9eb] prose-code:text-[#e7e9eb] prose-pre:text-[#e7e9eb] prose-ol:text-[#e7e9eb] prose-ul:text-[#e7e9eb] prose-li:text-[#e7e9eb] prose-li:marker:text-white prose-table:text-[#e7e9eb] prose-thead:text-[#e7e9eb] prose-tr:border-4 prose-tr:border-gray-300 prose-tr:text-[#e7e9eb] prose-th:border-2 prose-th:border-gray-300 prose-th:text-[#e7e9eb] prose-td:border-2 prose-td:border-gray-300 prose-td:text-[#e7e9eb] prose-img:text-[#e7e9eb] prose-video:text-[#e7e9eb] prose-hr:bg-gray-500 prose-hr:text-[#e7e9eb]"
			>
				{@html marked(ans.Description)}
			</div>
			<!-- share edit edited -->
			<div class="mt-2 flex flex-row">
				<Share />
				<Edit />
				<NetworkSignal />
				<!-- Gap -->
				<div class="grow" />
				<!-- MODIFIED BY -->
				{#if ans.EditedBy != ''}
					<!-- content here -->
					<div class=" flex h-20 w-60 flex-col overflow-hidden px-2">
						<p class="text-sm text-white">Edited : {ans.EditedBy}</p>
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
								<!-- else content here -->
							{/if}
							<div class=" flex flex-col">
								<div class=" ml-2 line-clamp-1">{EditedBy.UserName}</div>
								<div class=" ml-2 flex flex-row">
									{#if EditedBy.Badges.Reputation != 0}
										<RoundDot class="ml-1 mt-2 h-2 w-2 place-content-center  fill-white " />
										<p class="mx-1 text-white">{RoundNum(EditedBy.Badges.Reputation)}</p>
									{/if}
									{#if EditedBy.Badges.Gold != 0}
										<RoundDot class="ml-1 mt-2 h-2 w-2 place-content-center  fill-[#ffcc01] " />
										<p class="mx-1 text-[#ffcc01]">{RoundNum(EditedBy.Badges.Gold)}</p>
									{/if}
									{#if EditedBy.Badges.Silver != 0}
										<RoundDot class="ml-1 mt-2 h-2 w-2  place-content-center fill-[#b4b8bc] " />
										<p class="mx-1 text-[#b4b8bc]">{RoundNum(EditedBy.Badges.Silver)}</p>
									{/if}
									{#if EditedBy.Badges.Bronze != 0}
										<RoundDot class="ml-1 mt-2 h-2 w-2 place-content-center fill-[#d1a684] " />
										<p class="mx-1 text-[#d1a684]">{RoundNum(EditedBy.Badges.Bronze)}</p>
									{/if}
								</div>
							</div>
						</div>
					</div>
				{/if}
				<!-- ANSWERED BY -->
				<div class=" ml-2 flex h-20 w-56 flex-col overflow-hidden px-2">
					<p class="text-sm text-white">Answered : {ans.AnsweredTime}</p>
					<div class="flex h-12 w-full flex-row">
						{#if AnsGivenBy.UserImage != ''}
							<!-- svelte-ignore a11y-img-redundant-alt -->
							<img
								src={AnsGivenBy.UserImage}
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
							<div class=" ml-2 line-clamp-1">{AnsGivenBy.UserName}</div>
							<div class=" flex flex-row">
								{#if AnsGivenBy.Badges.Gold != 0}
									<RoundDot class="ml-1 mt-2 h-2 w-2 place-content-center  fill-[#ffcc01] " />

									<p class="mx-1 text-[#ffcc01]">{RoundNum(AnsGivenBy.Badges.Gold)}</p>
								{/if}
								{#if AnsGivenBy.Badges.Silver != 0}
									<RoundDot class="ml-1 mt-2 h-2 w-2  place-content-center fill-[#b4b8bc] " />

									<p class="mx-1 text-[#b4b8bc]">{RoundNum(AnsGivenBy.Badges.Silver)}</p>
								{/if}
								{#if AnsGivenBy.Badges.Bronze != 0}
									<RoundDot class="ml-1 mt-2 h-2 w-2 place-content-center fill-[#d1a684] " />

									<p class="mx-1 text-[#d1a684]">{RoundNum(AnsGivenBy.Badges.Bronze)}</p>
								{/if}
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
{:else}
	<div class=" mb-5 mr-[300px] mt-5 flex w-full flex-row pr-3 text-[#e7e9eb]">
		<Book />
	</div>
{/if}
