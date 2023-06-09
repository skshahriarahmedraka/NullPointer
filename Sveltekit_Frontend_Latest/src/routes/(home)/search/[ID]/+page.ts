

import type { PageLoad } from './$types';

// export const load = (async ({params,url,fetch}) => {
	export const load = (async ({params}) => {
	
	// const type :string = url.searchParams.get("type")
	// const start :string = url.searchParams.get("start")
	// const stop :string = url.searchParams.get("stop")
	// const order :string = url.searchParams.get("order")

	// let  QuestionList:QuesArrWithMetadataType= {} as QuesArrWithMetadataType;
	// if (typeof document !== 'undefined') {
		
	// 	const response = await fetch(`/api/search/q/${params.ID}?type=${type}&start=${start}&stop=${stop}&order=${order}`);
	// 	console.log("🚀 ~ file: +page.ts:33 ~ load ~ response:", response)
	// 	if (!response.ok) {
	// 		console.log(`❌ Failed to fetch params.ID /api/search/q/${params.ID}?type=${type}&start=${start}&stop=${stop}&order=${order} `);
	// 		// return {} 
	// 	}
	// 	QuestionList = await response.json();
	// 	console.log("🚀 ~ file: +page.ts:38 ~ load ~ QuestionList:", QuestionList)
		
	// }

	const SearchedString:string = params.ID;
	console.log("🚀 ~ file: +page.ts:28 ~ load ~ SearchedString:", SearchedString)

	return {
		// InfoCookieData
		SearchedString
	};
}) satisfies PageLoad;
