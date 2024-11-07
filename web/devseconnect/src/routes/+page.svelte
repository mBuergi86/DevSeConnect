<script lang="ts">
	import { CodeSolid, ArrowTrendUpSolid, UserGroupSolid } from 'svelte-awesome-icons';
</script>

<div class="container mx-auto flex w-full flex-col items-center justify-center gap-8 p-6">
	<div class="flex flex-col items-center gap-6 text-center">
		<h1 class="text-4xl font-bold text-white">
			Connect, Collaborate, and Grow in DevOps & Software Engineering
		</h1>
		<p class="text-lg text-white">
			Join the premier social platform designed exclusively for DevOps professionals and software
			engineers.
		</p>
		<a
			href="/register"
			class="flex h-12 w-[320px] items-center justify-center rounded-2xl bg-custom-gradient-light text-[24px] font-bold text-white hover:scale-110 hover:transition-all"
		>
			Join DevSeConnect Free
		</a>
	</div>

	<!-- Karten im Spaltenlayout, auch auf großen Bildschirmen -->
	<div class="flex w-full max-w-[800px] flex-col items-center gap-8">
		<!-- Card für "Share Knowledge" -->
		<div class="flex w-full flex-col items-center gap-4 rounded-lg bg-[#1F2937] p-6 text-white">
			<CodeSolid size="60" color="#22D3EE" />
			<div class="text-center">
				<h3 class="text-xl font-semibold">Share Knowledge</h3>
				<p class="text-base">Exchange code snippets, best practices, and industry insights</p>
			</div>
		</div>

		<!-- Card für "Build your Network" -->
		<div class="flex w-full flex-col items-center gap-4 rounded-lg bg-[#1F2937] p-6 text-white">
			<UserGroupSolid size="60" color="#22D3EE" />
			<div class="text-center">
				<h3 class="text-xl font-semibold">Build your Network</h3>
				<p class="text-base">Connect with like-minded professionals and industry leaders</p>
			</div>
		</div>

		<!-- Card für "Stay Updated" -->
		<div class="flex w-full flex-col items-center gap-4 rounded-lg bg-[#1F2937] p-6 text-white">
			<ArrowTrendUpSolid size="60" color="#22D3EE" />
			<div class="text-center">
				<h3 class="text-xl font-semibold">Stay Updated</h3>
				<p class="text-base">Keep up with the latest trends and technologies in DevOps</p>
			</div>
		</div>
	</div>
</div>
