<script lang="ts">
	// import { BlogDataType } from '$lib/store/types.ts';
	import type { BlogDataType } from '$lib/store/types';
	import Book from '$lib/Loading/book.svelte';
	import { onMount } from 'svelte';

	// import {QuestionData} from "../store/store"
	let Loading = false;
	onMount(() => {
		Loading = true;
	});
	export let BlogList: BlogDataType[] = [] as BlogDataType[];
	console.log('🚀 ~ file: index.svelte:11 ~ BlogList:', BlogList);
</script>

{#if Loading}
	<div class=" mt-2 max-h-full min-h-screen w-[1355px] bg-[#2d2d2d] pl-5">
		<!-- Big Blog -->
		<div
		on:click={()=>{}}
		on:keypress={()=>{}}
		 class="flex h-[580px] w-full flex-row gap-x-2">
			<div
				class="  mt-8 flex h-[515px] w-[985px] flex-col rounded-xl bg-[#202022] shadow-lg shadow-gray-600 transition-all duration-200 hover:scale-105 hover:cursor-pointer"
			>
				<img
					src={'https://149351115.v2.pressablecdn.com/wp-content/uploads/2023/05/050123-Stack-Overflow-AI-isnt-the-app-2-1200x630.jpg'}
					alt=""
					class=" h-[515px] w-[985px] rounded-xl object-cover"
				/>
			</div>
			<div
				class=" flex h-full w-[300px] flex-col text-[#fafafa] transition-all duration-200 hover:scale-105 hover:cursor-pointer"
			>
				<p class="font-Poppins ml-2 mt-5 text-3xl font-bold hover:text-blue-600">
					{"More on our AI future: building course recommendations and a new data platform"}
				</p>
				<p class=" font-OpenSans ml-3 mt-2 text-sm text-slate-400">
					{"Our senior data scientist goes deep in talking about how we built our course recommendation engine."}
				</p>

				<!-- {#if  BlogList[0].Tags.length ==0} -->
				<div class="flex flex-row flex-wrap">
					<li
						class="m-1 w-fit rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
					>
						{'Code For Living'}
					</li>
					<li
						class="m-1 w-fit rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
					>
						{'Code'}
					</li>
				</div>
				<!-- {:else} -->
				<!-- else content here -->
				<!-- {/if} -->
				<p class="  text-t">{"25 jan 2023"}</p>
				<div class=" grow" />
				<div class="mb-8 ml-2 flex h-10 w-full flex-row gap-2">
					<img
						src={'https://149351115.v2.pressablecdn.com/wp-content/uploads/2020/05/ak_big.png'}
						alt=""
						class=" h-10 w-10 rounded-xl"
					/>
					<div class="flex flex-col">
						<p class=" font-raleway text-sm font-semibold text-blue-600">{"David Gibson"}</p>
						<p class=" font-light text-gray-500">{"Senior Data Scientist"}</p>
					</div>
				</div>
			</div>
		</div>
		<!-- Small blogs  -->
		<div class="flex flex-row gap-2 flex-wrap ">
		{#each Array.from(Array(50+1).keys()).slice(1) as i}

		<div
		on:click={()=>{}}
		on:keypress={()=>{}}
			class="mt-8 flex h-[536px] w-[306px] items-center flex-col rounded-xl bg-[#2d2d2d] hover:shadow-lg hover:shadow-gray-600 transition-all duration-200  hover:cursor-pointer "
			>
			<img src={"https://149351115.v2.pressablecdn.com/wp-content/uploads/2023/05/052223-Stack-Overflow-Attention-is-your-greatest-resource-1200x630.jpg"} alt="" class="  h-[160px] w-[300px]  rounded-xl  object-cover" />
			<div class="text-[#fafafa] ">
				<p class=" mt-2 ml-3 font-OpenSans text-sm text-slate-400">
					{"Jul 27, 2022 3 min read"}
				</p>
				<p class="ml-2 mt-2 font-Poppins text-2xl font-bold line-clamp-5 font-sf-pro">
					{"Keep ‘em separated: Get better maintainability in web projects using the model-view-controller pattern "}
				</p>
				<p class="mt-2 mb-3 ml-3 w-[90%] text-sm text-slate-300 line-clamp-4 ">
					{"If you want to get the attention of a employer, marketing has the techniques to get you there. "}
				</p>
			</div>
			<div class=" grow" />
			<div class="mb-8 ml-4 flex h-10 w-full flex-row gap-2">
				<img
					src={'https://149351115.v2.pressablecdn.com/wp-content/uploads/2020/05/ak_big.png'}
					alt=""
					class=" h-10 w-10 rounded-xl"
				/>
				<div class="flex flex-col">
					<p class=" font-raleway text-sm font-semibold text-blue-600">{"David Gibson"}</p>
					<p class=" font-light text-gray-500">{"Senior Data Scientist"}</p>
				</div>
			</div>
		</div>
		{/each}
	</div>
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
