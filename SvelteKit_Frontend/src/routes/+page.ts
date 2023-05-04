export const ssr = false; 

export const prerender = true; 
export const csr = true; 


import type { CookieInfo1Type, UserDataType } from '$lib/store/types';
import { fetchUserData } from '$lib/store/fetch';	
import { UserData } from '$lib/store/store';
import { getCookieValue } from '$lib/store/utils';


import type { PageLoad } from './$types';

export const load = (async () => {
	const CookieValueInfo1: string = getCookieValue('Info1');

	const InfoCookieData = JSON.parse(atob(CookieValueInfo1)) as CookieInfo1Type;
	console.log('🚀 ~ file: +page.ts:24 ~ load ~ InfoCookieData:', InfoCookieData);
	let UserDataValue = {} as UserDataType;
	UserData.subscribe((value) => {
		UserDataValue = value;
	});
	if (UserDataValue.ID != InfoCookieData.UUID) {
		const GetUserData = await fetchUserData(InfoCookieData.UUID);
		console.log('🚀 ~ file: +page.ts:24 ~ InitializeData ~ GetUserData:', GetUserData);
		UserData.update(() => GetUserData);
	}

	return {
		InfoCookieData
	};
}) satisfies PageLoad;



