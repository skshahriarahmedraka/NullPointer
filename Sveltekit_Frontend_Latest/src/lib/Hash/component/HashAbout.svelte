<script lang="ts">
	import type { HashDataType } from "$lib/store/types";
	import { marked } from 'marked';


    export let HashData:HashDataType
    // your script goes here
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->


<div class="h-full w-full">

    {#if HashData.About.trim() != ''}
			<div
				class=" prose mt-5 min-w-full max-w-full overflow-hidden break-words bg-inherit p-5 text-[#e7e9eb] prose-headings:text-[#e7e9eb] prose-p:text-[#e7e9eb] prose-a:text-blue-500 prose-blockquote:border-sky-400 prose-blockquote:text-[#e7e9eb] prose-figure:text-white prose-figcaption:text-[#e7e9eb]   prose-strong:text-[#e7e9eb] prose-em:text-[#e7e9eb] prose-code:text-[#e7e9eb] prose-pre:text-[#e7e9eb] prose-ol:text-[#e7e9eb] prose-ul:text-[#e7e9eb] prose-li:text-[#e7e9eb] prose-li:marker:text-white prose-table:text-[#e7e9eb] prose-thead:text-[#e7e9eb] prose-tr:border-4 prose-tr:border-gray-300 prose-tr:text-[#e7e9eb] prose-th:border-2 prose-th:border-gray-300 prose-th:text-[#e7e9eb] prose-td:border-2 prose-td:border-gray-300 prose-td:text-[#e7e9eb] prose-img:text-[#e7e9eb] prose-video:text-[#e7e9eb]  prose-hr:bg-gray-500 prose-hr:text-[#e7e9eb]  "
			>
				{@html marked(HashData.About)}
			</div>
			<!-- post button -->
			<!-- <button
				class=" m-5 flex h-12 w-fit  items-center justify-center rounded-md bg-blue-500 px-3 hover:bg-blue-600 active:bg-blue-800 "
			>
				<p class="my-auto text-xl font-semibold text-gray-200">Post Your Question</p>
			</button> -->
		{/if}
</div>