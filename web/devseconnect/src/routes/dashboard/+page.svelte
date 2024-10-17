<script lang="ts">
	import * as Icon from 'svelte-awesome-icons';
	let { data } = $props();

	let posts = $state(data?.post);

	function calculateTimeAgo(createdAt) {
		const postTime = new Date(createdAt);
		const now = new Date();
		const diffInSeconds = Math.floor((now - postTime) / 1000);

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

<div class="container">
	<div class="aside-left">
		<div class="card">
			<div class="card-header">
				<h1>Top Communities</h1>
			</div>
			<div class="card-body">
				<ul>
					<li><a href="#DevOps">r/DevOps</a></li>
					<li><a href="#SoftwareEngineering">r/SoftwareEngineering</a></li>
					<li><a href="#Programming">r/Programming</a></li>
					<li><a href="#JavaScript">r/JavaScript</a></li>
					<li><a href="#Python">r/Python</a></li>
					<li><a href="#Java">r/Java</a></li>
				</ul>
			</div>
		</div>
	</div>
	<div class="aside-middle">
		<div class="form-container">
			<form method="POST" action="?/post">
				<div class="form-group">
					<input
						type="text"
						name="title"
						id="post"
						class="input-field"
						placeholder="Title"
						required
					/>
					<textarea name="content" id="post" class="textarea" placeholder="Create a Post" required
					></textarea>
					<button type="submit">Post</button>
				</div>
			</form>
		</div>
		{#each posts as post}
			<div class="card">
				<div class="card-header">
					<div class="post_icon_circle">
						<Icon.StarSolid size="60" color="#374151" />
					</div>
					<span>Posted by u/{post?.user?.username} - {calculateTimeAgo(post?.created_at)}</span>
				</div>
				<div class="card-body">
					<h2>{post?.title}</h2>
					<p>{post?.content}</p>
				</div>
				<div class="card-footer">
					<button><Icon.ThumbsUpSolid size="16" /></button><span>10</span><span>Likes</span>
					<button><Icon.CommentSolid size="16" /></button><span>5</span><span>Comments</span>
					<button><Icon.ShareNodesSolid size="16" /></button>
				</div>
			</div>
		{/each}
	</div>
	<div class="aside-right">
		<div class="card">
			<div class="card-header">
				<h3>Build your Network</h3>
			</div>
			<div class="card-body">
				<p>Connect with like-minded professionals and industry leaders</p>
			</div>
			<div class="card-footer">
				<button>Connect</button>
			</div>
		</div>
	</div>
</div>

<style lang="postcss">
	.container {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		justify-content: center;
		gap: 2rem;
		width: 1600px;
		height: 100%;
		margin: 0 auto;
		padding: 2rem;
	}

	.form-container {
		position: relative;
		display: flex;
		flex-direction: column;
		align-items: center;
		max-width: 800px;
		min-height: 250px;
		border-radius: 0.5rem;
		padding: 2rem;
		margin: 0 auto;
	}

	.input-field {
		min-width: 770px;
		max-width: 100%;
		padding: 1rem;
		border-radius: 0.25rem;
		outline: none;
		border: none;
		box-sizing: border-box;
		margin-bottom: 1rem;
		font-size: 1.2rem;
	}

	.textarea {
		min-width: 100%;
		max-width: 770px;
		min-height: 150px;
		padding: 1rem;
		border-radius: 0.25rem;
		outline: none;
		border: none;
		box-sizing: border-box;
		margin-bottom: 2.5rem;
		font-size: 1.2rem;
	}

	:global(.dark) .input-field::placeholder,
	:global(.dark) .textarea::placeholder {
		color: #ffffff;
		font-size: 1.25rem;
	}

	.input-field::placeholder,
	.textarea::placeholder {
		color: var(--form-placeholder-color);
		font-size: 1.25rem;
	}

	:global(.dark) .input-field,
	:global(.dark) .textarea {
		color: #ffffff;
		background: #374151;
	}

	.input-field,
	.textarea {
		color: #000000;
		background: #f3f4f6;
	}

	.form-container button {
		position: absolute;
		width: 100px;
		height: 40px;
		background: #2563eb;
		color: white;
		font-size: 1rem;
		outline: none;
		border: none;
		border-radius: 0.5rem;
		bottom: 1rem;
		right: 1rem;
	}

	.form-container button:hover {
		background: #1e40af;
		transition: background 0.3s ease;
	}

	:global(.dark) .form-container {
		background-color: var(--secondary-color-dark);
		color: var(--text-color-dark);
	}

	.form-container {
		background-color: var(--secondary-color-light);
		color: var(--text-color-light);
	}

	.aside-left > .card {
		display: flex;
		flex-direction: column;
		position: relative;
		min-width: 300px;
		padding: 1rem;
		border-radius: 0.5rem;
		background: transparent !important;
	}

	.aside-middle > .card {
		position: relative;
		width: 800px;
		min-height: 300px;
		border-radius: 0.5rem;
		padding: 1rem;
		margin-top: 1rem;
	}

	.aside-right > .card {
		display: flex;
		flex-direction: column;
		min-width: 400px;
		padding: 1rem;
		border-radius: 0.5rem;
	}

	:global(.dark) .card {
		background: var(--secondary-color-dark);
		color: var(--text-color-dark);
	}

	.card {
		background: var(--secondary-color-light);
		color: var(--text-color-light);
	}

	.card-header {
		display: flex;
		align-items: center;
		gap: 1rem;
	}

	.card-header span {
		font-size: 1rem;
	}

	.post_icon_circle {
		background: var(--primary-color);
		width: 60px;
		height: 60px;
		border-radius: 50%;
		padding: 10px;
	}

	.card-body {
		margin-bottom: 1rem;
	}

	.aside-left .card-body ul {
		list-style: none;
		padding: 0;
		margin: 0;
	}

	.aside-left .card-body li {
		margin: 1rem 1rem;
	}

	.aside-left .card-body a {
		text-decoration: none;
		color: var(--text-color-light);
	}

	.aside-left .card-body a:hover {
		color: #2563eb;
		transition: color 0.1s ease;
	}

	:global(.dark) .card-body a {
		color: var(--text-color-dark);
	}

	.card-footer {
		position: absolute;
		display: flex;
		gap: 0.5rem;
		bottom: 1rem;
	}

	.aside-right .card-footer {
		position: relative;
		display: flex;
		justify-content: end;
		bottom: 0rem;
	}

	.aside-right .card-footer button {
		background: #2563eb;
		color: white;
		border-radius: 0.5rem;
		border: none;
		padding: 0.5rem;
	}

	.aside-right .card-footer button:hover {
		background: #1e40af;
		transition: background 0.3s ease;
	}

	.aside-middle .card-footer button {
		background: none;
		border: none;
		cursor: pointer;
	}

	:global(.dark) .card-footer button {
		color: var(--text-color-dark);
	}

	.card-footer button {
		color: var(--text-color-light);
	}
</style>
