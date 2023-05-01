


<script lang="ts">
    
    // import Navbar from "$lib/Navbar/index.svelte"
    // import Collectives from "$lib/Collectives/index.svelte"
    import RelatedQues from "$lib/RelatedQues/index.svelte"
    // import Ques from "$lib/Ques/index.svelte"
    // import Ans from "$lib/Ans/index.svelte"
    // import {UserData,QuestionData,Loading} from "$lib/store/store"
    import LoadingSVG from "$lib/Loading/index.svelte"
    import SpaceCom from "$lib/SpaceCom/index.svelte"

import { onMount } from 'svelte';
import Navbar from '$lib/Navbar/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import Collectives from "$lib/Collectives/index.svelte"

// export let RelatedQuestionList:any

 

</script>

<style>
    /* your styles go here */
</style>


<div class="   flex   w-full flex-col  justify-center overflow-x-hidden overflow-y-hidden bg-[#181818] ">
	<Navbar />
	<div class="flex w-full flex-row justify-center   ">
		
		<!-- COLLECTIVE -->
		<Collectives />
		<!-- QUESTION LIST -->

        <p class="text-xl"> search result</p>
	</div>
	<Footer />
</div>






