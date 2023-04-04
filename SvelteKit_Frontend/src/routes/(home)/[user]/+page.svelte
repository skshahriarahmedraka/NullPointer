<script lang="ts">
	import UserInfo from '$lib/UserInfo/Title.svelte';
	import Stats from '$lib/UserInfo/Stats.svelte';
	import Abouts from '$lib/UserInfo/Abouts.svelte';
	// // import { UserData } from '$lib/store/store';
	import TopTags from '$lib/UserInfo/TopTags.svelte';
	import Badges from '$lib/UserInfo/Badges.svelte';
	import GroupList from '$lib/UserInfo/GroupList.svelte';
	import HashList from '$lib/UserInfo/HashList.svelte';

	let UserData: {
		UserID: string;
		UserName: string;
		Email: string;
		Password: string;
		UserImage: string;
		Badges: {
			Reputation: number;
			Gold: number;
			Silver: number;
			Bronze: number;
		};
		Follower: string[];
		Following: string[];
		Location: string;
		MembershipTime: string;
		LastSeen: string;
		Aboutme: string;
		Mysite: string;
		Github: string;
		Twitter: string;
		Linkedin: string;
		TopTags: {
			Name: string;
			Score: number;
			NumberOfPost: number;
		}[];

		SelectedPanel: string;
		AccountType: string;
	} = {
		UserID: 'skraka',
		UserName: 'Sk Shahriar Ahmed Raka',
		Email: 'skshahra@gmail.com',
		Password: '123456',
		// UserTitle string `json:"UserTitle"`
		UserImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
		Badges: {
			Reputation: 13452543,
			Gold: 999,
			Silver: 888,
			Bronze: 777
		},
		Follower: ['RKA', 'SHAHRIAR	'],
		Following: ['RKA', 'SHAHRIAR'],
		// Badges map[string]int
		Location: 'Dhaka, Bangladesh',
		MembershipTime: '3 year 5 Month',
		LastSeen: 'This Week',
		Aboutme:
			'A Curious Learner, Full-Stack Web Developer, Security Researcher\nHere are my skills and strengths:\n✓ Expert in Golang\n ✓ Expert in Fiber framework (using Golang) \n ✓ Expert in WebAssembly (using Golang)  Expert in Golang     ✓ Expert in Fiber framework (using Golang)    ✓ Expert in WebAssembly (using Golang) ✓ Expert in database design, development, optimization, and migration    (PostgreSQL, MySQL, MongoDB , Redis) ✓ Expert in ( Grpc, protocol buffer ) ✓ Experienced in ( WebSocket, Socket.io, WebRTC ) for real-time client and server applications ✓ Experienced in ( Svelte.js ) and some knowledge in ( TypeScript ) ✓ Good understanding of ( Docker, Bash, PowerShell, Git,    Nginx, Kubernetes )        Github : github.com/skshahriarahmedraka    Upwork : upwork.com/o/profiles/users/~0107ef3405bffbe57e    Linkedin : linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/  Telegram : t.me/shahriarraka ',
		Mysite: 'www.shahriarraka.me',
		Github: 'www.github.com/skshahriarahmedraka',
		Twitter: 'www.twitter.com/shahriarraka7',
		Linkedin: 'www.linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/',
		// "TopTags"    : ["Go","Rust","Python","Svelte","PostgreSQL"],
		TopTags: [
			{ Name: 'Go', Score: 12, NumberOfPost: 4 },
			{ Name: 'Rust', Score: 10, NumberOfPost: 6 },
			{ Name: 'Python', Score: 7, NumberOfPost: 4 },
			{ Name: 'Svelte', Score: 12, NumberOfPost: 4 },
			{ Name: 'PostgreSQL', Score: 12, NumberOfPost: 4 }
		],
		SelectedPanel: 'Profile',
		AccountType: 'regular'
	};

	// export let UserData:any
	// console.log("🚀 ~ file: index@root.svelte ~ line 47 ~ UserData", UserData)
</script>

<!-- settings -->
<div class=" mt-2 flex max-h-full  min-h-screen w-[1100px] flex-col bg-[#2d2d2d]   ">
	<!-- image & name & social media -->
	<UserInfo />
	<!-- Button : profile, activity, -->
	<div class="  flex w-full flex-row items-center justify-start gap-5 pl-10 ">
		<div
			on:click={() => (UserData.SelectedPanel = 'Profile')}
			on:keydown={() => {}}
			class="ml-3 -mr-3 h-9 w-24 rounded-xl {UserData.SelectedPanel === 'Profile'
				? 'bg-sky-500 hover:bg-blue-600'
				: 'border-2 border-blue-600 '}  flex items-center justify-center hover:cursor-pointer "
		>
			<p
				class=" my-auto  text-base font-semibold {UserData.SelectedPanel === 'Profile'
					? 'text-gray-200'
					: 'text-sky-500'}"
			>
				Profile
			</p>
		</div>
		<div
			on:click={() => (UserData.SelectedPanel = 'Activity')}
			on:keydown={() => {}}
			class="ml-3 -mr-3 h-9 w-24 rounded-xl {UserData.SelectedPanel === 'Activity'
				? 'bg-sky-500 hover:bg-blue-600'
				: 'border-2 border-blue-600 '}   flex items-center justify-center hover:cursor-pointer "
		>
			<p
				class=" my-auto  text-base font-semibold {UserData.SelectedPanel === 'Activity'
					? 'text-gray-200'
					: 'text-sky-500'}"
			>
				Activity
			</p>
		</div>
		<div
			on:click={() => (UserData.SelectedPanel = 'Settings')}
			on:keydown={() => {}}
			class="ml-3 -mr-3 h-9 w-24 rounded-xl {UserData.SelectedPanel === 'Settings'
				? 'bg-sky-500 hover:bg-blue-600'
				: 'border-2 border-blue-600 '}  flex items-center justify-center hover:cursor-pointer "
		>
			<p
				class=" my-auto  text-base font-semibold {UserData['SelectedPanel'] === 'Settings'
					? 'text-gray-200'
					: 'text-sky-500'}"
			>
				Settings
			</p>
		</div>
	</div>
	<!-- details about me -->
	{#if UserData['SelectedPanel'] === 'Profile'}
		<div class="mt-5 flex h-full w-full flex-row ">
			<div class="flex h-full w-[300px] flex-col ">
				<Stats />
				<TopTags />
				<GroupList />
				<HashList/>
			</div>
			<!-- Stats and Tags -->
			<!-- About -->
			<div class="flex h-full w-9/12 flex-col  ">
				<Abouts />
				<Badges />
			</div>
		</div>
		<!-- content here -->
	{:else if UserData['SelectedPanel'] === 'Activity'}
		<div class="">Activity</div>
	{:else}
		<div class="">Settings</div>
		<!-- else content here -->
	{/if}
</div>

<style>
	/* your styles go here */
</style>
