import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = (event) => {
	return error(404, 'Page not found');
};
