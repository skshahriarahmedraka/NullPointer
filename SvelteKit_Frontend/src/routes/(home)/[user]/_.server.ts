// // import { error, json } from '@sveltejs/kit';
// import * as jwt from 'jsonwebtoken';

// import { redirect } from '@sveltejs/kit';
// import type { PageServerLoad } from './$types';

// /** @type {import('./$types').PageServerLoad} */
// export const load: PageServerLoad = async ({ cookies, locals }) => {
// 	if (!locals.user.Authenticated) {
// 		console.log(
// 			'🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user',
// 			locals.user
// 		);
// 		throw redirect(302, '/login');
// 	}
// 	const MyCookie = cookies.get('Auth1') || '';
// 	const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;

	

// 	if (MyCookie != '') {
// 		interface tokeninterface{
// 			Email: string 
// 			Name: string,
// 			UserID: string,
// 			Accounttype: string,
// 			exp: number
// 		  }

// 		const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
// 		console.log('🚀 ~ file: +page.server.ts ~ line 136 ~ myfetch ~ decoded', decoded);
// 		if ((decoded as tokeninterface).Accounttype != 'admin') {
// 			throw redirect(302, '/');
// 		}

// 	}
// 	return {}
	
// };
