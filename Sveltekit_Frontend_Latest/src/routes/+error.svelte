<script lang="ts">
	import Animated404 from '$lib/errors/animated404.svelte';
	import Cat404 from '$lib/errors/cat404.svelte';
	import Robot404 from '$lib/errors/robot404.svelte';
</script>

<div class="   ">
	<button
    on:click={() => {
      window.location.href = '/';
    }}
		class=" absolute bottom-10 left-0 right-0 ml-auto mr-auto flex h-16 items-center justify-center rounded-3xl bg-[#3cbfc5] bg-opacity-40 p-5 text-center text-xl text-white "
	>
		Go Back Home
	</button>
	<!-- <Robot404/> -->
	<Cat404 />
</div>

<style>
	/* your styles go here */
</style>
	