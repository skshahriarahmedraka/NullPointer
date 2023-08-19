<script lang="ts">
	import HotQuesCom from '$lib/HotQues/index.svelte';
	import BlogListCom from '$lib/BlogList/index.svelte';
	import PublicQues from '$lib/PublicQues/PublicQues.svelte';
	import Navbar from '$lib/Navbar/index.svelte';
	import Collectives from '$lib/Collectives/index.svelte';
	import LoadingSVG from '$lib/Loading/index.svelte';

	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import PageNum from '$lib/PageNum/pageNum.svelte';
	// import type { PageData } from './$types';
	import { getCookieValue } from '$lib/store/utils';
	import type { CookieInfo1Type, UserDataType } from '$lib/store/types';
	import { fetchUserData } from '$lib/store/fetch';
	import { UserData } from '$lib/store/store';
	// import Filter from '$lib/SpaceCom/svgs/Filter.svelte';
	import QuesFilter from '$lib/QuesFilter/QuesFilter.svelte';
	import Option from '$lib/icons/option.svelte';
	// import { Buffer } from 'buffer';
	// export let data: PageData;
	// console.log('🚀 ~ file: +page.svelte:18 ~ data:', data);
	// const { InfoCookieData } = data;
	// console.log('🚀 ~ file: +page.svelte:19 ~ InfoCookieData:', InfoCookieData);

	// async function Necessary() {

	// }
	const HeadLogo = new URL('../../static/favicon.png', import.meta.url).href;

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
		if (UserDataValue.UserID != InfoCookieData.UserID) {
			const GetUserData = await fetchUserData(InfoCookieData.UserID);
			console.log('🚀 ~ file: +page.ts:24 ~ InitializeData ~ GetUserData:', GetUserData);
			UserData.update(() => GetUserData);
		}
		loadingState = true;
	});
</script>

<svelte:head>
	<title>Null Pointer</title>
	<link rel="icon" href={HeadLogo} />
</svelte:head>

{#if loadingState}
	<div
		class="   flex min-h-screen w-full flex-col justify-start overflow-x-hidden overflow-y-hidden bg-[#181818]"
	>
		<Navbar />
		<div
			class="flex w-full flex-row
		xs:h-fit xs:w-full xs:flex-col
		sm:h-fit sm:w-full sm:flex-col
		md:h-fit md:w-full md:flex-col md:items-center md:justify-center
		lg:h-fit lg:w-full lg:flex-row lg:items-start lg:justify-center
		xl:h-fit xl:w-full xl:flex-row xl:items-start xl:justify-center
		xxl:h-fit xxl:w-full xxl:flex-row xxl:items-start xxl:justify-center
		"
		>
			<!-- COLLECTIVE -->
			<Collectives class=" lg:hidden md:hidden sm:hidden xs:hidden  " />
			<!-- QUESTION LIST -->

			<!-- settings -->
			<div
				class="mb-2 mt-2 flex min-h-screen w-[1200px] flex-row bg-[#2d2d2d]
				xs:h-fit xs:w-full xs:flex-col
				sm:h-fit sm:w-full sm:flex-col
				md:h-fit md:w-fit md:flex-col
				lg:h-fit lg:w-fit lg:flex-col
				xl:h-fit xl:w-fit xl:flex-row
				xxl:h-fit xxl:w-fit xxl:flex-row
				"
			>
				<!-- questions -->
				<div class="flex w-[850px] flex-col xs:w-full sm:w-full">
					<!-- public ques & ask -->
					<dev class="   flex h-16 w-full flex-row">
						<Option class="m-2 h-6 w-6 xxl:hidden xl:hidden stroke-blue-600  " />
						<div class=" m-3 text-3xl text-[#e7e9eb] xs:text-xl sm:text-2xl">Public Questions</div>
						<div class="grow" />
						<div
							on:click={() => goto('/ask')}
							on:keydown={() => {}}
							class="mr-4 mt-2 flex h-12 w-28 items-center justify-center rounded-md bg-sky-500 hover:cursor-pointer hover:bg-blue-600"
						>
							<p class="  my-auto text-xl font-semibold text-gray-200 sm:text-lg">Ask</p>
						</div>
					</dev>
					<!-- filter -->
					<!-- <QuesFilter/> -->

					<!-- all Question list -->
					<PublicQues />
					<!-- <PageNum /> -->
				</div>
				<!-- right sidebar -->
				<div
					class="flex h-full w-[350px] flex-col sm:w-full sm:flex-row md:w-full md:flex-row lg:w-full lg:flex-row"
				>
					<!-- Blog -->
					<BlogListCom />
					<!-- hot ques -->
					<HotQuesCom />
				</div>
			</div>
		</div>
	</div>
	<Footer />
{:else}
	<LoadingSVG />
{/if}

<style>
</style>
