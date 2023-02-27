<script lang="ts">
	
	import Navbar from '$lib/Navbar/index.svelte';
	import Collectives from "$lib/Collectives/index.svelte"
	// import RelatedQues from "$lib/RelatedQues/index.svelte"
	// import Ques from "$lib/Ques/index.svelte"
	// import Ans from "$lib/Ans/index.svelte"
	// // import {UserData,QuestionData,Loading} from "$lib/store/store"
	import { afterUpdate } from 'svelte';

	let loadingState: boolean = false;
	afterUpdate(() => {
		loadingState = true;
	});
	import LoadingSVG from '$lib/Loading/index.svelte';
const HeadLogo = new URL("../../../../lib/icons/favicon.png", import.meta.url).href;
</script>

<!-- <slot /> -->
<svelte:head>
	<title>Null Pointer</title>
	<link rel="icon"  href={HeadLogo} />
</svelte:head>


{#if loadingState}
	<div class="flex   w-full flex-row justify-center bg-[#181818] z-40 overflow-y-scroll  ">
		<Navbar  />
		<!-- COLLECTIVE -->
		<Collectives  />
		<!-- QUESTION LIST -->
		

		<slot></slot>
    </div>
    
{:else}
	<LoadingSVG/>
{/if}