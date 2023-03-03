<script lang="ts">
	// import {QuestionData} from "../store/store"

	import AvatarDefault from "$lib/icons/avatarDefault.svg";

	export let ans: {
		ID: number;
		AnsweredTime: string;
		Modified: string;
		Upvote: number;
		Downvote: number;
		Bookmark: number;
		Accepted: boolean;
		AnsweredBy: {
			UserID: number;
		};
		EditedBy: {
			UserID: number;
		};
		Description: string;
		Comment: {
			ID: number;
			Upvote: number;
			Downvote: number;
			UserID: number;
			UserName: string;
			Time: string;
			Comment: string;
		}[];
	};
	// let x = {
	// 	ID: 1,
	// 	AnsweredTime: ' Aug 7, 2011 at 2:14 ',
	// 	Modified: ' Aug 7, 2011 at 2:14 ',

	// 	Upvote: 4,
	// 	Downvote: 32,
	// 	Bookmark: 23,
	// 	Accepted: true,

	// 	AnsweredBy: { UserID: 154123 },
	// 	EditedBy: { UserID: 154123 },

	// 	Description:
	// 		"In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.    In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.    In Go, a string is a primitive type, which means it is read-only, and every manipulation of it will create a new string. So if I want to concatenate strings many times without knowing the length of the resulting string, what's the best way to do it? The naive way would be: var s string for i := 0; i < 1000; i++ { s += getShortStringFromSomewhere() } return s but that does not seem very efficient.   ",
	// 	Comment: [
	// 		{
	// 			ID: 1,
	// 			Upvote: 4,
	// 			Downvote: 1,
	// 			UserID: 154123,
	// 			UserName: 'sk shahriar ahmed raka',
	// 			Time: ' Aug 7, 2011 at 2:14',
	// 			Comment:
	// 				'no Comment bossNote: This question and most answers seem to have been written before append() came into the language, which is a good solution for this.'
	// 		}
	// 	]
	// };
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
	let AnsGivenBy = {
		UserID: 'skraka',
		UserName: 'Sk Shahriar Ahmed Raka',

		UserImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
		Badges: { Reputation: 66813, Gold: 999025, Silver: 8880324, Bronze: 77789423 }
	};
	let ModifiedBy = {
		UserID: 'skraka',
		UserName: ' sheikh Ahmed Raka',

		UserImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
		Badges: { Reputation: 681385285, Gold: 8893785, Silver: 646234, Bronze: 77455345 }
	};
</script>

<div class=" mt-5 mr-[300px] mb-5 flex w-full flex-row pr-3 text-[#e7e9eb] ">
	<!-- Ratings -->
	<div class=" flex w-14 flex-col ">
		<!-- UP -->
		<svg
			class=" mx-2 h-8 w-8 fill-[#696f75] hover:fill-gray-400"
			xmlns="http://www.w3.org/2000/svg"
			width="16"
			height="16"
			fill="currentColor"
			viewBox="0 0 16 16"
		>
			<path
				d="m7.247 4.86-4.796 5.481c-.566.647-.106 1.659.753 1.659h9.592a1 1 0 0 0 .753-1.659l-4.796-5.48a1 1 0 0 0-1.506 0z"
			/>
		</svg>
		<div
			class=" mx-2 text-center text-xl {ans.Upvote - ans.Downvote < 0
				? ' text-red-500'
				: ' text-green-500'}"
		>
			{RoundNum(ans.Upvote - ans.Downvote)}
		</div>
		<!-- DOWN -->
		<svg
			class=" mx-2 h-8 w-8 fill-[#696f75] hover:fill-gray-400"
			xmlns="http://www.w3.org/2000/svg"
			width="16"
			height="16"
			fill="currentColor"
			viewBox="0 0 16 16"
		>
			<path
				d="M7.247 11.14 2.451 5.658C1.885 5.013 2.345 4 3.204 4h9.592a1 1 0 0 1 .753 1.659l-4.796 5.48a1 1 0 0 1-1.506 0z"
			/>
		</svg>
	</div>
	<div class="flex  flex-col">
		<!-- ANSWER -->
		<div class="">
			<!-- {@html marked(markdown)} -->
			<!-- <SvelteMarkdown {source} /> -->
			{ans.Description}
		</div>
		<!-- share edit edited -->
		<div class="mt-2 flex flex-row  ">
			<svg
				class="m-1 h-8 w-8"
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M15 8a3 3 0 10-2.977-2.63l-4.94 2.47a3 3 0 100 4.319l4.94 2.47a3 3 0 10.895-1.789l-4.94-2.47a3.027 3.027 0 000-.74l4.94-2.47C13.456 7.68 14.19 8 15 8z"
				/></svg
			>
			<svg
				class="m-1 h-8 w-8"
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M17.414 2.586a2 2 0 00-2.828 0L7 10.172V13h2.828l7.586-7.586a2 2 0 000-2.828z"
				/><path
					fill-rule="evenodd"
					d="M2 6a2 2 0 012-2h4a1 1 0 010 2H4v10h10v-4a1 1 0 112 0v4a2 2 0 01-2 2H4a2 2 0 01-2-2V6z"
					clip-rule="evenodd"
				/></svg
			>
			<svg
				class="m-1 h-8 w-8"
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M5 3a1 1 0 000 2c5.523 0 10 4.477 10 10a1 1 0 102 0C17 8.373 11.627 3 5 3z"
				/><path
					d="M4 9a1 1 0 011-1 7 7 0 017 7 1 1 0 11-2 0 5 5 0 00-5-5 1 1 0 01-1-1zM3 15a2 2 0 114 0 2 2 0 01-4 0z"
				/></svg
			>
			<!-- Gap -->
			<div class="grow" />
			<!-- MODIFIED BY -->
			{#if ans.Modified != ''}
				<!-- content here -->
				<div class=" flex h-20  w-60 flex-col overflow-hidden px-2 ">
					<p class="text-sm text-white ">Edited : {ans.Modified}</p>
					<div class="flex  h-12 w-full flex-row  ">
                        {#if ModifiedBy.UserImage != "" }
                             <img
                                 src="{ModifiedBy.UserImage}"
                                 alt="Modified By UserImage"
                                 class=" aspect-square  active:ring-offset-base-50 mt-1 h-10  w-10 cursor-pointer rounded-3xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
                             />
                        {:else}
                        <img
                        src="{AvatarDefault}"
                        alt="Modified By UserImage"
                        class=" aspect-square  active:ring-offset-base-50 mt-1 h-10  w-10 cursor-pointer rounded-full object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
                    />
                             <!-- else content here -->
                        {/if}
						<div class=" flex flex-col ">
							<div class=" ml-2  line-clamp-1">{ModifiedBy.UserName}</div>
							<div class=" ml-2  flex flex-row ">
								{#if ModifiedBy.Badges.Reputation != 0}
									<svg
										class="mt-2 ml-1 h-2 w-2 place-content-center  fill-white  "
										viewBox="0 0 512 512"
										xmlns="http://www.w3.org/2000/svg"
										><path
											d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
										/></svg
									>
									<p class="mx-1   text-white">{ RoundNum(ModifiedBy.Badges.Reputation)}</p>
								{/if}
                                {#if ModifiedBy.Badges.Gold != 0}
									<svg
										class="mt-2 ml-1 h-2 w-2 place-content-center  fill-[#ffcc01]  "
										viewBox="0 0 512 512"
										xmlns="http://www.w3.org/2000/svg"
										><path
											d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
										/></svg
									>
									<p class="mx-1   text-[#ffcc01]">{ RoundNum(ModifiedBy.Badges.Gold)}</p>
								{/if}
								{#if  ModifiedBy.Badges.Silver != 0}
									<svg
										class="mt-2 ml-1 h-2 w-2  place-content-center fill-[#b4b8bc]  "
										viewBox="0 0 512 512"
										xmlns="http://www.w3.org/2000/svg"
										><path
											d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
										/></svg
									>
									<p class="mx-1  text-[#b4b8bc]">{ RoundNum(ModifiedBy.Badges.Silver)}</p>
								{/if}
								{#if  ModifiedBy.Badges.Bronze != 0}
									<svg
										class="mt-2 ml-1 h-2 w-2 place-content-center fill-[#d1a684]  "
										viewBox="0 0 512 512"
										xmlns="http://www.w3.org/2000/svg"
										><path
											d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
										/></svg
									>
									<p class="mx-1  text-[#d1a684]">{ RoundNum(ModifiedBy.Badges.Bronze)}</p>
								{/if}
							</div>
						</div>
					</div>
				</div>
			{/if}
			<!-- ANSWERED BY -->
			<div class=" ml-2 flex h-20  w-56 flex-col overflow-hidden px-2">
				<p class="text-sm text-white ">Answered : {ans.AnsweredTime}</p>
				<div class="flex h-12 w-full flex-row ">
                    {#if AnsGivenBy.UserImage != ""}
                    <!-- svelte-ignore a11y-img-redundant-alt -->
                         <img
                             src="{AnsGivenBy.UserImage}"
                             alt="Answered By Image"
                             class=" aspect-square active:ring-offset-base-50 mt-1 h-10  w-10 cursor-pointer rounded-3xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
                         />
                    {:else}
                    <img
                    src="{AvatarDefault}"
                    alt="Modified By UserImage"
                    class=" aspect-square  active:ring-offset-base-50 mt-1 h-10  w-10 cursor-pointer rounded-full object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
                />
                    {/if}
					<div class=" flex flex-col ">
						<div class=" ml-2 line-clamp-1">{AnsGivenBy.UserName}</div>
						<div class=" flex flex-row">
							{#if AnsGivenBy.Badges.Gold != 0}
								<svg
									class="mt-2 ml-1 h-2 w-2 place-content-center  fill-[#ffcc01]  "
									viewBox="0 0 512 512"
									xmlns="http://www.w3.org/2000/svg"
									><path
										d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
									/></svg
								>
								<p class="mx-1   text-[#ffcc01]">{ RoundNum(AnsGivenBy.Badges.Gold) }</p>
							{/if}
							{#if AnsGivenBy.Badges.Silver  != 0}
								<svg
									class="mt-2 ml-1 h-2 w-2  place-content-center fill-[#b4b8bc]  "
									viewBox="0 0 512 512"
									xmlns="http://www.w3.org/2000/svg"
									><path
										d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
									/></svg
								>
								<p class="mx-1  text-[#b4b8bc]">{ RoundNum(AnsGivenBy.Badges.Silver)}</p>
							{/if}
							{#if  AnsGivenBy.Badges.Bronze != 0}
								<svg
									class="mt-2 ml-1 h-2 w-2 place-content-center fill-[#d1a684]  "
									viewBox="0 0 512 512"
									xmlns="http://www.w3.org/2000/svg"
									><path
										d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
									/></svg
								>
								<p class="mx-1  text-[#d1a684]">{ RoundNum(AnsGivenBy.Badges.Bronze)}</p>
							{/if}
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
