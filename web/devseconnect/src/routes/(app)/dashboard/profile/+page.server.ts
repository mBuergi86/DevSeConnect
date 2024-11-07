import type { PageServerLoad } from './$types';
import { URL_API } from '$env/static/private';

export const load: PageServerLoad = async ({ fetch, cookies }) => {
	const token = cookies.get('token');

	if (!token) {
		return { redirect: '/login' };
	}

	const fetchUser = async () => {
		const res = await fetch(`http://${URL_API}:1323/user`, {
			method: 'GET',
			headers: {
				Authorization: `Bearer ${token}`,
				'Content-Type': 'application/json'
			}
		});

		if (!res.ok) {
			throw new Error('Failed to fetch user data');
		}

		const user = await res.json();
		return user;
	};

	try {
		const user = await fetchUser();

		return {
			user
		};
	} catch (error) {
		console.error(error);
		return {
			error: 'Failed to load user or post data'
		};
	}
};
