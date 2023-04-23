
<script lang="ts">
import { goto } from "$app/navigation";
	import Amazon from "$lib/icons/Amazon.svelte";
	import Github from "$lib/icons/Github.svelte";
	import Google from "$lib/icons/google.svelte";
	import Twitter from "$lib/icons/Twitter.svelte";


	let LoginData = {
		Email: '',
		Password: ''
	};
	let b: boolean = false;
	let ErrorMsg = {
		
		Email: [false, 'invalid/wrong email'],
		Password: [false, 'Password is too short']
	};
	// let SuccessMsg:string ="Success"
	const ValidateEmail = (E: string) => {
		return String(E)
			.toLowerCase()
			.match(
				/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
			);
	};
	async function OnSubmit() {
		if (LoginData.Email.trim() === '') {
			ErrorMsg.Email[1] = 'Email is Empty';
			ErrorMsg.Email[0] = true;
		} else if (!ValidateEmail(LoginData.Email)) {
			ErrorMsg.Email[1] = 'invalid/wrong email or password';
			ErrorMsg.Email[0] = true;
		} 
		if (LoginData.Password.trim() === '') {
			ErrorMsg.Password[1] = 'Password is Empty';
			ErrorMsg.Password[0]=true
		} else if (LoginData.Password.length<8 || LoginData.Password.length>30){
			ErrorMsg.Password[1] = 'Password is Too short';
			ErrorMsg.Password[0]=true
		}
		 else {
			// await fetch("http://localhost:8080/login", {
			let res =await fetch('/api/login', {
				credentials: 'same-origin',
				method: 'POST',
				mode: 'cors',
				body: JSON.stringify(LoginData)
			})
			
			if (res.ok){
				let value= await res.json()
                console.log("🚀 ~ file: index.svelte ~ line 49 ~ OnSubmit ~ value : ", value)
				LoginData = {
					Email: '',
					Password: ''
				};
				// let s:string="/"+String(value.ID)
				goto("/");
				// goto();
			}else{
				console.log("🚀 ~ file: index.svelte ~ line 57 ~ OnSubmit ~ err : ")
				ErrorMsg.Email[1] = 'invalid/wrong email or password';
				ErrorMsg.Email[0] = true;
			}
			
	
		}
	}
</script>

<div class="grid h-screen w-full  place-items-center bg-[#181818] ">
	<a
		href="/register"
		on:click={() => {
			LoginData['Email'] = '';
			LoginData['Password'] = '';
		}}
		class="absolute  top-0 right-0  m-10"
	>
		<button
			class="  h-12 w-60 rounded-3xl bg-sky-600 text-xl font-bold text-gray-300 hover:bg-sky-600 active:bg-sky-500"
		>
			Register
		</button>
	</a>
	<div class="">
		<div class=" flex h-[420px] w-[500px] flex-col space-y-4 rounded-xl bg-[#2d2d2d]">
			<div
				class="mt-5 flex w-full items-center justify-center text-3xl font-semibold text-gray-200"
			>
				Sign In
			</div>
			<div class=" flex flex-row">
				<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					Email :
				</div>
				{#if ErrorMsg.Email[0] }
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
				bind:value={LoginData['Email']}
				type="text"
				class=" mx-4 mt-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<div class="flex flex-row">
				<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					Password :
				</div>
				{#if ErrorMsg.Password[0] }
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
						<p class="text-red-300">{ErrorMsg.Password[1]}</p>
					</div>
				{/if}
			</div>
			<input
				bind:value={LoginData['Password']}
				type="password"
				class=" mx-4 mt-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<div class="mt-8 ml-4 text-center text-xs font-semibold text-gray-500">
				By registering, you agree to Accord's <p
					class="inline cursor-pointer text-sm font-bold text-sky-600 hover:text-sky-300 active:text-sky-500"
				>
					Terms of Service
				</p>
				and
				<p
					class="inline cursor-pointer text-sm font-bold text-sky-600 hover:text-sky-300 active:text-sky-500"
				>
					privacy Policy
				</p>
			</div>

			<button
				on:click={OnSubmit}
				class=" m-4 h-12 w-60  self-center rounded-xl bg-sky-700 text-xl font-bold text-gray-300 hover:bg-sky-600 active:bg-sky-500"
			>
				Login
			</button>
		</div>
		<div class=" flex flex-row gap-1 pt-2  ">
			<Google />
			<!-- <svg class="h-12 w-12" xmlns="http://www.w3.org/2000/svg" width="12" height="12" focusable="false" viewBox="0 0 12 12"> <path fill="currentColor" d="M6 0a6 6 0 110 12A6 6 0 016 0zm0 .98C3.243.98 1 3.223 1 6a5.02 5.02 0 003.437 4.77.594.594 0 00.045.005c.203.01.279-.129.279-.25l-.007-.854c-1.39.303-1.684-.674-1.684-.674-.227-.58-.555-.734-.555-.734-.454-.312.034-.306.034-.306.365.026.604.288.708.43l.058.088c.446.767 1.17.546 1.455.418.046-.325.174-.546.317-.672-1.11-.127-2.277-.558-2.277-2.482 0-.548.195-.996.515-1.348l-.03-.085c-.064-.203-.152-.658.079-1.244l.04-.007c.124-.016.548-.013 1.335.522A4.77 4.77 0 016 3.408c.425.002.853.058 1.252.17.955-.65 1.374-.516 1.374-.516.272.692.1 1.202.05 1.33.32.35.513.799.513 1.347 0 1.93-1.169 2.354-2.283 2.478.18.155.34.462.34.93l-.006 1.378c0 .13.085.282.323.245A5.02 5.02 0 0011 6C11 3.223 8.757.98 6 .98z"/> </svg> -->
			<Github />
			
			<Amazon />
			<Twitter/>
		</div>
	</div>
</div>

<style>
	/* your styles go here */
	input:-webkit-autofill,
	input:-webkit-autofill:hover,
	input:-webkit-autofill:focus,
	textarea:-webkit-autofill,
	textarea:-webkit-autofill:hover,
	textarea:-webkit-autofill:focus,
	select:-webkit-autofill,
	select:-webkit-autofill:hover,
	select:-webkit-autofill:focus {
		border: 2px solid rgb(14 165 233 / var(--tw-border-opacity));
		-webkit-text-fill-color: #98999e;
		-webkit-box-shadow: 0 0 0px 1000px #303338 inset;
		transition: background-color 5000s ease-in-out 0s;
	}
</style>
