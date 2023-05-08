<script lang="ts">
	
	// import Navbar from '$lib/Navbar/index.svelte';
	// import Collectives from "$lib/Collectives/index.svelte"
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
	// import Footer from '$lib/Footer/footer.svelte';
const HeadLogo = new URL("../../../../lib/icons/favicon.png", import.meta.url).href;
</script>

<!-- <slot /> -->
<svelte:head>
	<title>Null Pointer</title>
	<link rel="icon"  href={HeadLogo} />
</svelte:head>

<!-- {#if loadingState} -->

	<!-- ////////////////////////////// -->
	<slot />
    
<!-- {:else} -->
	<!-- <LoadingSVG/> -->
<!-- {/if} -->