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
		throw new Error(`Failed to fetch user data for user ${UUID}.`);
	}
	const userData: UserDataType = await response.json();
	return userData;
}

async function fetchNotificationData(UUID: string): Promise<NotificationDataType[]> {
	const response = await fetch(`/api/user/${UUID}/notifications`);
	if (!response.ok) {
		throw new Error(`Failed to fetch /api/user/${UUID}/notifications.`);
	}
	const notificationData: NotificationDataType[] = await response.json();
	return notificationData;
}

async function fetchBlogList(UUID: string): Promise<BlogListType[]> {
	const response = await fetch(`/api/user/${UUID}/bloglist`);
	if (!response.ok) {
		throw new Error(`Failed to fetch /api/user/${UUID}/bloglist`);
	}
	const BlogListData: BlogListType[] = await response.json();
	return BlogListData;
}
async function fetchGroupList(UUID: string): Promise<GroupListType[]> {
	const response = await fetch(`/api/user/${UUID}/grouplist`);
	if (!response.ok) {
		throw new Error(`Failed to fetch /api/user/${UUID}/grouplist`);
	}
	const GroupListData: GroupListType[] = await response.json();
	return GroupListData;
}
async function fetchFavouriteHashList(UUID: string): Promise<FavouriteHashListType[]> {
	const response = await fetch(`/api/user/${UUID}/favouritehashlist`);
	if (!response.ok) {
		throw new Error(`Failed to fetch /api/user/${UUID}/favouritehashlist`);
	}
	const FavouriteHashListData: FavouriteHashListType[] = await response.json();
	return FavouriteHashListData;
}

async function fetchBlogData(BlogID: string): Promise<FavouriteHashListType[]> {
	const response = await fetch(`/api/blog/${BlogID}`);
	if (!response.ok) {
		throw new Error(`Failed to fetch /api/blog/${BlogID}`);
	}
	const BlogData: FavouriteHashListType[] = await response.json();
	return BlogData;
}

export {
	fetchUserData,
	fetchNotificationData,
	fetchBlogList,
	fetchGroupList,
	fetchFavouriteHashList,
	fetchBlogData
};
