<script lang="ts">
	// import { BlogDataType } from '$lib/store/types.ts';
	import type { HashArrWithMetadata, HashDataType } from '$lib/store/types';
	import Book from '$lib/Loading/book.svelte';
	import { onMount } from 'svelte';
	import Emptybox from '$lib/errors/emptybox.svelte';
	import { fetchHashArrWithMetadata, fetchUserFlairData } from '$lib/store/fetch';
	import UserAnonymous from '$lib/icons/UserAnonymous.svelte';
	import { goto } from '$app/navigation';

	// import {QuestionData} from "../store/store"
	// let Loading = false;

    let Hashlist :HashArrWithMetadata = {} as HashArrWithMetadata


	// let QuestionList:QuesArrWithMetadataType = {} as QuesArrWithMetadataType;
	// let fetchData = async () => {

	// }
	// fetchData()

	import ZtoA from '$lib/PublicQues/svgs/ZtoA.svelte';
	import PageNum from '$lib/PageNum/pageNum.svelte';
	import Verified from '$lib/SpaceCom/svgs/Verified.svelte';
	import BlueDot from '$lib/SpaceCom/svgs/BlueDot.svelte';

	let filterState: boolean =false
	let filterType: string ="follower"
	let pageNumStart: number=0;
	let pageNumNow: number =0;
	let pageNumEnd: number=0;
	$: console.log("pageNumNow:", pageNumNow)
	$: console.log("pageNumStart:", pageNumStart)
	$: console.log("pageNumEnd:", pageNumEnd)
	let contentPerPage:number=5
	let Loading:boolean=false
	onMount(async () => {
		filterState = true;
	 filterType = 'alpha';
     Hashlist = await fetchHashArrWithMetadata(
			filterType,
			pageNumNow*contentPerPage,
			pageNumNow*contentPerPage + contentPerPage,
			filterState ? -1 : 1
			);
		pageNumStart = 0;
		pageNumEnd = Math.floor(Hashlist.Metadata.Length / contentPerPage);
		Loading=true
	});
	async function ApplyFilter(filterType: string, filterState: boolean) {
		// time , views , unanswered , votes
		if (filterState) {
			Hashlist = await fetchHashArrWithMetadata(
				filterType,
				pageNumNow*contentPerPage,
				pageNumNow*contentPerPage + contentPerPage,
				-1
			);
		} else {
			Hashlist = await fetchHashArrWithMetadata(filterType, pageNumNow*contentPerPage, pageNumNow*contentPerPage + contentPerPage, 1);
		}
		
		pageNumStart = 0;
		pageNumEnd = Math.floor(Hashlist.Metadata.Length / contentPerPage);

	}


	$: pageNumNow , ApplyFilter(filterType, filterState)
	function ShortenNumber(i:number){
		let bigFig:string[]=["k","M","B","T"]
		let s:string=""
		for (let x=0;i>999;x++){
			if (i>999){
				i=i/1000
				s=bigFig[x]
			}
		}
		return String(i.toPrecision(3)+s)
	}
	function RoundNum(x: number) {
		if (x === null) {
			return 0;
		} else if (x >= 1000000000) {
			x /= 1000000000;
			return String(x.toFixed(2)) + 'B';
		} else if (x >= 1000000) {
			x /= 1000000;
			return String(x.toFixed(2)) + 'M';
		} else if (x >= 1000) {
			x /= 1000;
			return String(x.toFixed(2)) + 'K';
		} else {
			return String(x);
		}
	}
</script>

{#if Loading}
	<div class=" mt-2  min-h-screen w-[1355px] bg-[#2d2d2d] pl-5  flex flex-col">
        <div class="flex h-16 w-full flex-row">
            <div class="ml-5 self-center text-lg text-[#e7e9eb]">{Hashlist.Metadata.Length} Hashes</div>
            <div class="grow" />
            <div class=" flex h-9 flex-row justify-around self-center rounded-md border-2 border-[#7d858d]">
                <div
                    on:click={() => {
                        filterType='alpha'
                        ApplyFilter('alpha', filterState);
                        // CalculatePages();
                    }}
                    on:keypress={() => {}}
                    class="flex h-full items-center justify-center border-r-2 border-[#7d858d] px-2 text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] {filterType ===
                    'alpha'
                        ? 'bg-[#404245]'
                        : ''}  "
                >
                    Alphabetical
                </div>
                <div
                    on:click={() => {
                        filterType='follower'
                        ApplyFilter('follower', filterState);
                        // CalculatePages();
                    }}
                    on:keypress={() => {}}
                    class="flex h-full items-center justify-center border-r-2 border-[#7d858d] px-2 text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] {filterType ===
                    'follower'
                        ? 'bg-[#404245]'
                        : ''}"
                >
                    Followers
                </div>
                <div
                    on:click={() => {
                        filterType='popular'
                        ApplyFilter('popular', filterState);
                        // CalculatePages();
                    }}
                    on:keypress={() => {}}
                    class="flex h-full items-center justify-center border-r-2 border-[#7d858d] px-2 text-[#9cc1db] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245] {filterType ===
                    'popular'
                        ? 'bg-[#404245]'
                        : ''}"
                >
                    Popular
                </div>
                
                <!-- <div
                    class="flex h-full w-16 items-center justify-center text-[#e7e9eb] hover:cursor-pointer hover:bg-[#404245]"
                >
                    More
                </div> -->
            </div>
            <div
                on:click={() => {
                    filterState = !filterState;
                    ApplyFilter(filterType, filterState);
                    // CalculatePages();
                }}
                on:keydown={() => {}}
                class=" ml-4 flex h-9 w-24 flex-row items-center justify-center self-center rounded-md border-2 border-[#7d858d] hover:cursor-pointer hover:bg-[#3d4951] hover:text-[#9cc1db] active:bg-[#404245]"
            >
                <ZtoA class=" {filterState ? ' rotate-180' : ''}" />
                <div class="ml-1 font-semibold text-[#9cc1db]">order</div>
            </div>
        </div>

        <div class="">

		{#if Hashlist.HashList.length != 0}
			<!-- content here -->

			
			
			<div class="flex flex-row flex-wrap gap-2">
				{#each Hashlist.HashList as i}
				
						<div
						on:click={()=>{
							goto('/h/'+i.ID)	
						}}
						on:keypress={()=>{}}
					class=" flex h-48  w-64 flex-col rounded-xl border-2   border-blue-600 border-opacity-25  bg-slate-700 bg-opacity-30 fill-[#b4b8bc] text-[#e7e9eb] hover:cursor-pointer hover:bg-opacity-60"
					>
					<!-- SPACE WALL PHOTO -->
					
					<div
                    
					class="flex h-2/3 w-full items-end justify-start rounded-t-xl pl-2 "
					style=" background-image: url('{i.BannerImage}')"
                        
					>
						<!-- SPACE PROFILE PHOTO -->
						<img
							src="{i.Image }"
							alt=""
							class=" aspect-square active:ring-offset-base-50 h-14 w-14  cursor-pointer rounded-2xl  border-4 border-[#2d2d2d] object-cover transition-all duration-150 ease-linear hover:rounded-2xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
						/>
					</div>
					<div class="flex h-1/3 w-full flex-col ">
						<div class=" m-2 flex h-1/2 w-full flex-row text-base font-semibold ">
							<!-- SPACE NAME -->
							<p class="inline w-[92%] line-clamp-1">
								{i.Name}
							</p>
							<Verified/>
						</div>
						<div class=" h-1/2 w-full text-center text-sm">
							<p class="inline font-semibold text-white">{RoundNum(Number(i.NumOfFollower))}</p>
							<p class="inline font-semibold text-white">Follower</p>
							<BlueDot/>
							<p class="inline font-semibold text-white">{RoundNum(Number(i.NumOfQuestion))}</p>
							<p class="inline font-semibold text-white">Question</p>
						</div>
					</div>
				</div>
				{/each}
			</div>
			
		{:else}
			<div class=" flex h-full w-full items-center justify-center text-xl text-white">
				<div class="flex flex-col items-center justify-center">
					<Emptybox class="h-20 " />
					<p class=" font-raleway">Nothing Found </p>
				</div>
			</div>
		{/if}
    </div>

	</div>
{:else}
	<div
		class=" mt-2 flex max-h-full min-h-screen w-[1100px] items-center justify-center bg-[#2d2d2d] pl-5"
	>
		<Book />
	</div>
{/if}

<style>
</style>
