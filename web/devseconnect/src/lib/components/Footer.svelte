<script lang="ts">
	let { class: className = '' } = $props();
</script>

<div class={className}>
	<div class="container">
		<p>© 2021 DevSeConnect. All rights reserved.</p>
		<nav>
			<a href="/aboutme">About Me</a>
			<span>|</span>
			<a href="/privacy">Privacy Policy</a>
			<span>|</span>
			<a href="/terms">Terms of Service</a>
			<span>|</span>
			<a href="/contact">Contact</a>
		</nav>
	</div>
</div>

<style lang="postcss">
	.footer_group {
		font-size: clamp(0.875rem, 0.7917rem + 0.3333vw, 1.125rem);
	}

	.container {
		display: flex;
		justify-content: space-between;
		align-items: center;
		max-width: 1600px;
		margin: 0 auto;
		height: 80px;
	}

	a {
		width: 100vw;
		text-decoration: none;
		margin-left: 0.5rem;
		margin-right: 0.5rem;
	}

	a:hover {
		text-decoration: underline;
	}
</style>
