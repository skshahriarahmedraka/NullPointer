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
	BlogListType,
	GroupListType,
	FavouriteHashListType,
	QuestionDataType,
	UserFlairDataType,
	AnswerDataType
} from './types';

async function fetchUserData(UUID: string): Promise<UserDataType> {
	const response = await fetch(`/api/user/${UUID}`);
	if (!response.ok) {
		console.log(`Failed to fetch user data for user ${UUID}.`);

		return {} as UserDataType;
	}
	const userData: UserDataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:50 ~ fetchUserData ~ userData:", userData)
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
		return  [] as NotificationDataType[];
	}
	const notificationData: NotificationDataType[] = await response.json();
	console.log("🚀 ~ file: fetch.ts:61 ~ fetchNotificationData ~ notificationData:", notificationData)
	return notificationData;
}

async function fetchBlogList(UUID: string): Promise<BlogListType[]> {
	const response = await fetch(`/api/user/${UUID}/bloglist`);
	if (!response.ok) {
		console.log(`❌Failed to fetch /api/user/${UUID}/bloglist`);
		return [] as BlogListType[];
	}
	const BlogListData: BlogListType[] = await response.json();
	console.log("🚀 ~ file: fetch.ts:72 ~ fetchBlogList ~ BlogListData:", BlogListData)
	return BlogListData;
}
async function fetchGroupList(UUID: string): Promise<GroupListType[]> {
	const response = await fetch(`/api/user/${UUID}/grouplist`);
	if (!response.ok) {
		console.log(`❌Failed to fetch /api/user/${UUID}/grouplist`);
		return [] as GroupListType[];
	}
	const GroupListData: GroupListType[] = await response.json();
	console.log("🚀 ~ file: fetch.ts:82 ~ fetchGroupList ~ GroupListData:", GroupListData)
	return GroupListData;
}
async function fetchFavouriteHashList(UUID: string): Promise<FavouriteHashListType[]> {
	const response = await fetch(`/api/user/${UUID}/favouritehashlist`);
	if (!response.ok) {
		console.log(`❌Failed to fetch /api/user/${UUID}/favouritehashlist`);
		return [] as FavouriteHashListType[];
	}
	const FavouriteHashListData: FavouriteHashListType[] = await response.json();
	console.log("🚀 ~ file: fetch.ts:92 ~ fetchFavouriteHashList ~ FavouriteHashListData:", FavouriteHashListData)
	return FavouriteHashListData;
}

async function fetchBlogData(BlogID: string): Promise<FavouriteHashListType[]> {
	const response = await fetch(`/api/blog/${BlogID}`);
	if (!response.ok) {
		console.log(`❌Failed to fetch /api/blog/${BlogID}`);
		return [] as FavouriteHashListType[];
	}
	const BlogData: FavouriteHashListType[] = await response.json();
	console.log("🚀 ~ file: fetch.ts:103 ~ fetchBlogData ~ BlogData:", BlogData)
	return BlogData;
}
async function fetchUpdateUserData(UpdatedUserData : UserDataType): Promise<UserDataType> {
	const response = await fetch(`/api/updateuser`,{
		method: 'POST',
		body: JSON.stringify(UpdatedUserData),
		headers:{
		  'Content-Type': 'application/json'
		}
	  });
	if (!response.ok) {
		console.log(`Failed to fetch /api/updateuser UpdatedUserData`);
		return {} as UserDataType;
	}
	const UpdatedData:UserDataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:119 ~ fetchUpdateUserData ~ UpdatedData:", UpdatedData)
	return UpdatedData;
}

async function fetchAskQuestion(QuestionData: QuestionDataType) : Promise<{"InsertedID":string}> {
	const response = await fetch(`/api/q/askquestion`,{
		method: 'POST',
		body: JSON.stringify(QuestionData),
		headers:{
		  'Content-Type': 'application/json'
		}
	  });
	if (!response.ok) {
		console.log(`❌ Failed to fetch /api/q/askquestion QuestionData`);
		return {} as {"InsertedID":string};
	}
	const res:{"InsertedID":string}= await response.json();
	console.log("🚀 ~ file: fetch.ts:151 ~ fetchAskQuestion ~ res:", res)
	return res;
}

async function fetchQuestionData(QID: string) : Promise<QuestionDataType> {
	// const response = await (typeof event !== 'undefined' ? event.fetch : fetch)(`/api/q/${QID}`);
	const response = await fetch(`/api/q/${QID}`);
	if (!response.ok) {
		console.log(`❌ Failed to fetch /api/q/${QID} QuestionData`);
		return {} as QuestionDataType;
	}
	const QuestionDataResponse:QuestionDataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:164 ~ fetchQuestionData ~ QuestionDataResponse:", QuestionDataResponse)
	return QuestionDataResponse;
}

async function fetchPublicQuestionDataArr(type: string,start:number, stop :number ,order:number) : Promise<QuestionDataType[]> {
	const response = await fetch(`/api/public/q?type=${type}&start=${start}&stop=${stop}&order=${order}`);
	if (!response.ok) {
		console.log(`❌ /api/public/q?type=${type}&start=${start}&stop=${stop}&order=${order}`);
		return [] as QuestionDataType[];
	}
	const QuestionDataResponseArr:QuestionDataType[] = await response.json();
	console.log("🚀 ~ file: fetch.ts:173 ~ fetchPublicQuestionData ~ QuestionDataResponseArr:", QuestionDataResponseArr)
	return QuestionDataResponseArr;
}

async function fetchUserFlairData(UUID: string): Promise<UserFlairDataType> {
	const response = await fetch(`/api/user/flair/${UUID}`);
	if (!response.ok) {
		console.log("❌ Failed ~ file: fetch.ts:180 ~ fetchUserFlairData ~ /api/user/flair/${UUID}:", UUID)
	

		return { 
			UserID :  "",           
			UserURL    : "",             
				UserName : "",             
				UserImage : "" ,
				Badges  :   {
					Reputation : 0 ,
					Gold       : 0,
					Silver    : 0,
					Bronze     : 0,
				} ,
				Location   :    ""  ,           
				Aboutme     :   ""  ,            
			
		 } as UserFlairDataType;
	}
	const userFlairData: UserFlairDataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:187 ~ fetchUserFlairData ~ userFlairData:", userFlairData)
	return userFlairData;
	
}

async function fetchPostAnsData(QuesID:string,PostAnsData : AnswerDataType): Promise<AnswerDataType> {
	const response = await fetch(`/api/q/${QuesID}/answer`,{
		method: 'POST',
		body: JSON.stringify(PostAnsData),
		headers:{
		  'Content-Type': 'application/json'
		}
	  });
	if (!response.ok) {
		console.log(`Failed to fetch /api/q/answer/${QuesID} AnswerDataType`);
		return {} as AnswerDataType;
	}
	const UpdatedData:AnswerDataType = await response.json();
	console.log("🚀 ~ file: fetch.ts:220 ~ fetchPostAnsData ~ UpdatedData:", UpdatedData)
	return UpdatedData;
}


export {
	fetchUserData,
	fetchNotificationData,
	fetchBlogList,
	fetchGroupList,
	fetchFavouriteHashList,
	fetchBlogData,
	fetchUpdateUserData,
	fetchAskQuestion,
	fetchQuestionData,
	fetchPublicQuestionDataArr,
	fetchUserFlairData,
	fetchPostAnsData
};
