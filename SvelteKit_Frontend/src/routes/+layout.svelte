<script>
	import Footer from '$lib/Footer/footer.svelte';
import '../app.postcss';
	
	const HeadLogo = new URL("../../static/favicon.png", import.meta.url).href;
</script>


<svelte:head>
	<title>Null Pointer</title>
	<link rel="icon"  href={HeadLogo} />
</svelte:head>

<!-- <div class=" h-screen w-screen overflow-hidden "> -->
	<slot></slot>

<!-- </div> -->
