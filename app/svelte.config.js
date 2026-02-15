import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({
			pages: '../docs',
			assets: '../docs',
			fallback: undefined,
			precompress: false,
			strict: true
		}),
		paths: {
			base: process.env.PUBLIC_BASE_PATH || ''
		}
	}
};

export default config;
