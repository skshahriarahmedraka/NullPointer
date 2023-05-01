

<script lang="ts">
	// your script goes here
	import Navbar from '$lib/Navbar/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import Collectives from "$lib/Collectives/index.svelte"
</script>

<!-- markup (zero or more items) goes here -->

<div class="   flex   w-full flex-col  justify-center overflow-x-hidden overflow-y-hidden bg-[#181818] ">
	<Navbar />
	<div class="flex w-full flex-row justify-center   ">
		
		<!-- COLLECTIVE -->
		<Collectives />
		<!-- QUESTION LIST -->

		<h1 class=" flex flex-col text-white items-center justify-center ">Following</h1>
	</div>
	<Footer />
</div>



