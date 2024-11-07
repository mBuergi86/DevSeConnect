<script lang="ts">
	type Props = {
		method?: string;
		action?: string;
		first_name?: () => string;
		last_name?: () => string;
		username?: () => string;
		email?: () => string;
		confirmEmail?: () => string;
		password?: () => string;
		confirmPassword?: () => string;
		submit?: () => string;
	};
	const {
		method,
		action,
		first_name,
		last_name,
		username,
		email,
		confirmEmail,
		password,
		confirmPassword,
		submit
	}: Props = $props();
</script>

<div class="container mx-auto w-[1300px] max-w-xl rounded-lg bg-[#1F2937] p-6 shadow-lg">
	<form {method} {action} class="space-y-4">
		{#if first_name}
			<label for="first_name" class="block font-semibold text-white">{@render first_name()}</label>
			<input type="text" id="first_name" name="first_name" required class="w-full rounded-md p-2" />
		{/if}

		{#if last_name}
			<label for="last_name" class="block font-semibold text-white">{@render last_name()}</label>
			<input type="text" id="last_name" name="last_name" required class="w-full rounded-md p-2" />
		{/if}

		{#if username}
			<label for="username" class="block font-semibold text-white">{@render username()}</label>
			<input type="text" id="username" name="username" required class="w-full rounded-md p-2" />
		{/if}

		{#if email}
			<label for="email" class="block font-semibold text-white">{@render email()}</label>
			<input type="email" id="email" name="email" required class="w-full rounded-md p-2" />
		{/if}

		{#if confirmEmail}
			<label for="email" class="block font-semibold text-white">{@render confirmEmail()}</label>
			<input type="email" id="email" name="email" required class="w-full rounded-md p-2" />
		{/if}

		{#if password}
			<label for="password" class="block font-semibold text-white">{@render password()}</label>
			<input type="password" id="password" name="password" required class="w-full rounded-md p-2" />
		{/if}

		{#if confirmPassword}
			<label for="confirmPassword" class="block font-semibold text-white"
				>{@render confirmPassword()}</label
			>
			<input
				type="password"
				id="confirmPassword"
				name="confirmPassword"
				required
				class="w-full rounded-md p-2"
			/>
		{/if}

		<div class="form-actions flex gap-4 pt-4">
			<button type="reset" class="w-full rounded-md bg-red-500 p-2 text-white hover:bg-red-600"
				>Reset</button
			>
			<button type="submit" class="w-full rounded-md bg-blue-500 p-2 text-white hover:bg-blue-600"
				>{@render submit()}</button
			>
		</div>
	</form>
</div>
