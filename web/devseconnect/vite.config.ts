import { defineConfig } from 'vitest/config';
import { sveltekit } from '@sveltejs/kit/vite';
//import { Server } from 'socket.io';
//import type { ViteDevServer } from 'vite';

/*const webSocketServer = {
	name: 'webSocketServer',
	configureeServer(server: ViteDevServer) {
		if (!server.httpServer) return;

		const io = new Server(server.httpServer);

		io.on('connection', (socket) => {
			socket.emit('eventFromServer', 'Hello from server');
		});
	}
};*/

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
