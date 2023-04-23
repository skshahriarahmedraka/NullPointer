// export const ssr = false ; // all the code will be executed in client side

// export const prerender = false; // if true whole page will be generated as a html page in server side
// export const csr = true; // the component will only be rendered on the client-side, after the initial HTML page has been loaded.


import type { PageServerLoad } from './$types';
// import { redirect } from '@sveltejs/kit'
// import * as jwt from 'jsonwebtoken';
// import { UserProData } from '$lib/Store/store';
// import { redirect } from '@sveltejs/kit';


export const  load: PageServerLoad = async ({locals}) => {
	
    console.log(
        '🚀 ~ file: +page.server.ts ~ line 10 ~ constload:PageServerLoad= ~ user',
        locals.user
    );
	// if (!locals.user.Authenticated) {
	// 	throw redirect(302, '/login');
	// }

	//   console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies);
	//   console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies.get('Auth1'));
	//   const MyCookie = cookies.get('Auth1') || '';
	//   const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;
	//   let Userdata:any 
	//   if (MyCookie!= ''){
		  
	// 	  const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
	// 	  console.log("decoded: ",decoded);
	// 	//   let resdata
	// 	  console.log(`http://${process.env.GO_HOST}/user/${decoded.UserID}`)
	// 	  await fetch(`http://${process.env.GO_HOST}/user/${decoded.UserID}`,{
	// 		  mode:"no-cors"
	// 	  }).then((res)=>{
	// 		  return res.json()
	// 	  }).then((d)=>{
			  
	// 		Userdata=d
			
	// 	  })
  
		  
	//   }
	// return {
	// 	Userdata
	// }
}