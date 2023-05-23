// export const ssr = false; // all the code will be executed in client side

// export const prerender = false; // if true whole page will be generated as a html page in server side
// export const csr = true; // the component will only be rendered on the client-side, after the initial HTML page has been loaded.

// import { fetchUserData } from '$lib/Store/fetch';
import type { CookieInfo1Type, UserDataType } from '$lib/store/types';
import { fetchUserData } from '$lib/store/fetch';
import { UserData } from '$lib/store/store';
import { getCookieValue } from '$lib/store/utils';
// async function InitializeData() {

// 	return UserData;
// }

import type { PageLoad } from './$types';

export const load = (async () => {
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

	return {
		InfoCookieData
	};
}) satisfies PageLoad;
