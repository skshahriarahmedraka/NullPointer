<script lang="ts">
	import Book from "$lib/Loading/book.svelte";
	import { onMount } from "svelte";

	// import {QuestionData} from "../store/store"
	let Loading = false 
	onMount(()=>{
		Loading= true
	})
</script>

{#if Loading}
	<div class=" mt-2 max-h-full min-h-screen w-[1100px] bg-[#2d2d2d] pl-5">
		<p class="text-white"> blogs</p>
	</div>
{:else}
	<div
		class=" mt-2 flex max-h-full min-h-screen w-[1100px] items-center justify-center bg-[#2d2d2d] pl-5"
	>
		<Book />
	</div>
{/if}

<style>
</style>
