<script lang="ts">
	// import Title from './../../../../lib/UserInfo/Title.svelte';
	import Wyswyg from '$lib/WriteQues/wyswyg.svelte';
	import { UserData } from '$lib/store/store';
	import type { CookieInfo1Type, UserDataType } from '$lib/store/types';
	import { fetchUpdateUserData, fetchUserData } from '$lib/store/fetch';
	import Navbar from '$lib/Navbar/index.svelte';
	import Footer from '$lib/Footer/footer.svelte';
	import Collectives from '$lib/Collectives/index.svelte';
	import { getCookieValue } from '$lib/store/utils';
	import { onMount } from 'svelte';

	// export let $UserData:any
	let NewUserData: UserDataType = {
		UserID: $UserData.UserID,
		UserURL: $UserData.UserURL,
		UserName: $UserData['UserName'],
		Email: $UserData.Email,
		Password: '',
		UserTitle: $UserData['UserTitle'],
		UserImage: $UserData['UserImage'],
		Badges: $UserData.Badges,
		Follower: $UserData['Follower'],
		Following: $UserData['Following'],
		Location: $UserData['Location'],
		MembershipTime: $UserData['MembershipTime'],
		LastSeen: $UserData['LastSeen'],
		Aboutme: $UserData['Aboutme'],
		Mysite: $UserData['Mysite'],
		TopTags: $UserData['TopTags'],
		SelectedPanel: 'Profile',
		AccountType: $UserData.AccountType,
		Github: $UserData['Github'],
		Twitter: $UserData['Twitter'],
		Linkedin: $UserData['Linkedin']
	};
	let ErrorMsg = {
		UserURL: [false, 'UserURL is not Valid'],
		UserName: [false, 'UserName is not Valid'],
		Email: [false, 'invalid/wrong email'],
		Password: [false, 'Password is too short'],
		UserTitle: [false, 'UserTitle is not Valid'],
		// UserImage: [false, 'UserImage is not Valid'],
		Location: [false, 'Location is not Valid'],
		Aboutme: [false, 'Aboutme is not Valid'],
		Mysite: [false, 'Mysite is not Valid'],
		Github: [false, 'Github is not Valid'],
		Twitter: [false, 'Twitter is not Valid'],
		Linkedin: [false, 'Linkedin is not Valid']
	};
	let file: File | null = null;
	// let FileInput: HTMLInputElement;

	const handleFileChange = (event: Event) => {
		const inputElement = event.target as HTMLInputElement;
		file = inputElement.files?.[0] ?? null;
		handleSubmit();
	};

	const handleSubmit = async () => {
		if (!file) return;
		const formData = new FormData();
		formData.append('image', file);
		let ResObj: { ImageUrl: string } = {} as { ImageUrl: string };
		try {
			const response = await fetch('/api/uploadimage', {
				method: 'POST',
				body: formData
			});
			ResObj = await response.json();
			NewUserData.UserImage = ResObj.ImageUrl;
			console.log('🚀 ~ file: +page.svelte:42 ~ handleFileChange ~ NewUserData:', NewUserData);

			// handle server response
		} catch (error) {
			console.error(error);
		}
	};

	// 	import axios from 'axios';

	// // Function to split the file into chunks
	// function chunkFile(file: File, chunkSize: number): Blob[] {
	//   const chunks: Blob[] = [];
	//   let offset = 0;
	//   while (offset < file.size) {
	//     const chunk = file.slice(offset, offset + chunkSize);
	//     chunks.push(chunk);
	//     offset += chunk.size;
	//   }
	//   return chunks;
	// }

	// // Function to upload the file chunks
	// async function uploadChunks(chunks: Blob[], url: string) {
	//   for (let i = 0; i < chunks.length; i++) {
	//     const chunk = chunks[i];
	//     const formData = new FormData();
	//     formData.append('chunk', chunk);
	//     formData.append('index', i.toString());
	//     formData.append('total', chunks.length.toString());
	//     await axios.post(url, formData);
	//   }
	// }

	// // Function to handle the file input change event
	// async function handleFileChange(event: Event) {
	//   const input = event.target as HTMLInputElement;
	//   const file = input.files && input.files[0];
	//   if (file) {
	//     const chunks = chunkFile(file, 1024 * 1024); // Split the file into 1MB chunks
	//     const url = '/api/uploadimage'; // Replace with your server-side endpoint URL
	//     await uploadChunks(chunks, url); // Upload the file chunks
	//   }
	// }

	async function UpdateUserData() {
		fetchUpdateUserData(NewUserData);

		if (NewUserData.UserID != null) {
			const GetUserData = await fetchUserData(NewUserData.UserID);
			console.log('🚀 ~ file: +page.svelte:119 ~ UpdateUserData ~ GetUserData:', GetUserData);
			UserData.update(() => GetUserData);
		}
	}

	function DiscartChanges() {
		NewUserData = {
			UserID: $UserData.UserID,
			UserURL: $UserData.UserURL,
			UserName: $UserData['UserName'],
			Email: $UserData.Email,
			Password: '',
			UserTitle: $UserData['UserTitle'],
			UserImage: $UserData['UserImage'],
			Badges: $UserData.Badges,
			Follower: $UserData['Follower'],
			Following: $UserData['Following'],
			Location: $UserData['Location'],
			MembershipTime: $UserData['MembershipTime'],
			LastSeen: $UserData['LastSeen'],
			Aboutme: $UserData['Aboutme'],
			Mysite: $UserData['Mysite'],
			TopTags: $UserData['TopTags'],
			SelectedPanel: 'Profile',
			AccountType: $UserData.AccountType,
			Github: $UserData['Github'],
			Twitter: $UserData['Twitter'],
			Linkedin: $UserData['Linkedin']
		};
	}

	let loadingState: boolean = false;
	onMount(async () => {
		const CookieValueInfo1: string = getCookieValue('Info1');

		const InfoCookieData = JSON.parse(atob(CookieValueInfo1)) as CookieInfo1Type;
		console.log('🚀 ~ file: +page.ts:24 ~ load ~ InfoCookieData:', InfoCookieData);
		let UserDataValue = {} as UserDataType;
		UserData.subscribe((value) => {
			UserDataValue = value;
		});

		const GetUserData = await fetchUserData(InfoCookieData.UserID);
		console.log('🚀 ~ file: +page.ts:24 ~ InitializeData ~ GetUserData:', GetUserData);
		UserData.update(() => GetUserData);

		loadingState = true;
	});

	import LoadingSVG from '$lib/Loading/index.svelte';
	// const HeadLogo = new URL("../../../lib/icons/gPodcast.svg", import.meta.url).href;
	import { afterUpdate } from 'svelte';
	// import Footer from '$lib/Footer/footer.svelte';

	afterUpdate(() => {});
</script>

{#if loadingState}
	<div
		class="   flex w-full flex-col justify-center overflow-x-hidden overflow-y-hidden bg-[#181818]"
	>
		<Navbar />
		<div class="flex w-full flex-row justify-center">
			<!-- COLLECTIVE -->
			<Collectives />
			<!-- QUESTION LIST -->

			<div class=" mt-2 flex max-h-full min-h-screen w-[1100px] flex-col bg-[#2d2d2d] pb-44">
				<div class="mt-2 flex flex-row items-center gap-5">
					<img
						src={NewUserData.UserImage}
						alt=""
						class=" aspect-square active:ring-offset-base-50 m-5 h-32 w-32 cursor-pointer rounded-xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500 active:rounded-md active:ring active:ring-blue-600"
					/>
					<input
						on:change={handleFileChange}
						type="file"
						accept=".jpg, .jpeg, .png, .svg .webp"
						class=" to-gay-700 founded-full h-20
					w-96 cursor-pointer rounded-2xl
					bg-gradient-to-br from-gray-600 text-white/80
					shadow-xl shadow-gray-700/60 file:m-5
					 file:cursor-pointer file:rounded-full file:border-none file:bg-gradient-to-b file:from-blue-500 file:to-blue-600 file:px-6
					 file:py-3
					 file:text-white file:shadow-lg file:shadow-blue-600/50
				"
					/>
				</div>
				<div class="">
					<div class=" flex flex-row">
						<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300">
							UserURL :
						</div>
						{#if ErrorMsg.UserURL[0]}
							<!-- content here -->
							<div class=" ml-3 mt-2 inline-flex">
								<svg
									class="h-6 w-6 fill-red-500"
									fill="currentColor"
									viewBox="0 0 20 20"
									xmlns="http://www.w3.org/2000/svg"
									><path
										fill-rule="evenodd"
										d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
										clip-rule="evenodd"
									/></svg
								>
								<p class="text-red-300">{ErrorMsg.UserURL[1]}</p>
							</div>
						{/if}
					</div>
					<input
						bind:value={NewUserData.UserURL}
						type="text"
						class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500 active:border-gray-800"
					/>
				</div>
				<div class="">
					<div class=" flex flex-row">
						<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300">
							Username :
						</div>
						{#if ErrorMsg.UserName[0]}
							<!-- content here -->
							<div class=" ml-3 mt-2 inline-flex">
								<svg
									class="h-6 w-6 fill-red-500"
									fill="currentColor"
									viewBox="0 0 20 20"
									xmlns="http://www.w3.org/2000/svg"
									><path
										fill-rule="evenodd"
										d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
										clip-rule="evenodd"
									/></svg
								>
								<p class="text-red-300">{ErrorMsg.UserName[1]}</p>
							</div>
						{/if}
					</div>
					<input
						bind:value={NewUserData.UserName}
						type="text"
						class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500 active:border-gray-800"
					/>
				</div>
				<div class="">
					<div class=" flex flex-row">
						<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300">
							Title :
						</div>
						{#if ErrorMsg.UserTitle[0]}
							<!-- content here -->
							<div class=" ml-3 mt-2 inline-flex">
								<svg
									class="h-6 w-6 fill-red-500"
									fill="currentColor"
									viewBox="0 0 20 20"
									xmlns="http://www.w3.org/2000/svg"
									><path
										fill-rule="evenodd"
										d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
										clip-rule="evenodd"
									/></svg
								>
								<p class="text-red-300">{ErrorMsg.UserTitle[1]}</p>
							</div>
						{/if}
					</div>
					<input
						bind:value={NewUserData.UserTitle}
						type="text"
						class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500 active:border-gray-800"
					/>
				</div>
				<!-- <div class="relative">
				<select class="block appearance-none w-full bg-white border border-gray-400 hover:border-gray-500 px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none focus:shadow-outline">
				  <option>Location 1</option>
				  <option>Location 2</option>
				  <option>Location 3</option>
				</select>
				<div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
				  <svg class="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M10 12a2 2 0 100-4 2 2 0 000 4z"/></svg>
				</div>
			  </div> -->
				<!-- <div class="relative">
				<input id="location-input" type="text" class="block w-full py-2 pr-10 placeholder-gray-400 text-gray-700 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" placeholder="Search location...">
				<div class="absolute inset-y-0 right-0 flex items-center pr-3">
				  <svg class="h-4 w-4 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
					<path fill-rule="evenodd" d="M13.853 14.146a.5.5 0 01-.707 0l-3.98-3.98a4.5 4.5 0 115.657 0l.707.707a.5.5 0 010 .707l-3.98 3.98zM6.5 9a2.5 2.5 0 114.03 1.908l.227.318.318.227A2.5 2.5 0 016.5 9z" clip-rule="evenodd" />
				  </svg>
				</div>
			  </div>
			  <div id="location-results" class="  z-10 mt-1 w-full bg-white rounded-md shadow-lg">
				<ul id="location-list" class="py-2 text-sm">
					1Search results will be dynamically generated here
				  </ul><ul id="location-list" class="py-2 text-sm">
					2Search results will be dynamically generated here
				  </ul><ul id="location-list" class="py-2 text-sm">
					3Search results will be dynamically generated here
				  </ul><ul id="location-list" class="py-2 text-sm">
					4Search results will be dynamically generated here
				  </ul>
			  </div> -->

				<div class="">
					<div class=" flex flex-row">
						<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300">
							Location :
						</div>
						{#if ErrorMsg.Location[0]}
							<!-- content here -->
							<div class=" ml-3 mt-2 inline-flex">
								<svg
									class="h-6 w-6 fill-red-500"
									fill="currentColor"
									viewBox="0 0 20 20"
									xmlns="http://www.w3.org/2000/svg"
									><path
										fill-rule="evenodd"
										d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
										clip-rule="evenodd"
									/></svg
								>
								<p class="text-red-300">{ErrorMsg.Location[1]}</p>
							</div>
						{/if}
					</div>
					<input
						bind:value={NewUserData.Location}
						type="text"
						class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500 active:border-gray-800"
					/>
				</div>
				<div class="">
					<div class=" flex flex-row">
						<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300">
							About Me :
						</div>
						{#if ErrorMsg.Aboutme[0]}
							<!-- content here -->
							<div class=" ml-3 mt-2 inline-flex">
								<svg
									class="h-6 w-6 fill-red-500"
									fill="currentColor"
									viewBox="0 0 20 20"
									xmlns="http://www.w3.org/2000/svg"
									><path
										fill-rule="evenodd"
										d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
										clip-rule="evenodd"
									/></svg
								>
								<p class="text-red-300">{ErrorMsg.Aboutme[1]}</p>
							</div>
						{/if}
					</div>
					<Wyswyg bind:markdown={NewUserData.Aboutme} />
				</div>
				<div class="">
					<div class=" flex flex-row">
						<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300">
							Github :
						</div>
						{#if ErrorMsg.Github[0]}
							<!-- content here -->
							<div class=" ml-3 mt-2 inline-flex">
								<svg
									class="h-6 w-6 fill-red-500"
									fill="currentColor"
									viewBox="0 0 20 20"
									xmlns="http://www.w3.org/2000/svg"
									><path
										fill-rule="evenodd"
										d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
										clip-rule="evenodd"
									/></svg
								>
								<p class="text-red-300">{ErrorMsg.Github[1]}</p>
							</div>
						{/if}
					</div>
					<input
						bind:value={NewUserData.Github}
						type="text"
						class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500 active:border-gray-800"
					/>
				</div>
				<div class="">
					<div class=" flex flex-row">
						<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300">
							Twitter :
						</div>
						{#if ErrorMsg.Twitter[0]}
							<!-- content here -->
							<div class=" ml-3 mt-2 inline-flex">
								<svg
									class="h-6 w-6 fill-red-500"
									fill="currentColor"
									viewBox="0 0 20 20"
									xmlns="http://www.w3.org/2000/svg"
									><path
										fill-rule="evenodd"
										d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
										clip-rule="evenodd"
									/></svg
								>
								<p class="text-red-300">{ErrorMsg.Twitter[1]}</p>
							</div>
						{/if}
					</div>
					<input
						bind:value={NewUserData.Twitter}
						type="text"
						class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500 active:border-gray-800"
					/>
				</div>
				<div class="">
					<div class=" flex flex-row">
						<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300">
							Linkedin :
						</div>
						{#if ErrorMsg.Linkedin[0]}
							<!-- content here -->
							<div class=" ml-3 mt-2 inline-flex">
								<svg
									class="h-6 w-6 fill-red-500"
									fill="currentColor"
									viewBox="0 0 20 20"
									xmlns="http://www.w3.org/2000/svg"
									><path
										fill-rule="evenodd"
										d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
										clip-rule="evenodd"
									/></svg
								>
								<p class="text-red-300">{ErrorMsg.Linkedin[1]}</p>
							</div>
						{/if}
					</div>
					<input
						bind:value={NewUserData.Linkedin}
						type="text"
						class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500 active:border-gray-800"
					/>
				</div>
				<div class="">
					<div class=" flex flex-row">
						<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300">
							My WebSite :
						</div>
						{#if ErrorMsg.Mysite[0]}
							<!-- content here -->
							<div class=" ml-3 mt-2 inline-flex">
								<svg
									class="h-6 w-6 fill-red-500"
									fill="currentColor"
									viewBox="0 0 20 20"
									xmlns="http://www.w3.org/2000/svg"
									><path
										fill-rule="evenodd"
										d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
										clip-rule="evenodd"
									/></svg
								>
								<p class="text-red-300">{ErrorMsg.Mysite[1]}</p>
							</div>
						{/if}
					</div>
					<input
						bind:value={NewUserData.Mysite}
						type="text"
						class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500 active:border-gray-800"
					/>
				</div>

				<div class=" flex flex-row justify-center gap-5">
					<button
						on:click={UpdateUserData}
						class=" m-5 flex h-12 w-fit items-center justify-center rounded-md bg-blue-500 px-3 hover:bg-blue-600 active:bg-blue-800"
					>
						<p class="my-auto text-xl font-semibold text-gray-200">Save Info</p>
					</button>
					<button
						on:click={DiscartChanges}
						class=" m-5 flex h-12 w-fit items-center justify-center rounded-md border-2 border-red-600 bg-inherit px-3 text-red-600 hover:bg-red-500 hover:text-gray-200 active:bg-red-600"
					>
						<p class="my-auto text-xl font-semibold">Discart</p>
					</button>
				</div>
			</div>
		</div>
		<Footer />
	</div>
{:else}
	<LoadingSVG />
{/if}

<style>
</style>
