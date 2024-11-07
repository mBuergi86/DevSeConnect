<script lang="ts">
</script>

<main
	class="flex flex-grow flex-col items-center bg-gray-300 p-6 text-black dark:bg-gray-900 dark:text-white"
>
	<h2 class="mb-6 text-3xl font-bold text-black dark:text-white">Projektmanagement</h2>

	<!-- +page.svelte -->
	<form
		method="POST"
		action="?/projects"
		class="w-full max-w-lg space-y-6 rounded-lg bg-white p-6 shadow-lg dark:bg-gray-800"
	>
		<!-- Projekt erstellen -->
		<h3 class="text-2xl font-semibold text-black dark:text-white">Projekt erstellen</h3>
		<input
			name="projectTitle"
			placeholder="Projektname"
			class="w-full rounded border border-gray-300 p-2 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-700"
			required
		/>
		<textarea
			name="projectDescription"
			placeholder="Projektbeschreibung"
			class="w-full rounded border border-gray-300 p-2 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-700"
			required
		></textarea>
		<button
			type="submit"
			name="action"
			value="addProject"
			class="w-full rounded bg-green-600 px-4 py-2 text-white transition hover:bg-green-800"
		>
			Projekt hinzufügen
		</button>

		<!-- Aufgabe erstellen -->
		<h3 class="text-2xl font-semibold text-black dark:text-white">Aufgabe erstellen</h3>
		<select
			name="selectedProjectID"
			class="w-full rounded border border-gray-300 p-2 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-700"
			required
		>
			<option value="">Projekt auswählen</option>
			<!-- Dynamisch Projekte hier einfügen, wenn Projekte geladen sind -->
			// TODO: Add projects
		</select>
		<input
			name="taskTitle"
			placeholder="Aufgabentitel"
			class="w-full rounded border border-gray-300 p-2 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-700"
			required
		/>
		<textarea
			name="taskDescription"
			placeholder="Aufgabenbeschreibung"
			class="w-full rounded border border-gray-300 p-2 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-700"
			required
		></textarea>
		<button
			type="submit"
			name="action"
			value="addTask"
			class="w-full rounded bg-blue-600 px-4 py-2 text-white transition hover:bg-blue-800"
		>
			Aufgabe hinzufügen
		</button>
	</form>

	<!-- Projekt- und Aufgabenliste -->
	<div class="mt-8 w-full max-w-lg space-y-6">
		<h3 class="text-2xl font-semibold text-black dark:text-white">Projekte und Aufgaben</h3>
	</div>
</main>
