<script lang="ts">
	import logo_light from '$lib/assets/logo.png';
	import logo_dark from '$lib/assets/logo_transparent.png';
	import * as Icon from 'svelte-lucide';

	let { children } = $props();

	let isDark = $state(false);
	let userToggled = $state(false);

	$effect(() => {
		if (typeof window !== 'undefined') {
			const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');

			if (!userToggled) {
				isDark = mediaQuery.matches;
				applyTheme();
			}

			mediaQuery.addEventListener('change', (e) => {
				if (!userToggled) {
					isDark = e.matches;
					applyTheme();
				}
			});
		}

		applyTheme();
	});

	function applyTheme() {
		if (isDark) {
			window.document.documentElement.classList.add('dark');
		} else {
			window.document.documentElement.classList.remove('dark');
		}
	}

	function toggleTheme() {
		userToggled = true;
		isDark = !isDark;
		applyTheme();
	}
</script>

<header>
	<div class="container">
		{#if isDark}
			<img src={logo_dark} alt="DevSeConnect Logo Dark" class="logo" />
		{:else}
			<img src={logo_light} alt="DevSeConnect Logo Light" class="logo" />
		{/if}
		<div class="navigation-right">
			<nav>
				<ul>
					<li>
						<a href="/">
							<Icon.House size="40" />
						</a>
					</li>
					<li>
						<a href="projects">
							<Icon.Box size="40" />
						</a>
					</li>
					<li>
						<a href="/forum">
							<Icon.MessageSquare size="40" />
						</a>
					</li>
					<li>
						<a href="/chat">
							<Icon.MessageCircle size="40" />
						</a>
					</li>
					<li>
						<a href="/jobs">
							<Icon.Briefcase size="40" />
						</a>
					</li>
					<li>
						<a href="/tutorial">
							<Icon.BookOpen size="40" />
						</a>
					</li>
					<li>
						<a href="/notifications">
							<Icon.Bell size="40" />
						</a>
					</li>
					<li>
						<a href="/profile">
							<Icon.User size="40" />
						</a>
					</li>
					<li>
						<a href="/logout">
							<Icon.LogOut size="40" />
						</a>
					</li>
				</ul>
				<div class="toggleButton">
					<button class="togglebtn" onclick={toggleTheme}>
						{#if isDark}
							<Icon.Sun size="40" />
						{:else}
							<Icon.Moon size="40" />
						{/if}
					</button>
				</div>
			</nav>
		</div>
	</div>
</header>

{@render children?.()}

<style>
	.container {
		margin: 0 auto;
		display: flex;
		justify-content: space-between;
		align-items: center;
		max-width: 1600px;
		height: 100%;
	}

	.logo {
		width: auto;
		height: 65px;
	}

	.navigation-right {
		display: flex;
		align-items: center;
		gap: 1rem;
	}

	nav {
		display: flex;
		align-items: center;
	}

	nav ul {
		display: flex;
		gap: 1rem;
	}

	nav li {
		list-style: none;
	}

	nav a {
		text-decoration: none;
	}

	:global(.dark) nav a {
		color: var(--text-color-dark);
	}

	nav a {
		color: var(--text-color-light);
	}

	nav a:hover {
		color: #2563eb;
		transform: scale(1.2);
		transition: all 0.3s ease-in-out;
	}

	:global(.dark) header {
		width: 100%;
		height: 100px;
		background-color: var(--secondary-color-dark);
	}

	header {
		width: 100vw;
		height: 100px;
		background-color: var(--secondary-color-light);
	}

	.toggleButton {
		margin-left: 0.5rem;
	}

	.togglebtn {
		background: transparent;
		border: none;
		cursor: pointer;
	}

	:global(.dark) .togglebtn {
		color: var(--text-color-dark);
	}

	:global(.dark) .togglebtn:hover {
		color: yellow;
	}

	.togglebtn {
		color: var(--text-color-light);
		transform: scale(1.2);
		transition: all 0.3s ease-in-out;
	}

	.togglebtn:hover {
		color: #2563eb;
		transform: scale(1.2);
		transition: all 0.3s ease-in-out;
	}
</style>
