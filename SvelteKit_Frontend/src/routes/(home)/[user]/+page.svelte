<script lang="ts">
	import UserInfo from '$lib/UserInfo/Title.svelte';
	import Stats from '$lib/UserInfo/Stats.svelte';
	import Abouts from '$lib/UserInfo/Abouts.svelte';
	import { UserData } from '$lib/store/store';
	import TopTags from '$lib/UserInfo/TopTags.svelte';
	import Badges from '$lib/UserInfo/Badges.svelte';
	import GroupList from '$lib/UserInfo/GroupList.svelte';
	import HashList from '$lib/UserInfo/HashList.svelte';

	// export let UserData:any
	// console.log("🚀 ~ file: index@root.svelte ~ line 47 ~ UserData", UserData)
</script>

<!-- settings -->
<div class=" mt-2 flex max-h-full  min-h-screen w-[1100px] flex-col bg-[#2d2d2d]   ">
	<!-- image & name & social media -->
	<UserInfo />
	<!-- Button : profile, activity, -->
	<div class="  flex w-full flex-row items-center justify-start gap-5 pl-10 ">
		<div
			on:click={() => ($UserData['SelectedPanel'] = 'Profile')}
			on:keydown={() => {}}
			class="ml-3 -mr-3 h-9 w-24 rounded-xl {$UserData['SelectedPanel'] === 'Profile'
				? 'bg-sky-500 hover:bg-blue-600'
				: 'border-2 border-blue-600 '}  flex items-center justify-center hover:cursor-pointer "
		>
			<p
				class=" my-auto  text-base font-semibold {$UserData['SelectedPanel'] === 'Profile'
					? 'text-gray-200'
					: 'text-sky-500'}"
			>
				Profile
			</p>
		</div>
		<div
			on:click={() => ($UserData['SelectedPanel'] = 'Activity')}
			on:keydown={() => {}}
			class="ml-3 -mr-3 h-9 w-24 rounded-xl {$UserData['SelectedPanel'] === 'Activity'
				? 'bg-sky-500 hover:bg-blue-600'
				: 'border-2 border-blue-600 '}   flex items-center justify-center hover:cursor-pointer "
		>
			<p
				class=" my-auto  text-base font-semibold {$UserData['SelectedPanel'] === 'Activity'
					? 'text-gray-200'
					: 'text-sky-500'}"
			>
				Activity
			</p>
		</div>
		<div
			on:click={() => ($UserData['SelectedPanel'] = 'Settings')}
			on:keydown={() => {}}
			class="ml-3 -mr-3 h-9 w-24 rounded-xl {$UserData['SelectedPanel'] === 'Settings'
				? 'bg-sky-500 hover:bg-blue-600'
				: 'border-2 border-blue-600 '}  flex items-center justify-center hover:cursor-pointer "
		>
			<p
				class=" my-auto  text-base font-semibold {$UserData['SelectedPanel'] === 'Settings'
					? 'text-gray-200'
					: 'text-sky-500'}"
			>
				Settings
			</p>
		</div>
	</div>
	<!-- details about me -->
	{#if $UserData['SelectedPanel'] === 'Profile'}
		<div class="mt-5 flex h-full w-full flex-row ">
			<div class="flex h-full w-[300px] flex-col ">
				<Stats />
				<TopTags />
				<GroupList />
				<HashList/>
			</div>
			<!-- Stats and Tags -->
			<!-- About -->
			<div class="flex h-full w-9/12 flex-col  ">
				<Abouts />
				<Badges />
			</div>
		</div>
		<!-- content here -->
	{:else if $UserData['SelectedPanel'] === 'Activity'}
		<div class="">Activity</div>
	{:else}
		<div class="">Settings</div>
		<!-- else content here -->
	{/if}
</div>

<style>
	/* your styles go here */
</style>
