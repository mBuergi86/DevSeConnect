<script lang="ts">
	import '../app.css';
	import { page } from '$app/stores';
	import { derived } from 'svelte/store';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import { goto } from '$app/navigation';
	let { children } = $props();

	$effect(() => {
		const unsubscribe = page.subscribe(($page) => {
			const protectedRoutes = ['/dashboard'];
			const currentPath = $page.url.pathname;

			if ($page.data.isLoggedIn && protectedRoutes.includes(currentPath)) {
				goto('/dashboard');
			}
		});
		return unsubscribe;
	});

	const showNavigation = derived(page, ($page) => {
		return !$page.url.pathname.startsWith('/dashboard');
	});

	let navigationLinks = [
		{ url: '/login', name: 'Login', class: 'md:text-white text-black' },
		{
			url: '/register',
			name: 'Register',
			class:
				'justify-center w-24 bg-[#2563eb] md:text-white text-black rounded-lg hover:scale-110 hover:text-white'
		}
	];
</script>

{#if $showNavigation}
	<div class="flex min-h-screen flex-col bg-custom-gradient bg-no-repeat">
		<Header navigation={navigationLinks} />

		<div class="flex flex-grow items-center justify-center p-4">
			{@render children()}
		</div>

		<Footer />
	</div>
{:else}
	{@render children()}
{/if}
