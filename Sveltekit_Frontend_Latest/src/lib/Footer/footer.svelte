<script lang="ts">
	import { goto } from '$app/navigation';
	import Nullpointer from '$lib/icons/null pointer.svelte';
	import Stediablue from '$lib/icons/stediablue.svelte';
</script>



	<div class=" flex flex-row 

	h-80  
	xs:h-fit xs:overflow-y-scroll xs:p-2 xs:flex-col 
	sm:h-fit sm:grid sm:grid-cols-2 sm:grid-rows-3 sm:overflow-y-scroll sm:p-4  sm:items-start sm:flex-col 
	md:h-fit md:grid md:grid-cols-3 md:grid-rows-2 md:overflow-y-scroll md:p-4 md:justify-center md:items-start md:flex-col md:gap-4
	lg:h-80   lg:w-full lg:p-4 lg:flex-row lg:gap-4 lg:justify-around lg:items-start lg:overflow-hidden
	xl:h-80   xl:w-full xl:p-4 xl:flex-row xl:gap-4 xl:justify-around xl:items-start xl:overflow-hidden
	xxl:h-80   xxl:w-full xxl:p-4 xxl:flex-row xxl:gap-10 xxl:justify-center  xxl:items-start xxl:overflow-hidden
	pt-10
	    bg-[#232629]">
		<div class="   flex h-fit  w-fit flex-col items-start text-sm text-[#677079] ">
			<Nullpointer class=" ml-0 h-6  " />
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
			<p
				class=" cursor-pointer font-semibold text-gray-400 hover:text-gray-200"
				on:click={() => {goto('/b/write')}}
				on:keypress={() => {}}
			>
				Write Blog
			</p>
			<p
				class=" cursor-pointer font-semibold text-gray-400 hover:text-gray-200"
				on:click={() => {goto('/h/write')}}
				on:keypress={() => {}}
			>
				Write hash
			</p>
		</div>
		<div class="   h-fit w-fit flex-col items-start text-sm text-[#677079] ">
			<p class=" my-2 text-base font-semibold text-gray-400">Products</p>
			<p class="">Teams</p>
			<p class="">Advertising</p>
			<p class="">Collectives</p>
			<p class="">Talent</p>
		</div>
		<div class="   h-fit w-fit flex-col items-start text-sm text-[#677079] ">
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

		<div class="   h-fit w-fit flex-col items-start text-sm text-[#677079]">
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

		<div class="   h-fit w-96 xs:w-full flex-col items-start text-sm text-[#677079] ">
			<div class=" flex flex-wrap h-fit w-fit flex-row gap-2">
				<p class="">Blog</p>
				<p class="">Facebook</p>
				<p class="">Twitter</p>
				<p class="">LinkedIn</p>
				<p class="">Instagram</p>
			</div>

			<div class="mt-8  h-full   ">
				<Stediablue class="m-auto h-full " />
			</div>

			<p class="">
				Site design / logo © 2023 Null Pointer Inc; user contributions licensed under CC BY-SA. rev
				2023.3.3.43278
			</p>
		</div>
	</div>

<style>
	/* your styles go here */
</style>
