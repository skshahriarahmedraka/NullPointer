

<script lang="ts">
	// your script goes here
	import Navbar from '$lib/Navbar/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import Collectives from "$lib/Collectives/index.svelte"
import LoadingSVG from '$lib/Loading/index.svelte';
	import { onMount } from 'svelte';
	import { getCookieValue } from '$lib/store/utils';
	import type { CookieInfo1Type, UserDataType } from '$lib/store/types';
	import { fetchUserData } from '$lib/store/fetch';
	import { UserData } from '$lib/store/store';

let loadingState: boolean = false;
onMount(async () => {
		const CookieValueInfo1: string = getCookieValue('Info1');

		const InfoCookieData = JSON.parse(atob(CookieValueInfo1)) as CookieInfo1Type;
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

<!-- markup (zero or more items) goes here -->
{#if loadingState}
	<!-- ////////////////////////////// -->

    
	<div class="   flex   w-full flex-col  justify-center overflow-x-hidden overflow-y-hidden bg-[#181818] ">
		<Navbar />
	<div class="flex w-full flex-row justify-center   ">
		
		<!-- COLLECTIVE -->
		<Collectives />
		<!-- QUESTION LIST -->

		<h1 class=" flex flex-col text-white items-center justify-center ">Following</h1>
	</div>
	<Footer />
</div>



		{:else}
			<LoadingSVG/>
		{/if}
