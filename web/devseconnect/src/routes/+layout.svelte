<script lang="ts">
	let { children } = $props();
</script>

{@render children?.()}

<style>
	:global(:root) {
		--primary-color: #d1d5db;
		--secondary-color-dark: #1f2937;
		--secondary-color-light: white;
		--light-color: #d1d5db;
		--dark-color: #111827;
		--text-color-light: black;
		--text-color-dark: white;
		--form-placeholder-color: #4b5563;
	}
</style>
