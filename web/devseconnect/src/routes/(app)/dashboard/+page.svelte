<script lang="ts">
	import * as Icon from 'svelte-awesome-icons';
	let { data } = $props();

	let posts = $state(data?.post ?? []);

	function calculateTimeAgo(createdAt: string): string {
		const postTime = new Date(createdAt);
		const now = new Date();
		const diffInSeconds = Math.floor((now.getTime() - postTime.getTime()) / 1000);

		if (diffInSeconds < 60) {
			return `${diffInSeconds} Seconds ago`;
		} else if (diffInSeconds < 3600) {
			return `${Math.floor(diffInSeconds / 60)} Minutes ago`;
		} else if (diffInSeconds < 86400) {
			return `${Math.floor(diffInSeconds / 3600)} Hours ago`;
		} else {
			return `${Math.floor(diffInSeconds / 86400)} Days ago`;
		}
	}
</script>

<div class="flex flex-grow bg-gray-300 text-black dark:bg-gray-900 dark:text-white">
	<div class="container mx-auto grid grid-cols-1 gap-8 p-4 md:grid-cols-3">
		<!-- Left Sidebar -->
		<div class="h-2/6 rounded-lg bg-white p-6 shadow-lg dark:bg-gray-800">
			<h1 class="mb-4 text-xl font-bold dark:text-white">Top Communities</h1>
			<ul class="space-y-2">
				{#each ['DevOps', 'SoftwareEngineering', 'Programming', 'JavaScript', 'Python', 'Java'] as community}
					<li>
						<a href={`#${community}`} class="text-blue-600 hover:text-blue-800 dark:text-blue-400"
							>{`r/${community}`}</a
						>
					</li>
				{/each}
			</ul>
		</div>

		<!-- Middle Column (Main Content) -->
		<div class="space-y-6">
			<div class="h-auto rounded-lg bg-white p-6 shadow-lg dark:bg-gray-800">
				<form method="POST" action="?/post" enctype="multipart/form-data" class="space-y-4">
					<input
						type="text"
						name="title"
						placeholder="Title"
						required
						class="w-full rounded border p-2 focus:outline-none focus:ring-2 focus:ring-blue-600 dark:bg-gray-700 dark:text-white"
					/>

					<textarea
						name="content"
						placeholder="Create a Post"
						required
						class="h-32 w-full resize-none rounded border p-2 focus:outline-none focus:ring-2 focus:ring-blue-600 dark:bg-gray-700 dark:text-white"
					></textarea>

					<div class="flex items-center space-x-2">
						<Icon.UploadSolid size="20" color="#374151" class="dark:fill-white" />
						<input
							type="file"
							name="document"
							class="w-full rounded border p-2 focus:outline-none focus:ring-2 focus:ring-blue-600 dark:bg-gray-700 dark:text-white"
						/>
					</div>

					<!-- Foto-Upload mit Icon -->
					<div class="flex items-center space-x-2">
						<Icon.ImageSolid size="20" color="#374151" class="dark:fill-white" />
						<input
							type="file"
							name="photo"
							accept="image/*"
							class="w-full rounded border p-2 focus:outline-none focus:ring-2 focus:ring-blue-600 dark:bg-gray-700 dark:text-white"
						/>
					</div>

					<!-- Tags-Eingabe mit Icon -->
					<div class="flex items-center space-x-2">
						<Icon.TagSolid size="20" color="#374151" class="dark:fill-white" />
						<input
							type="text"
							name="tags"
							placeholder="Tags (z.B. Tech, News, Tutorial)"
							class="w-full rounded border p-2 focus:outline-none focus:ring-2 focus:ring-blue-600 dark:bg-gray-700 dark:text-white"
						/>
					</div>

					<button
						type="submit"
						class="w-full rounded bg-blue-600 px-4 py-2 text-white transition hover:bg-blue-700"
						>Post erstellen</button
					>
				</form>
			</div>

			{#each posts as post}
				<div class="rounded-lg bg-white p-6 shadow-lg dark:bg-gray-800">
					<div class="mb-4 flex items-center space-x-3">
						<div class="flex h-14 w-14 items-center justify-center rounded-full bg-gray-400">
							<Icon.StarSolid size="32" color="#374151" />
						</div>
						<span class="text-gray-500 dark:text-gray-300">
							Posted by u/{post?.user?.username} - {calculateTimeAgo(post?.created_at)}
						</span>
					</div>
					<h2 class="mb-2 text-2xl font-semibold">{post?.title}</h2>
					<p class="mb-4 text-gray-700 dark:text-gray-300">{post?.content}</p>

					{#if post?.tags}
						<div class="mb-4 flex flex-wrap gap-2">
							{#each post.tags.split(',') as tag}
								<span
									class="rounded bg-blue-100 px-2 py-1 text-sm text-blue-600 dark:bg-blue-700 dark:text-blue-200"
								>
									#{tag.trim()}
								</span>
							{/each}
						</div>
					{/if}

					{#if post?.documentUrl}
						<a
							href={post.documentUrl}
							target="_blank"
							class="text-blue-500 underline dark:text-blue-300"
						>
							Datei herunterladen
						</a>
					{/if}

					{#if post?.photoUrl}
						<img src={post.photoUrl} alt="Foto" class="mt-4 h-48 w-full rounded-lg object-cover" />
					{/if}

					<div class="flex items-center space-x-4 text-gray-500 dark:text-gray-400">
						<button class="flex items-center space-x-1">
							<Icon.ThumbsUpSolid size="16" /><span>10</span><span>Likes</span>
						</button>
						<button class="flex items-center space-x-1">
							<Icon.CommentSolid size="16" /><span>5</span><span>Comments</span>
						</button>
						<button class="flex items-center space-x-1">
							<Icon.ShareNodesSolid size="16" /><span>Share</span>
						</button>
					</div>
				</div>
			{/each}
		</div>
		<!-- Right Sidebar -->
		<div class="h-2/6 rounded-lg bg-white p-6 shadow-lg dark:bg-gray-800">
			<h3 class="mb-2 text-xl font-semibold dark:text-white">Build your Network</h3>
			<p class="mb-4 text-gray-700 dark:text-gray-300">
				Connect with like-minded professionals and industry leaders
			</p>
			<button class="rounded bg-blue-600 px-4 py-2 text-white transition hover:bg-blue-700"
				>Connect</button
			>
		</div>
	</div>
</div>
