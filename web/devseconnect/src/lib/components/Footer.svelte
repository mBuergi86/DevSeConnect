<script lang="ts">
	let { class: className = '' } = $props();
	let _class = $state(className);
</script>

<footer class={_class}>
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
</footer>

<style lang="postcss">
	.footer {
		font-size: clamp(0.875rem, 0.7917rem + 0.3333vw, 1.125rem);
	}

	:global(.dark) .footer,
	:global(.dark) .footer a {
		background-color: var(--secondary-color-dark) !important;
		color: var(--text-color-dark) !important;
	}

	.footer,
	.footer a {
		background-color: var(--secondary-color-light) !important;
		color: var(--text-color-light) !important;
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
