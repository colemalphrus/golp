import App from './svelte/hello.svelte';

const app = new App({
	target: document.getElementById("aboutUsBody"),
	props: {}
});


export default app;