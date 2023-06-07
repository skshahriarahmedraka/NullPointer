


<script lang="ts">
    import WriteQues from '$lib/WriteQues/index.svelte'
    import Navbar from '$lib/Navbar/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import Collectives from "$lib/Collectives/index.svelte"
	import LoadingSVG from '$lib/Loading/index.svelte';
	import { onMount } from 'svelte';
	import { getCookieValue } from '$lib/store/utils';
	import type { CookieInfo1Type, UserDataType } from '$lib/store/types';
	import { UserData } from '$lib/store/store';
	import { fetchUserData } from '$lib/store/fetch';
	import HashWrite from '$lib/Hash/HashWrite.svelte';

    // import Navbar from "$lib/Navbar/index.svelte"
    // import Collectives from "$lib/Collectives/index.svelte"
    // import LoadingSVG from "$lib/Loading/index.svelte"
    // import {Loading} from "$lib/store/store"
    // export let HotQues:any
    let loadingState: boolean = false;
onMount(async () => {
		const CookieValueInfo1: string = getCookieValue('Info1');

		const InfoCookieData = JSON.parse(atob(CookieValueInfo1)) as CookieInfo1Type;
		console.log('🚀 ~ file: +page.ts:24 ~ load ~ InfoCookieData:', InfoCookieData);
		let UserDataValue = {} as UserDataType;
		UserData.subscribe((value) => {
			UserDataValue = value;
		});
		// if (UserDataValue.UserID != InfoCookieData.UserID) {
			const GetUserData = await fetchUserData(InfoCookieData.UserID);
			console.log('🚀 ~ file: +page.ts:24 ~ InitializeData ~ GetUserData:', GetUserData);
			UserData.update(() => GetUserData);
		// }
		loadingState = true;
	});
</script>

<style>
</style>





    
<!-- <div class="bg-[#181818]   w-full flex flex-row justify-center  pb-10 "> -->
        
    <!-- <Navbar/> -->
    <!-- COLLECTIVE -->
    <!-- <Collectives/> -->
    <!-- QUESTION LIST -->
    <!-- <Ques/> -->
{#if loadingState}

<div class="   flex   w-full flex-col  justify-center overflow-x-hidden overflow-y-hidden bg-[#181818] ">
    <Navbar />
    <div class="flex w-full flex-row justify-center   ">
        
        <!-- COLLECTIVE -->
        <!-- <Collectives /> -->
        <!-- QUESTION LIST -->

        <!-- <WriteQues /> -->
        <HashWrite/>
    </div>
    <Footer />
</div>
    

{:else}
<LoadingSVG/>
{/if}

        <!-- RELATED QUESTION -->
        <!-- <RelatedQues/> -->

    <!-- </div> -->

