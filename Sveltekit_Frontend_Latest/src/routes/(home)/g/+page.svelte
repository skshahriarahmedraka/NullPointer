


<script lang="ts">
    
    // import Navbar from "$lib/Navbar/index.svelte"
    // import Collectives from "$lib/Collectives/index.svelte"
    import RelatedQues from "$lib/RelatedQues/index.svelte"
    // import Ques from "$lib/Ques/index.svelte"
    // import Ans from "$lib/Ans/index.svelte"
    // import {UserData,QuestionData,Loading} from "$lib/store/store"
    import LoadingSVG from "$lib/Loading/index.svelte"
    import SpaceCom from "$lib/SpaceCom/index.svelte"
	import Navbar from '$lib/Navbar/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import Collectives from "$lib/Collectives/index.svelte"
import { onMount } from 'svelte';

// export let RelatedQuestionList:any

 

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

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->


<!-- {#if RelatedQuestionListLoading} -->
    
{#if loadingState}

	<!-- ////////////////////////////// -->

    
    
    
<div class="   flex   w-full flex-col  justify-center overflow-x-hidden overflow-y-hidden bg-[#181818] ">
	<Navbar />
	<div class="flex w-full flex-row justify-center   ">
		
		<!-- COLLECTIVE -->
		<Collectives />
		<!-- QUESTION LIST -->

        <SpaceCom/>
	</div>
	<Footer />
</div>
{:else}
    <LoadingSVG/>
{/if}
        <!-- SPACES -->
        

        <!-- RELATED QUESTION -->
        <!-- <RelatedQues {RelatedQuestionList}/> -->


<!-- {:else}
    <LoadingSVG/>
{/if} -->

    








