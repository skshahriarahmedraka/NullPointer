<script lang="ts">
	import Book from '$lib/Loading/book.svelte';
	import { fetchUserFlairData } from '$lib/store/fetch';
	import type { BlogDataType, UserFlairDataType } from '$lib/store/types';
	import { onMount } from 'svelte';
	import { marked } from 'marked';
	import UserAnonymous from '$lib/icons/UserAnonymous.svelte';

	// import {QuestionData} from "../store/store"
	export let BlogData: BlogDataType;
	console.log('🚀 ~ file: index.svelte:12 ~ BlogData:', BlogData);
	let Loading = false;
	let BlogWrittenBy: UserFlairDataType[] = [] as UserFlairDataType[];
	// let EditedBy: UserFlairDataType;

	// async function GetAskbyData(ID: string) {
	// }
	// async function GetEditedByData(ID: string) {
	// }

	onMount(async () => {
		if (BlogData.WrittenBy.length != 0) {
			for (let index = 0; index < BlogData.WrittenBy.length; index++) {
				let temp = await fetchUserFlairData(BlogData.WrittenBy[index]);
				BlogWrittenBy.push(temp);
			}
			// GetAskbyData(QuestionData.QuesAskedBy);
		}
		Loading = true;
	});
</script>

{#if Loading}
	<div class=" mt-2 flex max-h-full min-h-screen w-[1400px] flex-col gap-5 bg-transparent">
		<!-- <p class=""> Blog Title </p> -->
		<div class="flex h-[360px] w-full flex-row">
			<!-- blog details -->
			<div class="flex h-full w-1/2 flex-col justify-start gap-2">
				<p class=" font-OpenSans mt-2 text-sm text-slate-400">
					{new Date(BlogData.WrittenTime).toLocaleDateString()}
				</p>
				<p class=" font-sf-pro text-4xl font-bold text-[#e7e8eb] hover:text-blue-600">
					{BlogData.Title}
				</p>
				<p class="  font-raleway text-sm text-slate-400">
					{BlogData.MetaTitle}
				</p>
				<div class=" grow" />
				<div class="  flex h-10 w-full flex-row gap-2">
					{#each BlogWrittenBy as i}
						<!-- content here -->
						<div class="flex h-10 w-fit flex-row gap-2">
							{#if i.UserImage != ''}
								<!-- content here -->
								<img
									src={i.UserImage}
									alt=""
									class=" aspect-square active:ring-offset-base-50 mx-2 h-10 w-10 cursor-pointer rounded-lg object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500 active:rounded-md active:ring active:ring-blue-600"
								/>
							{:else}
								<!-- else content here -->
								<UserAnonymous
									class=" aspect-square active:ring-offset-base-50 mx-2  h-10 w-10  cursor-pointer  object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
								/>
							{/if}
							<!-- <img
						src={'https://149351115.v2.pressablecdn.com/wp-content/uploads/2020/05/ak_big.png'}
						alt=""
						class=" h-10 w-10 rounded-xl"
						/> -->
							<div class="flex flex-col">
								<p class=" text-sm font-semibold text-blue-600 font-sf-pro">{i.UserName}</p>
								<p class=" font-poppins font-light text-gray-500">{i.UserTitle}</p>
							</div>
						</div>
					{/each}
				</div>
			</div>
			<!-- blog image -->
			<div class=" h-full w-1/2">
				<img src={BlogData.Image} alt="" class=" h-full w-full" />
			</div>
		</div>
		<div class=" flex flex-col items-center text-base text-[#e7e8eb]">
			<div
				class=" font-poppins prose mt-5 min-w-full max-w-full overflow-hidden break-words bg-inherit p-5 text-[#e7e9eb] prose-headings:text-[#e7e9eb] prose-p:text-[#e7e9eb] prose-a:text-blue-500 prose-blockquote:border-sky-400 prose-blockquote:text-[#e7e9eb] prose-figure:text-white prose-figcaption:text-[#e7e9eb] prose-strong:text-[#e7e9eb] prose-em:text-[#e7e9eb] prose-code:text-[#e7e9eb] prose-pre:text-[#e7e9eb] prose-ol:text-[#e7e9eb] prose-ul:text-[#e7e9eb] prose-li:text-[#e7e9eb] prose-li:marker:text-white prose-table:text-[#e7e9eb] prose-thead:text-[#e7e9eb] prose-tr:border-4 prose-tr:border-gray-300 prose-tr:text-[#e7e9eb] prose-th:border-2 prose-th:border-gray-300 prose-th:text-[#e7e9eb] prose-td:border-2 prose-td:border-gray-300 prose-td:text-[#e7e9eb] prose-img:text-[#e7e9eb] prose-video:text-[#e7e9eb] prose-hr:bg-gray-500 prose-hr:text-[#e7e9eb]"
			>
				{@html marked(BlogData.Description)}
			</div>
			<!-- <p class=" w-[950px] h-fit">

				{`Since my last quarterly update, companies across nearly every sector have experienced significant transformation—whether it’s a more aggressive focus on profitability or a shift in product strategy due to the acceleration of generative AI (GenAI). Thematically, however, one thing has remained the same: companies are committed to driving productivity and efficiency throughout their organizations. At Stack Overflow, we continue to help our customers and community deliver both.

In the last quarter of our fiscal year, which ended on March 31, that meant announcing the availability of Stack Overflow for Teams in the Microsoft Azure marketplace while launching Topic Collectives and Staging Ground Beta 2 on our public platform. But most significantly, we accelerated our AI efforts internally and look forward to sharing more this summer. I’m excited to see how we leverage our domain focus and special community-driven blend of trust, recognition, and accuracy to GenAI.`}
</p> -->
			<div class=" flex h-fit w-[950px] flex-row mb-3">
				<p class=" text-lg font-medium text-[#e7e8eb] ">{"Tags "}</p>
				<div class="  h-10  flex w-full flex-row flex-wrap">
					{#each BlogData.Tags as tag}
						<div
							class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
						>
							{tag}
						</div>
					{/each}
				</div>
				<!-- <p class=" font-semibold text-lg text-blue-600 ">Golang</p>
		<p class=" font-semibold text-lg text-blue-600 ">Rust</p>
		<p class=" font-semibold text-lg text-blue-600 ">Elixir</p> -->
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
