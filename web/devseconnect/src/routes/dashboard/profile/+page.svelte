<script lang="ts">
	import defaultImage from '$lib/assets/images/default.png?enhanced';

	let { data } = $props();
	let imagePath = $state('');

	$effect(() => {
		if (typeof window !== 'undefined' && data?.user?.profile_picture) {
			imagePath = `./${data.user.profile_picture}`;
		} else {
			imagePath = defaultImage;
		}
	});
</script>

<main>
	<div class="profile-card">
		{#if imagePath}
			<enhanced:img src={imagePath} alt="Profile_picture" class="profile-picture" />
		{:else}
			<enhanced:img src={defaultImage} alt="Profile_picture" class="profile-picture" />
		{/if}
		<h1>{data.user.username}</h1>
		<h3>Bio</h3>
		<p>{data.user.bio}</p>

		<div class="profile-info">
			<h3>Email:</h3>
			<span>{data.user.email}</span>

			<h3>Name:</h3>
			<span>{data.user.first_name} {data.user.last_name}</span>
		</div>
	</div>
</main>

<style>
	:global(.dark) {
		--background-color: #111827;
		--text-color: #ffffff;
		--card-background: #1f2937;
	}

	:global(html) {
		--background-color: #d1d5db;
		--text-color: #000000;
		--card-background: #ffffff;
	}

	main {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100%;
		background-color: var(--background-color);
		color: var(--text-color);
	}

	.profile-card {
		background-color: var(--card-background);
		padding: 2rem;
		border-radius: 15px;
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
		text-align: center;
		max-width: 500px;
		width: 100%;
	}

	.profile-picture {
		width: 150px;
		height: 150px;
		border-radius: 50%;
		object-fit: cover;
		margin-bottom: 1rem;
	}

	h1 {
		font-size: 2rem;
		margin-bottom: 1rem;
	}

	p {
		font-size: 1.2rem;
		margin-bottom: 1.5rem;
	}

	.profile-info {
		text-align: left;
	}

	.profile-info h3 {
		font-size: 1.1rem;
		margin: 0.5rem 0;
	}

	.profile-info span {
		display: block;
		margin-bottom: 0.5rem;
	}
</style>
