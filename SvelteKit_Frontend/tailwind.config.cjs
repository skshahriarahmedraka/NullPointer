const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			fontFamily: {
				'poppins': ['Poppins', 'sans-serif'],
				"sf-pro": ['SF Pro Display', 'sans-serif'],
				"raleway": ['Raleway', 'sans-serif']
			  },
		},
		
	},

	plugins: [
		require('tailwindcss-question-mark'),
		require('@tailwindcss/line-clamp'),
		require('@tailwindcss/typography'),
		require('@tailwindcss/aspect-ratio'),
		require('@tailwindcss/forms'),
		require('prettier-plugin-tailwindcss')
		
	]
};

module.exports = config;











