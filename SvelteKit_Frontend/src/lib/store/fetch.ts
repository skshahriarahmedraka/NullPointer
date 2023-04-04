
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

import type { UserDataType,NotificationDataType } from './types';
async function fetchUserData(userId: string): Promise<UserDataType> {
    const response = await fetch(`/api/users/${userId}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch user data for user ${userId}.`);
    }
    const userData: UserDataType = await response.json();
    return userData;
  }

async function fetchNotificationData(userId: string): Promise<NotificationDataType[]> {
    const response = await fetch(`/api/users/${userId}/notifications`);
    if (!response.ok) {
      throw new Error(`Failed to fetch notification data for user ${userId}.`);
    }
    const notificationData: NotificationDataType[] = await response.json();
    return notificationData;
  }

export {fetchUserData, fetchNotificationData}