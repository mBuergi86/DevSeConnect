<script lang="ts">
	import type { Snippet } from 'svelte';

	let { children, prefix } = $props();
	prefix = prefix as Snippet;
</script>

<div>
	<h1>Welcome to SvelteKit</h1>
	<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
	{@render children?.()}
	{@render prefix?.()}
</div>
