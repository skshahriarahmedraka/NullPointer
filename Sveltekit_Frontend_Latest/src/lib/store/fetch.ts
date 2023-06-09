// function fetchWithAutoRetry(fetcher, maxRetryCount:number) {
//     return new Promise((resolve, reject) => {
//       let retries = 0;
//       const caller = () =>
//         fetcher().then((data) => {
//             resolve(data);
//           })
//           .catch((error) => {
//             if (retries < maxRetryCount) {
//               retries++;
//               caller();
//             } else {
//               reject(error);
//             }
//           });
//       retries = 1;
//       caller();
//     });
//   }

//   const fetchSouravProfile = async () => {
//     console.log("Fetching..");
//     const rawResponse = await fetch("https://api.github.com/users/sourav-singhh");
//     const jsonResponse = await rawResponse.json();
//     console.log(jsonResponse);
//     return jsonResponse;
//   };

//   fetchWithAutoRetry(fetchSouravProfile, 5);

// fetch user data using user id

// import { UserData } from './store';
import type {
	UserDataType,
	NotificationDataType,
	// BlogListType,
	GroupListType,
	FavouriteHashListType,
	QuestionDataType,
	UserFlairDataType,
	AnswerDataType,
	BlogDataType,
	QuesArrWithMetadataType,
	HashDataType,
	HashArrWithMetadata,
	BlogArrWithMetadataType
} from './types';

async function fetchUserData(UUID: string): Promise<UserDataType> {
	const response = await fetch(`/api/user/${UUID}`);
	if (!response.ok) {
		console.log(`Failed to fetch user data for user ${UUID}.`);

		return {} as UserDataType;
	}
	const userData: UserDataType = await response.json();
	console.log('🚀 ~ file: fetch.ts:50 ~ fetchUserData ~ userData:', userData);
	return userData;
	// const response = await fetch(`/api/user/${UUID}`);
	// if (!response.ok) {
	// 	console.log(`Failed to fetch user data for user ${UUID}.`);

	// 	return {} as UserDataType;
	// }
	// const userData: UserDataType = await response.json();
	// console.log("🚀 ~ file: fetch.ts:50 ~ fetchUserData ~ userData:", userData)
	// return new Promise<UserDataType>((resolve) => {
	// 	setTimeout(() => {

	// 	  resolve(userData);
	// 	}, 2000);
	//   });
}

async function fetchNotificationData(UUID: string): Promise<NotificationDataType[]> {
	const response = await fetch(`/api/user/${UUID}/notifications`);
	if (!response.ok) {
		console.log(`❌Failed to fetch /api/user/${UUID}/notifications`);
		return [] as NotificationDataType[];
	}
	const notificationData: NotificationDataType[] = await response.json();
	console.log(
		'🚀 ~ file: fetch.ts:61 ~ fetchNotificationData ~ notificationData:',
		notificationData
	);
	return notificationData;
}

async function fetchGroupList(UUID: string): Promise<GroupListType[]> {
	const response = await fetch(`/api/user/${UUID}/grouplist`);
	if (!response.ok) {
		console.log(`❌Failed to fetch /api/user/${UUID}/grouplist`);
		return [] as GroupListType[];
	}
	const GroupListData: GroupListType[] = await response.json();
	console.log('🚀 ~ file: fetch.ts:82 ~ fetchGroupList ~ GroupListData:', GroupListData);
	return GroupListData;
}
async function fetchFavouriteHashList(UUID: string): Promise<FavouriteHashListType[]> {
	const response = await fetch(`/api/user/${UUID}/favouritehashlist`);
	if (!response.ok) {
		console.log(`❌Failed to fetch /api/user/${UUID}/favouritehashlist`);
		return [] as FavouriteHashListType[];
	}
	const FavouriteHashListData: FavouriteHashListType[] = await response.json();
	console.log(
		'🚀 ~ file: fetch.ts:92 ~ fetchFavouriteHashList ~ FavouriteHashListData:',
		FavouriteHashListData
		);
		return FavouriteHashListData;
	}
	async function fetchBlogList(
		type: string,
		start: number,
		stop: number,
		order: number
	): Promise<BlogDataType[]> {
		const response = await fetch(
			`/api/b/list?type=${type}&start=${start}&stop=${stop}&order=${order}`
		);
		if (!response.ok) {
			console.log(`❌ /api/b/list?type=${type}&start=${start}&stop=${stop}&order=${order}`);
			// return [] as BlogDataType[];
		}
		const BlogDataList: BlogDataType[] = await response.json();
		console.log("🚀 ~ file: fetch.ts:100 ~ BlogDataList:", BlogDataList)
	
		return BlogDataList ;
	}

async function fetchBlogData(BlogID: string): Promise<BlogDataType> {
	const response = await fetch(`/api/b/${BlogID}`);
	if (!response.ok) {
		console.log(`❌Failed to fetch /api/b/${BlogID}`);
		return {} as BlogDataType;
	}
	const BlogData: BlogDataType = await response.json();
	console.log('🚀 ~ file: fetch.ts:103 ~ fetchBlogData ~ BlogData:', BlogData);
	return BlogData;
}
async function fetchPostBlog(BlogData : BlogDataType): Promise<BlogDataType> {
	const response = await fetch(`/api/b/write`,{
		method: 'POST',
		body: JSON.stringify(BlogData),
	});
	if (!response.ok) {
		console.log(`❌Failed to fetch /api/b/write`,BlogData);
		return {} as BlogDataType;
	}
	const ResBlogData: BlogDataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:148 ~ fetchPostBlog ~ ResBlogData:", ResBlogData)
	return ResBlogData;
}
async function fetchUpdateUserData(UpdatedUserData: UserDataType): Promise<UserDataType> {
	const response = await fetch(`/api/updateuser`, {
		method: 'POST',
		body: JSON.stringify(UpdatedUserData),
		headers: {
			'Content-Type': 'application/json'
		}
	});
	if (!response.ok) {
		console.log(`Failed to fetch /api/updateuser UpdatedUserData`);
		return {} as UserDataType;
	}
	const UpdatedData: UserDataType = await response.json();
	console.log('🚀 ~ file: fetch.ts:119 ~ fetchUpdateUserData ~ UpdatedData:', UpdatedData);
	return UpdatedData;
}

async function fetchAskQuestion(QuestionData: QuestionDataType): Promise<{ InsertedID: string }> {
	const response = await fetch(`/api/q/askquestion`, {
		method: 'POST',
		body: JSON.stringify(QuestionData),
		headers: {
			'Content-Type': 'application/json'
		}
	});
	if (!response.ok) {
		console.log(`❌ Failed to fetch /api/q/askquestion QuestionData`);
		return {} as { InsertedID: string };
	}
	const res: { InsertedID: string } = await response.json();
	console.log('🚀 ~ file: fetch.ts:151 ~ fetchAskQuestion ~ res:', res);
	return res;
}

async function fetchCreateHash(HashData: HashDataType): Promise<{ InsertedID: string }> {
	const response = await fetch(`/api/h/write`, {
		method: 'POST',
		body: JSON.stringify(HashData),
		headers: {
			'Content-Type': 'application/json'
		}
	});
	if (!response.ok) {
		console.log(`❌ Failed to fetch /api/q/askquestion QuestionData`);
		return {} as { InsertedID: string };
	}
	const res: { InsertedID: string } = await response.json();
	console.log('🚀 ~ file: fetch.ts:151 ~ fetchAskQuestion ~ res:', res);
	return res;
}

async function fetchHashData(QID: string): Promise<HashDataType> {
	// const response = await (typeof event !== 'undefined' ? event.fetch : fetch)(`/api/q/${QID}`);
	const response = await fetch(`/api/h/${QID}`);
	if (!response.ok) {
		console.log(`❌ Failed to fetch /api/h/${QID} HashData`);
		return {} as HashDataType;
	}
	const HashData:HashDataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:212 ~ fetchHashData ~ HashData:", HashData)
	
	return HashData;
}

async function fetchQuestionData(QID: string): Promise<QuestionDataType> {
	// const response = await (typeof event !== 'undefined' ? event.fetch : fetch)(`/api/q/${QID}`);
	const response = await fetch(`/api/q/${QID}`);
	if (!response.ok) {
		console.log(`❌ Failed to fetch /api/q/${QID} QuestionData`);
		return {} as QuestionDataType;
	}
	const QuestionDataResponse: QuestionDataType = await response.json();
	console.log(
		'🚀 ~ file: fetch.ts:164 ~ fetchQuestionData ~ QuestionDataResponse:',
		QuestionDataResponse
	);
	return QuestionDataResponse;
}

async function fetchPublicQuestionDataArr(
	type: string,
	start: number,
	stop: number,
	order: number
): Promise<QuestionDataType[]> {
	const response = await fetch(
		`/api/public/q?type=${type}&start=${start}&stop=${stop}&order=${order}`
	);
	console.log(` /api/public/q?type=${type}&start=${start}&stop=${stop}&order=${order}`);
	if (!response.ok) {
		console.log(`❌ /api/public/q?type=${type}&start=${start}&stop=${stop}&order=${order}`);
		return [] as QuestionDataType[];
	}
	const QuestionDataResponseArr: QuestionDataType[] = await response.json();
	console.log(
		'🚀 ~ file: fetch.ts:173 ~ fetchPublicQuestionData ~ QuestionDataResponseArr:',
		QuestionDataResponseArr
	);
	return QuestionDataResponseArr;
}

async function fetchQuesArrWithMetadata(
	type: string,
	start: number,
	stop: number,
	order: number
): Promise<QuesArrWithMetadataType> {
	const response = await fetch(
		`/api/public/meta/q?type=${type}&start=${start}&stop=${stop}&order=${order}`
	);
	console.log("🚀 ~ file: fetch.ts:231 ~ response:", response)
	console.log(`QuesArrWithMetadataType /api/public/meta/q?type=${type}&start=${start}&stop=${stop}&order=${order}`);
	if (!response.ok) {
		console.log(`❌QuesArrWithMetadataType /api/public/meta/q?type=${type}&start=${start}&stop=${stop}&order=${order}`);
		return {} as QuesArrWithMetadataType;
	}
	const QuestionDataResponseArr: QuesArrWithMetadataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:236 ~ QuesArrWithMetadataType:", QuestionDataResponseArr)
	
	return QuestionDataResponseArr;
}

async function fetchHashArrWithMetadata(
	type: string,
	start: number,
	stop: number,
	order: number
): Promise<HashArrWithMetadata> {
	const response = await fetch(
		`/api/h/list?type=${type}&start=${start}&stop=${stop}&order=${order}`
	);
	console.log("🚀 ~ file: fetch.ts:231 ~ response:", response)
	console.log(`HashArrWithMetadata /api/h/list?type=${type}&start=${start}&stop=${stop}&order=${order}`);
	if (!response.ok) {
		console.log(`HashArrWithMetadata /api/h/list?type=${type}&start=${start}&stop=${stop}&order=${order}`);
		return {} as HashArrWithMetadata;
	}
	const HashDataResponseArr: HashArrWithMetadata = await response.json();
	console.log("🚀 ~ file: fetch.ts:292 ~ HashDataResponseArr:", HashDataResponseArr)
	
	return HashDataResponseArr;
}

async function fetchUserFlairData(UUID: string): Promise<UserFlairDataType> {
	const response = await fetch(`/api/user/flair/${UUID}`);
	if (!response.ok) {
		console.log(
			'❌ Failed ~ file: fetch.ts:180 ~ fetchUserFlairData ~ /api/user/flair/${UserID}:',
			UUID
		);

		return {
			UserID: '',
			UserURL: '',
			UserName: '',
			UserImage: '',
			Badges: {
				Reputation: 0,
				Gold: 0,
				Silver: 0,
				Bronze: 0
			},
			Location: '',
			Aboutme: ''
		} as UserFlairDataType;
	}
	const userFlairData: UserFlairDataType = await response.json();
	console.log('🚀 ~ file: fetch.ts:187 ~ fetchUserFlairData ~ userFlairData:', userFlairData);
	return userFlairData;
}

async function fetchPostAnsData(
	QuesID: string,
	PostAnsData: AnswerDataType
): Promise<AnswerDataType> {
	const response = await fetch(`/api/q/${QuesID}/answer`, {
		method: 'POST',
		body: JSON.stringify(PostAnsData),
		headers: {
			'Content-Type': 'application/json'
		}
	});
	if (!response.ok) {
		console.log(`Failed to fetch /api/q/answer/${QuesID} AnswerDataType`,PostAnsData);
		return {} as AnswerDataType;
	}
	const UpdatedData: AnswerDataType = await response.json();
	console.log('🚀 ~ file: fetch.ts:220 ~ fetchPostAnsData ~ UpdatedData:', UpdatedData);
	return UpdatedData;
}

async function fetchAnsData(UUID: string): Promise<AnswerDataType> {
	const response = await fetch(`/api/q/answer/${UUID}`);
	if (!response.ok) {
		console.log("🚀 ~ file: fetch.ts:250 ~ fetchAnsData ~ /api/q/answer/${UUID}:", UUID)
		

		return {
			
		} as AnswerDataType;
	}
	const AnsData: AnswerDataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:258 ~ fetchAnsData ~ AnsData:", AnsData)
	return AnsData;
}


async function fetchSearchQuesArrWithMetadata(
	SearchedString: string,
	type: string,
	start: number,
	stop: number,
	order: number
): Promise<QuesArrWithMetadataType> {
	const response = await fetch(
		`/api/search/q/${SearchedString}?type=${type}&start=${start}&stop=${stop}&order=${order}`
	);
	console.log("🚀 ~ file: fetch.ts:231 ~ response:", response)
	console.log(`Searched QuesArrWithMetadataType /api/search/q/${SearchedString}?type=${type}&start=${start}&stop=${stop}&order=${order}`);
	if (!response.ok) {
		console.log(`❌Searched QuesArrWithMetadataType  /api/search/q/${SearchedString}?type=${type}&start=${start}&stop=${stop}&order=${order}`);
		return {} as QuesArrWithMetadataType;
	}
	const QuestionDataResponseArr: QuesArrWithMetadataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:236 ~ QuesArrWithMetadataType:", QuestionDataResponseArr)
	
	return QuestionDataResponseArr;
}

async function fetchSearchBlogArrWithMetadata(
	SearchedString: string,
	type: string,
	start: number,
	stop: number,
	order: number
): Promise<BlogArrWithMetadataType> {
	const response = await fetch(
		`/api/search/b/${SearchedString}?type=${type}&start=${start}&stop=${stop}&order=${order}`
	);
	console.log("🚀 ~ file: fetch.ts:393 ~ response:", response)
	console.log(`Searched BlogArrWithMetadataType /api/search/b/${SearchedString}?type=${type}&start=${start}&stop=${stop}&order=${order}`);
	if (!response.ok) {
		console.log(`❌Searched BlogArrWithMetadataType  /api/search/b/${SearchedString}?type=${type}&start=${start}&stop=${stop}&order=${order}`);
		return {} as BlogArrWithMetadataType;
	}
	const blogDataResponseArr: BlogArrWithMetadataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:401 ~ blogDataResponseArr:", blogDataResponseArr)
	
	return blogDataResponseArr;
}

async function fetchSearchhashArrWithMetadata(
	SearchedString: string,
	type: string,
	start: number,
	stop: number,
	order: number
): Promise<HashArrWithMetadata> {
	const response = await fetch(
		`/api/search/h/${SearchedString}?type=${type}&start=${start}&stop=${stop}&order=${order}`
	);
	console.log(`Searched HashArrWithMetadataType /api/search/h/${SearchedString}?type=${type}&start=${start}&stop=${stop}&order=${order}`);
	if (!response.ok) {
		console.log(`❌Searched HashArrWithMetadataType  /api/search/h/${SearchedString}?type=${type}&start=${start}&stop=${stop}&order=${order}`);
		return {} as HashArrWithMetadata;
	}
	const HashDataResponseArr: HashArrWithMetadata = await response.json();
	console.log("🚀 ~ file: fetch.ts:422 ~ HashDataResponseArr:", HashDataResponseArr)
	
	return HashDataResponseArr;
}


async function fetchHashViewData(HashViewData: HashViewDataType)  {
	const response = await fetch(
		`/follow/hash`,{
			method: 'POST',
			body: JSON.stringify(HashViewData),
		}
	);
	console.log("🚀 ~ file: fetch.ts:435 ~ fetchHashViewData ~ response:", response)
	if (!response.ok) {
		console.log(`❌/follow/hash`);
		return {} as HashViewDataType;
	}
	const res: {message:boolean} = await response.json();
	
	return res;
}

export {
	fetchUserData,
	fetchUpdateUserData,
	fetchUserFlairData,

	fetchBlogList,
	fetchBlogData,
	fetchPostBlog,
	
	fetchAskQuestion,
	fetchQuestionData,
	fetchAnsData,
	fetchPostAnsData,
	fetchPublicQuestionDataArr,
	fetchQuesArrWithMetadata,
	
	fetchNotificationData,
	fetchGroupList,
	fetchFavouriteHashList,
	
	fetchCreateHash,
	fetchHashData,
	fetchHashArrWithMetadata,
	
	fetchSearchQuesArrWithMetadata,
	fetchSearchBlogArrWithMetadata,
	fetchSearchhashArrWithMetadata,

	fetchHashViewData,
};
