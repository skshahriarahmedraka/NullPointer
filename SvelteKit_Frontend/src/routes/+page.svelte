<script lang="ts">
	import HotQuesCom from '$lib/HotQues/index.svelte';
	import BlogListCom from '$lib/BlogList/index.svelte';
	import PublicQues from '$lib/PublicQues/index.svelte';
	import Navbar from '$lib/Navbar/index.svelte';
	import Collectives from '$lib/Collectives/index.svelte';
	import LoadingSVG from '$lib/Loading/index.svelte';

	import { goto } from '$app/navigation';
	import { afterUpdate } from 'svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import PageNum from '$lib/PageNum/pageNum.svelte';
	import type { PageData } from './$types';
	import { UserData } from '$lib/store/store';
	import type { UserDataType } from '$lib/store/types';

	// export let data: PageData 
	// let GetUserData:UserDataType = data.GetUserData
	// UserData.update((data) => {
	// 	data = GetUserData;
	// 	return data;
	// });
	let loadingState: boolean = false;
	afterUpdate(() => {
		loadingState = true;
	});
</script>

{#if loadingState}
	<div class="   flex   w-full flex-col  justify-start overflow-x-hidden overflow-y-hidden bg-[#181818] min-h-screen ">
		<Navbar />
		<div class="flex w-full flex-row justify-center   ">
			
			<!-- COLLECTIVE -->
			<Collectives />
			<!-- QUESTION LIST -->

			<!-- settings -->
			<div class=" mt-2 mb-2 min-h-screen   w-[1200px]  bg-[#2d2d2d]  ">
				<div class="flex flex-row ">
					<!-- questions -->
					<div class="flex w-[850px] flex-col">
						<!-- public ques & ask -->
						<div class="   flex h-16  w-full flex-row ">
							<div class=" m-3 text-3xl text-[#e7e9eb] ">Public Questions</div>
							<div class="grow" />
							<div
								on:click={() => goto('/ask')}
								on:keydown={() => {}}
								class="mr-4  mt-2 flex h-12  w-28 items-center justify-center rounded-md bg-sky-500 hover:cursor-pointer hover:bg-blue-600 "
							>
								<p class="  my-auto text-xl font-semibold text-gray-200 ">Ask</p>
							</div>
						</div>
						<!-- filter -->
						<div class="flex h-16  w-full flex-row">
							<div class="ml-5 self-center text-lg text-[#e7e9eb] ">{22587351} Questions</div>
							<div class="grow" />
							<div
								class=" flex h-9  w-96 flex-row self-center rounded-md border-2 border-[#7d858d]"
							>
								<div
									class="flex h-full w-20 items-center justify-center border-r-2 border-[#7d858d] text-[#e7e9eb] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] "
								>
									Newest
								</div>
								<div
									class="flex h-full w-16 items-center justify-center border-r-2 border-[#7d858d] text-[#e7e9eb] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] "
								>
									Active
								</div>
								<div
									class="flex h-full w-20 items-center justify-center border-r-2 border-[#7d858d] text-[#e7e9eb] hover:cursor-pointer hover:bg-[#3d4951]  hover:text-[#9cc1db] active:bg-[#404245] "
								>
									Bountied
								</div>
								<div
									class="flex h-full w-28 items-center justify-center border-r-2 border-[#7d858d] text-[#e7e9eb] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] "
								>
									Unanswered
								</div>
								<div
									class="flex h-full  w-16 items-center justify-center text-[#e7e9eb] hover:cursor-pointer hover:bg-[#404245] "
								>
									More
								</div>
							</div>
							<div
								class=" ml-4 flex  h-9 w-24 flex-row items-center justify-center self-center rounded-md  border-2 border-[#7d858d] "
							>
								<svg
									class="h-6 w-6  stroke-[#9cc1db]"
									fill="none"
									stroke="currentColor"
									viewBox="0 0 24 24"
									xmlns="http://www.w3.org/2000/svg"
									><path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"
									/></svg
								>
								<div class="ml-1 font-semibold text-[#9cc1db] ">Filter</div>
							</div>
						</div>

						<!-- all Question list -->
						<PublicQues />
						<PageNum/>
					</div>
					<!-- right sidebar -->
					<div class="flex h-full w-[350px] flex-col ">
						<!-- Blog -->
						<BlogListCom />
						<!-- hot ques -->
						<HotQuesCom />
					</div>
				</div>
			</div>
		</div>
		<Footer />
	</div>
{:else}
	<LoadingSVG />
{/if}

<style>
</style>
