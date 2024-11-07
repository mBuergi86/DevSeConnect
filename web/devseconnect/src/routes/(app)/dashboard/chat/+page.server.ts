import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';
import { URL_API } from '$env/static/private';

export const load: PageServerLoad = async ({ cookies, fetch }) => {
	const token = cookies.get('token');

	if (!token) {
		return redirect(302, '/login');
	}

	try {
		const res = await fetch(`${URL_API}/user`, {
			headers: {
				Authorization: `Bearer ${token}`
			}
		});

		if (!res.ok) {
			throw new Error('Failed to fetch user data');
		}

		const user = await res.json();

		return {
			user
		};
	} catch (error) {
		console.error('Error loading user data:', error);
		return {
			error: 'Failed to load user data'
		};
	}
};
