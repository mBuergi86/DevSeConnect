import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (cookies) => {
	const token = cookies.cookies.get('token');

	if (token) {
		console.info('User logged out');
		cookies.cookies.delete('token', { path: '/' });
		return redirect(302, '/home');
	}

	return {};
};
