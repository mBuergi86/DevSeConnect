import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	// get cookies from browser
	const token = event.cookies.get('token');

	if (token) {
		// if token exists, add it to the headers
		event.locals.token = token;
	} else {
		// if token does not exist, remove it from the headers
		event.locals.token = '';
	}

	// load page as normal
	return await resolve(event);
};
