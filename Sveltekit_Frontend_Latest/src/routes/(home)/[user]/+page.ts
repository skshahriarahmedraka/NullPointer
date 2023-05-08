// export const ssr = false; // all the code will be executed in client side

// export const prerender = false; // if true whole page will be generated as a html page in server side
// export const csr = true; // the component will only be rendered on the client-side, after the initial HTML page has been loaded.

// import { fetchUserData } from '$lib/Store/fetch';
// import type { CookieInfo1Type, UserDataType } from '$lib/store/types';
// import { fetchUserData } from '$lib/store/fetch';
// import { UserData } from '$lib/store/store';
// import { getCookieValue } from '$lib/store/utils';
// async function InitializeData() {

// 	return UserData;
// }

import type { PageLoad } from './$types';

export const load = (async () => {
	
	return {
		// InfoCookieData
	};
}) satisfies PageLoad;
