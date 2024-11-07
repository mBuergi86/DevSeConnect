// +page.server.ts
import { fail } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	addProject: async ({ request }) => {
		const formData = await request.formData();
		const projectTitle = formData.get('projectTitle') as string;
		const projectDescription = formData.get('projectDescription') as string;

		if (!projectTitle || !projectDescription) {
			return fail(400, { missingFields: true });
		}

		console.log('Projekt hinzufügen:', { projectTitle, projectDescription });

		return { success: true };
	},

	addTask: async ({ request }) => {
		const formData = await request.formData();
		const selectedProjectID = formData.get('selectedProjectID') as string;
		const taskTitle = formData.get('taskTitle') as string;
		const taskDescription = formData.get('taskDescription') as string;

		if (!selectedProjectID || !taskTitle || !taskDescription) {
			return fail(400, { missingFields: true });
		}

		console.log('Aufgabe hinzufügen:', { selectedProjectID, taskTitle, taskDescription });

		return { success: true };
	}
};
