<script>
	import '../app.postcss';
	
	const HeadLogo = new URL("../../static/favicon.png", import.meta.url).href;
</script>


<svelte:head>
	<title>Null Pointer</title>
	<link rel="icon"  href={HeadLogo} />
</svelte:head>

		<slot></slot>
