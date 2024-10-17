import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ cookies }) => {
	const isLoggedIn = cookies.get('token') ? true : false;

	if (!isLoggedIn) {
		redirect(302, '/home');
	}

	try {
		return {
			isLoggedIn
		};
	} catch (error) {
		console.error(error);
		return {
			error: 'Failed to load user data'
		};
	}
};
