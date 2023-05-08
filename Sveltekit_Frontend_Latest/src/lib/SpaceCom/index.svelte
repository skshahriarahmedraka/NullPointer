<script lang="ts">
	// import {QuestionData} from "../store/store"
	// import RelatedQues from "../RelatedQues/index.svelte"
	// import Ans from "../Ans/index.svelte"
	// import BgImg from "./img/dr.jpg"

import BlueDot from "./svgs/BlueDot.svelte";
import Filter from "./svgs/Filter.svelte";
import RoundPlus from "./svgs/RoundPlus.svelte";
import Sperkel from "./svgs/Sperkel.svelte";
import Verified from "./svgs/Verified.svelte";
let DefaultUser="img/u.jpg"
const DefaultUserImageUrl = new URL('./img/u.jpg', import.meta.url).href
const DefaultBannerUrl = new URL('./img/bro.jpg', import.meta.url).href

	let DropDownData = {
		Value: 'Recent Activity',
		Show: false,
		'a-z': true,
		List: ['Recent Activity', 'Newest', 'Highest score', 'Most frequent', 'bounty']
	};

	function DropDownClick() {
		DropDownData['Show'] = !DropDownData['Show'];
	}
	function DropDownBlur() {
		if (DropDownData['Show']) {
			DropDownData['Show'] = false;
		}
	}

    let SpaceList=[
        // ["SpaceName","verified", "ProfileImage","BannerImage", "SpaceReputation","SpaceMember"],
        ["Rust Programming hub",false, "","", 1234143,871382],
        ["Golang Programming Zone",true, "","", 47823098249,3423478243],
        ["Svelte Socity",true, "","", 28347932,237842378],
        ["laravel framework",false, "","", 979,3422],
        ["Cpp coder lab",true, "","", 345,345435],
        ["C programming made easy",true, "","", 435,5435343],
        ["javascript datastructure and algorighm ",true, "","https://res.cloudinary.com/dqo0ssnti/image/upload/v1642488377/sample.jpg", 278422434,23423234],
        ["New to Typescript",false, "","", 9999,97788],
        ["Backend programmers",false, "","", 678678,678],
        ["fiber framework",false, "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png","", 4878,267837],
        ["Gin framwork made easy",true, "","", 78,783],
        ["Data Structure and Algorighm",true, "","", 43543,787823],
        ["C# software and game development",false, "","", 78243,23478],
        ["github and gitlab",true, "","", 78621,9898],
        ["devOps pro max",true, "","", 78243,2378],
        ["linux",true, "","", 8,78],
        ["Azure cloud",true, "","", 7,5],
    ]
	let TrendingSpaceList=[
        ["Java programming online",true, "","", 34,234],
        ["spring boot framework",true, "","", 7567,5],
        ["PHP coding",false, "","", 85677,6785],
        ["Competitive programming",true, "","", 756343,5],
        ["Graph Algorighm",true, "","", 7565678,556756],
        ["C programming",true, "","", 6,556567],
        ["kernel developer",true, "","", 6787,567677],
        ["System Developer",true, "","", 7,55484],
        ["Encription decryption",false, "","", 357,50],
        ["Golang interview question",true, "","", 56787,65],
        ["software security",true, "","", 7675875678,54568456567],
        ["linux developer",true, "","", 7678678,584678],

	]

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

</script>

<div class=" mt-12 flex max-h-full  min-h-screen w-[1100px] flex-row bg-[#2d2d2d]">
	<div class=" flex w-9/12 flex-col ">
		<!-- TITLE AND FILTER -->
		<div class="m-4 flex h-16 w-full flex-row  ">
			<div class=" mt-3 text-2xl font-bold text-[#e7e9eb] ">My Spaces</div>
			<div class="grow" />
			<!-- DROP dOWN -->
			<div class="m-2 mr-4 ">
				<div class=" relative z-10 inline-block ">
					<button
						on:blur={DropDownBlur}
						on:click={DropDownClick}
						class="inline-flex items-center rounded border-2 border-[#688fac]  p-1 pl-2 text-lg font-medium text-[#e7e9eb]"
					>
						<Filter/>
						<span class="mr-1 ">{DropDownData['Value']}</span>
					</button>
					{#if DropDownData['Show']}
						<ul
							class=" absolute w-full rounded border-2 border-[#688fac] bg-slate-700 text-lg font-medium  text-[#e7e9eb]    "
						>
							{#each DropDownData['List'] as i}
								<button class=" p-1 pl-2 hover:cursor-pointer hover:bg-slate-500  ">{i}</button>
							{/each}
						</ul>
					{/if}
				</div>
			</div>
		</div>
		<!-- Create or descover spaces -->
		<div class="mx-5 flex h-8 w-full flex-row gap-3 ">
			<!-- CREATE -->
			<div
				class=" flex items-center justify-center rounded-3xl  border-2 border-blue-600 pr-3 pl-2 text-sm text-blue-600 hover:cursor-pointer hover:bg-sky-300 hover:bg-opacity-10"
			>
				<RoundPlus/>
				<div class="ml-1 font-semibold">Create Space</div>
			</div>
			<!-- DISCOVER -->
			<div
				class=" flex items-center justify-center rounded-3xl border-2 border-blue-600 pr-3 pl-2 text-sm text-blue-600 hover:cursor-pointer hover:bg-sky-300 hover:bg-opacity-10"
			>
				<Sperkel/>
				<div class="ml-1 font-semibold">Discover Spaces</div>
			</div>
		</div>
		<!-- SPACES -->
		<div class=" m-5 flex w-full flex-row flex-wrap  gap-5  ">
			{#each SpaceList as i}
				<div
					class=" flex h-48  w-64 flex-col rounded-xl border-2   border-blue-600 border-opacity-25  bg-slate-700 bg-opacity-30 fill-[#b4b8bc] text-[#e7e9eb] hover:cursor-pointer hover:bg-opacity-60"
					>
					<!-- SPACE WALL PHOTO -->
					
					<div
                    
					class="flex h-2/3 w-full items-end justify-start rounded-t-xl pl-2 "
					style=" background-image: url('{i[3]===""? DefaultBannerUrl : i[3] }')"
                        
					>
						<!-- SPACE PROFILE PHOTO -->
						<img
							src="{i[2]===""? DefaultUserImageUrl : String(i[2]) }"
							alt=""
							class=" aspect-square active:ring-offset-base-50 h-14 w-14  cursor-pointer rounded-2xl  border-4 border-[#2d2d2d] object-cover transition-all duration-150 ease-linear hover:rounded-2xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
						/>
					</div>
					<div class="flex h-1/3 w-full flex-col ">
						<div class=" m-2 flex h-1/2 w-full flex-row text-base font-semibold ">
							<!-- SPACE NAME -->
							<p class="inline w-[92%] line-clamp-1">
								{i[0]}
							</p>
							<Verified/>
						</div>
						<div class=" h-1/2 w-full text-center text-sm">
							<p class="inline font-semibold text-white">{ShortenNumber(Number(i[4]))}</p>
							<p class="inline font-semibold text-white">Reputation</p>
							<BlueDot/>
							<p class="inline font-semibold text-white">{ShortenNumber(Number(i[5]))}</p>
							<p class="inline font-semibold text-white">Members</p>
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
	<!-- Trending Spaces -->
	<div class="flex  w-3/12 flex-col items-center justify-start gap-4 ">
		<div class=" mt-3 text-2xl font-bold text-sky-400 ">Trending Spaces</div>

		{#each TrendingSpaceList as i}
				<div
					class=" flex h-48  w-64 flex-col rounded-xl border-2   border-blue-600 border-opacity-25  bg-slate-700 bg-opacity-30 fill-[#b4b8bc] text-[#e7e9eb] hover:cursor-pointer hover:bg-opacity-60"
					>
					<!-- SPACE WALL PHOTO -->
					
					<div
                    
					class="flex h-2/3 w-full items-end justify-start rounded-t-xl pl-2 "
					style=" background-image: url('{i[3]===""? DefaultBannerUrl : i[3] }')"
                        
					>
						<!-- SPACE PROFILE PHOTO -->
						<img
							src="{i[2]===""? DefaultUserImageUrl : String(i[2]) }"
							alt=""
							class=" aspect-square active:ring-offset-base-50 h-14 w-14  cursor-pointer rounded-2xl  border-4 border-[#2d2d2d] object-cover transition-all duration-150 ease-linear hover:rounded-2xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
						/>
					</div>
					<div class="flex h-1/3 w-full flex-col ">
						<div class=" m-2 flex h-1/2 w-full flex-row text-base font-semibold ">
							<!-- SPACE NAME -->
							<p class="inline w-[92%] line-clamp-1">
								{i[0]}
							</p>
							<Verified/>
						</div>
						<div class=" h-1/2 w-full text-center text-sm">
							<p class="inline font-semibold text-white">{ShortenNumber(Number(i[4]))}</p>
							<p class="inline font-semibold text-white">Reputation</p>
							<BlueDot/>
							<p class="inline font-semibold text-white">{ShortenNumber(Number(i[5]))}</p>
							<p class="inline font-semibold text-white">Members</p>
						</div>
					</div>
				</div>
			{/each}
	</div>
</div>

<style>
   
    
</style>
