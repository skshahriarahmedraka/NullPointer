<script lang="ts">
	import CameraPlus from '$lib/icons/cameraPlus.svelte';

	// import { Following } from '$lib/Navbar/svgs/following.svelte';
	import type { HashDataType } from '$lib/store/types';
	import { tick } from 'svelte';
	import HashAbout from './component/HashAbout.svelte';
	import HashBlog from './component/HashBlog.svelte';
	import HashQues from './component/HashQues.svelte';
	import { marked } from 'marked';


	import BoldIcon from '$lib/icons/writtingIcon/boldIcon.svelte';
	import ItalicIcon from '$lib/icons/writtingIcon/ItalicIcon.svelte';
	import UnderlineIcon from '$lib/icons/writtingIcon/UnderlineIcon.svelte';
	import DeletedIcon from '$lib/icons/writtingIcon/DeletedIcon.svelte';
	import CodeIcon from '$lib/icons/writtingIcon/CodeIcon.svelte';
	import SuperscriptIcon from '$lib/icons/writtingIcon/SuperscriptIcon.svelte';
	import SubscriptIcon from '$lib/icons/writtingIcon/SubscriptIcon.svelte';
	import OrderedListIcon from '$lib/icons/writtingIcon/OrderedListIcon.svelte';
	import UnorderedList from '$lib/icons/writtingIcon/UnorderedList.svelte';
	import UndoIcon from '$lib/icons/writtingIcon/UndoIcon.svelte';
	import RedoIcon from '$lib/icons/writtingIcon/RedoIcon.svelte';
	import LinkIcon from '$lib/icons/writtingIcon/LinkIcon.svelte';
	import ImageIcon from '$lib/icons/writtingIcon/ImageIcon.svelte';
	import TableIcon from '$lib/icons/writtingIcon/TableIcon.svelte';
	import AlignLeft from '$lib/icons/writtingIcon/AlignLeft.svelte';
	import AlignRight from '$lib/icons/writtingIcon/AlignRight.svelte';
	import AlignCenter from '$lib/icons/writtingIcon/AlignCenter.svelte';
	import { fade } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import { fetchCreateHash } from '$lib/store/fetch';

	// your script goes here
	let ShowLinkedText: boolean = false;
	let ShowLinkedImage: boolean = false;
	let CashedLinkImage: string = 'https://';

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
		handleBannerSubmit();
	};
	const handleBannerSubmit = async () => {
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
			HashData.BannerImage = ResObj.ImageUrl;
			console.log("🚀 ~ file: HashWrite.svelte:69 ~ handleBannerSubmit ~ HashData.BannerImage:", HashData.BannerImage)

		} catch (error) {
			console.error(error);
		}
	};

    const OnImageSelected = (event: Event) => {
		const inputElement = event.target as HTMLInputElement;
		file = inputElement.files?.[0] ?? null;
		handleImageSubmit();
	};
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
			HashData.Image = ResObj.ImageUrl;
			console.log("🚀 ~ file: HashWrite.svelte:93 ~ handleImageSubmit ~ HashData.Image:", HashData.Image)

			// handle server response
		} catch (error) {
			console.error(error);
		}
	};

	const OnAboutDataImageSelected = (event: Event) => {
		const inputElement = event.target as HTMLInputElement;
		file = inputElement.files?.[0] ?? null;
		AboutImageSubmit();
	};

	const AboutImageSubmit = async () => {
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
	let ShowTableInfo = false;
	async function OnclickTable() {
		ShowTableInfo = !ShowTableInfo;
		if (ShowTableInfo) {
			ShowLinkedImage = false;
			ShowLinkedText = false;
		}
	}

	let ErrorMsg = {
		HashName: [false, 'Hash  Name is not Valid'],
		HashMetaTitle: [false, 'Hash Title is not Valid'],
		HashAbout: [false, 'Hash About is not Valid'],
		
	};
	function ValidHashTitle(s:string): boolean{
		if (s.length < 5) {
			return false;
		} else {
			return true
		}
	}
	function ValidHashDescription(s:string): boolean{
		if (s.length < 15) {
			return false;
		} else {
			return true
		}
	}
	async function OnSubmitHash() {
		HashData.ID=""
		HashData.NumOfFollower=0
		HashData.NumOfQuestion=0
		HashData.NumOfBlog=0
		HashData.NumOfAnswer=0
		HashData.About=markdown
		
		
		if (HashData.Name.trim() === '') {
			ErrorMsg.HashName[1] = 'Title is empty';
			ErrorMsg.HashName[0] = true;
		} else if (!ValidHashTitle(HashData.MetaTitle)) {
			ErrorMsg.HashMetaTitle[1] = 'MetaTitle is too small';
			ErrorMsg.HashMetaTitle[0] = true;
		} else {
			ErrorMsg.HashMetaTitle[0] = false;
		}
		
		
		 if (!ValidHashDescription(HashData.About)) {
			ErrorMsg.HashAbout[1] = 'About is too short ';
			ErrorMsg.HashAbout[0] = true;
		} else {
			ErrorMsg.HashAbout[0] = false;
		}
		
		console.log("🚀 ~ file: HashWrite.svelte:209 ~ OnSubmitHash ~ ErrorMsg.HashAbout[0]:", ErrorMsg.HashAbout[0])
		console.log("🚀 ~ file: HashWrite.svelte:210 ~ OnSubmitHash ~ ErrorMsg.HashMetaTitle[0]:", ErrorMsg.HashMetaTitle[0])
		console.log("🚀 ~ file: HashWrite.svelte:210 ~ OnSubmitHash ~ ErrorMsg.HashName[0]:", ErrorMsg.HashName[0])
		if (!ErrorMsg.HashAbout[0] && !ErrorMsg.HashName[0] && !ErrorMsg.HashMetaTitle[0]) {
			

			const res:{"InsertedID":string} = await fetchCreateHash(HashData)
			console.log("🚀 ~ file: HashWrite.svelte:212 ~ OnSubmitHash ~ res:", res)
			if (res.InsertedID!==''){
				goto('/h/'+res.InsertedID)

			}
			// await fetch('/api/register', {
			// 	// credentials: 'same-origin',
			// 	method: 'POST',
			// 	// mode: 'cors',
			// 	body: JSON.stringify(RegisterData)
			// })
			// 	.then((response) => response.json())
			// 	.then((value) => {
			// 		console.log("🚀 ~ file: index.svelte ~ line 71 ~ .then ~ value : ", value)
			// 		// console.log("🚀 ~ file: index.svelte ~ line 71 ~ .then ~ value.ID : ", value.ID)
			// 		ErrorMsg.UserName[1] = 'Success';
			// 		ErrorMsg.UserName[0] = true;
			// 		RegisterData = {
			// 			UserName: '',
			// 			Email: '',
			// 			Password: ''
			// 		};
		
            //         // console.log("🚀 ~ file: index.svelte ~ line 73 ~ .then ~ s : ", s)
            //         // console.log("🚀 ~ file: index.svelte ~ line 73 ~ .then ~ s typeof : ", typeof(s))
			// 		goto("/");
			// 	});
		}
	}
	async function OnResetHash(){
		HashData= {
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
	}

	
	
	}
	async function OnCancelHash() {
		OnResetHash()
		goto('/');
	}


	let tArea: any = null;
	let markdown: string = '';
	let CashedLink: string = 'https://';

	async function HandleSelectedStr(
		e: MouseEvent & { currentTarget: EventTarget & HTMLButtonElement },
		s: string
	) {
		// if (e.key !="Tab") return ;

		e.preventDefault();
		let { selectionStart, selectionEnd, value } = tArea;
		if (selectionStart === selectionEnd) {
			if (!(s === 'BreakLine')) {
				tArea.focus();
				return;
			}
		}
		let selection: string = value.slice(selectionStart, selectionEnd);
		let replacement: string = '';

		switch (s) {
			case 'bold': {
				if (
					value.slice(selectionStart - 2, selectionStart) === '**' &&
					value.slice(selectionEnd, selectionEnd + 2) === '**'
				) {
					// replacement= value.slice(0,selectionStart-2)+selection
					markdown = value.slice(0, selectionStart - 2) + selection + value.slice(selectionEnd + 2);
					await tick();
					tArea.selectionStart = selectionStart - 2;
					tArea.selectionEnd = selectionEnd - 2;
					// tArea.disable=false
					tArea.focus();
				} else {
					replacement = '**' + selection + '**';
					markdown = value.slice(0, selectionStart) + replacement + value.slice(selectionEnd);
					await tick();
					tArea.selectionStart = selectionStart + 2;
					tArea.selectionEnd = selectionEnd + 2;
					// tArea.disable=false
					tArea.focus();
				}
				break;
			}
			case 'italic': {
				if (
					value.slice(selectionStart - 1, selectionStart) === '*' &&
					value.slice(selectionEnd, selectionEnd + 1) === '*'
				) {
					// replacement= value.slice(0,selectionStart-2)+selection
					markdown = value.slice(0, selectionStart - 1) + selection + value.slice(selectionEnd + 1);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart - 1;
					tArea.selectionEnd = selectionEnd - 1;
				} else {
					replacement = '*' + selection + '*';
					markdown = value.slice(0, selectionStart) + replacement + value.slice(selectionEnd);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart + 1;
					tArea.selectionEnd = selectionEnd + 1;
				}
				break;
			}
			case 'underline': {
				if (
					value.slice(selectionStart - 3, selectionStart) === '<u>' &&
					value.slice(selectionEnd, selectionEnd + 4) === '</u>'
				) {
					// replacement= value.slice(0,selectionStart-2)+selection
					markdown = value.slice(0, selectionStart - 3) + selection + value.slice(selectionEnd + 4);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart - 3;
					tArea.selectionEnd = selectionEnd - 3;
				} else {
					replacement = '<u>' + selection + '</u>';
					markdown = value.slice(0, selectionStart) + replacement + value.slice(selectionEnd);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart + 3;
					tArea.selectionEnd = selectionEnd + 3;
				}
				break;
			}
			case 'StrikeOut': {
				if (
					value.slice(selectionStart - 3, selectionStart) === '<s>' &&
					value.slice(selectionEnd, selectionEnd + 4) === '</s>'
				) {
					// replacement= value.slice(0,selectionStart-2)+selection
					markdown = value.slice(0, selectionStart - 3) + selection + value.slice(selectionEnd + 4);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart - 3;
					tArea.selectionEnd = selectionEnd - 3;
				} else {
					replacement = '<s>' + selection + '</s>';
					markdown = value.slice(0, selectionStart) + replacement + value.slice(selectionEnd);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart + 3;
					tArea.selectionEnd = selectionEnd + 3;
				}
				break;
			}
			case 'SuperScript': {
				if (
					value.slice(selectionStart - 5, selectionStart) === '<sup>' &&
					value.slice(selectionEnd, selectionEnd + 6) === '</sup>'
				) {
					// replacement= value.slice(0,selectionStart-2)+selection
					markdown = value.slice(0, selectionStart - 5) + selection + value.slice(selectionEnd + 6);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart - 5;
					tArea.selectionEnd = selectionEnd - 5;
				} else {
					replacement = '<sup>' + selection + '</sup>';
					markdown = value.slice(0, selectionStart) + replacement + value.slice(selectionEnd);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart + 5;
					tArea.selectionEnd = selectionEnd + 5;
				}
				break;
			}
			case 'SubScript': {
				if (
					value.slice(selectionStart - 5, selectionStart) === '<sub>' &&
					value.slice(selectionEnd, selectionEnd + 6) === '</sub>'
				) {
					// replacement= value.slice(0,selectionStart-2)+selection
					markdown = value.slice(0, selectionStart - 5) + selection + value.slice(selectionEnd + 6);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart - 5;
					tArea.selectionEnd = selectionEnd - 5;
				} else {
					replacement = '<sub>' + selection + '</sub>';
					markdown = value.slice(0, selectionStart) + replacement + value.slice(selectionEnd);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart + 5;
					tArea.selectionEnd = selectionEnd + 5;
				}
				break;
			}
			case 'BreakLine': {
				// replacement = '<sup>' + selection + '</sup>';
				markdown = value.slice(0, selectionStart) + '<br/>' + value.slice(selectionEnd);
				await tick();
				tArea.focus();
				tArea.selectionStart = selectionStart;
				tArea.selectionEnd = selectionEnd;

				break;
			}
			case 'code': {
				markdown =
					value.slice(0, selectionStart) +
					'```\n' +
					selection +
					'\n```' +
					value.slice(selectionEnd);
				await tick();
				tArea.focus();
				tArea.selectionStart = selectionStart + 3;
				tArea.selectionEnd = selectionEnd + 3;
				break;
			}
			case 'OrderedList': {
				console.log('🚀 ~ file: index.svelte:252 ~ AddOrdered: OrderedList');
				let AddOrdered: boolean = false;
				//check ordered or not
				let i: number;
				let arr: string[] = selection.split('\n');

				let interate = arr.map((s) => {
					if (!(/[0-9]/.test(s[0]) && s.slice(1, 3) === '. ')) {
						AddOrdered = true;
					}
				});
				interate;

				if (AddOrdered) {
					// add order number
					let i: number;
					let count: number = 1;
					for (i = 0; true; i++) {
						if (i > selection.length) {
							break;
						}
						if (i === 0) {
							selection = `${count}. ` + selection;
							count += 1;
						} else if (selection[i] === '\n') {
							selection = selection.slice(0, i + 1) + `${count}. ` + selection.slice(i + 3);
							count += 1;
						}
					}
					// console.log("🚀 ~ file: index.svelte ~ line 260 ~ selection", selection)
					markdown = value.slice(0, selectionStart) + selection + value.slice(selectionEnd);

					// markdown=selection
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart;
					tArea.selectionEnd = selectionEnd + count + 1;
				} else {
					// remove Order
					let count: number = 0;
					let arr: string[] = selection.split('\n');

					let x: string = '';

					// for (count=0;count<arr.length;count++){
					// 	if (count===arr.length-1){
					// 		x+=s.slice(3)+
					// 	count+=1
					// 	}
					// }

					let interate = arr.map((s) => {
						if (count === arr.length - 1) {
							x += s.slice(3);
							count += 1;
						} else {
							x += s.slice(3) + '\n';
							count += 1;
						}
					});
					interate;
					markdown = value.slice(0, selectionStart) + x + value.slice(selectionEnd);

					// markdown=selection
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart;
					tArea.selectionEnd = selectionEnd - 3 * count;
				}
				break;
			}
			case 'UnorderedList': {
				let AddOrdered: boolean = true;
				//check ordered or not
				let i: number;
				let arr: string[] = selection.split('\n');

				let interate = arr.map((s) => {
					if (s.slice(0, 2) === '- ') {
						AddOrdered = false;
					}
				});
				interate;

				if (AddOrdered) {
					// add order number
					let i: number;
					let count: number = 1;
					for (i = 0; i <= selection.length; i++) {
						if (i === 0) {
							selection = `- ` + selection;
							count += 1;
						} else if (selection[i] === '\n') {
							selection = selection.slice(0, i + 1) + `- ` + selection.slice(i + 1);
							count += 1;
						}
					}
					// console.log("🚀 ~ file: index.svelte ~ line 260 ~ selection", selection)
					markdown = value.slice(0, selectionStart) + selection + value.slice(selectionEnd);

					// markdown=selection
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart;
					tArea.selectionEnd = selectionEnd + count + 1;
				} else {
					// remove Order
					let count: number = 0;
					let arr: string[] = selection.split('\n');

					let x: string = '';

					// for (count=0;count<arr.length;count++){
					// 	if (count===arr.length-1){
					// 		x+=s.slice(3)+
					// 	count+=1
					// 	}
					// }

					let interate = arr.map((s) => {
						if (count === arr.length - 1) {
							x += s.slice(2);
							count += 1;
						} else {
							x += s.slice(2) + '\n';
							count += 1;
						}
					});
					interate;
					markdown = value.slice(0, selectionStart) + x + value.slice(selectionEnd);

					// markdown=selection
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart;
					tArea.selectionEnd = selectionEnd - 2 * count;
				}
				break;
			}
			case 'Undo': {
				console.log('🚀 ~ Undo is called');

				tArea.focus();
				document.execCommand('undo');
				tArea.focus();

				break;
			}
			case 'Redo': {
				console.log('🚀 ~ redo is called');

				tArea.focus();
				document.execCommand('redo');
				tArea.focus();

				break;
			}
			case 'Link': {
				markdown =
					value.slice(0, selectionStart) +
					'[' +
					selection +
					']' +
					`(${CashedLink})` +
					value.slice(selectionEnd);
				await tick();
				tArea.focus();
				tArea.selectionStart = selectionStart + 1;
				tArea.selectionEnd = selectionEnd + 1;
				ShowLinkedText = !ShowLinkedText;
				CashedLink = 'https://';
				break;
			}
			case 'ImageLink': {
				markdown =
					value.slice(0, selectionStart) +
					'![' +
					selection +
					']' +
					`(${CashedLink})` +
					value.slice(selectionEnd);
				await tick();
				tArea.focus();
				tArea.selectionStart = selectionStart + 2;
				tArea.selectionEnd = selectionEnd + 2;
				ShowLinkedImage = !ShowLinkedImage;
				CashedLinkImage = 'https://';
				break;
			}

			// 				<div align="left">raka</div>
			// <div align="center">raka</div>
			// <div align="right">raka</div>
			case 'AlignLeft': {
				if (
					value.slice(selectionStart - 18, selectionStart) === '<div align="left">' &&
					value.slice(selectionEnd, selectionEnd + 6) === '</div>'
				) {
					// replacement= value.slice(0,selectionStart-2)+selection
					markdown =
						value.slice(0, selectionStart - 18) + selection + value.slice(selectionEnd + 6);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart - 18;
					tArea.selectionEnd = selectionEnd - 18;
				} else {
					replacement = '<div align="left">' + selection + '</div>';
					markdown = value.slice(0, selectionStart) + replacement + value.slice(selectionEnd);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart + 18;
					tArea.selectionEnd = selectionEnd + 18;
				}
				break;
			}
			case 'AlignCenter': {
				if (
					value.slice(selectionStart - 20, selectionStart) === '<div align="center">' &&
					value.slice(selectionEnd, selectionEnd + 6) === '</div>'
				) {
					// replacement= value.slice(0,selectionStart-2)+selection
					markdown =
						value.slice(0, selectionStart - 20) + selection + value.slice(selectionEnd + 6);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart - 20;
					tArea.selectionEnd = selectionEnd - 20;
				} else {
					replacement = '<div align="center">' + selection + '</div>';
					markdown = value.slice(0, selectionStart) + replacement + value.slice(selectionEnd);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart + 20;
					tArea.selectionEnd = selectionEnd + 20;
				}
				break;
			}
			case 'AlignRight': {
				if (
					value.slice(selectionStart - 19, selectionStart) === '<div align="right">' &&
					value.slice(selectionEnd, selectionEnd + 6) === '</div>'
				) {
					// replacement= value.slice(0,selectionStart-2)+selection
					markdown =
						value.slice(0, selectionStart - 19) + selection + value.slice(selectionEnd + 6);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart - 19;
					tArea.selectionEnd = selectionEnd - 19;
				} else {
					replacement = '<div align="right">' + selection + '</div>';
					markdown = value.slice(0, selectionStart) + replacement + value.slice(selectionEnd);
					await tick();
					tArea.focus();
					tArea.selectionStart = selectionStart + 19;
					tArea.selectionEnd = selectionEnd + 19;
				}
				break;
			}
		}
	}

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
		<!-- <CameraPlus class=" border-2 border-slate-400 rounded-xl  h-56 w-56  stroke-1 " /> -->
		
		<button class="  h-full w-full self-end" on:click={() => ThisBannerInput.click()}>
			<CameraPlus
				class=" h-full w-full rounded-xl  border-2 border-slate-200 stroke-slate-400 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
			/>
				<input
					bind:this={ThisBannerInput}
					on:change={(e) => OnBannerSelected(e)}
					type="file"
					class=" hidden h-full w-full "
					accept=".jpg, .jpeg, .png, .svg .webp"
				/>
			</button>
		{/if}
	</div>
	<div class="flex flex-row gap-2">
		<div class=" x mx-2 h-48 w-48 ">
			<!-- <img src={HashData.Image} alt="" class=" h-full w-full object-contain" /> -->
            {#if HashData.Image != ''}
			<!-- banner image -->
			<div class="h-full w-full">
				<svg
					on:click={() => {
						HashData.Image = '';
					}}
					on:keydown={() => {}}
					class="absolute left-5 top-5 h-8 w-8  stroke-white hover:cursor-pointer hover:fill-gray-500"
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
					bind:this={ThisImage}
					on:error={() => {
						ThisImage.src = HashData.Image;
					}}
					alt="ProfileImage"
					class=" aspect-[18/10] h-full w-full self-center rounded-lg object-cover"
				/>
			</div>
		{:else}
		<!-- <CameraPlus class=" border-2 border-slate-400 rounded-xl  h-56 w-56  stroke-1 " /> -->
		
		<button class="   h-full w-full self-end" on:click={() => ThisImageInput.click()}>
			<CameraPlus
				class=" h-full w-full rounded-xl  border-2 border-slate-200 stroke-slate-400 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
			/>
				<input
					bind:this={ThisImageInput}
					on:change={(e) => OnImageSelected(e)}
					type="file"
					class="  hidden h-full w-full"
					accept=".jpg, .jpeg, .png, .svg .webp"
				/>
			</button>
		{/if}

		</div>
		<div class=" flex h-48 w-full flex-col justify-evenly gap-2">
			<input
				bind:value={HashData.Name}
				type="text"
				class=" mx-2 mt-3 h-12 w-[95%] rounded-lg border-2 border-[#24262b] bg-[#303338] p-2 font-sf-pro text-3xl font-medium text-[#e7e9eb] outline-none focus:border-sky-500 active:border-gray-800"
			/>
			<!-- <p class="" >{}</p> -->
			<div class=" flex h-full w-full flex-row">
				<div class=" flex h-full w-full flex-col justify-evenly">
					<!-- <p class=" font-poppins text-slate-300">
						{HashData.MetaTitle}
					</p> -->
					<textarea
					bind:value={HashData.MetaTitle}
					
					class=" font-poppins text-slate-300 mt-3 h-full w-[95%] rounded-lg border-2 border-[#24262b] bg-[#303338] p-2  text-lg font-medium   outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
					<!-- <div class=" flex flex-row gap-4">
						<p class="font-raleway text-[#c4c8cc]">{RoundNum(HashData.NumOfQuestion)} Question</p>
						<p class="font-raleway text-[#c4c8cc]">{RoundNum(HashData.NumOfAnswer)} Answers</p>
						<p class="font-raleway text-[#c4c8cc]">{RoundNum(HashData.NumOfFollower)} Followers</p>
					</div> -->
				</div>
			</div>
		</div>
	</div>
	<hr class=" mx-2 bg-slate-400 opacity-30" />

	<div class=" flex h-20 w-full flex-row items-center justify-start gap-3">
		
		<button
			
			class=" ml-4 rounded-2xl px-3 py-1 font-poppins text-base border-2 border-sky-500  text-[#e7e9eb]">About</button
		>
		<!-- <button class=" px-3 py-1 font-poppins text-lg rounded-2xl { seletedTab==="questions" ? "bg-sky-500" : "border-2 border-sky-500" } text-[#e7e9eb]  ">Questions</button> -->
	</div>
	<hr class=" mx-2 bg-slate-400 opacity-30" />
	<!-- Write about the Hash -->
	<div class=" flex  flex-col justify-center items-center h-fit min-h-[50vh] w-full" >
		<div
			class=" mt-3 flex h-fit w-[95%] flex-row flex-wrap gap-2 self-center hover:cursor-pointer "
		>
			<!-- BOLD -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'bold')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500 active:bg-gray-400"
			>
				<BoldIcon />
			</button>
			<!-- ITALIC -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'italic')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<ItalicIcon />
			</button>
			<!-- UNDERLINE -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'underline')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<UnderlineIcon />
			</button>
			<!-- DELETED TEXT -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'StrikeOut')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<DeletedIcon />
			</button>
			<!-- code -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'code')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<CodeIcon />
			</button>
			<!-- breakline -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'BreakLine')}
				class="flex h-9 w-9  items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit p-1 text-lg text-slate-400 hover:bg-gray-500 hover:text-gray-300"
			>
				{'br/'}
			</button>
			<!-- SUPERSCRIPT -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'SuperScript')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<SuperscriptIcon />
			</button>
			<!-- SUBSCRIPT -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'SubScript')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<SubscriptIcon />
			</button>
			<!-- ORDERED LIST -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'OrderedList')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500 "
			>
				<OrderedListIcon />
			</button>
			<!-- UNORDERED LIST -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'UnorderedList')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<UnorderedList />
			</button>
			<!-- UNDO -->
			<button
				on:click={(e) => {
					HandleSelectedStr(e, 'Undo');
				}}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<UndoIcon />
			</button>
			<!-- REDO -->
			<button
				on:click={(e) => {
					HandleSelectedStr(e, 'Redo');
				}}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<RedoIcon />
			</button>
			<!-- LINK -->
			<button
				on:click={(e) => {
					ShowLinkedText = !ShowLinkedText;
					if (ShowLinkedText) {
						ShowLinkedImage = false;
					}
				}}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<LinkIcon />
			</button>
			<!-- image -->
			<button
				on:click={(e) => {
					HandleSelectedStr(e, 'image');
					ShowLinkedImage = !ShowLinkedImage;
					if (ShowLinkedImage) {
						ShowLinkedText = false;
					}
				}}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<ImageIcon />
			</button>
			<!-- unlink -->
			<!-- <button
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<svg
					class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
					viewBox="0 0 640 512"
					xmlns="http://www.w3.org/2000/svg"
					><path
						d="M485.1 354.9l113.5-113.5c55.21-55.21 55.21-144.7 0-199.9C570.1 13.8 534.8 0 498.6 0s-72.36 13.8-99.96 41.41l-43.36 43.36c15.11 8.012 29.47 17.58 41.91 30.02c3.146 3.146 5.898 6.518 8.742 9.838l37.96-37.96C458.5 72.05 477.1 64 498.6 64s40.1 8.047 54.71 22.66c14.61 14.61 22.66 34.04 22.66 54.71s-8.049 40.1-22.66 54.71l-119 119l-30.09-23.59c21.49-51.28 12.12-112.4-29.63-154.1C346.1 109.8 310.8 96 274.6 96c-29.6 0-58.93 9.752-83.83 28.23L38.81 5.109C34.41 1.672 29.19 0 24.03 0c-7.125 0-14.19 3.156-18.91 9.187c-8.188 10.44-6.375 25.53 4.062 33.7l591.1 463.1c10.5 8.203 25.56 6.328 33.69-4.078c8.188-10.44 6.375-25.53-4.062-33.7L485.1 354.9zM350.8 249.6L244.3 166.2C253.8 162.2 264 160 274.6 160c20.67 0 40.1 8.049 54.71 22.66c14.62 14.61 22.66 34.04 22.66 54.71C352 241.5 351.4 245.6 350.8 249.6zM234 387.4l-37.96 37.96C181.5 439.1 162 448 141.4 448c-20.67 0-40.1-8.047-54.71-22.66c-14.61-14.61-22.66-34.04-22.66-54.71s8.049-40.1 22.66-54.71l84.83-84.83L120.7 191.3L41.41 270.7c-55.21 55.21-55.21 144.7 0 199.9C69.01 498.2 105.2 512 141.4 512c36.18 0 72.36-13.8 99.96-41.41l43.36-43.36c-15.11-8.012-29.47-17.58-41.91-30.02C239.6 394.1 236.9 390.7 234 387.4zM265.4 374.6C293 402.2 329.2 416 365.4 416c11.98 0 23.84-2.082 35.51-5.111L224.6 272.7C223.9 309.5 237.3 346.5 265.4 374.6z"
					/></svg
				>
			</button> -->
			<!-- Table -->
			<button
				on:click={() => {
					OnclickTable();
				}}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<TableIcon />
			</button>
			<!-- ALIGN LEFT -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'AlignLeft')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<AlignLeft />
			</button>
			<!-- ALIGN RIGHT -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'AlignRight')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<AlignRight />
			</button>
			<!-- ALIGN CENTER -->
			<button
				on:click={(e) => HandleSelectedStr(e, 'AlignCenter')}
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<AlignCenter />
			</button>
			<!-- MOVE TO RIGHT -->
			<!-- <button
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="16"
					height="16"
					fill="currentColor"
					class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
					viewBox="0 0 16 16"
				>
					<path
						fill-rule="evenodd"
						d="M6 8a.5.5 0 0 0 .5.5h5.793l-2.147 2.146a.5.5 0 0 0 .708.708l3-3a.5.5 0 0 0 0-.708l-3-3a.5.5 0 0 0-.708.708L12.293 7.5H6.5A.5.5 0 0 0 6 8zm-2.5 7a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 1 0v13a.5.5 0 0 1-.5.5z"
					/>
				</svg>
			</button> -->
			<!-- MOVE TO LEFT  -->
			<!-- <button
				class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="16"
					height="16"
					fill="currentColor"
					class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
					viewBox="0 0 16 16"
				>
					<path
						fill-rule="evenodd"
						d="M12.5 15a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 1 0v13a.5.5 0 0 1-.5.5zM10 8a.5.5 0 0 1-.5.5H3.707l2.147 2.146a.5.5 0 0 1-.708.708l-3-3a.5.5 0 0 1 0-.708l3-3a.5.5 0 1 1 .708.708L3.707 7.5H9.5a.5.5 0 0 1 .5.5z"
					/>
				</svg>
			</button> -->
		</div>
		{#if ShowLinkedText}
			<div in:fade out:fade class=" mt-3 ml-6 flex  h-20 w-[95%] flex-col gap-2 text-gray-300  ">
				<div class="  w-full text-lg font-semibold ">Insert Hyperlink</div>
				<div class="flex-fow flex gap-2">
					<input
						bind:value={CashedLink}
						type="text"
						class="h-12 w-[70%] rounded-lg border-[1px] border-sky-300 bg-inherit px-4  "
					/>
					<button
						on:click={(e) => {
							HandleSelectedStr(e, 'Link');
						}}
						class="rounded-md border-2 border-sky-600 bg-sky-500 px-2 font-semibold text-gray-200 hover:bg-blue-500  "
						>Add Link</button
					>
					<button
						on:click={() => {
							ShowLinkedText = false;
							CashedLink = 'https://';
						}}
						class=" text-base text-red-500 hover:text-red-400 active:text-red-700">Cancel</button
					>
				</div>
			</div>
		{:else if ShowLinkedImage}
			<div class="mt-2 flex flex-row items-center gap-5 ">
				<img
					src={CashedLinkImage}
					alt=""
					class=" aspect-square active:ring-offset-base-50 m-5 h-32  w-32  cursor-pointer rounded-xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
				/>
				<input
					on:change={OnAboutDataImageSelected}
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
				<button
					on:click={(e) => {
						HandleSelectedStr(e, 'ImageLink');
					}}
					class="rounded-md border-2 border-sky-600 bg-sky-500 px-2 font-semibold text-gray-200 hover:bg-blue-500  "
					>Add Link</button
				>
				<button
					on:click={() => {
						ShowLinkedImage = false;
						CashedLinkImage = 'https://';
					}}
					class=" text-base text-red-500 hover:text-red-400 active:text-red-700">Cancel</button
				>
			</div>
		{:else if ShowTableInfo}
			<div class=" mx-4 flex  flex-col flex-wrap justify-between  text-[#e7e9eb]">
				<p class=" ">
					<b
						>Create tables using the <a
							href="https://github.github.com/gfm/#tables-extension-"
							class=" text-blue-600 underline">GitHub-flavored markdown format</a
						></b
					>
				</p>
				<pre class=" my-2 flex-wrap bg-transparent p-4">| A header | Another header |
| -------- | -------------- |
| First    | row            |
| Second   | row            |</pre>

				<p class="">
					A header row is required and must be followed by a separator row with the same number of
					cells. Cells are separated by a pipe (<code>|</code>) symbol.
				</p>

				<p class="">
					<b
						>Set the alignment of a table column by placing a <code>:</code> on the left, right, or both
						sides of a separator in the separator line</b
					>
				</p>
				<pre class=" flex-wrap bg-transparent p-4">| left | center | right |
|:---- |:------:| -----:|
| One  | Two    | Three |</pre>
			</div>
		{/if}

		<textarea
			bind:value={markdown}
			bind:this={tArea}
			id="tArea"
			class=" scrol3 b mt-3  h-fit min-h-[500px] w-[95%] self-center overflow-scroll overflow-x-hidden overflow-y-scroll  rounded-md border-2 border-gray-500 bg-inherit p-5 text-[#e7e9eb]   outline-0  focus:border-sky-500 active:outline-0"
			placeholder="Write Here"
		/>
		<!-- Render Question Description -->
		{#if markdown.trim() != ''}
			<div
				class=" prose mt-5 min-w-full max-w-full overflow-hidden break-words bg-inherit p-5 text-[#e7e9eb] prose-headings:text-[#e7e9eb] prose-p:text-[#e7e9eb] prose-a:text-blue-500 prose-blockquote:border-sky-400 prose-blockquote:text-[#e7e9eb] prose-figure:text-white prose-figcaption:text-[#e7e9eb]   prose-strong:text-[#e7e9eb] prose-em:text-[#e7e9eb] prose-code:text-[#e7e9eb] prose-pre:text-[#e7e9eb] prose-ol:text-[#e7e9eb] prose-ul:text-[#e7e9eb] prose-li:text-[#e7e9eb] prose-li:marker:text-white prose-table:text-[#e7e9eb] prose-thead:text-[#e7e9eb] prose-tr:border-4 prose-tr:border-gray-300 prose-tr:text-[#e7e9eb] prose-th:border-2 prose-th:border-gray-300 prose-th:text-[#e7e9eb] prose-td:border-2 prose-td:border-gray-300 prose-td:text-[#e7e9eb] prose-img:text-[#e7e9eb] prose-video:text-[#e7e9eb]  prose-hr:bg-gray-500 prose-hr:text-[#e7e9eb]  "
			>
				{@html marked(markdown)}
			</div>
			<!-- post button -->
			<!-- <button
				class=" m-5 flex h-12 w-fit  items-center justify-center rounded-md bg-blue-500 px-3 hover:bg-blue-600 active:bg-blue-800 "
			>
				<p class="my-auto text-xl font-semibold text-gray-200">Post Your Question</p>
			</button> -->
		{/if}
	</div>
	<div class=" flex flex-row  justify-center gap-5">
		<button
		on:click={() => {OnSubmitHash()}}
			class=" m-5 flex h-12 w-fit  items-center justify-center rounded-md bg-blue-500 px-3 hover:bg-blue-600 active:bg-blue-800 "
		>
			<p class="my-auto text-xl font-semibold text-gray-200">Create Hash</p>
		</button>
		<button
		on:click={() => {OnResetHash()}}
			class=" m-5 flex h-12 w-fit  items-center justify-center rounded-md bg-blue-500 px-3 hover:bg-blue-600 active:bg-blue-800 "
		>
			<p class="my-auto text-xl font-semibold text-gray-200">Reset</p>
		</button>
		<button
			on:click={() => {OnCancelHash()}}
			class=" m-5 flex h-12 w-fit items-center justify-center  rounded-md border-2 border-red-600 bg-inherit px-3 text-red-600 hover:bg-red-500 hover:text-gray-200 active:bg-red-600 "
		>
			<p class="my-auto text-xl font-semibold  ">Discart</p>
		</button>
	</div>
</div>

<style>
    /* your styles go here */
</style>