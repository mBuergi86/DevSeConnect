<script lang="ts">
	import { onMount } from 'svelte';
	import { writable, get, type Writable } from 'svelte/store';

	interface Message {
		username: string;
		content: string;
	}

	export let data: { user: { username: string } };

	let messages: Writable<Message[]> = writable([]);
	let newMessage = writable('');
	let connectionStatus = writable('Verbunden');
	let socket: WebSocket | null = null;

	onMount(() => {
		socket = new WebSocket('ws://localhost:1323/ws');

		socket.onopen = () => {
			connectionStatus.set('Verbunden');
		};

		socket.onmessage = (event) => {
			const msg: Message = JSON.parse(event.data);
			messages.update((msgs) => [...msgs, msg]);
		};

		socket.onclose = () => {
			connectionStatus.set('Getrennt. Versuche, neu zu verbinden...');
			reconnectWebSocket();
		};

		return () => socket?.close();
	});

	function reconnectWebSocket() {
		setTimeout(() => {
			socket = new WebSocket('ws://localhost:1323/ws');
		}, 5000);
	}

	function sendMessage() {
		if (socket?.readyState === WebSocket.OPEN) {
			const message = { username: data.user.username, content: get(newMessage) };
			socket.send(JSON.stringify(message));
			newMessage.set('');
		}
	}
</script>

<main
	class="flex flex-grow flex-col items-center bg-gray-300 p-4 text-black dark:bg-gray-900 dark:text-white"
>
	<h2 class="mb-4 text-xl font-bold">Echtzeit-Chat</h2>

	<div class="mb-2 text-sm text-gray-600 dark:text-gray-400">
		Status: {$connectionStatus}
	</div>

	<div
		class="mb-4 h-[600px] w-full max-w-4xl overflow-y-scroll rounded bg-white p-4 dark:bg-gray-800"
	>
		{#each $messages as message}
			<p><strong>{message.username}:</strong> {message.content}</p>
		{/each}
	</div>

	<textarea
		bind:value={$newMessage}
		class="mb-2 w-full max-w-4xl rounded border p-2 dark:bg-gray-700 dark:text-white"
		placeholder="Nachricht eingeben..."
		onkeypress={(e) => e.key === 'Enter' && sendMessage()}
	></textarea>
	<div class="flex w-full max-w-4xl justify-end">
		<button
			onclick={sendMessage}
			class="rounded bg-blue-600 p-2 text-white transition hover:bg-blue-700"
		>
			Senden
		</button>
	</div>
</main>
