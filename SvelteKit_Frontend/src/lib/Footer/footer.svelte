<script lang="ts">
	import Nullpointer from '$lib/icons/null pointer.svelte';
	import Stediablue from '$lib/icons/stediablue.svelte';
</script>

<div class=" flex h-screen w-screen flex-col bg-slate-800 ">
	<div class=" grow" />

	<div class=" flex h-80  w-full flex-row items-start justify-center gap-2 bg-[#232629]">
		<div class=" my-10 flex h-fit  w-52 flex-col items-start text-sm text-[#677079] ">
			<Nullpointer class=" ml-0 h-6 w-auto" />
			<!-- <p class=" font-medium">Stack Overflow</p> -->
			<p class="">Questions</p>
			<p class="">Help</p>
			<p
				class=" cursor-pointer font-semibold text-gray-400 hover:text-gray-200"
				on:click={() => {}}
				on:keypress={() => {}}
			>
				Modarator Pannel
			</p>
			<p
				class=" cursor-pointer font-semibold text-gray-400 hover:text-gray-200"
				on:click={() => {}}
				on:keypress={() => {}}
			>
				Admin Pannel
			</p>
		</div>
		<div class=" my-10 h-fit w-40 flex-col items-start text-sm text-[#677079] ">
			<p class=" my-2 text-base font-semibold text-gray-400">Products</p>
			<p class="">Teams</p>
			<p class="">Advertising</p>
			<p class="">Collectives</p>
			<p class="">Talent</p>
		</div>
		<div class=" my-10 h-fit w-48 flex-col items-start text-sm text-[#677079] ">
			<p class="my-2 text-base font-semibold text-gray-400">Company</p>
			<p class="">About</p>
			<p class="">Press</p>
			<p class="">Work Here</p>
			<p class="">Legal</p>
			<p class="">Privacy Policy</p>
			<p class="">Terms of Service</p>
			<p class="">Contact Us</p>
			<p class="">Cookie Settings</p>
			<p class="">Cookie Policy</p>
		</div>

		<div class=" my-10 h-fit w-72 flex-col items-start text-sm text-[#677079]">
			<p class="my-2 text-base font-semibold text-gray-400">Stack Exchange Network</p>
			<p class="">Technology</p>
			<p class="">Culture & recreation</p>
			<p class="">Life & arts</p>
			<p class="">Science</p>
			<p class="">Professional</p>
			<p class="">Business</p>
			<p class="">API</p>
			<p class="">Data</p>
		</div>

		<div class=" my-10 h-fit w-80 flex-col items-start text-sm text-[#677079] ">
			<div class=" flex h-fit w-full flex-row gap-2">
				<p class="">Blog</p>
				<p class="">Facebook</p>
				<p class="">Twitter</p>
				<p class="">LinkedIn</p>
				<p class="">Instagram</p>
			</div>

			<div class="mt-8  h-full  w-auto ">
				<Stediablue class="m-auto h-full w-auto " />
			</div>

			<p class="">
				Site design / logo © 2023 Null Pointer Inc; user contributions licensed under CC BY-SA. rev
				2023.3.3.43278
			</p>
		</div>
	</div>
</div>

<style>
	/* your styles go here */
</style>
