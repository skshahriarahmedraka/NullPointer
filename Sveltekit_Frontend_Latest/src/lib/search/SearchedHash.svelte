<script lang="ts">
	import { goto } from '$app/navigation';
	import PageNum from '$lib/PageNum/pageNum.svelte';
	import { onMount } from "svelte";
	import Emptybox from "$lib/errors/emptybox.svelte";

	// your script goes here
	const rustLogo = new URL('../images/rust.png', import.meta.url).href;
	const awsLogo = new URL('../images/go.png', import.meta.url).href;
	export let filterType: string = 'time';
	console.log("🚀 ~ file: SearchedHash.svelte:9 ~ filterType:", filterType)
	export let pageNumStart: number = 0;
	export  let pageNumNow: number = 0;
	export  let pageNumEnd: number = 0;
	$: console.log("pageNumNow:", pageNumNow)
	$: console.log("pageNumStart:", pageNumStart)
	$: console.log("pageNumEnd:", pageNumEnd)
	let Loading:boolean=false 
	onMount(()=>{
		Loading= true
	})
</script>

<!-- markup (zero or more items) goes here -->
{#if Loading}


<div class=" mb-2 mt-2 h-full w-full bg-[#2d2d2d]">
	{#each Array(4) as _, i}
		<div
			on:click={() => {
				// goto(`/q/${i.ID}`);
			}}
			on:keydown={() => {}}
			class=" min-h-32 mx-2 flex h-40 w-full cursor-pointer flex-row border-t-2 border-[#404245] py-2 hover:bg-slate-800 hover:bg-opacity-50"
		>
			<div class="my-2 flex h-full w-60 flex-col items-center justify-center gap-1">
				<img src={awsLogo} alt="" class="" />
			</div>
			<div class="flex flex-col m-3">
				<p class=" h-12 w-full text-2xl font-poppins text-[#e7e9eb] ">
					{"Golang"}
				</p>
				<p class=" text-slate-400  font-sf-pro">{`
					C# (pronounced "see sharp") is a high-level, statically typed, multi-paradigm programming language developed by Microsoft. C# code usually targets Microsoft's .NET family of tools and run-times, which…
				`}</p>
			</div>
		</div>
	{/each}
<!-- <PageNum bind:pageNumStart={pageNumStart} bind:pageNumNow={pageNumNow} bind:pageNumEnd={pageNumEnd} /> -->

</div>
{:else}
	<div class=" flex h-full w-full items-center justify-center text-xl text-white">
		<div class="flex flex-col items-center justify-center">
			<Emptybox class="h-20 " />
			<p class=" font-raleway">error</p>
		</div>
	</div>
{/if}