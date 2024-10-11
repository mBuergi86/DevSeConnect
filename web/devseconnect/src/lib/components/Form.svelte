<script lang="ts">
	let { firstName, lastName, username, password, confirmPassword, submit } = $props();

	let firstNameInput: string = $state('');
	let lastNameInput: string = $state('');
	let usernameInput: string = $state('');
	let passwordInput: string = $state('');
	let confirmPasswordInput: string = $state('');
</script>

<div class="container">
	<form method="POST">
		{#if firstName}
			<label for="firstName">{@render firstName()}</label>
			<input type="text" id="firstName" name="firstName" required bind:value={firstNameInput} />
		{/if}
		{#if lastName}
			<label for="lastName">{@render lastName()}</label>
			<input type="text" id="lastName" name="lastName" required bind:value={lastNameInput} />
		{/if}
		{#if username}
			<label for="username">{@render username()}</label>
			<input type="text" id="username" name="username" required bind:value={usernameInput} />
		{/if}
		{#if password}
			<label for="password">{@render password()}</label>
			<input type="password" id="password" name="password" required bind:value={passwordInput} />
		{/if}
		{#if confirmPassword}
			<label for="confirmPassword">{@render confirmPassword()}</label>
			<input
				type="password"
				id="confirmPassword"
				name="confirmPassword"
				required
				bind:value={confirmPasswordInput}
			/>
		{/if}
		{#if submit}
			<div class="form-actions">
				<button type="reset">Reset</button>
				<button type="submit">{@render submit()}</button>
			</div>
		{/if}
	</form>
</div>

<style lang="postcss">
	:global(:root) {
		--green-button: #10b981;
		--red-button: #ef4444;
		--yellow-button: #f59e0b;
		--background-color: #d9d9d9;
		--border-color: #d1d5db;
		--color: #ffffff;
	}

	.container {
		display: flex;
		justify-content: center;
		align-items: center;
		margin-top: 3rem;
		margin-bottom: 3rem;
		height: 100%;
	}

	form {
		display: flex;
		flex-direction: column;
		width: 600px;
		max-width: calc(100% - 2rem);
		margin: 0 auto;
		background-color: white;
		padding: 2rem;
		border-radius: 8px;
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
	}

	form > * {
		margin: 0.75rem 0;
	}

	form > input {
		background-color: var(--background-color);
		border: 1px solid var(--border-color);
		border-radius: 5px;
		padding: 0.75rem;
		font-size: 1rem;
		outline: none;
	}

	form > input:focus {
		border-color: var(--yellow-button);
		box-shadow: 0 0 0 2px rgba(76, 175, 80, 0.2);
	}

	.form-actions {
		display: flex;
		justify-content: space-around;
		gap: 1rem;
		margin-top: 1.5rem;
	}

	form > .form-actions > button {
		padding: 0.5rem 3rem;
		border: none;
		border-radius: 5px;
		cursor: pointer;
		font-size: 1rem;
		transition: background-color 0.3s ease;
	}

	form > .form-actions > button[type='submit'] {
		background-color: var(--green-button);
		color: white;
	}

	form > .form-actions > button[type='reset'] {
		background-color: var(--red-button);
		color: white;
	}

	form > .form-actions > button[type='submit']:hover {
		background-color: #388e3c;
		scale: 1.1;
		transition: scale 0.3s ease;
	}

	form > .form-actions > button[type='reset']:hover {
		background-color: #d32f2f;
		scale: 1.1;
		transition: scale 0.3s ease;
	}
</style>
