import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (cookies) => {
	const token = cookies.cookies.get('token');

	if (token) {
		cookies.cookies.delete('token', { path: '/' });
		return redirect(302, '/');
	}

	return {};
};
