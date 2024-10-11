<script lang="ts">
	import Form from '$lib/components/Form.svelte';
	let { children } = $props();
</script>

<div>
	<Form>
		{#snippet firstName()}
			First Name
		{/snippet}
		{#snippet lastName()}
			Last Name
		{/snippet}
		{#snippet username()}
			Username
		{/snippet}
		{#snippet password()}
			Password
		{/snippet}
		{#snippet confirmPassword()}
			Confirm Password
		{/snippet}
		{#snippet submit()}
			Register
		{/snippet}
	</Form>
</div>
{@render children?.()}
