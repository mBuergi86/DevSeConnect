<script lang="ts">
	import logo from '$lib/assets/images/logo.png';
	interface NavigationLink {
		url: string;
		name: string;
		class?: string;
	}

	let { navigation }: { navigation: NavigationLink[] } = $props();
	let menuOpen = $state(false);

	function handleOutsieClick(event: MouseEvent) {
		const menu = document.getElementById('menu');
		if (menuOpen && !menu?.contains(event.target as Node)) {
			menuOpen = false;
		}
	}

	$effect(() => {
		if (menuOpen) {
			document.addEventListener('click', handleOutsieClick);
		} else {
			document.removeEventListener('click', handleOutsieClick);
		}

		return () => {
			document.removeEventListener('click', handleOutsieClick);
		};
	});

	function menuClose() {
		menuOpen = false;
	}
</script>

<header>
	<div class="container mx-auto mt-2 flex h-24 flex-row items-center justify-between px-4 md:px-0">
		<a href="/" class="flex-shrink-0">
			<img src={logo} alt="DevSeConnect Logo" class="h-16 w-auto md:h-20" />
		</a>

		<nav class="hidden md:flex">
			<ul class="flex flex-row gap-5">
				{#each navigation as link}
					<li>
						<a
							href={link.url}
							class="flex h-10 items-center hover:text-[#2563eb] hover:transition-all {link.class}"
							>{link.name}</a
						>
					</li>
				{/each}
			</ul>
		</nav>

		<div class="flex md:hidden">
			<button
				class="text-gray-700"
				onclick={(event) => {
					event.stopPropagation();
					menuOpen = !menuOpen;
				}}
				aria-label="Button"
			>
				<svg class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M4 6h16M4 12h16m-7 6h7"
					/>
				</svg>
			</button>
		</div>

		{#if menuOpen}
			<div class="fixed inset-0 z-10 flex items-start justify-end bg-black bg-opacity-50 md:hidden">
				<nav
					id="menu"
					class="mr-4 mt-20 flex h-auto w-full max-w-sm flex-col rounded-lg bg-white p-4 shadow-lg"
				>
					<ul class="flex flex-col gap-3">
						{#each navigation as link}
							<li>
								<a
									href={link.url}
									class="block bg-white px-4 py-2 text-black {link.class}"
									onclick={menuClose}
								>
									{link.name}
								</a>
							</li>
						{/each}
					</ul>
				</nav>
			</div>
		{/if}
	</div>
</header>
