<script lang="ts">
	import Footer from '$lib/components/Footer.svelte';
	import logo from '$lib/assets/images/logo.png';
	import {
		Sun,
		Moon,
		House,
		Box,
		MessageSquare,
		MessageCircle,
		Briefcase,
		BookOpen,
		Bell,
		User,
		LogOut
	} from 'svelte-lucide';
	import { onMount } from 'svelte';
	let { children } = $props();

	let dashboardLinks = [
		{ url: '/dashboard', name: 'Home', icon: House },
		{ url: '/dashboard/projects', name: 'Projects', icon: Box },
		{ url: '/dashboard/forum', name: 'Forum', icon: MessageSquare },
		{ url: '/dashboard/chat', name: 'Chat', icon: MessageCircle },
		{ url: '/dashboard/jobs', name: 'Jobs', icon: Briefcase },
		{ url: '/dashboard/tutorial', name: 'Tutorials', icon: BookOpen },
		{ url: '/dashboard/notifications', name: 'Notifications', icon: Bell },
		{ url: '/dashboard/profile', name: 'Profile', icon: User },
		{ url: '/logout', name: 'Logout', icon: LogOut }
	];

	let isDark = $state(true);

	onMount(() => {
		const theme = localStorage.getItem('theme');
		isDark =
			theme === 'dark' || (!theme && window.matchMedia('(prefers-color-scheme: dark)').matches);
	});

	function toggleTheme() {
		isDark = !isDark;
		localStorage.setItem('theme', isDark ? 'dark' : 'light');
	}

	$effect(() => {
		document.documentElement.classList.toggle('dark', isDark);
	});
</script>

<div class="flex min-h-screen flex-col">
	<header class="flex h-24 w-full justify-center bg-gray-200 px-6 shadow-lg dark:bg-[#1F2937]">
		<div class="container flex w-full flex-row items-center justify-between">
			<!-- Logo -->
			<div class="hidden md:block">
				<a href="/dashboard" class="text-gray-900dark:text-white text-2xl font-bold">
					<img src={logo} alt="DevSeConnect" class="h-20 w-auto" />
				</a>
			</div>
			<div class="flex items-center gap-4">
				{#each dashboardLinks as { url, name, icon: Icon }}
					<a
						href={url}
						class="flex items-center gap-1 text-gray-900 hover:text-blue-600 dark:text-white"
					>
						<Icon size="24" />
						<span class="hidden md:inline">{name}</span>
					</a>
				{/each}
			</div>

			<!-- Dark/Light Mode Toggle -->
			<button
				onclick={toggleTheme}
				class="text-gray-900 hover:text-blue-400 dark:text-white dark:hover:text-yellow-400"
			>
				{#if isDark}
					<Sun size="24" />
				{:else}
					<Moon size="24" />
				{/if}
			</button>
		</div>
	</header>

	<!-- Centered Main Content -->
	{@render children()}

	<Footer class="bg-gray-200 !text-black dark:bg-[#1F2937] dark:text-white" />
</div>
