import { redirect, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { fail } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ locals }) => {
	// redirect user if logged in
	if (locals.token) {
		return redirect(302, '/dashboard');
	}
};

export const actions: Actions = {
	login: async ({ fetch, request, cookies }) => {
		const form = await request.formData();
		const username = form.get('username');
		const password = form.get('password');

		try {
			const res = await fetch('http://localhost:1323/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ username, password })
			});

			const data = await res.json();

			if (res.status === 200) {
				cookies.set('token', data.token, { path: '/' });
				return redirect(302, '/dashboard');
			} else {
				return fail(422, {
					message: data.message
				});
			}
		} catch (error) {
			console.error(error);
			return fail(500, {
				message: 'Internal server error'
			});
		}
	}
};
