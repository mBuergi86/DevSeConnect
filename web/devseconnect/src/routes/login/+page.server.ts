import { redirect, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { fail } from '@sveltejs/kit';
import { URL_API, VITE_DEV_DOMAIN, VITE_PROD_DOMAIN } from '$env/static/private';
import { goto } from '$app/navigation';

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
			const res = await fetch(`http://${URL_API}:1323/login`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ username, password })
			});

			const data = await res.json();

			if (res.status === 200) {
				const domain = import.meta.env.DEV ? VITE_PROD_DOMAIN : VITE_DEV_DOMAIN;

				cookies.set('token', data.token, {
					path: '/',
					maxAge: 86400,
					sameSite: 'lax',
					httpOnly: true,
					secure: !import.meta.env.DEV,
					domain
				});
				goto('/dashboard', { replaceState: true });
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
