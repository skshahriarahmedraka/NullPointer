<script lang="ts">
	import LoadingSVG from '$lib/Loading/index.svelte';
	import Navbar from '$lib/Navbar/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';
	// import Collectives from "$lib/Collectives/index.svelte"
	import { onMount } from 'svelte';

	// export let RelatedQuestionList:any

	import { getCookieValue } from '$lib/store/utils';
	import type { BlogDataType, CookieInfo1Type, UserDataType } from '$lib/store/types';
	import { fetchUserData } from '$lib/store/fetch';
	import { UserData } from '$lib/store/store';
	import HashHome from '$lib/Hash/HashHome.svelte';

	let loadingState: boolean = false;
	// export let data
	// BlogList =data.BlogList
	onMount(async () => {
		const CookieValueInfo1: string = getCookieValue('Info1');

		const InfoCookieData = JSON.parse(atob(CookieValueInfo1)) as CookieInfo1Type;
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

		// let BlogDataList = [] as BlogDataType[];
		// BlogList =await  fetchBlogList("time", 0, 10, 1);
		loadingState = true;
	});
</script>

<!-- markup (zero or more items) goes here -->

<!-- {#if RelatedQuestionListLoading} -->

{#if loadingState}
	<!-- ////////////////////////////// -->

	<div
		class="   flex w-full flex-col justify-center overflow-x-hidden overflow-y-hidden bg-[#2d2d2d]"
	>
		<Navbar />
		<div class="flex w-full flex-row justify-center">
			<!-- COLLECTIVE -->
			<!-- svelte-ignore missing-declaration -->
			<!-- <Collectives /> -->
			<!-- QUESTION LIST -->

			<!-- <BlogHomeCom {BlogList} /> -->
			<HashHome />
		</div>
		<Footer />
	</div>
{:else}
	<LoadingSVG />
{/if}
<!-- SPACES -->

<!-- RELATED QUESTION -->
<!-- <RelatedQues {RelatedQuestionList}/> -->

<!-- {:else}
    <LoadingSVG/>
{/if} -->

<style>
	/* your styles go here */
</style>
