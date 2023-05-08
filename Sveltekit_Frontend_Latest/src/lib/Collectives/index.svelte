<script lang="ts">
	import { SpaceList } from '../store/store';

	// import {} from "../store/store"

	// export let FavoriteHash:any
	// export let SpaceList:any
	const rustLogo = new URL('../images/rust.png', import.meta.url).href;
	const awsLogo = new URL('../images/go.png', import.meta.url).href;

	let FOCUSID = '';
	// let HashFOCUSID:string=""
	function SetFocusID(i: string) {
		FOCUSID = i;
	}
	// function SetHashFOCUSID(i:string){
	//     HashFOCUSID=i
	// }

	const GroupList: {
		ID: string;
		Name: string;
		Image: string;
		Description: string;
		CoverImage: string;
		NumberOfMember: number;
	}[] = [
		{
			ID: '11',
			Name: 'Data Stucture and Algorithm',
			Image:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png',
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4',
			NumberOfMember: 10823890
		},
		{
			ID: '21',
			Name: 'Go programming',
			Image:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png',
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4',
			NumberOfMember: 10823890
		},
		{
			ID: '31',
			Name: 'Rust Programming',
			Image:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png',
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4',
			NumberOfMember: 10823890
		},
		{
			ID: '41',
			Name: 'SvelteJS and SvelteKit SvelteJS and SvelteKitSvelteJS and SvelteKit',
			Image:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png',
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4',
			NumberOfMember: 10823890
		},
		{
			ID: '41',
			Name: 'SvelteJSandSvelteKitSvelteJSandSvelteKitSvelteJSandSvelteKit',
			Image:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png',
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4',
			NumberOfMember: 10823890
		}
	];

	const FavoriteHash: {
		ID: string;
		Name: string;
		Image: string;
		Description: string;
		CoverImage: string;
	}[] = [
		{
			ID: '1',
			Name: 'AWS AWS AWS AWS AWS AWS',
			Image: awsLogo,
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4'
		},
		{
			ID: '2',
			Name: 'rust',
			Image: rustLogo,
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4'
		},
		{
			ID: '3',
			Name: 'go',
			Image: awsLogo,
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4'
		},
		{
			ID: '4',
			Name: 'jenkins',
			Image: rustLogo,
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4'
		}
	];
</script>

<div class=" mt-2  flex w-[165px] flex-col  ">
	<!-- My Favorite hash Tags -->
	<div class="h-fit w-full">
		<div
			class="mb-2 mt-4 h-10 w-full rounded-l-lg border-2 border-blue-500 pt-2 text-center font-semibold text-sky-600"
		>
			Favorites
		</div>
		<div class=" flex flex-col ">
			{#if FavoriteHash.length != 0}
				<!-- content here -->
				{#each FavoriteHash as i}
					<div
						on:click={() => SetFocusID(i.ID)}
						on:keypress={() => {}}
						class="  my-1 flex h-10 w-full flex-row  items-center justify-start gap-2  rounded-l-xl  text-sm  text-[#e7e9eb] hover:cursor-pointer hover:bg-[#3d3d3d]  {i.ID ===
						FOCUSID
							? ' border-y-2 border-r-4 border-l-2 border-y-[#2d2d2d] border-r-blue-600 border-l-[#2d2d2d] bg-[#2d2d2d]  '
							: ''}"
					>
						<img class="ml-1 h-6 w-6 object-cover" src={i.Image} alt="" />
						<!-- <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14"></path></svg> -->
						<!-- <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 4a1 1 0 011-1h4a1 1 0 010 2H6.414l2.293 2.293a1 1 0 11-1.414 1.414L5 6.414V8a1 1 0 01-2 0V4zm9 1a1 1 0 010-2h4a1 1 0 011 1v4a1 1 0 01-2 0V6.414l-2.293 2.293a1 1 0 11-1.414-1.414L13.586 5H12zm-9 7a1 1 0 012 0v1.586l2.293-2.293a1 1 0 111.414 1.414L6.414 15H8a1 1 0 010 2H4a1 1 0 01-1-1v-4zm13-1a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 010-2h1.586l-2.293-2.293a1 1 0 111.414-1.414L15 13.586V12a1 1 0 011-1z" clip-rule="evenodd"></path></svg> -->

						<div class="break-words text-xs  line-clamp-2  ">
							{i.Name}
						</div>
					</div>
				{/each}
			{:else}
				<!-- else content here -->
				<div class=" h-full w-full text-xl text-white ">No Favorite</div>
			{/if}
		</div>
	</div>

	<!-- <div class=" w-full h-60 bg-slate-500 m-2 "></div> -->
	<div
		class="mb-2 h-10   w-full rounded-l-lg border-2 border-blue-500 pt-2 text-center font-semibold text-sky-600"
	>
		Groups
	</div>
	{#if GroupList.length != 0}
		{#each GroupList as i}
			<div
				on:click={() => SetFocusID(i.ID)}
				on:keypress={() => {}}
				class="   my-1 flex h-10 w-full  flex-row  items-center justify-start gap-2  rounded-l-xl   pr-1 text-sm  text-[#e7e9eb] hover:cursor-pointer hover:bg-[#3d3d3d]  {i.ID ===
				FOCUSID
					? ' border-y-2 border-r-4 border-l-2 border-y-[#2d2d2d] border-r-blue-600 border-l-[#2d2d2d] bg-[#2d2d2d]  '
					: ''}"
			>
				<img class="ml-1 h-6 w-6 object-cover" src={i.Image} alt="" />
				<!-- <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14"></path></svg> -->
				<!-- <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 4a1 1 0 011-1h4a1 1 0 010 2H6.414l2.293 2.293a1 1 0 11-1.414 1.414L5 6.414V8a1 1 0 01-2 0V4zm9 1a1 1 0 010-2h4a1 1 0 011 1v4a1 1 0 01-2 0V6.414l-2.293 2.293a1 1 0 11-1.414-1.414L13.586 5H12zm-9 7a1 1 0 012 0v1.586l2.293-2.293a1 1 0 111.414 1.414L6.414 15H8a1 1 0 010 2H4a1 1 0 01-1-1v-4zm13-1a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 010-2h1.586l-2.293-2.293a1 1 0 111.414-1.414L15 13.586V12a1 1 0 011-1z" clip-rule="evenodd"></path></svg> -->
	
				<div class="break-words text-xs  line-clamp-2  ">
					{i.Name}
				</div>
			</div>
		{/each}
	{:else}
		<!-- else content here -->
		<div class=" h-full w-full text-xl text-white ">No Group</div>

	{/if}
</div>

<style>
	/* your styles go here */
</style>
