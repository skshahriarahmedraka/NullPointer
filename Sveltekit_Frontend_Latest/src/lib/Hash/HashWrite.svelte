<script lang="ts">
	import CameraPlus from '$lib/icons/cameraPlus.svelte';

	// import { Following } from '$lib/Navbar/svgs/following.svelte';
	import type { HashDataType } from '$lib/store/types';
	import HashAbout from './component/HashAbout.svelte';
	import HashBlog from './component/HashBlog.svelte';
	import HashQues from './component/HashQues.svelte';

	// your script goes here

	let seletedTab = 'questions';
	let FollowingStatus = false;
	let HashData: HashDataType = {
		ID: '',
		Name: '',
		MetaTitle: '',
		Image: '',
		BannerImage: '',
		NumOfFollower: 0,
		NumOfQuestion: 0,
		NumOfBlog: 0,
		NumOfAnswer: 0,
		About: ''
	};
	const HeadLogo = new URL('../../static/favicon.png', import.meta.url).href;
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

	// Image Upload
	let file: File | null = null;
	let ThisImageInput: HTMLInputElement;
	let ThisBannerInput: HTMLInputElement;
	let ThisImage: HTMLImageElement;
	let ThisBannerImage: HTMLImageElement;
	// let FileInput: HTMLInputElement;

	const OnBannerSelected = (event: Event) => {
		const inputElement = event.target as HTMLInputElement;
		file = inputElement.files?.[0] ?? null;
		handleSubmit();
	};
	let CashedLinkImage: string = 'https://';
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
			CashedLinkImage = ResObj.ImageUrl;
			console.log('🚀 ~ file: index.svelte:106 ~ handleSubmit ~ CashedLinkImage:', CashedLinkImage);

			// handle server response
		} catch (error) {
			console.error(error);
		}
	};

    const OnImageSelected = (event: Event) => {
		const inputElement = event.target as HTMLInputElement;
		file = inputElement.files?.[0] ?? null;
		handleImageSubmit();
	};
	let CashedLinkImage: string = 'https://';
	const handleImageSubmit = async () => {
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
			CashedLinkImage = ResObj.ImageUrl;
			console.log('🚀 ~ file: index.svelte:106 ~ handleSubmit ~ CashedLinkImage:', CashedLinkImage);

			// handle server response
		} catch (error) {
			console.error(error);
		}
	};
</script>

<!-- <div class=" mt-2 flex max-h-full min-h-screen w-[1400px] flex-col gap-5 bg-transparent"> -->

<svelte:head>
	<title>{'Add Hash'}</title>
	<link rel="icon" href={HeadLogo} />
</svelte:head>

<div class=" mb-2 mt-2 flex min-h-screen w-[1200px] flex-col gap-2 bg-[#2d2d2d]">
	<div class=" flex h-64 w-full items-center justify-center p-3">
		<!-- <img src="{HashData.BannerImage}" alt="" class=" h-full w-full object-cover"> -->
		{#if HashData.BannerImage != ''}
			<!-- banner image -->
			<div class="h-full w-full">
				<svg
					on:click={() => {
						HashData.BannerImage = '';
					}}
					on:keydown={() => {}}
					class="absolute left-5 top-5 h-8 w-8 stroke-white hover:cursor-pointer hover:fill-gray-500"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
					/></svg
				>
				<img
					src={HashData.BannerImage}
					bind:this={ThisBannerImage}
					on:error={() => {
						ThisBannerImage.src = HashData.BannerImage;
					}}
					alt="ProfileImage"
					class=" aspect-[18/10] h-full w-full self-center rounded-lg object-cover"
				/>
			</div>
		{:else}
			<CameraPlus
				class=" h-full w-full rounded-xl  border-2 border-slate-200 stroke-slate-400 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
			/>
			<!-- <CameraPlus class=" border-2 border-slate-400 rounded-xl  h-56 w-56  stroke-1 " /> -->

			<button class=" absolute h-full w-full self-end" on:click={() => ThisBannerInput.click()}>
				<input
					bind:this={ThisBannerInput}
					on:change={(e) => OnBannerSelected(e)}
					type="file"
					class=" hidden h-10 w-10"
					accept=".jpg, .jpeg, .png, .svg .webp"
				/>
			</button>
		{/if}
	</div>
	<div class="flex flex-row gap-2">
		<div class=" x mx-2 h-48 w-48">
			<!-- <img src={HashData.Image} alt="" class=" h-full w-full object-contain" /> -->
            {#if HashData.Image != ''}
			<!-- banner image -->
			<div class="h-full w-full">
				<svg
					on:click={() => {
						HashData.Image = '';
					}}
					on:keydown={() => {}}
					class="absolute left-5 top-5 h-8 w-8 stroke-white hover:cursor-pointer hover:fill-gray-500"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
					/></svg
				>
				<img
					src={HashData.Image}
					bind:this={ThisBannerImage}
					on:error={() => {
						ThisBannerImage.src = HashData.BannerImage;
					}}
					alt="ProfileImage"
					class=" aspect-[18/10] h-full w-full self-center rounded-lg object-cover"
				/>
			</div>
		{:else}
			<CameraPlus
				class=" h-full w-full rounded-xl  border-2 border-slate-200 stroke-slate-400 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
			/>
			<!-- <CameraPlus class=" border-2 border-slate-400 rounded-xl  h-56 w-56  stroke-1 " /> -->

			<button class=" absolute h-full w-full self-end" on:click={() => ThisImageInput.click()}>
				<input
					bind:this={ThisImageInput}
					on:change={(e) => OnImageSelected(e)}
					type="file"
					class=" hidden h-10 w-10"
					accept=".jpg, .jpeg, .png, .svg .webp"
				/>
			</button>
		{/if}

		</div>
		<div class=" flex h-48 w-full flex-col justify-evenly gap-2">
			<input
				bind:value={HashData.Name}
				type="text"
				class=" mx-4 mt-3 h-12 w-[95%] rounded-lg border-2 border-[#24262b] bg-[#303338] p-2 font-sf-pro text-3xl font-medium text-[#e7e9eb] outline-none focus:border-sky-500 active:border-gray-800"
			/>
			<!-- <p class="" >{}</p> -->
			<div class=" flex h-full w-full flex-row">
				<div class=" flex h-full w-full flex-col justify-evenly">
					<p class=" font-poppins text-slate-300">
						{HashData.MetaTitle}
					</p>
					<div class=" flex flex-row gap-4">
						<p class="font-raleway text-[#c4c8cc]">{RoundNum(HashData.NumOfQuestion)} Question</p>
						<p class="font-raleway text-[#c4c8cc]">{RoundNum(HashData.NumOfAnswer)} Answers</p>
						<p class="font-raleway text-[#c4c8cc]">{RoundNum(HashData.NumOfFollower)} Followers</p>
					</div>
				</div>
			</div>
		</div>
	</div>
	<hr class=" mx-2 bg-slate-400 opacity-30" />

	<div class=" flex h-20 w-full flex-row items-center justify-start gap-3">
		<button
			on:click={() => {
				seletedTab = 'questions';
			}}
			class=" rounded-2xl px-3 py-1 font-poppins text-base {seletedTab === 'questions'
				? 'bg-sky-500'
				: 'border-2 border-sky-500'} ml-3 text-[#e7e9eb]">Questions</button
		>
		<button
			on:click={() => {
				seletedTab = 'blogs';
			}}
			class=" rounded-2xl px-3 py-1 font-poppins text-base {seletedTab === 'blogs'
				? 'bg-sky-500'
				: 'border-2 border-sky-500'}  text-[#e7e9eb]">Blogs</button
		>
		<button
			on:click={() => {
				seletedTab = 'about';
			}}
			class=" rounded-2xl px-3 py-1 font-poppins text-base {seletedTab === 'about'
				? 'bg-sky-500'
				: 'border-2 border-sky-500'}  text-[#e7e9eb]">About</button
		>
		<!-- <button class=" px-3 py-1 font-poppins text-lg rounded-2xl { seletedTab==="questions" ? "bg-sky-500" : "border-2 border-sky-500" } text-[#e7e9eb]  ">Questions</button> -->
	</div>
	<hr class=" mx-2 bg-slate-400 opacity-30" />
	<div class=" h-fit min-h-[50vh] w-full" />
</div>

<style>
    /* your styles go here */
</style>