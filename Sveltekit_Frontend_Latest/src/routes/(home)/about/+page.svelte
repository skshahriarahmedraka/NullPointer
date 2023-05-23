<script lang="ts">
	import { fetchUserData } from '$lib/store/fetch';
	import { UserData } from '$lib/store/store';
	import type { CookieInfo1Type, UserDataType } from '$lib/store/types';
	import { getCookieValue } from '$lib/store/utils';
	import { onMount } from 'svelte';

	// your script goes here
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

<style>
	/* your styles go here */
</style>
