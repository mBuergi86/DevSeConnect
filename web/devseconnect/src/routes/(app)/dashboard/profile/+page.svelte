<script lang="ts">
	import defaultImage from '$lib/assets/images/code_icon.png';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	let imagePath = $state('');

	const user = data.user;

	$effect(() => {
		if (typeof window !== 'undefined' && user?.profile_picture) {
			imagePath = `./${user.profile_picture}`;
		} else {
			imagePath = defaultImage;
		}
	});
</script>

<main
	class="flex flex-grow items-center justify-center bg-gray-300 text-black dark:bg-gray-900 dark:text-white"
>
	<div class="w-full max-w-md rounded-lg bg-white p-8 text-center shadow-lg dark:bg-gray-800">
		{#if imagePath}
			<img
				src={imagePath}
				alt="Profile_picture"
				class="mx-auto mb-4 h-36 w-36 rounded-full object-cover"
			/>
		{:else}
			<img
				src={defaultImage}
				alt="Profile_picture"
				class="mx-auto mb-4 h-36 w-36 rounded-full object-cover"
			/>
		{/if}

		<h1 class="mb-4 text-3xl font-bold">{user.username}</h1>
		<h3 class="mb-2 text-xl font-semibold">Bio</h3>
		{#if user.bio}
			<p class="mb-6 text-lg">{user.bio}</p>
		{:else}
			<p class="mb-6 text-lg text-gray-500">No bio available</p>
		{/if}

		<div class="space-y-2 text-left">
			<div>
				<h3 class="text-lg font-medium">Email:</h3>
				<span class="block text-gray-700 dark:text-gray-300">{user.email}</span>
			</div>
			<div>
				<h3 class="text-lg font-medium">Name:</h3>
				<span class="block text-gray-700 dark:text-gray-300"
					>{user.first_name} {user.last_name}</span
				>
			</div>
		</div>
	</div>
</main>
