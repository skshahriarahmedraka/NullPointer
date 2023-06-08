// export const ssr = false; // all the code will be executed in client side

// export const prerender = false; // if true whole page will be generated as a html page in server side
// export const csr = true; // the component will only be rendered on the client-side, after the initial HTML page has been loaded.

// import { fetchUserData } from '$lib/Store/fetch';
import type {  HashDataType } from '$lib/store/types';
// import { fetchBlogData } from '$lib/store/fetch';
// import { UserData } from '$lib/store/store';
// import { getCookieValue } from '$lib/store/utils';
// async function InitializeData() {

// 	return UserData;
// }

import type { PageLoad } from './$types';

export const load = (async ({ params, fetch }) => {


	let HashData = {} as HashDataType;
	// BlogData = await fetchBlogData(params.ID);

	if (typeof document !== 'undefined') {
		const response = await fetch(`/api/h/${params.ID}`);
	if (!response.ok) {
		console.log(`❌ Failed to fetch /api/h/${params.ID} HashData`);
		return {} as HashDataType;
	}
	HashData = await response.json();
	console.log("🚀 ~ file: fetch.ts:212 ~ fetchHashData ~ HashData:", HashData)
		// return BlogData;
	}

	return {
		// InfoCookieData,
		HashData
	};
}) satisfies PageLoad;
