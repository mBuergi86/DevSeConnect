<script lang="ts">
	import type { Snippet } from 'svelte';

	let { children, prefix } = $props();
	prefix = prefix as Snippet;
	import { CodeSolid, ArrowTrendUpSolid, UserGroupSolid } from 'svelte-awesome-icons';
</script>

<div>
	<div class="container">
		<div>
			<h1>Connect, Collaborate, and Grow in DevOps & Software Engineering</h1>
			<p>
				Join the premier social platform designed exclusively for DevOps professionals and software
				engineers.
			</p>
			<a href="/register" class="register"> Join DevSeConnect Free </a>
		</div>
		<div class="card">
			<CodeSolid size="60" color="#22D3EE" />
			<h3>Share Knowledge</h3>
			<p>Exchange code snippets, best practices, and industry insights</p>
		</div>
		<div class="card">
			<UserGroupSolid size="60" color="#22D3EE" />
			<h3>Build your Network</h3>
			<p>Connect with like-minded professionals and industry leaders</p>
		</div>
		<div class="card">
			<ArrowTrendUpSolid size="60" color="#22D3EE" />
			<h3>Stay Updated</h3>
			<p>Keep up with the latest trends and technologies in DevOps</p>
		</div>
	</div>
	{@render children?.()}
	{@render prefix?.()}
</div>

<style lang="postcss">
	.container {
		display: flex;
		justify-content: center;
		align-items: center;
		flex-wrap: wrap;
		text-align: center;
		color: var(--text-color-dark);
	}

	.container h1 {
		font-size: clamp(2.125rem, 2.0417rem + 0.3333vw, 2.375rem);
		font-weight: 700;
	}

	.container p {
		font-size: clamp(1.125rem, 1.0417rem + 0.3333vw, 1.375rem);
	}

	.container a {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 300px;
		height: 45px;
		border-radius: 50px;
		background: linear-gradient(90deg, #22d3ee 20%, #2563eb 79.5%);
		font-size: clamp(1.125rem, 1.0417rem + 0.3333vw, 1.375rem);
		font-weight: 700;
		margin-top: 3rem;
		margin-bottom: 3rem;
		margin-left: auto;
		margin-right: auto;
	}

	.container a:hover {
		scale: 1.1;
		transition: scale 0.3s ease;
	}

	.container .card {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		margin-top: 1rem;
		margin-bottom: 2rem;
		width: calc(100% - 12rem);
		max-width: 1000px;
		height: 150px;
		border-radius: 8px;
		background: #1f2937;
		padding: 1rem;
	}

	.container .card h3 {
		font-size: clamp(1.125rem, 1.0417rem + 0.3333vw, 1.375rem);
		font-weight: 700;
	}

	.container .card p {
		font-size: clamp(0.875rem, 0.7917rem + 0.3333vw, 1.125rem);
	}

	@media (max-width: 480px) {
		.container {
			padding: 0 !important;
		}
		.container h1 {
			font-size: clamp(1.5rem, 1.375rem + 0.1667vw, 1.75rem);
		}

		.container p {
			font-size: clamp(1rem, 0.9167rem + 0.3333vw, 1.25rem);
		}

		.container a {
			width: 100%;
			max-width: 250px;
		}

		.container .card {
			width: calc(100% - 3rem);
			height: 250px;
		}
	}

	@media (max-width: 768px) {
		.container {
			padding: 2rem;
		}

		.container h1 {
			font-size: clamp(1.5rem, 1.375rem + 0.1667vw, 1.75rem);
		}

		.container p {
			font-size: clamp(1rem, 0.9167rem + 0.3333vw, 1.25rem);
		}

		.container a {
			width: 100%;
			max-width: 250px;
		}

		.container .card {
			width: calc(100% - 8rem);
			height: 250px;
		}
	}
</style>
