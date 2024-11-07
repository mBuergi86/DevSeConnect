import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ cookies }) => {
	const isLoggedIn = cookies.get('token') ? true : false;

	if (!isLoggedIn) {
		return redirect(302, '/');
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
