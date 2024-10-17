<script lang="ts">
	import '../../app.css';
	import logo from '$lib/assets/logo_secondary.png';
	import Footer from '$lib/components/Footer.svelte';

	let { children: slot } = $props();
</script>

<header>
	<div class="container">
		<a href="/"><img src={logo} alt="DevSeConnect" class="logo" /></a>
		<nav>
			<a href="/home/login" class="login">Login</a>
			<a href="/home/register" class="register">Register</a>
		</nav>
	</div>
</header>
<main>
	{@render slot?.()}
</main>

<Footer class="h-20 w-full bg-[#1f2937] text-white" />

<style>
	header {
		padding: 1rem;
		color: var(--text-color-dark);
		font-size: clamp(0.875rem, 0.7917rem + 0.3333vw, 1.125rem);
	}

	nav {
		display: flex;
		justify-content: center;
		align-items: center;
	}

	nav a {
		color: var(--text-color-dark);
		margin: 0 1rem;
	}

	nav .login:hover {
		color: #2563eb;
		transition: color 0.3s ease;
	}

	nav .register {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100px;
		height: 40px;
		border-radius: 50px;
		background: #2563eb;
	}

	nav .register:hover {
		scale: 1.1;
		transition: scale 0.3s ease;
	}
</style>
