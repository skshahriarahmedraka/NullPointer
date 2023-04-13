export const ssr = false; // all the code will be executed in client side

export const prerender = false; // if true whole page will be generated as a html page in server side
export const csr = true; // the component will only be rendered on the client-side, after the initial HTML page has been loaded.

// import { fetchUserData } from '$lib/Store/fetch';
import type { CookieInfo1Type } from '$lib/store/types';
import { fetchUserData } from '$lib/store/fetch';
import { UserData } from '$lib/store/store';

// async function InitializeData() {

	
// 	return UserData;
// }

import type { PageLoad } from './$types';
 
export const load = (async () => {
    
	const CookieValueInfo1:string = getCookieValue('Info1');


	const jsonObject = JSON.parse(atob(CookieValueInfo1)) as CookieInfo1Type;
	console.log("🚀 ~ file: +page.ts:24 ~ load ~ jsonObject:", jsonObject)

	const GetUserData = await fetchUserData(jsonObject.UUID);
	console.log("🚀 ~ file: +page.ts:24 ~ InitializeData ~ GetUserData:", GetUserData)
	UserData.update(()=>GetUserData)
  return {
  };
}) satisfies PageLoad;


function getCookieValue(cookieName:string) {
	// Split the cookies into an array
	const cookies:string[] = document.cookie.split(';');
	
	// Loop through the array to find the cookie with the given name
	for (let i = 0; i < cookies.length; i++) {
	  const cookie:string = cookies[i].trim();
	  
	  // Check if the cookie name matches the given name
	  if (cookie.indexOf(cookieName + '=') == 0) {
		// Return the cookie value
		return cookie.substring(cookieName.length + 1, cookie.length);
	  }
	}
	
	// If the cookie is not found, return null
	return "";
  }
  