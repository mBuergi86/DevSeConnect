<script lang="ts">
	import '../app.css';
	import logo from '$lib/assets/logo_secondary.png';
</script>

<header>
	<div class="container">
		<img src={logo} alt="DevSeConnect" class="logo" />
		<nav>
			<a href="/login" class="login">Login</a>
			<a href="/register" class="register">Register</a>
		</nav>
	</div>
</header>

<main>
	<slot />
</main>

<footer>
	<div class="container">
		<p>© 2021 DevSeConnect. All rights reserved.</p>
		<nav>
			<a href="/aboutme">About Me</a>
			<a href="/privacy">Privacy Policy</a>
			<a href="/terms">Terms of Service</a>
		</nav>
	</div>
</footer>

<style>
	:global(:root) {
		--primary-color: #d1d5db;
		--secondary-color-dark: #1f2937;
		--secondary-color-light: white;
		--system-color: linear-gradient(to left, #f3f4f6 50%, #111827 50%);
		--light-color: #f3f4f6;
		--dark-color: #111827;
		--text-color-light: black;
		--text-color-dark: white;
	}
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

	footer {
		width: 100%;
		background-color: var(--secondary-color-dark);
		color: var(--text-color-dark);
		font-size: clamp(0.875rem, 0.7917rem + 0.3333vw, 1.125rem);
	}

	footer > .container {
		height: 80px;
	}

	footer a {
		color: var(--text-color-dark);
	}
</style>
