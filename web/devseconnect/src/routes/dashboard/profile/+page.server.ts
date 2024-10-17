import type { PageServerLoad } from '../$types';

export const load: PageServerLoad = async ({ fetch, cookies }) => {
	const token = cookies.get('token');

	const fetchUser = async () => {
		const res = await fetch('http://localhost:1323/user', {
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
			error: 'Failed to load user data'
		};
	}
};
