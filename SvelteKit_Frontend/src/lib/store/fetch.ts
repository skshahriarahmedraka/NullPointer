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
	FavouriteHashListType
} from './types';

async function fetchUserData(UUID: string): Promise<UserDataType> {
	const response = await fetch(`/api/user/${UUID}`);
	if (!response.ok) {
		console.log(`Failed to fetch user data for user ${UUID}.`);

		return {} as UserDataType;
	}
	const userData: UserDataType = await response.json();
	return userData;
}

async function fetchNotificationData(UUID: string): Promise<NotificationDataType[]> {
	const response = await fetch(`/api/user/${UUID}/notifications`);
	if (!response.ok) {
		console.log(`Failed to fetch /api/user/${UUID}/notifications`);
		return  [] as NotificationDataType[];
	}
	const notificationData: NotificationDataType[] = await response.json();
	return notificationData;
}

async function fetchBlogList(UUID: string): Promise<BlogListType[]> {
	const response = await fetch(`/api/user/${UUID}/bloglist`);
	if (!response.ok) {
		console.log(`Failed to fetch /api/user/${UUID}/bloglist`);
		return [] as BlogListType[];
	}
	const BlogListData: BlogListType[] = await response.json();
	return BlogListData;
}
async function fetchGroupList(UUID: string): Promise<GroupListType[]> {
	const response = await fetch(`/api/user/${UUID}/grouplist`);
	if (!response.ok) {
		console.log(`Failed to fetch /api/user/${UUID}/grouplist`);
		return [] as GroupListType[];
	}
	const GroupListData: GroupListType[] = await response.json();
	return GroupListData;
}
async function fetchFavouriteHashList(UUID: string): Promise<FavouriteHashListType[]> {
	const response = await fetch(`/api/user/${UUID}/favouritehashlist`);
	if (!response.ok) {
		console.log(`Failed to fetch /api/user/${UUID}/favouritehashlist`);
		return [] as FavouriteHashListType[];
	}
	const FavouriteHashListData: FavouriteHashListType[] = await response.json();
	return FavouriteHashListData;
}

async function fetchBlogData(BlogID: string): Promise<FavouriteHashListType[]> {
	const response = await fetch(`/api/blog/${BlogID}`);
	if (!response.ok) {
		console.log(`Failed to fetch /api/blog/${BlogID}`);
		return [] as FavouriteHashListType[];
	}
	const BlogData: FavouriteHashListType[] = await response.json();
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
	return UpdatedData;
}

export {
	fetchUserData,
	fetchNotificationData,
	fetchBlogList,
	fetchGroupList,
	fetchFavouriteHashList,
	fetchBlogData,
	fetchUpdateUserData
};
