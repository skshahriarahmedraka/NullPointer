<script lang="ts">
	import { goto } from '$app/navigation';
	import PageNum from '$lib/PageNum/pageNum.svelte';
	import { onMount } from "svelte";
	import Emptybox from "$lib/errors/emptybox.svelte";
	import { StoredSearchedString } from '$lib/store/store';
	import { fetchSearchhashArrWithMetadata } from '$lib/store/fetch';
	import type { HashArrWithMetadata } from '$lib/store/types';
	import Verified from '$lib/SpaceCom/svgs/Verified.svelte';
	import BlueDot from '$lib/SpaceCom/svgs/BlueDot.svelte';

	// your script goes here
	const rustLogo = new URL('../images/rust.png', import.meta.url).href;
	const awsLogo = new URL('../images/go.png', import.meta.url).href;


	export let filterType: string;
	export let filterState:boolean

	console.log("🚀 ~ file: SearchedHash.svelte:12 ~ filterState:", filterState)
	console.log("🚀 ~ file: SearchedHash.svelte:9 ~ filterType:", filterType)
	let pageNumStart: number = 0;
	  let pageNumNow: number = 0;
	  let pageNumEnd: number = 0;
	$: console.log("pageNumNow:", pageNumNow)
	$: console.log("pageNumStart:", pageNumStart)
	$: console.log("pageNumEnd:", pageNumEnd)
	let Loading:boolean=false 
	let contentPerPage: number = 5;
    let HashList :HashArrWithMetadata = {} as HashArrWithMetadata

	let SearchedString= ''

	onMount(async () => {
		filterType= 'alpha'
		filterState = true
		StoredSearchedString.subscribe((value) => {
			console.log("🚀 ~ file: SearchedHash.svelte:34 ~ StoredSearchedString.subscribe ~ value:", value)
			
			SearchedString = value;
		});
		if ( SearchedString.trim() != "" ){

			HashList = await fetchSearchhashArrWithMetadata(
				SearchedString,
				filterType,
				pageNumNow * contentPerPage,
				pageNumNow * contentPerPage + contentPerPage,
				filterState ? -1 : 1
				);
				console.log("🚀 ~ file: SearchedHash.svelte:43 ~ onMount ~ HashList:", HashList)
		}
		pageNumStart = 0;
		// pageNumNow = pageNumNow * contentPerPage;
		pageNumEnd = Math.floor(HashList.Metadata.Length / contentPerPage);
		Loading = true;
	});
	async function ApplyFilter(filterType: string, filterState: boolean) {
		// time , views , unanswered , votes
		// let SearchedString= ''
		StoredSearchedString.subscribe((value) => {
			console.log("🚀 ~ file: searchedQues.svelte:65 ~ StoredSearchedString.subscribe ~ value:", value)
			SearchedString = value;
		});
		
		if (filterState && SearchedString.trim() != "" ) {
			HashList = await fetchSearchhashArrWithMetadata(
				SearchedString,
				filterType,
				pageNumNow * contentPerPage,
				pageNumNow * contentPerPage + contentPerPage,
				-1
			);
		} else if ( SearchedString.trim() != "" ) {
			HashList = await fetchSearchhashArrWithMetadata(
				SearchedString,
				filterType,
				pageNumNow * contentPerPage,
				pageNumNow * contentPerPage + contentPerPage,
				1
			);
		}

		pageNumStart = 0;
		// pageNumNow = pageNumNow * contentPerPage;
		pageNumEnd = Math.floor(HashList.Metadata.Length / contentPerPage);
	}
	$: pageNumNow, ApplyFilter(filterType, filterState);
	$: $StoredSearchedString, ApplyFilter(filterType, filterState)

	function RoundNum(x: number) {
		if (x === null) {
			return 0;
		} else if (x >= 1000000000) {
			x /= 1000000000;
			return String(x.toFixed(2)) + 'B';
		} else if (x >= 1000000) {
			x /= 1000000;
			return String(x.toFixed(2)) + 'M';
		} else if (x >= 1000) {
			x /= 1000;
			return String(x.toFixed(2)) + 'K';
		} else {
			return String(x);
		}
	}
</script>

{#if Loading}


<div class=" mb-2 mt-2 h-full w-full bg-[#2d2d2d]">
	<!-- {#each Array(4) as _, i}
		<div
			on:click={() => {
				goto(`/q/${i.ID}`);
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
	{/each} -->
	<div class="">

		{#if HashList.HashList.length != 0}
			<!-- content here -->

			
			
			<div class="flex flex-row flex-wrap gap-2">
				{#each HashList.HashList as i}
				
						<div
						on:click={()=>{
							goto('/h/'+i.ID)	
						}}
						on:keypress={()=>{}}
					class=" flex h-48  w-64 flex-col rounded-xl border-2   border-blue-600 border-opacity-25  bg-slate-700 bg-opacity-30 fill-[#b4b8bc] text-[#e7e9eb] hover:cursor-pointer hover:bg-opacity-60"
					>
					<!-- SPACE WALL PHOTO -->
					
					<div
                    
					class="flex h-2/3 w-full items-end justify-start rounded-t-xl pl-2 "
					style=" background-image: url('{i.BannerImage}')"
                        
					>
						<!-- SPACE PROFILE PHOTO -->
						<img
							src="{i.Image }"
							alt=""
							class=" aspect-square active:ring-offset-base-50 h-14 w-14  cursor-pointer rounded-2xl  border-4 border-[#2d2d2d] object-cover transition-all duration-150 ease-linear hover:rounded-2xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
						/>
					</div>
					<div class="flex h-1/3 w-full flex-col ">
						<div class=" m-2 flex h-1/2 w-full flex-row text-base font-semibold ">
							<!-- SPACE NAME -->
							<p class="inline w-[92%] line-clamp-1">
								{i.Name}
							</p>
							<Verified/>
						</div>
						<div class=" h-1/2 w-full text-center text-sm">
							<p class="inline font-semibold text-white">{RoundNum(Number(i.NumOfFollower))}</p>
							<p class="inline font-semibold text-white">Follower</p>
							<BlueDot/>
							<p class="inline font-semibold text-white">{RoundNum(Number(i.NumOfQuestion))}</p>
							<p class="inline font-semibold text-white">Question</p>
						</div>
					</div>
				</div>
				{/each}
			</div>
			
		{:else}
			<div class=" flex h-full w-full items-center justify-center text-xl text-white">
				<div class="flex flex-col items-center justify-center">
					<Emptybox class="h-20 " />
					<p class=" font-raleway">Nothing Found </p>
				</div>
			</div>
		{/if}
    </div>


<PageNum bind:pageNumStart={pageNumStart} bind:pageNumNow={pageNumNow} bind:pageNumEnd={pageNumEnd} />

</div>
{:else}
	<div class=" flex h-full w-full items-center justify-center text-xl text-white">
		<div class="flex flex-col items-center justify-center">
			<Emptybox class="h-20 " />
			<p class=" font-raleway">error</p>
		</div>
	</div>
{/if}