const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		screens: {
      // bootstrap default is used
			xs: { min: '1px', max: '575px' },
			sm: { min: '576px', max: '767px' },
			md: { min: '768px', max: '991px' },
			lg: { min: '992px', max: '1199' },
			xl: { min: '1200px', max: '1399' },
			xxl: { min: '1400px' }
		},
		extend: {
			fontFamily: {
				poppins: ['Poppins', 'sans-serif'],
				'sf-pro': ['SF Pro Display', 'sans-serif'],
				raleway: ['Raleway', 'sans-serif']
			}
		}
	},

	plugins: [
		require('tailwindcss-question-mark'),
		require('@tailwindcss/typography'),
		require('@tailwindcss/aspect-ratio'),
		require('@tailwindcss/forms'),
		require('prettier-plugin-tailwindcss')
	]
};

module.exports = config;
