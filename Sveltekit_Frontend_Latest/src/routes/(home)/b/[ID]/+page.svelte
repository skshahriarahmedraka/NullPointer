<script lang="ts">
	// import Navbar from "$lib/Navbar/index.svelte"
	// import Collectives from "$lib/Collectives/index.svelte"
	// import RelatedQues from "$lib/RelatedQues/index.svelte"
	import Ques from '$lib/Ques/index.svelte';
	import Navbar from '$lib/Navbar/index.svelte';
	import Collectives from '$lib/Collectives/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';

	// import Ans from "$lib/Ans/index.svelte"
	// import {UserData,QuestionData,Loading} from "$lib/store/store"
	import LoadingSVG from '$lib/Loading/index.svelte';

	import type { PageData } from './$types';
	import type { CookieInfo1Type, QuestionDataType, UserDataType } from '$lib/store/types';
	import { onMount } from 'svelte';
	import { getCookieValue } from '$lib/store/utils';
	import { fetchUserData } from '$lib/store/fetch';
	import { UserData } from '$lib/store/store';
	import Book from '$lib/Loading/book.svelte';

	export let data: PageData;
	console.log("🚀 ~ file: +page.svelte:23 ~ data:", data)

    let Loading = false;

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
		
        Loading = true;
	});
	const HeadLogo = new URL('$lib/icons/favicon.png', import.meta.url).href;

	
</script>

<svelte:head>
	<title>{"Blog Title"}</title>
	<link rel="icon" href={HeadLogo} />
</svelte:head>
<!-- markup (zero or more items) goes here -->

{#if Loading}
	<!-- content here -->
	<!-- <h1 class="">blog</h1> -->

    
    <div class="   flex   w-full flex-col  justify-center overflow-x-hidden overflow-y-hidden bg-[#181818] ">
        <Navbar />
        <div class="flex w-full flex-row justify-center   ">
            
            <!-- COLLECTIVE -->
            <Collectives />
            <!-- QUESTION LIST -->
    
            <Ques  QuestionData={QuestionData} />
        </div>
        <Footer />
    </div>
{:else}
	<Book />
{/if}


