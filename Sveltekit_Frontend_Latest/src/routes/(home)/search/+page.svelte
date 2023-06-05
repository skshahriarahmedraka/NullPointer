<script lang="ts">
	import Navbar from '$lib/Navbar/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import Collectives from '$lib/Collectives/index.svelte';

	import LoadingSVG from '$lib/Loading/index.svelte';

	import { onMount } from 'svelte';
	import { getCookieValue } from '$lib/store/utils';
	import type { CookieInfo1Type, UserDataType } from '$lib/store/types';
	import { fetchUserData } from '$lib/store/fetch';
	import { UserData } from '$lib/store/store';
	import HotQuesCom from '$lib/HotQues/index.svelte';
	import BlogListCom from '$lib/BlogList/index.svelte';

	import Search from '$lib/search/search.svelte';

	let loadingState: boolean = false;
	onMount(async () => {
		const CookieValueInfo1: string = getCookieValue('Info1');

		const InfoCookieData = JSON.parse(atob(CookieValueInfo1)) as CookieInfo1Type;
		console.log('🚀 ~ file: +page.ts:24 ~ load ~ InfoCookieData:', InfoCookieData);
		let UserDataValue = {} as UserDataType;
		UserData.subscribe((value) => {
			UserDataValue = value;
		});
		if (UserDataValue.UserID != InfoCookieData.UserID) {
			const GetUserData = await fetchUserData(InfoCookieData.UserID);
			console.log('🚀 ~ file: +page.ts:24 ~ InitializeData ~ GetUserData:', GetUserData);
			UserData.update(() => GetUserData);
		}
		loadingState = true;
	});
</script>

{#if loadingState}
	<!-- ////////////////////////////// -->

	<div
		class="   flex w-full flex-col justify-center overflow-x-hidden overflow-y-hidden bg-[#181818]"
	>
		<Navbar />
		<div class="flex w-full flex-row justify-center">
			<!-- COLLECTIVE -->
			<Collectives />
			<!-- QUESTION LIST -->

			<!-- settings -->
			<div class=" mb-2 mt-2 min-h-screen w-[1200px] bg-[#2d2d2d]   ">
				<div class="flex flex-row">
					<!-- questions -->
					<div class="flex w-[850px] flex-col">
						<!-- public ques & ask -->
						
						<!-- filter -->
						<!-- <QuesFilter/> -->

						<!-- all Question list -->
						<Search/>
						<!-- <PageNum /> -->
					</div>
					<!-- right sidebar -->
					<div class="flex h-full w-[350px] flex-col">
						<!-- Blog -->
						<!-- svelte-ignore missing-declaration -->
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
