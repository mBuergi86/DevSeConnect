<script lang="ts">
	let { children } = $props();
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
</script>

<Header />

<main>
	{@render children?.()}
</main>

<Footer />

<style>
	:global(html),
	:global(body) {
		margin: 0;
		padding: 0;
		box-sizing: border-box;
		height: 100vh;
		font-family: 'Inter', sans-serif;
	}

	*,
	*::before,
	*::after {
		box-sizing: inherit;
	}

	:global(body) {
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		background-color: var(--light-color);
		color: black;
	}

	main {
		flex-grow: 1;
	}

	:global(.dark) :global(body) {
		background-color: var(--dark-color);
		color: white;
	}
</style>
