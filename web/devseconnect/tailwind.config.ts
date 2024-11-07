import aspectRatio from '@tailwindcss/aspect-ratio';
import containerQueries from '@tailwindcss/container-queries';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';
import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	darkMode: 'selector',
	important: true,
	theme: {
		extend: {
			backgroundImage: {
				'custom-gradient': 'linear-gradient(125.25deg, #111827 43.07%, #00216a 78.7%)',
				'custom-gradient-light': 'linear-gradient(90deg, #22D3EE 20%, #2563EB 79.5%)'
			}
		}
	},

	plugins: [typography, forms, containerQueries, aspectRatio]
} satisfies Config;
