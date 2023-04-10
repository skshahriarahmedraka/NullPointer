import * as jwt from 'jsonwebtoken';
import type { PageServerLoad } from './$types';
import type { UserDataType } from '$lib/store/types';
import { redirect } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ cookies, parent ,locals}) => {
	await parent();
	if (!locals.user.Authenticated) {
		console.log(
			'🚀 ~ file: +page.server.ts ~ line 10 ~ constload:PageServerLoad= ~ user',
			locals.user.Authenticated
		);
		throw redirect(302, '/login');
	}
	const MyCookie = cookies.get('Auth1') || '';
	const JWT_AUTH_KEY: string = process.env.JWT_AUTH_KEY || '';
	let GetUserData: UserDataType={} as UserDataType;
	if (MyCookie != '') {
		interface tokeninterface {
			Email: string;
			Name: string;
			UserID: string;
			Accounttype: string;
			exp: number;
		}
		const decoded = jwt.verify(MyCookie, JWT_AUTH_KEY) as tokeninterface;

		await fetch(`/api/user/${decoded.UserID}`, {
			method: 'GET'
		})
			.then((res) => {
				return res.json();
			})
			.then((data) => {
				GetUserData = data;
				console.log(GetUserData);
			})
			.catch((err) => {
				console.log(err);
			});
	}

	return {
		GetUserData
	};
};
