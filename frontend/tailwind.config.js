const flowbitePlugin = require('flowbite/plugin');

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{html,js,svelte}', 
    './node_modules/flowbite-svelte/**/*.{html,js,svelte}'
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        // Customize your colors here if needed
      },
    },
  },
  plugins: [flowbitePlugin],
};