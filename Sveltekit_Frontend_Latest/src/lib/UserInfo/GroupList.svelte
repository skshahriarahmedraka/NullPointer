<script lang="ts">
	// import { GroupList } from '$lib/UserInfo/GroupList.svelte';
	import GroupCom from './component/groupCom.svelte';
	import Groupdecoration from "./svgs/groupdecoration.svelte";

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
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg',
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
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg',
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

    
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->
<div class=" text-[#e7e9eb] text-2xl ml-4 ">
    Groups 
</div>
<div class="w-auto h-auto  m-4 rounded-xl border-[1px]  border-gray-500 p-2 flex flex-col gap-1  ">
    {#if GroupList.length !=0}
         <!-- content here -->
         {#each  GroupList as i}
              <GroupCom GroupData={i}/>
         {/each}
    {:else}
         <!-- else content here -->
        <div class="text-xl text-gray-500 h-24 grid place-content-center">No Group</div>
    {/if}

</div>