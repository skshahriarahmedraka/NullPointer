<script lang="ts">
	import { marked } from 'marked';
	import { fade, fly } from 'svelte/transition';
	import { tick } from 'svelte';
	import HotQuesCom from '$lib/HotQues/index.svelte';

	let ShowLinkedText: boolean = false;
	export let markdown: string = '';
	let QuestionTitle: string = '';
	// let obj;
	let TagObj: { Input: string; Focus: boolean; List: string[]; Suggested: string[] } = {
		Input: '' as string,
		Focus: false as boolean,
		List: [] as string[],
		Suggested: [] as string[]
	};
	// console.log('🚀 ~ file: index.svelte ~ line 27 ~ TagObj.Input : ', TagObj.Input);

	let GroupsObj: { Input: string; Focus: boolean; List: string[]; Suggested: string[] } = {
		Input: '' as string,
		Focus: false as boolean,
		List: [] as string[],
		Suggested: [] as string[]
	};

	// TAGS
	function onKeyPressTags(e: { keyCode: any }) {
		switch (e.keyCode) {
			case 32:
				onKeyTag();
				break;
			case 13:
				onKeyTag();
				break;
		}
	}
	function onKeyTag() {
		if (TagObj.Input.trim() != ('' as String)) {
			TagObj.List = [...TagObj.List, TagObj.Input.trim()];
			TagObj.Input = '' as string;
		}
	}

	// GROUP
	function onKeyPressGroup(e: { keyCode: any }) {
		switch (e.keyCode) {
			case 32:
				onKeyGroup();
				break;
			case 13:
				onKeyGroup();
				break;
		}
	}
	function onKeyGroup() {
		if (GroupsObj.Input.trim() != '') {
			GroupsObj.List = [...GroupsObj.List, GroupsObj.Input.trim()];
			GroupsObj.Input = '';
		}
	}
	let CashedLink: string = 'https://';
	let tArea: any = null;
	async function HandleSelectedStr(
		e: MouseEvent & { currentTarget: EventTarget & HTMLButtonElement },
		s: string
	) {
		// if (e.key !="Tab") return ;

		e.preventDefault();
		const { selectionStart, selectionEnd, value } = tArea;
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
			}
			case 'unrderedList': {
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
			}
			case 'link': {
				markdown =
					value.slice(0, selectionStart) +
					'[' +
					selection +
					']' +
					`(${CashedLink})` +
					value.slice(selectionEnd);
				await tick();
				tArea.focus();
				tArea.selectionStart = selectionStart + 3;
				tArea.selectionEnd = selectionEnd + 3;
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

<!-- <h1>Markdown Editor</h1> -->
<!-- <TextArea bind:value={markdown} minRows={4} maxRows={40} /> -->

<!-- Question -->
<div class=" mt-5 flex  max-h-full  w-9/12 flex-col  ">
	<!-- bold-italic buttons -->
	<div class=" mt-3 flex h-fit w-[95%] flex-row flex-wrap gap-2 self-center hover:cursor-pointer ">
		<!-- bold -->
		<!-- <button
				on:click={(e) => HandleSelectedStr(e, 'bold')}
				class="h-full w-12 rounded-lg bg-slate-500 text-white hover:bg-slate-200 focus:bg-stone-500"
				>Bold</button
			>
			<button
				on:click={(e) => HandleSelectedStr(e, 'italic')}
				class="h-full w-12 rounded-lg bg-slate-500 text-white hover:bg-slate-200 focus:bg-stone-500"
				>italic</button
			>
			<button
				on:click={(e) => HandleSelectedStr(e, 'blockquote')}
				class="h-full w-12 rounded-lg bg-slate-500 text-white hover:bg-slate-200 focus:bg-stone-500"
				>blockquote</button
			>
			<button
				on:click={(e) => HandleSelectedStr(e, 'codesample')}
				class="h-full w-12 rounded-lg bg-slate-500 text-white hover:bg-slate-200 focus:bg-stone-500"
				>code</button
			> -->
		<!-- BOLD -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'bold')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500 active:bg-gray-400"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2 hover:fill-gray-300 "
				viewBox="0 0 384 512"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M321.1 242.4C340.1 220.1 352 191.6 352 160c0-70.59-57.42-128-128-128L32 32.01c-17.67 0-32 14.31-32 32s14.33 32 32 32h16v320H32c-17.67 0-32 14.31-32 32s14.33 32 32 32h224c70.58 0 128-57.41 128-128C384 305.3 358.6 264.8 321.1 242.4zM112 96.01H224c35.3 0 64 28.72 64 64s-28.7 64-64 64H112V96.01zM256 416H112v-128H256c35.3 0 64 28.71 64 63.1S291.3 416 256 416z"
				/></svg
			>
		</button>
		<!-- ITALIC -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'italic')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
				viewBox="0 0 384 512"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M384 64.01c0 17.69-14.31 32-32 32h-58.67l-133.3 320H224c17.69 0 32 14.31 32 32s-14.31 32-32 32H32c-17.69 0-32-14.31-32-32s14.31-32 32-32h58.67l133.3-320H160c-17.69 0-32-14.31-32-32s14.31-32 32-32h192C369.7 32.01 384 46.33 384 64.01z"
				/></svg
			>
		</button>
		<!-- UNDERLINE -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'underline')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
				viewBox="0 0 448 512"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M416 448H32c-17.69 0-32 14.31-32 32s14.31 32 32 32h384c17.69 0 32-14.31 32-32S433.7 448 416 448zM48 64.01H64v160c0 88.22 71.78 159.1 160 159.1s160-71.78 160-159.1v-160h16c17.69 0 32-14.32 32-32s-14.31-31.1-32-31.1l-96-.0049c-17.69 0-32 14.32-32 32s14.31 32 32 32H320v160c0 52.94-43.06 95.1-96 95.1S128 276.1 128 224v-160h16c17.69 0 32-14.31 32-32s-14.31-32-32-32l-96 .0049c-17.69 0-32 14.31-32 31.1S30.31 64.01 48 64.01z"
				/></svg
			>
		</button>
		<!-- DELETED TEXT -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'StrikeOut')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
				viewBox="0 0 512 512"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M332.2 319.9c17.22 12.17 22.33 26.51 18.61 48.21c-3.031 17.59-10.88 29.34-24.72 36.99c-35.44 19.75-108.5 11.96-186-19.68c-16.34-6.686-35.03 1.156-41.72 17.53s1.188 35.05 17.53 41.71c31.75 12.93 95.69 35.37 157.6 35.37c29.62 0 58.81-5.156 83.72-18.96c30.81-17.09 50.44-45.46 56.72-82.11c3.998-23.27 2.168-42.58-3.488-59.05H332.2zM488 239.9l-176.5-.0309c-15.85-5.613-31.83-10.34-46.7-14.62c-85.47-24.62-110.9-39.05-103.7-81.33c2.5-14.53 10.16-25.96 22.72-34.03c20.47-13.15 64.06-23.84 155.4 .3438c17.09 4.531 34.59-5.654 39.13-22.74c4.531-17.09-5.656-34.59-22.75-39.12c-91.31-24.18-160.7-21.62-206.3 7.654C121.8 73.72 103.6 101.1 98.09 133.1C89.26 184.5 107.9 217.3 137.2 239.9L24 239.9c-13.25 0-24 10.75-24 23.1c0 13.25 10.75 23.1 24 23.1h464c13.25 0 24-10.75 24-23.1C512 250.7 501.3 239.9 488 239.9z"
				/></svg
			>
		</button>
		<!-- code -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'code')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-6 w-6 fill-slate-400"
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path
					fill-rule="evenodd"
					d="M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z"
					clip-rule="evenodd"
				/></svg
			>
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
			<svg
				class="h-9 w-9 fill-slate-400 p-2 hover:fill-gray-300"
				viewBox="0 0 512 512"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M480 160v-128c0-11.09-5.75-21.37-15.17-27.22C455.4-1.048 443.6-1.548 433.7 3.39l-32 16c-15.81 7.906-22.22 27.12-14.31 42.94C392.1 73.55 404.3 80.01 416 80.01v80c-17.67 0-32 14.31-32 32s14.33 32 32 32h64c17.67 0 32-14.31 32-32S497.7 160 480 160zM320 128c17.67 0 32-14.31 32-32s-14.33-32-32-32l-32-.0024c-10.44 0-20.23 5.101-26.22 13.66L176 200.2L90.22 77.67C84.23 69.11 74.44 64.01 64 64.01L32 64.01c-17.67 0-32 14.32-32 32s14.33 32 32 32h15.34L136.9 256l-89.6 128H32c-17.67 0-32 14.31-32 32s14.33 31.1 32 31.1l32-.0024c10.44 0 20.23-5.086 26.22-13.65L176 311.8l85.78 122.5C267.8 442.9 277.6 448 288 448l32 .0024c17.67 0 32-14.31 32-31.1s-14.33-32-32-32h-15.34l-89.6-128l89.6-127.1H320z"
				/></svg
			>
		</button>
		<!-- SUBSCRIPT -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'SubScript')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
				viewBox="0 0 512 512"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M480 448v-128c0-11.09-5.75-21.38-15.17-27.22c-9.422-5.875-21.25-6.344-31.14-1.406l-32 16c-15.81 7.906-22.22 27.12-14.31 42.94c5.609 11.22 16.89 17.69 28.62 17.69v80c-17.67 0-32 14.31-32 32s14.33 32 32 32h64c17.67 0 32-14.31 32-32S497.7 448 480 448zM320 128c17.67 0 32-14.31 32-32s-14.33-32-32-32l-32-.0024c-10.44 0-20.23 5.101-26.22 13.66L176 200.2L90.22 77.67C84.23 69.11 74.44 64.01 64 64.01L32 64.01c-17.67 0-32 14.32-32 32s14.33 32 32 32h15.34L136.9 256l-89.6 128H32c-17.67 0-32 14.31-32 32s14.33 31.1 32 31.1l32-.0024c10.44 0 20.23-5.086 26.22-13.65L176 311.8l85.78 122.5C267.8 442.9 277.6 448 288 448l32 .0024c17.67 0 32-14.31 32-31.1s-14.33-32-32-32h-15.34l-89.6-128l89.6-127.1H320z"
				/></svg
			>
		</button>
		<!-- ORDERED LIST -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'OrderedList')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2 hover:fill-gray-300 "
				viewBox="0 0 512 512"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M480 224H224C206.3 224 192 238.3 192 256s14.33 32 32 32h256c17.67 0 32-14.33 32-32S497.7 224 480 224zM224 128h256c17.67 0 32-14.33 32-32s-14.33-32-32-32H224C206.3 64 192 78.33 192 96S206.3 128 224 128zM480 384H224c-17.67 0-32 14.33-32 32s14.33 32 32 32h256c17.67 0 32-14.33 32-32S497.7 384 480 384zM40 224h80C133.3 224 144 213.3 144 199.1s-10.75-24-24-24h-16V55.99c0-8.594-4.594-16.53-12.06-20.81C84.5 30.91 75.28 30.95 67.91 35.25l-32 18.67C24.47 60.61 20.59 75.3 27.28 86.75C33.13 96.85 45.31 101.1 56 97.29v78.71h-16c-13.25 0-24 10.75-24 24S26.75 224 40 224zM136.3 432h-49.44l36.4-32.48c29.92-25.83 33.54-71.36 8.049-101.5C118.6 282.1 100.7 273.8 80.89 272.2C61.27 270.7 41.93 276.9 26.92 289.9L14.42 300.7c-10.04 8.656-11.17 23.81-2.527 33.86c8.672 10.06 23.8 11.14 33.78 2.516l12.51-10.79c5.24-4.531 12.1-6.672 18.96-6.156c6.926 .5469 13.2 3.734 17.62 8.969c8.703 10.28 7.486 25.26-3.057 34.36l-83.7 74.67c-7.424 6.625-9.982 17.16-6.457 26.47C5.119 473.8 14.01 480 23.96 480h112.4c13.23 0 23.96-10.75 23.96-23.1C160.3 442.8 149.5 432 136.3 432z"
				/></svg
			>
		</button>
		<!-- UNORDERED LIST -->
		<button
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2 hover:fill-gray-300 "
				viewBox="0 0 512 512"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M48 208C21.49 208 0 229.5 0 256s21.49 48 48 48S96 282.5 96 256S74.51 208 48 208zM48 368C21.49 368 0 389.5 0 416s21.49 48 48 48S96 442.5 96 416S74.51 368 48 368zM48 48C21.49 48 0 69.49 0 96s21.49 48 48 48S96 122.5 96 96S74.51 48 48 48zM192 128h288c17.67 0 32-14.33 32-31.1S497.7 64 480 64H192C174.3 64 160 78.33 160 95.1S174.3 128 192 128zM480 224H192C174.3 224 160 238.3 160 256s14.33 32 32 32h288c17.67 0 32-14.33 32-32S497.7 224 480 224zM480 384H192c-17.67 0-32 14.33-32 32s14.33 32 32 32h288c17.67 0 32-14.33 32-32S497.7 384 480 384z"
				/></svg
			>
		</button>
		<!-- UNDO -->
		<button
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
				viewBox="0 0 512 512"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M480 256c0 123.4-100.5 223.9-223.9 223.9c-48.86 0-95.19-15.58-134.2-44.86c-14.14-10.59-17-30.66-6.391-44.81c10.61-14.09 30.69-16.97 44.8-6.375c27.84 20.91 61 31.94 95.89 31.94C344.3 415.8 416 344.1 416 256s-71.67-159.8-159.8-159.8C205.9 96.22 158.6 120.3 128.6 160H192c17.67 0 32 14.31 32 32S209.7 224 192 224H48c-17.67 0-32-14.31-32-32V48c0-17.69 14.33-32 32-32s32 14.31 32 32v70.23C122.1 64.58 186.1 32.11 256.1 32.11C379.5 32.11 480 132.6 480 256z"
				/></svg
			>
		</button>
		<!-- REDO -->
		<button
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
				viewBox="0 0 512 512"
				transform="scale (-1, 1)"
				transform-origin="center"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M480 256c0 123.4-100.5 223.9-223.9 223.9c-48.86 0-95.19-15.58-134.2-44.86c-14.14-10.59-17-30.66-6.391-44.81c10.61-14.09 30.69-16.97 44.8-6.375c27.84 20.91 61 31.94 95.89 31.94C344.3 415.8 416 344.1 416 256s-71.67-159.8-159.8-159.8C205.9 96.22 158.6 120.3 128.6 160H192c17.67 0 32 14.31 32 32S209.7 224 192 224H48c-17.67 0-32-14.31-32-32V48c0-17.69 14.33-32 32-32s32 14.31 32 32v70.23C122.1 64.58 186.1 32.11 256.1 32.11C379.5 32.11 480 132.6 480 256z"
				/></svg
			>
		</button>
		<!-- LINK -->
		<button
			on:click={(e) => {
				ShowLinkedText = !ShowLinkedText;
			}}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
				style="enable-background:new 0 0 24 24;"
				version="1.1"
				viewBox="0 0 24 24"
				xml:space="preserve"
				xmlns="http://www.w3.org/2000/svg"
				xmlns:xlink="http://www.w3.org/1999/xlink"
				><style type="text/css">
					.st0 {
						opacity: 0.2;
						fill: none;
						stroke: #000000;
						stroke-width: 5e-2;
						stroke-miterlimit: 10;
					}
				</style><g id="grid_system" /><g id="_icons"
					><path
						d="M15.6,3.3c-1.4,0-2.6,0.5-3.6,1.5L9.8,7c-0.4,0.4-0.4,1,0,1.4s1,0.4,1.4,0l2.2-2.2c1.2-1.2,3.2-1.2,4.4,0   c0.6,0.6,0.9,1.4,0.9,2.2c0,0.8-0.3,1.6-0.9,2.2l-2.2,2.2c-0.4,0.4-0.4,1,0,1.4c0.2,0.2,0.5,0.3,0.7,0.3s0.5-0.1,0.7-0.3l2.2-2.2   c1-1,1.5-2.2,1.5-3.6c0-1.4-0.5-2.6-1.5-3.6C18.2,3.8,17,3.3,15.6,3.3z"
					/><path
						d="M8.4,11.3c0.4-0.4,0.4-1,0-1.4s-1-0.4-1.4,0L4.8,12c-1,1-1.5,2.2-1.5,3.6c0,1.4,0.5,2.6,1.5,3.6c1,1,2.2,1.5,3.6,1.5   s2.6-0.5,3.6-1.5l2.2-2.2c0.4-0.4,0.4-1,0-1.4s-1-0.4-1.4,0l-2.2,2.2c-1.2,1.2-3.2,1.2-4.4,0c-0.6-0.6-0.9-1.4-0.9-2.2   c0-0.8,0.3-1.6,0.9-2.2L8.4,11.3z"
					/><path
						d="M9.1,14.9c0.2,0.2,0.5,0.3,0.7,0.3s0.5-0.1,0.7-0.3l4.3-4.3c0.4-0.4,0.4-1,0-1.4s-1-0.4-1.4,0l-4.3,4.3   C8.7,13.8,8.7,14.5,9.1,14.9z"
					/></g
				></svg
			>
		</button>
		<!-- image -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'image')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-6 w-6 fill-slate-400 stroke-1 "
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path
					fill-rule="evenodd"
					d="M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm12 12H4l4-8 3 6 2-4 3 6z"
					clip-rule="evenodd"
				/></svg
			>
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
			on:click={(e) => HandleSelectedStr(e, 'Table')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-6 w-6 fill-slate-400 "
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
				xmlns="http://www.w3.org/2000/svg"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M3 10h18M3 14h18m-9-4v8m-7 0h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"
				/></svg
			>
		</button>
		<!-- ALIGN LEFT -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'AlignLeft')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2 hover:fill-gray-300"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path d="M1 1h18v2H1V1zm0 8h18v2H1V9zm0 8h18v2H1v-2zM1 5h12v2H1V5zm0 8h12v2H1v-2z" /></svg
			>
		</button>
		<!-- ALIGN RIGHT -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'AlignRight')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path d="M1 1h18v2H1V1zm0 8h18v2H1V9zm0 8h18v2H1v-2zM7 5h12v2H7V5zm0 8h12v2H7v-2z" /></svg
			>
		</button>
		<!-- ALIGN CENTER -->
		<button
			on:click={(e) => HandleSelectedStr(e, 'AlignCenter')}
			class="flex h-9 w-9 items-center justify-center rounded-lg border-[1px] border-gray-500 bg-inherit hover:bg-gray-500"
		>
			<svg
				class="h-9 w-9 fill-slate-400 p-2  hover:fill-gray-300"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path d="M1 1h18v2H1V1zm0 8h18v2H1V9zm0 8h18v2H1v-2zM4 5h12v2H4V5zm0 8h12v2H4v-2z" /></svg
			>
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
					on:click={() => (ShowLinkedText = false)}
					class=" text-base text-red-500 hover:text-red-400 active:text-red-700">Cancel</button
				>
			</div>
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
