<script>
	import  LoadingSVG  from '$lib/Loading/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';
import '../app.postcss';
	
	const HeadLogo = new URL("../../static/favicon.png", import.meta.url).href;

	import { onMount } from 'svelte';
  import { writable } from 'svelte/store';
  import { beforeNavigate, afterNavigate } from '$app/navigation';

  // set up a store to keep track of the loading state
  const isLoading = writable(false);

  onMount(() => {
    // add event listeners to the router
    beforeNavigate(() => {
      isLoading.set(true);
    });

    afterNavigate(() => {
      isLoading.set(false);
    });
  });
</script>


<svelte:head>
	<title>Null Pointer</title>
	<link rel="icon"  href={HeadLogo} />
</svelte:head>

<!-- <div class=" h-screen w-screen overflow-hidden "> -->


  <slot></slot>



<!-- show page content -->


<!-- </div> -->
