<script lang="ts">
	import Form from '$lib/components/Forms.svelte';
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
