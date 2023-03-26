<script lang="ts">
	// import { HashList } from '$lib/UserInfo/HashList.svelte';
	import GroupCom from './component/groupCom.svelte';
	import HashCom from './component/hashCom.svelte';
	import Groupdecoration from "./svgs/groupdecoration.svelte";
	const rustLogo = new URL('../images/rust.png', import.meta.url).href;
	const awsLogo = new URL('../images/go.png', import.meta.url).href;

	const FavoriteHash: {
		ID: string;
		Name: string;
		Image: string;
		Description: string;
		CoverImage: string;
		NumberOfMember: number;
	}[] = [
		{
			ID: '1',
			Name: 'AWS AWS AWS AWS AWS AWS',
			Image: awsLogo,
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4',
			NumberOfMember: 10823890,
			
		},
		{
			ID: '2',
			Name: 'rust',
			Image: rustLogo,
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4',
				NumberOfMember: 10823890,
		
			},
		{
			ID: '3',
			Name: 'go',
			Image: awsLogo,
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4',
				NumberOfMember: 10823890,
		},
		{
			ID: '4',
			Name: 'jenkins',
			Image: rustLogo,
			Description: 'Null Pointer is a group of people who are interested in prog',
			CoverImage:
				'https://res.cloudinary.com/dqo0ssnti/video/upload/v1642488412/samples/elephants.mp4',
			NumberOfMember: 10823890,

		}
	];

    
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->
<div class=" text-[#e7e9eb] text-2xl ml-4 ">
    Groups 
</div>
<div class="w-auto h-auto  m-4 rounded-xl border-[1px]  border-gray-500 p-2 flex flex-col gap-1  ">
    {#if FavoriteHash.length !=0}
         <!-- content here -->
         {#each  FavoriteHash as i}
		 <HashCom HashData={i}/>
         {/each}
    {:else}
         <!-- else content here -->
        <div class="text-xl text-gray-500 h-24 grid place-content-center">No Group</div>
    {/if}

</div>