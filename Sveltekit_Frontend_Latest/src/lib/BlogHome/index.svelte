<script lang="ts">
	// import { BlogDataType } from '$lib/store/types.ts';
	import type { BlogDataType } from '$lib/store/types';
	import Book from '$lib/Loading/book.svelte';
	import { onMount } from 'svelte';
	import Emptybox from '$lib/errors/emptybox.svelte';
	import { fetchUserFlairData } from '$lib/store/fetch';
	import UserAnonymous from '$lib/icons/UserAnonymous.svelte';
	import { goto } from '$app/navigation';

	// import {QuestionData} from "../store/store"
	let Loading = false;
	onMount(() => {
		Loading = true;
	});
	export let BlogList: BlogDataType[] = [] as BlogDataType[];
	console.log('🚀 ~ file: index.svelte:11 ~ BlogList:', BlogList);
	let PosterBlog: BlogDataType = BlogList[0];
</script>

{#if Loading}
	<div class=" mt-2  min-h-screen w-[1355px] bg-[#2d2d2d] pl-5">
		{#if BlogList.length != 0}
			<!-- content here -->

			<!-- Big Blog -->
			<div
				on:click={() => { goto(`/b/${PosterBlog.ID}`) }}
				on:keypress={() => {}}
				class="flex h-[580px] w-full flex-row gap-x-2"
			>
				<div
					class="  mt-8 flex h-[515px] w-[985px] flex-col rounded-xl bg-[#202022] shadow-lg shadow-gray-600 transition-all duration-200 hover:scale-105 hover:cursor-pointer"
				>
					<img src={PosterBlog.Image} alt="" class=" h-[515px] w-[985px] rounded-xl object-cover" />
				</div>
				<div
					class=" flex h-full w-[300px] flex-col text-[#fafafa] transition-all duration-200 hover:scale-105 hover:cursor-pointer"
				>
					<p class="font-Poppins ml-2 mt-5 text-3xl font-bold hover:text-blue-600">
						{PosterBlog.Title}
					</p>
					<p class=" font-OpenSans ml-3 mt-2 text-sm text-slate-400">
						{PosterBlog.MetaTitle}
					</p>

					<div class="  min-h-10 flex max-h-20 w-full flex-row flex-wrap">
						{#each PosterBlog.Tags as tag}
							<div
								class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
							>
								{tag}
							</div>
						{/each}
					</div>

					<p class="  font-OpenSans mt-2 ml-3 text-sm text-slate-400">
						{new Date(PosterBlog.WrittenTime).toLocaleDateString()}
					</p>
					<div class=" grow" />
					<div class="mb-8 ml-2 flex h-10 w-full flex-row gap-2">
						{#each PosterBlog.WrittenBy as i}
							<!-- content here -->
							{#await fetchUserFlairData(i)}
								<p class=" text-white">...waiting</p>
							{:then flairUserData}
							{#if flairUserData.UserImage.trim() != ''}
							<!-- content here -->
							<img
								src={flairUserData.UserImage}
								alt=""
								class=" aspect-square active:ring-offset-base-50 mx-2 h-10 w-10 cursor-pointer rounded-lg object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500 active:rounded-md active:ring active:ring-blue-600"
							/>
							{:else}
							<!-- else content here -->
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
							{/await}
						{/each}
					</div>
				</div>
			</div>
			<!-- Small blogs  -->
			{#if BlogList.slice(1).length >0}
			
			<div class="flex flex-row flex-wrap gap-2">
				{#each BlogList.slice(1) as i}
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
			</div>
			{:else }
			<!-- only one blog exist -->
				<div class="flex flex-col h-48 w-full items-center justify-center    ">
					<Emptybox class="h-20 " />
					<p class=" font-raleway text-white">No More Blog Found</p>
				</div>
			{/if}
		{:else}
			<div class=" flex h-full w-full items-center justify-center text-xl text-white">
				<div class="flex flex-col items-center justify-center">
					<Emptybox class="h-20 " />
					<p class=" font-raleway">Nothing Found </p>
				</div>
			</div>
		{/if}
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
