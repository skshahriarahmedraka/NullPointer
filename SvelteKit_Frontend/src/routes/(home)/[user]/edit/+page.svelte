<script lang="ts">
	import Wyswyg from '$lib/WriteQues/wyswyg.svelte';
	import { UserData } from '$lib/store/store';
	import type { UserDataType } from '$lib/store/types';

	// export let $UserData:any
	let NewUserData: UserDataType = {
		UserID: $UserData.UserID,
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
		Email: [false, 'not a valid Email address']
	};
</script>

<div class=" mt-2 flex max-h-full  min-h-screen w-[1100px] flex-col bg-[#2d2d2d] pb-44  ">
	<div class="mt-2 flex flex-row items-center gap-5 ">
		<img
			src={$UserData.UserImage}
			alt=""
			class=" aspect-square active:ring-offset-base-50 m-5 h-32  w-32  cursor-pointer rounded-xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
		/>
		<input
			type="file"
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
			<div class=" ml-8  mt-2 self-start text-left text-xl font-semibold text-gray-300 ">
				Username :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={NewUserData['UserName']}
			type="text"
			class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<div class="">
		<div class=" flex flex-row">
			<div class=" ml-8 mt-2 self-start text-left text-xl  font-semibold text-gray-300 ">
				Title :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={NewUserData['UserTitle']}
			type="text"
			class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<div class="">
		<div class=" flex flex-row">
			<div class=" ml-8 mt-2 self-start text-left  text-xl font-semibold text-gray-300 ">
				Location :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={NewUserData['Location']}
			type="text"
			class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<div class="">
		<div class=" flex flex-row">
			<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300 ">
				About Me :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<Wyswyg bind:markdown={NewUserData['Aboutme']} />
	</div>
	<div class="">
		<div class=" flex flex-row">
			<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300 ">
				Github :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={NewUserData['Github']}
			type="text"
			class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<div class="">
		<div class=" flex flex-row">
			<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300 ">
				Twitter :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={NewUserData['Twitter']}
			type="text"
			class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<div class="">
		<div class=" flex flex-row">
			<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300 ">
				Linkedin :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={NewUserData['Linkedin']}
			type="text"
			class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<div class="">
		<div class=" flex flex-row">
			<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300 ">
				My WebSite :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={NewUserData['Mysite']}
			type="text"
			class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>

	<div class="">
		<div class=" flex flex-row">
			<div class=" ml-8 mt-2 self-start text-left text-xl font-semibold text-gray-300 ">
				Username :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={NewUserData['UserName']}
			type="text"
			class=" ml-20 mt-2 h-12 w-[515px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<div class=" flex flex-row  justify-center gap-5">
		<button
			class=" m-5 flex h-12 w-fit  items-center justify-center rounded-md bg-blue-500 px-3 hover:bg-blue-600 active:bg-blue-800 "
		>
			<p class="my-auto text-xl font-semibold text-gray-200">Save Info</p>
		</button>
		<button
			class=" m-5 flex h-12 w-fit items-center justify-center  rounded-md border-2 border-red-600 bg-inherit px-3 text-red-600 hover:bg-red-500 hover:text-gray-200 active:bg-red-600 "
		>
			<p class="my-auto text-xl font-semibold  ">Discart</p>
		</button>
	</div>
</div>

<style>
</style>
