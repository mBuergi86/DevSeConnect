<script lang="ts">
	import Form from '$lib/components/Forms.svelte';
</script>

<div>
	<Form method="POST" action="?/register">
		{#snippet first_name()}
			First Name
		{/snippet}
		{#snippet last_name()}
			Last Name
		{/snippet}
		{#snippet email()}
			E-Mail
		{/snippet}
		{#snippet confirmEmail()}
			Confirm E-Mail
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
