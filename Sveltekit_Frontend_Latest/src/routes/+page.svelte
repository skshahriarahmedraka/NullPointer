<script lang="ts">
	import HotQuesCom from '$lib/HotQues/index.svelte';
	import BlogListCom from '$lib/BlogList/index.svelte';
	import PublicQues from '$lib/PublicQues/PublicQues.svelte';
	import Navbar from '$lib/Navbar/index.svelte';
	import Collectives from '$lib/Collectives/index.svelte';
	import LoadingSVG from '$lib/Loading/index.svelte';

	import { goto } from '$app/navigation';
	import {  onMount } from 'svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import PageNum from '$lib/PageNum/pageNum.svelte';
	// import type { PageData } from './$types';
	import { getCookieValue } from '$lib/store/utils';
	import type { CookieInfo1Type, UserDataType } from '$lib/store/types';
	import { fetchUserData } from '$lib/store/fetch';
	import { UserData } from '$lib/store/store';
	// import Filter from '$lib/SpaceCom/svgs/Filter.svelte';
	import QuesFilter from '$lib/QuesFilter/QuesFilter.svelte';
	// import { Buffer } from 'buffer';
	// export let data: PageData;
	// console.log('🚀 ~ file: +page.svelte:18 ~ data:', data);
	// const { InfoCookieData } = data;
	// console.log('🚀 ~ file: +page.svelte:19 ~ InfoCookieData:', InfoCookieData);

	// async function Necessary() {

	// }
	const HeadLogo = new URL("../../static/favicon.png", import.meta.url).href;

	let loadingState: boolean = false;
	onMount(async () => {
		let CookieValueInfo1: string = getCookieValue('Info1');
		console.log('🚀 ~ file: +page.svelte:30 ~ onMount ~ CookieValueInfo1:', CookieValueInfo1);
		let InfoCookieData = JSON.parse(atob(CookieValueInfo1)) as CookieInfo1Type;

		// let InfoCookieData = JSON.parse(Buffer.from(CookieValueInfo1, 'base64').toString('utf-8')) as CookieInfo1Type;
		console.log('🚀 ~ file: +page.ts:24 ~ load ~ InfoCookieData:', InfoCookieData);
		let UserDataValue = {} as UserDataType;
		UserData.subscribe((value) => {
			UserDataValue = value;
		});
		if (UserDataValue.ID != InfoCookieData.UUID) {
			const GetUserData = await fetchUserData(InfoCookieData.UUID);
			console.log('🚀 ~ file: +page.ts:24 ~ InitializeData ~ GetUserData:', GetUserData);
			UserData.update(() => GetUserData);
		}
		loadingState = true;
	});

	
</script>

<svelte:head>
	<title>Null Pointer</title>
	<link rel="icon"  href={HeadLogo} />
</svelte:head>

{#if loadingState}
	<div
		class="   flex min-h-screen w-full flex-col justify-start overflow-x-hidden overflow-y-hidden bg-[#181818]"
	>
		<Navbar />
		<div class="flex w-full flex-row justify-center">
			<!-- COLLECTIVE -->
			<Collectives />
			<!-- QUESTION LIST -->

			<!-- settings -->
			<div class=" mb-2 mt-2 min-h-screen w-[1200px] bg-[#2d2d2d]   ">
				<div class="flex flex-row">
					<!-- questions -->
					<div class="flex w-[850px] flex-col">
						<!-- public ques & ask -->
						<div class="   flex h-16 w-full flex-row">
							<div class=" m-3 text-3xl text-[#e7e9eb]">Public Questions</div>
							<div class="grow" />
							<div
								on:click={() => goto('/ask')}
								on:keydown={() => {}}
								class="mr-4 mt-2 flex h-12 w-28 items-center justify-center rounded-md bg-sky-500 hover:cursor-pointer hover:bg-blue-600"
							>
								<p class="  my-auto text-xl font-semibold text-gray-200">Ask</p>
							</div>
						</div>
						<!-- filter -->
						<QuesFilter/>

						<!-- all Question list -->
						<PublicQues />
						<PageNum />
					</div>
					<!-- right sidebar -->
					<div class="flex h-full w-[350px] flex-col">
						<!-- Blog -->
						<BlogListCom />
						<!-- hot ques -->
						<HotQuesCom />
					</div>
				</div>
			</div>
		</div>
		<Footer />
	</div>
{:else}
	<LoadingSVG />
{/if}

<style>
</style>
