<script lang="ts">
	let { class: className = '' } = $props();
	let _class = $state(className);
</script>

<footer class="mt-auto flex min-h-20 w-full items-center bg-[#1F2937] {_class}">
	<div
		class="container mx-auto flex h-full flex-col items-center justify-center gap-3 p-4 md:flex-row md:justify-between"
	>
		<p class="text-center text-black dark:text-white {_class}">
			© 2024 DevSeConnect. All rights reserved.
		</p>
		<nav>
			<ul class="flex flex-col items-center gap-2 md:flex-row md:gap-5">
				<li>
					<a
						href="/about"
						class="text-black hover:text-[#2563eb] hover:transition-colors dark:text-white {_class}"
					>
						About me
					</a>
				</li>
				<li>
					<a
						href="/privacy"
						class="text-black hover:text-[#2563eb] hover:transition-colors dark:text-white {_class}"
					>
						Privacy Policy
					</a>
				</li>
				<li>
					<a
						href="/terms"
						class="text-black hover:text-[#2563eb] hover:transition-colors dark:text-white {_class}"
					>
						Terms of Service
					</a>
				</li>
			</ul>
		</nav>
	</div>
</footer>
