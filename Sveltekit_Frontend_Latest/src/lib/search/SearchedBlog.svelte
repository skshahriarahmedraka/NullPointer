<script lang="ts">
	import { goto } from "$app/navigation";
	import PageNum from "$lib/PageNum/pageNum.svelte";
	import Emptybox from "$lib/errors/emptybox.svelte";
	import UserAnonymous from "$lib/icons/UserAnonymous.svelte";
	import { fetchBlogList, fetchUserFlairData } from "$lib/store/fetch";
	import type { BlogDataType } from "$lib/store/types";
	import { onMount } from "svelte";

    // your script goes here
	export let filterType: string;
	export let filterState:boolean
	console.log("🚀 ~ file: SearchedBlog.svelte:12 ~ filterType:", filterType)
	 let pageNumStart: number = 0;
	  let pageNumNow: number = 0;
	  let pageNumEnd: number = 0;
	let contentPerPage:number=10

	$: console.log("pageNumNow:", pageNumNow)
	$: console.log("pageNumStart:", pageNumStart)
	$: console.log("pageNumEnd:", pageNumEnd)
	let Loading = false;
	onMount(async () => {
		BlogList =await  fetchBlogList(filterType,
			pageNumNow*contentPerPage,
			pageNumNow*contentPerPage + contentPerPage,
			filterState ? -1 : 1);
	
	pageNumStart = 0;
		// pageNumNow = pageNumNow*contentPerPage;
		pageNumEnd = Math.ceil(BlogList.length / contentPerPage);
		Loading=true
	});
	async function ApplyFilter(filterType: string, filterState: boolean) {
		// time , views , unanswered , votes
		BlogList =await  fetchBlogList(filterType,
			pageNumNow*contentPerPage,
			pageNumNow*contentPerPage + contentPerPage,
			filterState ? -1 : 1);
		console.log("🚀 ~ file: SearchedBlog.svelte:41 ~ ApplyFilter ~ BlogList:", BlogList)

		
		pageNumStart = 0;
	}


	$: pageNumNow , ApplyFilter(filterType, filterState)
	 let BlogList: BlogDataType[] = [] as BlogDataType[];
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<!-- Small blogs  -->
{#if Loading}

{#if BlogList.length >0}
			
<div class="flex flex-row flex-wrap gap-2 justify-around">
	{#each BlogList as i}
		<div
			on:click={() => {goto(`/b/${i.ID}`) }}
			on:keypress={() => {}}
			class="mt-8 flex h-[536px] w-[306px] flex-col items-center rounded-xl bg-[#2d2d2d] transition-all duration-200 hover:cursor-pointer hover:shadow-lg hover:shadow-gray-600"
		>
			<img
				src={i.Image}
				alt=""
				class="  h-[160px] w-[300px] rounded-xl object-cover"
			/>
			<div class="text-[#fafafa]">
				<p class=" font-OpenSans ml-3 mt-2 text-sm text-slate-400">
					{new Date(i.WrittenTime).toLocaleDateString()}
				</p>
				<p class="font-Poppins ml-2 mt-2 line-clamp-5 font-sf-pro text-2xl font-bold">
					{i.Title}
				</p>
				<p class="mb-3 ml-3 mt-2 line-clamp-4 w-[90%] text-sm text-slate-300">
					{i.MetaTitle}
				</p>
			</div>
			<div class=" grow" />
				{#await fetchUserFlairData(i.WrittenBy[0])}
					<p class=" text-white">...waiting</p>
				{:then flairUserData}
				{#if flairUserData.UserImage.trim() != ''}
			
				<img
					src={flairUserData.UserImage}
					alt=""
					class=" aspect-square active:ring-offset-base-50 mx-2 h-10 w-10 cursor-pointer rounded-lg object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500 active:rounded-md active:ring active:ring-blue-600"
				/>
				{:else}
					<UserAnonymous
						class=" aspect-square active:ring-offset-base-50 mx-2  h-10 w-10  cursor-pointer  object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
					/>
				{/if}
					<!-- <img src={flairUserData.UserImage} alt="" class=" h-10 w-10 rounded-xl" /> -->
					<div class="flex flex-col">
						<p class=" font-raleway text-sm font-semibold text-blue-600">
							{flairUserData.UserName}
						</p>
						<p class=" font-light text-gray-500">{flairUserData.UserTitle}</p>
					</div>
				{:catch error}
					<p style="color: red">{error.message}</p>
				{/await}					</div>
	{/each}
<PageNum bind:pageNumStart={pageNumStart} bind:pageNumNow={pageNumNow} bind:pageNumEnd={pageNumEnd} />

	<!-- <PageNum bind:pageNumStart={pageNumStart} bind:pageNumNow={pageNumNow} bind:pageNumEnd={pageNumEnd} /> -->
</div>
{:else }
	<div class="flex flex-col h-48 w-full items-center justify-center    ">
		<Emptybox class="h-20 " />
		<p class=" font-raleway text-white">No More Blog Found</p>
	</div>
{/if}
{:else}
	<div class=" flex h-full w-full items-center justify-center text-xl text-white">
		<div class="flex flex-col items-center justify-center">
			<Emptybox class="h-20 " />
			<p class=" font-raleway">error</p>
		</div>
	</div>
{/if}
