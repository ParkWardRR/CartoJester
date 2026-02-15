<script lang="ts">
	import "../app.css";
	import Navbar from "$lib/components/Navbar.svelte";
	import SearchPalette from "$lib/components/SearchPalette.svelte";
	import { initTheme } from "$lib/stores";
	import { onMount } from "svelte";

	let { children } = $props();

	onMount(() => {
		initTheme();
	});
</script>

<div class="app-shell">
	<Navbar />
	<main class="app-main">
		{@render children()}
	</main>
</div>
<SearchPalette />

<style>
	.app-shell {
		display: flex;
		flex-direction: column;
		height: 100dvh;
		overflow: hidden;
	}
	.app-main {
		flex: 1;
		overflow: hidden;
		display: flex;
		flex-direction: column;
	}
</style>
