
// import { fetchUserData } from '$lib/Store/fetch';
import type { QuestionDataType } from '$lib/store/types';
// import { fetchQuestionData } from '$lib/store/fetch';
// import {  UserData } from '$lib/store/store';
// import { getCookieValue } from '$lib/store/utils';
// async function InitializeData() {

// 	return UserData;
// }

import type { PageLoad } from './$types';

export const load = (async ({fetch,params}) => {
	// const CookieValueInfo1: string = getCookieValue('Info1');

	// const InfoCookieData = JSON.parse(atob(CookieValueInfo1)) as CookieInfo1Type;
	// console.log('🚀 ~ file: +page.ts:24 ~ load ~ InfoCookieData:', InfoCookieData);
	// // let UserDataValue = {} as UserDataType;
	// // UserData.subscribe((value) => {
	// // 	UserDataValue = value;
	// // });
	// let QuestionData:QuestionDataType = {} as QuestionDataType;
	// // if (UserDataValue.ID != InfoCookieData.UUID) {

		
	// 		const GetUserData:UserDataType = await fetchUserData(InfoCookieData.UUID);
	// 		console.log("🚀 ~ file: +page.ts:32 ~ fetch1 ~ GetUserData:", GetUserData)
	// 		// console.log('🚀 ~ file: +page.ts:24 ~ InitializeData ~ GetUserData:', GetUserData);
	// 		UserData.update(() => GetUserData);
		
		
		
	// }
	let  QuestionData:QuestionDataType= {} as QuestionDataType;
	if (typeof document !== 'undefined') {

		const response = await fetch(`/api/q/${params.QID}`);
		if (!response.ok) {
			console.log(`❌ Failed to fetch params.QID /api/q/${params.QID} QuestionData`);
			// return {} 
		}
		 QuestionData = await response.json();
		console.log("🚀 ~ file: +page.ts:42 ~ load ~ QuestionData:", QuestionData)
		
	}

		// QuestionData= await fetchQuestionData(params.QID) 
	
	

	return {
		// InfoCookieData,
		QuestionData
	};
}) satisfies PageLoad;
