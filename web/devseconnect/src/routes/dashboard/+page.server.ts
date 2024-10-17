import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies }) => {
	const token = cookies.get('token');

	if (!token) {
		return { redirect: '/login' };
	}

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

	const fetchPosts = async () => {
		const res = await fetch('http://localhost:1323/post', {
			method: 'GET',
			headers: {
				Authorization: `Bearer ${token}`,
				'Content-Type': 'application/json'
			}
		});

		if (!res.ok) {
			return [];
		}

		const post = await res.json();
		return post;
	};

	try {
		const user = await fetchUser();
		const post = await fetchPosts();

		return {
			user,
			post
		};
	} catch (error) {
		console.error(error);
		return {
			error: 'Failed to load user or post data'
		};
	}
};

export const actions: Actions = {
	post: async ({ fetch, request, cookies }) => {
		const form = await request.formData();
		const title = form.get('title');
		const content = form.get('content');
		const token = cookies.get('token');

		if (!title || !content) {
			return fail(422, {
				message: 'Title and content are required'
			});
		}

		try {
			const res = await fetch('http://localhost:1323/posts', {
				method: 'POST',
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ title, content })
			});
			if (!res.ok) {
				return new Error('Failed to create post');
			}
			return {
				success: true,
				message: 'Post created successfully',
				redirect: '/dashboard'
			};
		} catch (error) {
			console.error(error);
			return fail(500, {
				message: 'Internal server error'
			});
		}
	}
};
