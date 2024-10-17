<script lang="ts">
	import Form from '$lib/components/Form.svelte';
</script>

<div>
	<Form method="POST" action="?/login">
		{#snippet username()}
			Username
		{/snippet}

		{#snippet password()}
			Password
		{/snippet}

		{#snippet submit()}
			Login
		{/snippet}
	</Form>
</div>
