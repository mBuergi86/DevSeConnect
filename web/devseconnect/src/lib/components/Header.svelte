<script lang="ts">
	import logo_light from '$lib/assets/logo.png';
	import logo_dark from '$lib/assets/logo_transparent.png';

	let { children } = $props();

	let theme = $state('system');
	let dropdownOpen = $state(false);

	$effect(() => {
		if (typeof window !== 'undefined') {
			theme = localStorage.getItem('theme') || 'system'; // Lade gespeichertes Thema
			applyTheme();

			const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
			const handleChange = () => {
				if (theme === 'system') {
					applyTheme();
				}
			};
			mediaQuery.addEventListener('change', handleChange);

			return () => {
				mediaQuery.removeEventListener('change', handleChange);
			};
		}
	});

	function applyTheme() {
		if (typeof window !== 'undefined') {
			// Entferne beide Klassen ("light" und "dark")
			document.documentElement.classList.remove('light', 'dark');

			// Wenn das Theme auf "system" steht, wende das System-Thema an
			if (theme === 'system') {
				const dark = window.matchMedia('(prefers-color-scheme: dark)').matches;
				document.documentElement.classList.add(dark ? 'dark' : 'light');
			} else {
				// Ansonsten füge die Klasse direkt hinzu ("light" oder "dark")
				document.documentElement.classList.add(theme);
			}
		}
	}

	function setTheme(newTheme: string) {
		theme = newTheme;
		localStorage.setItem('theme', newTheme);
		dropdownOpen = false;
	}
</script>

<header>
	<div class="container">
		{#if theme === 'dark' || (theme === 'system' && typeof window !== 'undefined' && window.matchMedia('(prefers-color-scheme: dark)').matches)}
			<img src={logo_dark} alt="DevSeConnect Logo Dark" class="logo" />
		{:else}
			<img src={logo_light} alt="DevSeConnect Logo Light" class="logo" />
		{/if}
		<div class="navigation-right">
			<nav>
				<a href="/">Home</a>
				<a href="/about">About</a>
				<a href="/blog">Blog</a>
			</nav>
			<div class="dropdown" class:open={dropdownOpen}>
				<button class="dropbtn" onclick={() => (dropdownOpen = !dropdownOpen)}>
					{#if theme === 'system'}🖥️
					{:else if theme === 'light'}☀️
					{:else}🌙
					{/if}
				</button>
				<div class="dropdown-content">
					<button onclick={() => setTheme('system')} class:active={theme === 'system'}>🖥️</button>
					<button onclick={() => setTheme('light')} class:active={theme === 'light'}>☀️</button>
					<button onclick={() => setTheme('dark')} class:active={theme === 'dark'}>🌙</button>
				</div>
			</div>
		</div>
	</div>
</header>

{@render children?.()}

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

	:global(.dark) {
		background-color: var(--dark-color);
		color: white;
	}

	:global(.light) {
		background-color: var(--light-color);
		color: black;
	}

	:global(.dark) header {
		padding: 1rem;
		width: 100%;
		height: 100px;
		background-color: var(--secondary-color-dark);
	}

	:global(.light) header {
		padding: 1rem;
		width: 100vw;
		height: 100px;
		background-color: var(--secondary-color-light);
	}

	.dropdown {
		position: relative;
		display: inline-block;
	}

	.dropbtn {
		background-color: var(--primary-color);
		color: white;
		width: 60px;
		height: 60px;
		border: none;
		cursor: pointer;
		border-radius: 50%;
		display: flex;
		justify-content: center;
		align-items: center;
		transition: background-color 0.3s;
		font-size: 20px;
	}

	.dropdown-content {
		opacity: 0;
		position: absolute;
		background-color: transparent;
		z-index: 1;
		margin-top: 0.5rem;
		right: 0.6rem;
		transition: opacity 0.5s ease-in-out;
		pointer-events: none;
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.dropdown.open .dropdown-content {
		opacity: 1;
		pointer-events: auto;
	}

	.dropdown-content button {
		color: black;
		display: flex;
		justify-content: center;
		align-items: center;
		width: 50px;
		height: 50px;
		border-radius: 50%;
		margin: 5px 0;
		font-size: 20px;
		border: none;
		cursor: pointer;
		opacity: 0;
		transform: scale(0.8);
		transition:
			transform 0.3s,
			opacity 0.5s;
	}

	.dropdown.open .dropdown-content button:nth-child(1) {
		background: var(--system-color);
		opacity: 1;
		transform: scale(1);
		transition-delay: 0.1s;
	}

	.dropdown.open .dropdown-content button:nth-child(2) {
		background-color: var(--light-color);
		opacity: 1;
		transform: scale(1);
		transition-delay: 0.3s;
	}

	.dropdown.open .dropdown-content button:nth-child(3) {
		background-color: var(--dark-color);
		opacity: 1;
		transform: scale(1);
		transition-delay: 0.5s;
	}

	.dropdown.open .dropdown-content button:hover {
		transform: scale(1.2);
		transition: transform 0.2s ease-out;
	}

	.dropdown:not(.open) .dropdown-content button:nth-child(3) {
		opacity: 0;
		transform: scale(0.8);
		transition-delay: 0.5s;
	}

	.dropdown:not(.open) .dropdown-content button:nth-child(1) {
		opacity: 0;
		transform: scale(0.8);
		transition-delay: 0.3s;
	}

	.dropdown:not(.open) .dropdown-content button:nth-child(2) {
		opacity: 0;
		transform: scale(0.8);
		transition-delay: 0.1s;
	}

	:global(.dark) .dropdown-content button.active {
		border: 2px solid var(--text-color-dark);
	}

	:global(.light) .dropdown-content button.active {
		border: 2px solid var(--text-color-light);
	}
</style>
