<script lang="ts">
	import { page } from '$app/stores';
</script>

<div>
	<h1>404 - Page Not Found</h1>
	{#if $page.error}
		<div class="mt-4 border-y-2 p-4">
			{#if $page?.status}
				<p>Page status: {$page?.status}</p>
			{/if}
			{#if $page?.error?.message}
				<p>Error message: {$page?.error?.message}</p>
			{/if}
		</div>
	{/if}
</div>
