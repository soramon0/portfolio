import forms from '@tailwindcss/forms';

/** @type {import('tailwindcss').Config} */
module.exports = {
	darkMode: 'class',
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			maxWidth: {
				'8xl': '90rem'
			}
		}
	},
	plugins: [forms]
};
