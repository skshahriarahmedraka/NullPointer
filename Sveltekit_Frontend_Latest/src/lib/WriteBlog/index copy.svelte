<script lang="ts">
	import Book from '$lib/Loading/book.svelte';
	import type { BlogDataType } from '$lib/store/types';
	import { onMount } from 'svelte';

	// import {QuestionData} from "../store/store"
	let Loading = false;
	onMount(() => {
		Loading = true;
	});
	export let BlogData: BlogDataType;
	console.log('🚀 ~ file: index.svelte:12 ~ BlogData:', BlogData);
</script>

{#if Loading}
	<div class=" mt-2 max-h-full min-h-screen w-[1400px] bg-transparent flex flex-col gap-5">
		<!-- <p class=""> Blog Title </p> -->
		<div class="flex h-[360px] w-full flex-row">
			<!-- blog details -->
			<div class="flex h-full w-1/2 flex-col justify-start gap-2">
				<p class=" font-OpenSans mt-2 text-sm text-slate-400">
					{'Jul 27, 2022 3 min read'}
				</p>
				<p class="font-Poppins text-4xl font-bold text-[#e7e8eb] hover:text-blue-600">
					{'More on our AI future: building course recommendations and a new data platform'}
				</p>
				<p class=" font-OpenSans text-sm text-slate-400">
					{'Our senior data scientist goes deep in talking about how we built our course recommendation engine.'}
				</p>
				<div class=" grow" />
				<div class="  flex h-10 w-full flex-row gap-2">
					<img
						src={'https://149351115.v2.pressablecdn.com/wp-content/uploads/2020/05/ak_big.png'}
						alt=""
						class=" h-10 w-10 rounded-xl"
					/>
					<div class="flex flex-col">
						<p class=" font-raleway text-sm font-semibold text-blue-600">{'David Gibson'}</p>
						<p class=" font-light text-gray-500">{'Senior Data Scientist'}</p>
					</div>
				</div>
			</div>
			<!-- blog image -->
			<div class=" h-full w-1/2">
				<img
					src={'https://149351115.v2.pressablecdn.com/wp-content/uploads/2023/05/050123-Stack-Overflow-AI-isnt-the-app-2-1200x630.jpg'}
					alt=""
					class=" h-full w-full"
				/>
			</div>
		</div>
		<div class=" text-base text-[#e7e8eb] flex flex-col items-center  ">
			<p class=" w-[950px] h-fit">

				{`Since my last quarterly update, companies across nearly every sector have experienced significant transformation—whether it’s a more aggressive focus on profitability or a shift in product strategy due to the acceleration of generative AI (GenAI). Thematically, however, one thing has remained the same: companies are committed to driving productivity and efficiency throughout their organizations. At Stack Overflow, we continue to help our customers and community deliver both.

In the last quarter of our fiscal year, which ended on March 31, that meant announcing the availability of Stack Overflow for Teams in the Microsoft Azure marketplace while launching Topic Collectives and Staging Ground Beta 2 on our public platform. But most significantly, we accelerated our AI efforts internally and look forward to sharing more this summer. I’m excited to see how we leverage our domain focus and special community-driven blend of trust, recognition, and accuracy to GenAI.`}
</p>
	<div class=" w-[950px] h-fit flex flex-row">
		<p class=" font-medium text-lg text-[#e7e8eb]">Tags</p>
		<p class=" font-semibold text-lg text-blue-600 ">Golang</p>
		<p class=" font-semibold text-lg text-blue-600 ">Rust</p>
		<p class=" font-semibold text-lg text-blue-600 ">Elixir</p>
	</div>
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
