


<script lang="ts">
    import WriteQues from '$lib/WriteQues/index.svelte'
    import Navbar from '$lib/Navbar/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import Collectives from "$lib/Collectives/index.svelte"
    // import Navbar from "$lib/Navbar/index.svelte"
    // import Collectives from "$lib/Collectives/index.svelte"
    // import LoadingSVG from "$lib/Loading/index.svelte"
    // import {Loading} from "$lib/store/store"
    // export let HotQues:any

</script>

<style>
</style>





    
<!-- <div class="bg-[#181818]   w-full flex flex-row justify-center  pb-10 "> -->
        
    <!-- <Navbar/> -->
    <!-- COLLECTIVE -->
    <!-- <Collectives/> -->
    <!-- QUESTION LIST -->
    <!-- <Ques/> -->
    <div class="   flex   w-full flex-col  justify-center overflow-x-hidden overflow-y-hidden bg-[#181818] ">
        <Navbar />
        <div class="flex w-full flex-row justify-center   ">
            
            <!-- COLLECTIVE -->
            <Collectives />
            <!-- QUESTION LIST -->
    
            <WriteQues />
        </div>
        <Footer />
    </div>
        

        <!-- RELATED QUESTION -->
        <!-- <RelatedQues/> -->

    <!-- </div> -->

