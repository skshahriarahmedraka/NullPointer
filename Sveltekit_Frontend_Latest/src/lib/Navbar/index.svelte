<script lang="ts">
	// import { UserData } from '$lib/store/store';
	import { goto } from '$app/navigation';
	import K8s from '$lib/Navbar/svgs/k8s.svelte';
	import Following from '$lib/Navbar/svgs/following.svelte';
	import Spaces from './svgs/Spaces.svelte';
	import QuestionForYou from './svgs/QuestionForYou.svelte';
	import Search from './svgs/Search.svelte';
	import GoldDot from './svgs/GoldDot.svelte';
	import SilverDot from './svgs/SilverDot.svelte';
	import BronzeDot from './svgs/BronzeDot.svelte';
	import NotificationIcon from './svgs/NotificationIcon.svelte';
	import Svelte from './svgs/svelte.svelte';
	import Flutter from './svgs/Flutter.svelte';
	import Vue from './svgs/vue.svelte';
	import Menubar from './svgs/Menubar.svelte';
	import WhiteDot from './svgs/WhiteDot.svelte';
	import GPodcast from '$lib/icons/gPodcast.svelte';
	import Nullpointer from '$lib/icons/null pointer.svelte';
	import Stedia from '$lib/icons/stedia.svelte';
	import StediaRed from '$lib/icons/stediaRed.svelte';
	import Stediablue from '$lib/icons/stediablue.svelte';
	import { UserData, StoredSearchedString } from '$lib/store/store';

	console.log('🚀 ~ file: navBar.svelte:24 ~ UserData:', UserData);

	let SearchedStr: string = '';
	let UserdataLoading: boolean = true;
	function handleSearchKeyPress(event: any) {
		if (event.key === 'Enter') {
			if ($StoredSearchedString.trim().length > 0) {
				if (typeof document !== 'undefined') {
					// window.location.href = `/search/${$StoredSearchedString.trim()}`
				}
				goto(`/search/${$StoredSearchedString.trim()}`);
			}
		}
	}
	// let UserData: {
	// 	UserID: string;
	// 	UserName: string;
	// 	Email: string;
	// 	Password: string;
	// 	UserImage: string;
	// 	Badges: {
	// 		Reputation: number;
	// 		Gold: number;
	// 		Silver: number;
	// 		Bronze: number;
	// 	};
	// 	Follower: string[];
	// 	Following: string[];
	// 	Location: string;
	// 	MembershipTime: string;
	// 	LastSeen: string;
	// 	Aboutme: string;
	// 	Mysite: string;
	// 	Github: string;
	// 	Twitter: string;
	// 	Linkedin: string;
	// 	TopTags: {
	// 		Name: string;
	// 		Score: number;
	// 		NumberOfPost: number;
	// 	}[];

	// 	SelectedPanel: string;
	// 	AccountType: string;
	// } = {
	// 	UserID: 'skraka',
	// 	UserName: 'Sk Shahriar Ahmed Raka',
	// 	Email: 'skshahra@gmail.com',
	// 	Password: '123456',
	// 	// UserTitle string `json:"UserTitle"`
	// 	UserImage:
	// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
	// 	Badges: {
	// 		Reputation: 13452543,
	// 		Gold: 999,
	// 		Silver: 888,
	// 		Bronze: 777
	// 	},
	// 	Follower: ['RKA', 'SHAHRIAR	'],
	// 	Following: ['RKA', 'SHAHRIAR'],
	// 	// Badges map[string]int
	// 	Location: 'Dhaka, Bangladesh',
	// 	MembershipTime: '3 year 5 Month',
	// 	LastSeen: 'This Week',
	// 	Aboutme:
	// 		'A Curious Learner, Full-Stack Web Developer, Security Researcher\nHere are my skills and strengths:\n✓ Expert in Golang\n ✓ Expert in Fiber framework (using Golang) \n ✓ Expert in WebAssembly (using Golang)  Expert in Golang     ✓ Expert in Fiber framework (using Golang)    ✓ Expert in WebAssembly (using Golang) ✓ Expert in database design, development, optimization, and migration    (PostgreSQL, MySQL, MongoDB , Redis) ✓ Expert in ( Grpc, protocol buffer ) ✓ Experienced in ( WebSocket, Socket.io, WebRTC ) for real-time client and server applications ✓ Experienced in ( Svelte.js ) and some knowledge in ( TypeScript ) ✓ Good understanding of ( Docker, Bash, PowerShell, Git,    Nginx, Kubernetes )        Github : github.com/skshahriarahmedraka    Upwork : upwork.com/o/profiles/users/~0107ef3405bffbe57e    Linkedin : linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/  Telegram : t.me/shahriarraka ',
	// 	Mysite: 'www.shahriarraka.me',
	// 	Github: 'www.github.com/skshahriarahmedraka',
	// 	Twitter: 'www.twitter.com/shahriarraka7',
	// 	Linkedin: 'www.linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/',
	// 	// "TopTags"    : ["Go","Rust","Python","Svelte","PostgreSQL"],
	// 	TopTags: [
	// 		{ Name: 'Go', Score: 12, NumberOfPost: 4 },
	// 		{ Name: 'Rust', Score: 10, NumberOfPost: 6 },
	// 		{ Name: 'Python', Score: 7, NumberOfPost: 4 },
	// 		{ Name: 'Svelte', Score: 12, NumberOfPost: 4 },
	// 		{ Name: 'PostgreSQL', Score: 12, NumberOfPost: 4 }
	// 	],
	// 	SelectedPanel: 'Profile',
	// 	AccountType: 'regular'
	// };

	// let NavbarOpts = ['Profile', 'Followers', 'Followings', 'Other Networks', 'settings', 'logout'];
	// let NotificationOpts = [
	// 	'thanks for your ans',
	// 	'new Question new Question new Question new Question new Question',
	// 	'reply of your question ',
	// 	'thanks for your ans',
	// 	'new Question new Question new Question new Question new Question',
	// 	'reply of your question ',
	// 	'thanks for your ans',
	// 	'new Question new Question new Question new Question new Question',
	// 	'reply of your question ',
	// 	'thanks for your ans',
	// 	'new Question new Question new Question new Question new Question',
	// 	'reply of your question ',
	// 	'thanks for your ans',
	// 	'new Question new Question new Question new Question new Question',
	// 	'reply of your question ',
	// 	'thanks for your ans',
	// 	'new Question new Question new Question new Question new Question',
	// 	'reply of your question ',
	// 	'thanks for your ans',
	// 	'new Question new Question new Question new Question new Question',
	// 	'reply of your question ',
	// 	'thanks for your ans',
	// 	'new Question new Question new Question new Question new Question',
	// 	'reply of your question '
	// ];

	// let NotificationOpts2 = [
	// 	[
	// 		'answer',
	// 		'5 jul at 1:23',
	// 		'How to ensure translation with I18n is working in production?',
	// 		'either add the pass keyword to an empty function, also function args cannot have .s in them only valid names and * or ** followed by a name'
	// 	],
	// 	[
	// 		'comment',
	// 		'29 aug at 3:12',
	// 		'How build lib for 386 arch with cgo on windows?',
	// 		" have a golang library that builds and works well on Linux, MacOs and Windows. The problem comes when I'm trying to build it for 386 on the amd64 Windows VM"
	// 	],
	// 	['reward', '29 aug at 3:12', 20, 'Golang pass sensitive data as argument using docker'],
	// 	[
	// 		'answer',
	// 		'30 jan at 22:23',
	// 		'How to access a method from the extended abstract class when using it as a type constraint in a method',
	// 		"Except this doesn't work because it isn't assumed that T extends Animal so there I can't access it.What can I do make this getAnimalSound method working?"
	// 	],
	// 	[
	// 		'comment',
	// 		'29 aug at 3:12',
	// 		'What I would do in your place is to use overriding. However, it also only works on non static methods. (untested but should work)',
	// 		"Yeah, sorry I was a bit out of practice. However, you don't need a type parameter at all! You just say that the thing you want to have the sound from is an animal, doesn't matter which. If you give it a dog, dart will call the getSound method inside of the dog. I've updated my answer accordingly"
	// 	],
	// 	['reward', '29 aug at 3:12', -5, 'https and websocket handler with different deadlines'],
	// 	[
	// 		'reward',
	// 		'29 aug at 3:12',
	// 		-20,
	// 		'How can I add entire Array of structures into mysql database using GORM'
	// 	],
	// 	[
	// 		'comment',
	// 		'24 mar at 2:43',
	// 		'Encoding with image/jpeg cause image saturation / wrong pixels',
	// 		'I don\'t speak "Go" at all, but you appear to be trying to get and set RGBA values on JPEG images '
	// 	],
	// 	[
	// 		'comment',
	// 		'23 sep at 23;54',
	// 		"Can't install golangci-lint locally",
	// 		'I am dockerized a golang application and i am trying to access the application.'
	// 	],
	// 	[
	// 		'answer',
	// 		'5 jul at 1:23',
	// 		'How to ensure translation with I18n is working in production?',
	// 		'either add the pass keyword to an empty function, also function args cannot have .s in them only valid names and * or ** followed by a name'
	// 	],
	// 	[
	// 		'comment',
	// 		'29 aug at 3:12',
	// 		'How build lib for 386 arch with cgo on windows?',
	// 		" have a golang library that builds and works well on Linux, MacOs and Windows. The problem comes when I'm trying to build it for 386 on the amd64 Windows VM"
	// 	],
	// 	['reward', '29 aug at 3:12', 20, 'Golang pass sensitive data as argument using docker'],
	// 	[
	// 		'answer',
	// 		'30 jan at 22:23',
	// 		'How to access a method from the extended abstract class when using it as a type constraint in a method',
	// 		"Except this doesn't work because it isn't assumed that T extends Animal so there I can't access it.What can I do make this getAnimalSound method working?"
	// 	],
	// 	[
	// 		'comment',
	// 		'29 aug at 3:12',
	// 		'What I would do in your place is to use overriding. However, it also only works on non static methods. (untested but should work)',
	// 		"Yeah, sorry I was a bit out of practice. However, you don't need a type parameter at all! You just say that the thing you want to have the sound from is an animal, doesn't matter which. If you give it a dog, dart will call the getSound method inside of the dog. I've updated my answer accordingly"
	// 	],
	// 	['reward', '29 aug at 3:12', -5, 'https and websocket handler with different deadlines'],
	// 	[
	// 		'reward',
	// 		'29 aug at 3:12',
	// 		-20,
	// 		'How can I add entire Array of structures into mysql database using GORM'
	// 	],
	// 	[
	// 		'comment',
	// 		'24 mar at 2:43',
	// 		'Encoding with image/jpeg cause image saturation / wrong pixels',
	// 		'I don\'t speak "Go" at all, but you appear to be trying to get and set RGBA values on JPEG images '
	// 	],
	// 	[
	// 		'comment',
	// 		'23 sep at 23;54',
	// 		"Can't install golangci-lint locally",
	// 		'I am dockerized a golang application and i am trying to access the application.'
	// 	]
	// ];
	// let $UserData : UserDataType
	// UserData.subscribe((data) => {
	// 	$UserData = data;
	// });
	import type { NavShowType, NotificationDataType, UserDataType } from '$lib/store/types';
	import UserAnonymous from '$lib/icons/UserAnonymous.svelte';
	import Hash from '$lib/icons/hash.svelte';
	import Blog from '$lib/icons/blog.svelte';
	const NotificationData: NotificationDataType[] = [
		{
			type: 'answer',
			time: '5 jul at 1:23',
			title: 'How to ensure translation with I18n is working in production?',
			description:
				'either add the pass keyword to an empty function, also function args cannot have .s in them only valid names and * or ** followed by a name'
		},
		{
			type: 'comment',
			time: '29 aug at 3:12',
			title: 'How build lib for 386 arch with cgo on windows?',
			description:
				" have a golang library that builds and works well on Linux, MacOs and Windows. The problem comes when I'm trying to build it for 386 on the amd64 Windows VM"
		},
		{
			type: 'reward',
			time: '29 aug at 3:12',
			point: 20,
			title: 'Golang pass sensitive data as argument using docker'
		},
		{
			type: 'answer',
			time: '30 jan at 22:23',
			title:
				'How to access a method from the extended abstract class when using it as a type constraint in a method',
			description:
				"Except this doesn't work because it isn't assumed that T extends Animal so there I can't access it.What can I do make this getAnimalSound method working?"
		},
		{
			type: 'comment',
			time: '29 aug at 3:12',
			title:
				'What I would do in your place is to use overriding. However, it also only works on non static methods. (untested but should work)',
			description:
				"Yeah, sorry I was a bit out of practice. However, you don't need a type parameter at all! You just say that the thing you want to have the sound from is an animal, doesn't matter which. If you give it a dog, dart will call the getSound method inside of the dog. I've updated my answer accordingly"
		},
		{
			type: 'reward',
			time: '29 aug at 3:12',
			point: -20,
			title: 'How can I add entire Array of structures into mysql database using GORM'
		}
	];

	let NavbarOptsOpen = false;
	let NotificationOptsOpen = false;

	//   $: (NavbarOptsOpen ? (NotificationOptsOpen=false) : (NotificationOptsOpen=true))
	function AllOtherFalse() {
		NavbarOptsOpen = false;
		NotificationOptsOpen = false;
	}

	function ShortenNumber(i: number) {
		let bigFig: string[] = ['k', 'M', 'B', 'T'];
		let s: string = '';
		for (let x = 0; i > 999; x++) {
			if (i > 999) {
				i = i / 1000;
				s = bigFig[x];
			}
		}
		return String(i.toPrecision(1) + s);
	}
	let showComponent: NavShowType = {
		Logo: true,
		Hash: true,
		Blog: true,
		SearchInput: true,
		SearchIcon: true,
		Notification: true,
		Profile: true,
		Options: true
	};
	// 	if (window.innerWidth < 576) {
	// 		showComponent ={
	// 			"Logo": true,
	// 			"Hash": true,
	// 			"Blog":true ,
	// 			"SearchInput":false,
	// 			"SearchIcon":true,
	// 		"Notification":false,
	// 		"Profile" : false,
	// 		"Options":true,
	// 	}
	// }
	// 	else if (window.innerWidth < 767) {
	// 		showComponent ={
	// 			"Logo": true,
	// 			"Hash": true,
	// 			"Blog":true ,
	// 			"SearchInput":false,
	// 			"SearchIcon":true,
	// 		"Notification":true,
	// 		"Profile" : true,
	// 		"Options":true,
	// 	}}
	// 	else if (window.innerWidth < 991) {
	// 		showComponent ={
	// 			"Logo": true,
	// 			"Hash": true,
	// 			"Blog":true ,
	// 			"SearchInput":false,
	// 			"SearchIcon":true,
	// 		"Notification":true,
	// 		"Profile" : true,
	// 		"Options":true,
	// 	}}
	// 	else if (window.innerWidth < 1199) {
	// 		showComponent ={
	// 			"Logo": true,
	// 			"Hash": true,
	// 			"Blog":true ,
	// 			"SearchInput":false,
	// 			"SearchIcon":true,
	// 		"Notification":true,
	// 		"Profile" : true,
	// 		"Options":true,
	// 	}}
	// if (window.innerWidth < 1399) {
	// 		showComponent ={
	// 			"Logo": true,
	// 			"Hash": true,
	// 			"Blog":true ,
	// 			"SearchInput":false,
	// 			"SearchIcon":true,
	// 		"Notification":true,
	// 		"Profile" : true,
	// 		"Options":true,
	// 	}
	// }
	console.log('🚀 ~ file: index.svelte:310 ~ showComponent:', showComponent);
</script>

<div
	class=" top-0 flex h-[53px] w-full flex-row items-center justify-center overflow-hidden border-b-2 border-solid border-[#32353a] bg-[#262626] xs:justify-around sm:justify-around md:justify-around lg:gap-2 xl:gap-2"
>
	<!-- <div class=" grow" /> -->
	<!-- Company logo -->
	{#if showComponent.Logo}
		<button
			class="flex h-[53px] w-fit flex-row items-center justify-center gap-1 rounded-lg px-3 hover:bg-slate-800 xs:px-0"
			on:click={() => {
				goto('/');
			}}
		>
			<Stediablue class=" h-14  w-20   " />
			<!-- <GPodcast  /> -->
			<Nullpointer class="h-7  xs:hidden sm:hidden  " />
			<!-- <p class="font-serif text-xl leading-5 text-sky-500">Null Pointer</p> -->
		</button>
	{/if}
	<!-- hash -->
	{#if showComponent.Hash}
		<button class="rounded-lg px-2 hover:bg-slate-800" on:click={() => goto('/h')}>
			<Hash class=" h-8 fill-sky-600" />
		</button>
	{/if}
	<!-- following -->
	<!-- <Following /> -->
	<!-- spaces -->
	<!-- <button class="rounded-lg hover:bg-slate-800" on:click={() => goto('/g')}>
			<Spaces />
		</button> -->
	<!-- question for you -->
	{#if showComponent.Blog}
		<button
			class="w-fit rounded-lg pb-1 hover:bg-slate-800"
			on:click={() => {
				goto('/b');
			}}
		>
			<Blog class="h-8 w-fit fill-sky-600" />
			<!-- <QuestionForYou /> -->
		</button>
	{/if}
	<!-- search bar  -->
	<div class="mx-2 flex justify-center text-gray-200 xs:hidden sm:hidden md:hidden">
		<div class=" lg:w-[450px] xl:min-w-[500px] xl:max-w-[550px]">
			<div class="input-group relative flex w-full flex-row items-stretch">
				<!-- search input -->
				{#if showComponent.SearchInput}
					<input
						on:keypress={handleSearchKeyPress}
						bind:value={$StoredSearchedString}
						type="search"
						class="form-control relative m-0 block min-w-0 max-w-full flex-auto rounded border border-solid border-gray-800 bg-gray-600 bg-clip-padding px-3 py-1.5 text-base font-normal transition ease-in-out focus:border-blue-600 focus:bg-gray-600 focus:outline-none"
						placeholder="Search"
						aria-label="Search"
						aria-describedby="button-addon2"
					/>
				{/if}
				<!-- search button -->
			</div>
		</div>
	</div>
	{#if showComponent.SearchIcon}
		<button
			on:click={() => {
				goto(`/search/${$StoredSearchedString.trim()}`);
			}}
			class="  btn inline-block items-center rounded bg-sky-500 px-6 py-2.5 text-xs font-medium uppercase leading-tight text-white shadow-md transition duration-150 ease-in-out hover:bg-blue-600 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg xs:px-3"
			type="button"
			id="button-addon2"
		>
			<Search />
		</button>
	{/if}
	<!-- User Profile data -->
	{#if showComponent.Profile}
		{#if $UserData.UserID != ''}
			<div
				class=" flex flex-row items-center justify-center rounded-lg hover:cursor-pointer hover:bg-gray-700 hover:bg-opacity-50 lg:mx-2"
				on:click={() => {
					goto(`/${$UserData.UserID}`);
				}}
				on:keydown={(e) => {}}
			>
				<!-- User -->
				{#if $UserData.UserImage != ''}
					<!-- content here -->
					<img
						src={$UserData.UserImage}
						alt=""
						class=" aspect-square active:ring-offset-base-50 h-10 w-10 cursor-pointer rounded-lg object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500 active:rounded-md active:ring active:ring-blue-600"
					/>
				{:else}
					<!-- else content here -->
					<UserAnonymous
						class=" aspect-square active:ring-offset-base-50  h-10 w-10  cursor-pointer  object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
					/>
				{/if}
				<!-- user Credit -->
				<div
					class="max-w-36 flex h-fit min-w-min flex-row items-center justify-center xs:hidden sm:hidden"
				>
					{#if $UserData.Badges.Reputation != 0}
						<p class="mx-1 font-semibold text-white">
							{ShortenNumber($UserData.Badges.Reputation)}
						</p>
					{:else}
						<p class="mx-1 font-semibold text-white">
							{'1 Reputation'}
						</p>
					{/if}

					{#if $UserData.Badges.Gold !== 0}
						<GoldDot />
						<p class="mx-1 text-[#ffcc01]">{ShortenNumber($UserData.Badges.Gold)}</p>
					{/if}
					{#if $UserData.Badges.Silver !== 0}
						<SilverDot />
						<p class="mx-1 my-4 text-[#b4b8bc]">{ShortenNumber($UserData.Badges.Silver)}</p>
					{/if}
					{#if $UserData.Badges.Bronze !== 0}
						<BronzeDot />
						<p class="mx-1 my-4 text-[#d1a684]">{ShortenNumber($UserData.Badges.Bronze)}</p>
					{/if}
				</div>
			</div>

			<!-- notification -->
			<div class="   ">
				<button
					class=" flex h-full w-10 items-center justify-center rounded-lg hover:bg-slate-800 focus:outline-none xs:hidden sm:hidden md:hidden"
					on:blur={() => {
						NotificationOptsOpen = false;
					}}
					on:click={() => {
						if (NotificationOptsOpen) {
							NotificationOptsOpen = false;
						} else {
							AllOtherFalse();
							NotificationOptsOpen = true;
						}
					}}
				>
					<NotificationIcon />
				</button>
				{#if NotificationOptsOpen}
					<div
						class=" scrol3 fixed z-30 float-left -ml-[340px] mt-2 flex max-h-screen min-h-fit w-96 flex-col overflow-scroll rounded-lg border-2 border-gray-600 bg-gray-500 text-gray-200"
					>
						{#each NotificationData as i}
							<button
								class="flex h-fit w-full flex-row border-0 border-b-[1px] border-gray-700 bg-[#2d2d2d] px-4 py-2 text-[#3196e3] transition-all duration-150 ease-linear hover:bg-[#404245]"
							>
								<!-- <svg class="h-6 w-6" version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 98.1 118" style="enable-background:new 0 0 98.1 118;" xml:space="preserve" > <style type="text/css"> .st0 { fill: #ff3e00; } .st1 { fill: #ffffff; } </style> <path class="st0" d="M91.8,15.6C80.9-0.1,59.2-4.7,43.6,5.2L16.1,22.8C8.6,27.5,3.4,35.2,1.9,43.9c-1.3,7.3-0.2,14.8,3.3,21.3  c-2.4,3.6-4,7.6-4.7,11.8c-1.6,8.9,0.5,18.1,5.7,25.4c11,15.7,32.6,20.3,48.2,10.4l27.5-17.5c7.5-4.7,12.7-12.4,14.2-21.1  c1.3-7.3,0.2-14.8-3.3-21.3c2.4-3.6,4-7.6,4.7-11.8C99.2,32.1,97.1,22.9,91.8,15.6" /> <path class="st1" d="M40.9,103.9c-8.9,2.3-18.2-1.2-23.4-8.7c-3.2-4.4-4.4-9.9-3.5-15.3c0.2-0.9,0.4-1.7,0.6-2.6l0.5-1.6l1.4,1  c3.3,2.4,6.9,4.2,10.8,5.4l1,0.3l-0.1,1c-0.1,1.4,0.3,2.9,1.1,4.1c1.6,2.3,4.4,3.4,7.1,2.7c0.6-0.2,1.2-0.4,1.7-0.7L65.5,72  c1.4-0.9,2.3-2.2,2.6-3.8c0.3-1.6-0.1-3.3-1-4.6c-1.6-2.3-4.4-3.3-7.1-2.6c-0.6,0.2-1.2,0.4-1.7,0.7l-10.5,6.7  c-1.7,1.1-3.6,1.9-5.6,2.4c-8.9,2.3-18.2-1.2-23.4-8.7c-3.1-4.4-4.4-9.9-3.4-15.3c0.9-5.2,4.1-9.9,8.6-12.7l27.5-17.5  c1.7-1.1,3.6-1.9,5.6-2.5c8.9-2.3,18.2,1.2,23.4,8.7c3.2,4.4,4.4,9.9,3.5,15.3c-0.2,0.9-0.4,1.7-0.7,2.6l-0.5,1.6l-1.4-1  c-3.3-2.4-6.9-4.2-10.8-5.4l-1-0.3l0.1-1c0.1-1.4-0.3-2.9-1.1-4.1c-1.6-2.3-4.4-3.3-7.1-2.6c-0.6,0.2-1.2,0.4-1.7,0.7L32.4,46.1  c-1.4,0.9-2.3,2.2-2.6,3.8s0.1,3.3,1,4.6c1.6,2.3,4.4,3.3,7.1,2.6c0.6-0.2,1.2-0.4,1.7-0.7l10.5-6.7c1.7-1.1,3.6-1.9,5.6-2.5  c8.9-2.3,18.2,1.2,23.4,8.7c3.2,4.4,4.4,9.9,3.5,15.3c-0.9,5.2-4.1,9.9-8.6,12.7l-27.5,17.5C44.8,102.5,42.9,103.3,40.9,103.9" /> </svg>
							<div class="">
								<div
									class=" h-fit  w-full border-0 py-2 px-4 text-[#3196e3] transition-all duration-150  ease-linear   "
								>
									{i}
								</div>
							</div> -->
								<!-- type: string;
							time: string;
							title: string;
							description: string;
							point?: number; -->
								{#if i.type === 'answer'}
									<Svelte />
									<div class=" flex w-80 flex-col gap-1">
										<div class="StackTextColor flex flex-row font-mono">
											<div class="">{i.type}</div>
											<div class=" grow" />
											<div class="">{i.time}</div>
										</div>
										<div class=" line-clamp-1 text-left">
											{i.title}
										</div>
										<div class=" StackTextColor line-clamp-2 text-left leading-tight">
											{i.description}
										</div>
									</div>
								{:else if i.type === 'comment'}
									<Flutter />
									<div class=" flex w-80 flex-col gap-1">
										<div class="StackTextColor flex flex-row font-mono">
											<div class="">{i.type}</div>
											<div class=" grow" />
											<div class="">{i.time}</div>
										</div>
										<div class=" line-clamp-1 text-left">
											{i.title}
										</div>
										<div class=" StackTextColor line-clamp-2 text-left leading-tight">
											{i.description}
										</div>
									</div>
								{:else if i.type === 'reward'}
									<Vue />
									<div class=" flex w-80 flex-col gap-1">
										<div class="StackTextColor flex flex-row font-mono">
											{#if Number(i.point) < 0}
												<div class=" mx-1 text-xs font-semibold text-red-500">{i.point}</div>
											{:else}
												<div class=" mx-1 font-semibold text-green-400">+{i.point}</div>
											{/if}
											<div class="">{i.type}</div>
											<div class=" grow" />
											<div class="">{i.time}</div>
										</div>
										<div class=" line-clamp-1 text-left">
											{i.title}
										</div>
									</div>
								{/if}
							</button>
						{/each}
					</div>
				{/if}
			</div>

			<!-- menu -->
			<div class="">
				<button
					class=" flex h-full w-10 items-center justify-center overflow-hidden rounded-lg hover:bg-slate-800 focus:outline-none"
					on:blur={() => {
						NavbarOptsOpen = false;
					}}
					on:click={() => {
						if (NavbarOptsOpen) {
							NavbarOptsOpen = false;
						} else {
							AllOtherFalse();
							NavbarOptsOpen = true;
						}
					}}
				>
					<Menubar />
				</button>
				{#if NavbarOptsOpen}
					<div
						class=" fixed z-30 float-left -ml-[340px] mt-2 flex max-h-screen min-h-fit w-96 flex-col overflow-scroll rounded-lg border-2 border-gray-600 bg-[#2d2d2d] text-gray-200"
					>
						<div class="mb-1 ml-2 mt-3 flex flex-row">
							{#if $UserData.UserImage != ''}
								<!-- content here -->
								<img
									src={$UserData.UserImage}
									alt=""
									class=" aspect-square active:ring-offset-base-50 mx-2 h-10 w-10 cursor-pointer rounded-lg object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500 active:rounded-md active:ring active:ring-blue-600"
								/>
							{:else}
								<!-- else content here -->
								<UserAnonymous
									class=" aspect-square active:ring-offset-base-50 mx-2  h-10 w-10  cursor-pointer  object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
								/>
							{/if}
							<!-- <img
							src={$UserData['UserImage']}
							alt=""
							class=" aspect-square active:ring-offset-base-50 mx-2  h-16 w-16  cursor-pointer rounded-xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
						/> -->
							<div class="flex flex-col items-center justify-center">
								<div class=" line-clamp-1 text-2xl font-light">{$UserData.UserName}</div>
								<p class=" font-semibold">
									{ShortenNumber($UserData.Badges.Reputation)} Reputation <WhiteDot />
									{ShortenNumber($UserData.Follower.length)} Follower
								</p>
							</div>
						</div>
						<!-- <button
						class="flex h-fit w-full  flex-row border-0 border-b-[1px] border-gray-700 bg-[#2d2d2d]  py-2 px-4 text-[#3196e3] transition-all duration-150  ease-linear  hover:bg-[#404245] "
					>
						profile
					</button> -->
						<button
							class=" h-fit w-full border-0 border-b-[1px] border-gray-700 bg-[#2d2d2d] px-4 py-2 text-xl text-[#3196e3] transition-all duration-150 ease-linear hover:bg-[#404245]"
						>
							My followers
						</button>
						<button
							class=" h-fit w-full border-0 border-b-[1px] border-gray-700 bg-[#2d2d2d] px-4 py-2 text-xl text-[#3196e3] transition-all duration-150 ease-linear hover:bg-[#404245]"
						>
							My Following
						</button>
						<button
							class=" h-fit w-full border-0 border-b-[1px] border-gray-700 bg-[#2d2d2d] px-4 py-2 text-xl text-[#3196e3] transition-all duration-150 ease-linear hover:bg-[#404245]"
						>
							My Stats
						</button>
						<button
							class=" h-fit w-full border-0 border-b-[1px] border-gray-700 bg-[#2d2d2d] px-4 py-2 text-xl text-[#3196e3] transition-all duration-150 ease-linear hover:bg-[#404245]"
						>
							HeapOverflow Stats
						</button>
						<button
							class=" h-fit w-full border-0 border-b-[1px] border-gray-700 bg-[#2d2d2d] px-4 py-2 text-xl text-[#3196e3] transition-all duration-150 ease-linear hover:bg-[#404245]"
						>
							Settings
						</button>

						<button
							on:click={() => {
								goto('/api/logout');
							}}
							class=" h-fit w-full border-0 border-b-[1px] border-gray-700 bg-[#2d2d2d] px-4 py-2 text-xl text-red-500 transition-all duration-150 ease-linear hover:bg-[#404245]"
						>
							Logout
						</button>
					</div>
				{/if}
			</div>
		{:else}
			<!-- else content here -->
			<div class="flex h-full flex-row items-center justify-center gap-2">
				<UserAnonymous class="h-10 w-10 fill-blue-600" />
				<button
					on:click={() => {
						goto('/login');
					}}
					class=" my-2 rounded-md border-[1px] border-blue-500 bg-blue-500 bg-opacity-25 px-2 py-1 font-medium text-blue-600 hover:bg-opacity-0"
					>Sign In</button
				>
				<button
					on:click={() => {
						goto('/register');
					}}
					class=" my-2 rounded-md border-[1px] border-blue-500 bg-blue-700 bg-opacity-50 px-2 py-1 font-medium text-gray-100 hover:bg-opacity-0"
					>Sign Up</button
				>
			</div>
		{/if}
	{/if}

	<!-- <div class=" grow" /> -->
</div>
