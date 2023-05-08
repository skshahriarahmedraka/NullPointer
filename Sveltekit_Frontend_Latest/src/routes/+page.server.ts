


import type { PageServerLoad } from './$types';

import { redirect } from '@sveltejs/kit';

export const load: PageServerLoad = async ({  parent ,locals}) => {
	await parent();
	if (!locals.user.Authenticated) {
		console.log(
			'🚀 ~ file: +page.server.ts ~ line 10 ~ constload:PageServerLoad= ~ user',
			locals.user.Authenticated
		);
		throw redirect(302, '/login');
	}


	return {
		
	};
};
