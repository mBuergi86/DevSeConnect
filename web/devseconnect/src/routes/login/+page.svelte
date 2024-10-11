<script lang="ts">
	import Form from '$lib/components/Form.svelte';
	let { children } = $props();
</script>

<div>
	{#snippet username()}
		Username
	{/snippet}
	{#snippet password()}
		Password
	{/snippet}
	{#snippet submit()}
		Login
	{/snippet}
	<Form {username} {password} {submit} />
</div>
{@render children?.()}
