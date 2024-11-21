import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ cookies }) => {
	const isLoggedIn = cookies.get('token') ? true : false;

	return { isLoggedIn };
};
