<script lang="ts">
	// import { Following } from '$lib/Navbar/svgs/following.svelte';
	import type { HashDataType } from '$lib/store/types';
	import { onMount } from 'svelte';
	import HashAbout from './component/HashAbout.svelte';
	import HashBlog from './component/HashBlog.svelte';
	import HashQues from './component/HashQues.svelte';

	// your script goes here
	export let HashData: HashDataType = {} as HashDataType;
	onMount(async () => {});

	let seletedTab = 'questions';
	let FollowingStatus = false;
	// let HashData: HashDataType= {
	//     "ID": "1",
	//     "Name": "AWS Collective",
	//     "MetaTitle": "Amazon Web Services (AWS) is the world’s most comprehensive and broadly adopted cloud platform, offering over 200 fully featured services from data centers globally. The AWS Collective is a community-driven site with resources for developers.",
	//     Image : "https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1920px-Go_Logo_Blue.svg.png",
	// BannerImage : "https://cdn.sstatic.net/Sites/stackoverflow/Img/subcommunities/aws-header.png?v=0925b21e52ea",
	// NumOfFollower : 13417878 ,
	// NumOfQuestion : 7846781789,
	// NumOfBlog : 148798,
	// NumOfAnswer : 8978147807,
	// About : ""
	// }
	const HeadLogo = new URL('../../static/favicon.png', import.meta.url).href;
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
</script>

<!-- <div class=" mt-2 flex max-h-full min-h-screen w-[1400px] flex-col gap-5 bg-transparent"> -->

<svelte:head>
	<title>{HashData.Name}</title>
	<link rel="icon" href={HeadLogo} />
</svelte:head>

<div class=" mb-2 mt-2 flex min-h-screen w-[1200px] flex-col gap-2 bg-[#2d2d2d]">
	<div class=" flex h-64 w-full items-center justify-center p-3">
		<img src={HashData.BannerImage} alt="" class=" h-full w-full object-cover" />
	</div>
	<div class="flex flex-row gap-2">
		<div class=" x mx-2 h-48 w-48">
			<img src={HashData.Image} alt="" class=" h-full w-full object-contain" />
		</div>
		<div class=" flex h-48 w-full flex-col justify-evenly gap-2">
			<p class=" font-sf-pro text-3xl text-[#e7e9eb]">{HashData.Name}</p>
			<div class=" flex h-full w-full flex-row">
				<div class=" flex h-full w-full flex-col justify-evenly">
					<p class=" font-poppins text-slate-300">
						{HashData.MetaTitle}
					</p>
					<div class=" flex flex-row gap-4">
						<p class="font-raleway text-[#c4c8cc]">{RoundNum(HashData.NumOfQuestion)} Question</p>
						<p class="font-raleway text-[#c4c8cc]">{RoundNum(HashData.NumOfAnswer)} Answers</p>
						<p class="font-raleway text-[#c4c8cc]">{RoundNum(HashData.NumOfFollower)} Followers</p>
					</div>
				</div>
				<div class="flex h-full w-36 items-center justify-center">
					{#if FollowingStatus}
						<button
							on:click={() => {
								FollowingStatus = !FollowingStatus;
							}}
							class=" h-10 w-28 rounded-3xl bg-blue-500 font-poppins font-semibold text-[#e7e9eb]"
							>Follow</button
						>
					{:else}
						<button
							on:click={() => {
								FollowingStatus = !FollowingStatus;
							}}
							class=" h-10 w-28 rounded-3xl bg-green-500 font-poppins font-semibold text-[#e7e9eb]"
							>Following</button
						>
					{/if}
				</div>
			</div>
		</div>
	</div>
	<hr class=" mx-2 bg-slate-400 opacity-30" />

	<div class=" flex h-20 w-full flex-row items-center justify-start gap-3">
		<button
			on:click={() => {
				seletedTab = 'questions';
			}}
			class=" rounded-2xl px-3 py-1 font-poppins text-base {seletedTab === 'questions'
				? 'bg-sky-500'
				: 'border-2 border-sky-500'} ml-3 text-[#e7e9eb]">Questions</button
		>
		<button
			on:click={() => {
				seletedTab = 'blogs';
			}}
			class=" rounded-2xl px-3 py-1 font-poppins text-base {seletedTab === 'blogs'
				? 'bg-sky-500'
				: 'border-2 border-sky-500'}  text-[#e7e9eb]">Blogs</button
		>
		<button
			on:click={() => {
				seletedTab = 'about';
			}}
			class=" rounded-2xl px-3 py-1 font-poppins text-base {seletedTab === 'about'
				? 'bg-sky-500'
				: 'border-2 border-sky-500'}  text-[#e7e9eb]">About</button
		>
		<!-- <button class=" px-3 py-1 font-poppins text-lg rounded-2xl { seletedTab==="questions" ? "bg-sky-500" : "border-2 border-sky-500" } text-[#e7e9eb]  ">Questions</button> -->
	</div>
	<hr class=" mx-2 bg-slate-400 opacity-30" />
	<div class=" h-fit min-h-[50vh] w-full">
		{#if seletedTab === 'questions'}
			<!-- content here -->
			<HashQues {HashData} />
		{:else if seletedTab === 'about'}
			<HashAbout {HashData} />
		{:else if seletedTab === 'blogs'}
			<HashBlog {HashData} />
		{/if}
	</div>
</div>

<style>
	/* your styles go here */
</style>
