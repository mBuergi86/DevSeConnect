import { redirect, fail } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';
import { URL_API } from '$env/static/private';

export const load: PageServerLoad = async ({ locals, parent }) => {
	await parent();
	if (locals.token) {
		redirect(302, '/dashboard');
	}
	return {};
};

export const actions: Actions = {
	register: async ({ fetch, request }) => {
		const form = await request.formData();
		const first_name = form.get('first_name') as string;
		const last_name = form.get('last_name') as string;
		const username = form.get('username') as string;
		const email = form.get('email') as string;
		const confirmEmail = form.get('confirmEmail') as string;
		const password = form.get('password') as string;
		const confirmPassword = form.get('confirmPassword') as string;

		if (
			!first_name ||
			!last_name ||
			!username ||
			!email ||
			!confirmEmail ||
			!password ||
			!confirmPassword
		) {
			return fail(400, { message: 'All fields are required.' });
		}

		if (email !== confirmEmail) {
			return fail(400, { message: 'Emails do not match.' });
		}

		if (password !== confirmPassword) {
			return fail(400, { message: 'Passwords do not match.' });
		}

		try {
			const response = await fetch(`http://${URL_API}:1323/register`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					first_name,
					last_name,
					username,
					email,
					password
				})
			});

			if (response.ok) {
				redirect(302, '/login');
			} else {
				const errorData = await response.json();
				return fail(response.status, { message: errorData.message || 'Registration failed' });
			}
		} catch (error) {
			return fail(500, { message: 'Unexpected server error', error });
		}
	}
};
