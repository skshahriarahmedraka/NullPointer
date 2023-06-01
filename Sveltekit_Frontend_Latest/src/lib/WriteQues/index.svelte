<script lang="ts">
	
	import { UserData } from '$lib/store/store';
	import { marked } from 'marked';
	import { fade, fly } from 'svelte/transition';
	import { tick } from 'svelte';
	import HotQuesCom from '$lib/HotQues/index.svelte';
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
	import type { QuestionDataType } from '$lib/store/types';
	import { fetchAskQuestion } from '$lib/store/fetch';
	import { goto } from '$app/navigation';

	// export let HotQues: any;
	let ShowLinkedText: boolean = false;
	let ShowLinkedImage: boolean = false;
	let markdown: string = '';
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
	let CashedLinkImage: string = 'https://';
	let tArea: any = null;

	// Image Upload
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

	let QuestionData: QuestionDataType = {
		ID: '',
		QuesTitle: QuestionTitle,
		QuesViewed: 0,
		QuesUpvote: 0,
		QuesDownvote: 0,
		QuesVotes: [] as {
		ID : string;
		UserID : string;
		Vote : number;
		VoteTime : string;
	}[],
		QuesBookmark: [] as  {
        ID: string;
        UserID: string;
        BookmarkTime: string;
    }[],
		QuesTags: TagObj.List,
		QuesGroup : GroupsObj.List,
		QuesAnsAccepted: "",

		QuesAskedBy: $UserData.UserID,
		QuesAskedTime: new Date().toISOString(),

		QuesEditedBy: "",
		QuesEditedTime: new Date().toISOString(),

		QuesDescription: markdown ,
		QuesComment: [] as string[],
		Answers: [] as {
    ID: string;
    UpVote: number;
    DownVote: number;
    AnsweredBy: string;
    Comment: string[];
}[]
	};
	let ErrorMsg = {
		QuesTitle: [false, 'QuesTitle is not Valid'],
		QuesTags: [false, 'QuesTags are empty'],
		QuesDescription: [false, 'QuesDescription is too short']
	};
	function ValidQuesTitle(s:string): boolean{
		if (s.length < 5) {
			return false;
		} else {
			return true
		}
	}
	function ValidQuesDescription(s:string): boolean{
		if (s.length < 15) {
			return false;
		} else {
			return true
		}
	}
	async function OnSubmitQuestion() {
		QuestionData = {
		ID: '',
		QuesTitle: QuestionTitle,
		QuesViewed: 0,
		QuesUpvote: 0,
		QuesDownvote: 0,
		QuesVotes: [] as {
		ID : string;
		UserID : string;
		Vote : number;
		VoteTime : string;
	}[],
		QuesBookmark: [] as  {
        ID: string;
        UserID: string;
        BookmarkTime: string;
    }[],
		QuesTags: TagObj.List,
		QuesGroup : GroupsObj.List,
		QuesAnsAccepted: "",

		QuesAskedBy: $UserData.UserID,
		QuesAskedTime: new Date().toISOString(),

		QuesEditedBy: "",
		QuesEditedTime: new Date().toISOString(),

		QuesDescription: markdown ,
		QuesComment: [] as string[],
		Answers: [] as {
    ID: string;
    UpVote: number;
    DownVote: number;
    AnsweredBy: string;
    Comment: string[];
}[]
	}
		if (QuestionData.QuesTitle.trim() === '') {
			ErrorMsg.QuesTitle[1] = 'Title is empty';
			ErrorMsg.QuesTitle[0] = true;
		} else if (!ValidQuesTitle(QuestionData.QuesTitle)) {
			ErrorMsg.QuesTitle[1] = 'Title is too small';
			ErrorMsg.QuesTitle[0] = true;
		} else {
			ErrorMsg.QuesTitle[0] = false;
		}
		if (QuestionData.QuesTags.length <1) {
			ErrorMsg.QuesTags[1] = 'Tags are atleast 1';
			ErrorMsg.QuesTags[0] = true;
		}else {
			ErrorMsg.QuesTags[0] = false;
		}
		
		 if (!ValidQuesDescription(QuestionData.QuesDescription)) {
			ErrorMsg.QuesDescription[1] = 'Description is too short ';
			ErrorMsg.QuesDescription[0] = true;
		} else {
			ErrorMsg.QuesDescription[0] = false;
		}
		
		if (!ErrorMsg.QuesTitle[0] && !ErrorMsg.QuesTags[0] && !ErrorMsg.QuesDescription[0]) {
			

			const res:{"InsertedID":string} = await fetchAskQuestion(QuestionData)
			console.log("🚀 TagObj~ file: index.svelte:656 ~ OnSubmitQuestion ~ res:", res)
			if (res.InsertedID!==''){
				goto('/q/'+res.InsertedID)

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
	async function OnResetQuestion(){

		QuestionData = {
		ID: '',
		QuesTitle: "",
		QuesViewed: 0,
		QuesUpvote: 0,
		QuesDownvote: 0,
		QuesVotes: [] as {
		ID : string;
		UserID : string;
		Vote : number;
		VoteTime : string;
	}[],
		QuesBookmark: [] as  {
        ID: string;
        UserID: string;
        BookmarkTime: string;
    }[],
		QuesTags: [] as string[],
		QuesGroup : [] as string[],
		QuesAnsAccepted: "",
	
		QuesAskedBy: $UserData.UserID,
		QuesAskedTime: "",
	
		QuesEditedBy: '',
		QuesEditedTime: '',
	
		QuesDescription: "" ,
		QuesComment: [] as string[],
		Answers: [] as {
    ID: string;
    UpVote: number;
    DownVote: number;
    AnsweredBy: string;
    Comment: string[];
}[]
	};
	}
	async function OnCancelQuestion() {
		OnResetQuestion()
		goto('/');
	}
	let ThisTagInput: any;
	function RemoveTagFromList(s: number) {
		TagObj.List.splice(s, 1);
		ThisTagInput.focus();
	}
	let ThisGroupInput: any;
	function RemoveGroupFromList(s: number) {
		GroupsObj.List.splice(s, 1);
		ThisGroupInput.focus();
	}
</script>

<!-- <h1>Markdown Editor</h1> -->
<!-- <TextArea bind:value={markdown} minRows={4} maxRows={40} /> -->

<div class=" mt-12 flex max-h-full  min-h-screen w-[1100px]  flex-row bg-[#2d2d2d] ">
	<!-- Question -->
	<div class=" mt-5 flex  max-h-full min-h-screen w-9/12 flex-col  ">
		<!-- Ques Title -->
		<div class=" mx-5 text-left text-lg font-semibold text-gray-300   ">Title :</div>
		<div class=" mx-5 text-left text-sm font-semibold text-gray-500   ">
			Be specific and imagine you're asking a question to another person
		</div>
		<input
			bind:value={QuestionTitle}
			type="text"
			class=" mx-4 mt-3 h-12 w-[95%] rounded-lg border-2 border-[#24262b] bg-[#303338] p-2  text-lg font-medium text-[#98999e]  outline-none  focus:border-sky-500 active:border-gray-800 "
		/>

		<!-- Ques Description -->

		<div class=" mx-5 mt-3 text-left text-lg font-semibold text-gray-300   ">Description :</div>
		<div class=" mx-5 text-left text-sm font-semibold text-gray-500   ">
			Include all the information someone would need to answer your question
		</div>
		<!-- bold-italic buttons -->
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
		<!-- Tags -->
		<div class=" mx-5 mt-5 text-left text-lg font-semibold text-gray-300   ">Tags :</div>
		<div class=" mx-5 text-left text-sm font-semibold text-gray-500   ">
			Add up to 5 tags to describe what your question is about
		</div>

		<div
			class="mx-4 mt-3 h-fit w-[95%] rounded-lg border-2 border-[#24262b] bg-[#303338] p-2  text-lg font-medium text-[#98999e]  outline-0   {TagObj.Focus
				? 'border-sky-500'
				: ''}  "
		>
			<ul class="flex flex-row flex-wrap gap-2">
				{#each TagObj.List as t,i}
					<li
						class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
					>
						{t}
						<svg
						on:click={() => {
							RemoveTagFromList(i);
						}}
						on:keypress={() => {}}
						class="inline-flex h-6 w-6 hover:stroke-red-500"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
						xmlns="http://www.w3.org/2000/svg"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M6 18L18 6M6 6l12 12"
						/></svg
					>
					</li>
				{/each}
				<input
				bind:this={ThisTagInput}
					on:keypress={onKeyPressTags}
					maxlength="20"
					bind:value={TagObj.Input}
					on:focus={() => {
						TagObj.Focus = true;
					}}
					on:blur={() => {
						TagObj.Focus = true;
					}}
					type="text"
					class=" inline w-fit grow border-0  bg-inherit outline-0 focus:outline-none  "
				/>
			</ul>
		</div>

		<!-- <input
			bind:value={QuestionTitle}
			type="text"
			class=" mx-4 mt-3 h-12 w-[95%] rounded-lg border-2 border-[#24262b] bg-[#303338] p-2  text-lg font-medium text-[#98999e]  outline-none  focus:border-sky-500 active:border-gray-800 "
		/> -->

		<!-- Groups -->
		<div class=" mx-5 mt-5 text-left text-lg font-semibold text-gray-300   ">Group (Optional):</div>
		<div class=" mx-5 text-left text-sm font-semibold text-gray-500   ">
			Add question to a group for like minded people to find
		</div>
		<div
			class="mx-4 mt-3 h-fit w-[95%] rounded-lg border-2 border-[#24262b] bg-[#303338] p-2  text-lg font-medium text-[#98999e]  outline-none   {GroupsObj.Focus
				? 'border-sky-500'
				: ''} active:border-gray-800 "
		>
			<ul class="flex flex-row flex-wrap gap-2">
				{#each GroupsObj.List as t,i}
					<li
						class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
					>
						{t}
						<svg
						on:click={() => {
							RemoveGroupFromList(i);
						}}
						on:keypress={() => {}}
						class="inline-flex h-6 w-6 hover:stroke-red-500"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
						xmlns="http://www.w3.org/2000/svg"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M6 18L18 6M6 6l12 12"
						/></svg
					>
					</li>
				{/each}
				<input
					bind:this={ThisGroupInput}
					on:keypress={onKeyPressGroup}
					maxlength="20"
					bind:value={GroupsObj.Input}
					type="text"
					on:focus={() => {
						GroupsObj.Focus = true;
					}}
					on:blur={() => {
						GroupsObj.Focus = false;
					}}
					class=" inline w-fit grow border-0  bg-inherit shadow-none outline-0 focus:border-0 focus:outline-none"
				/>
			</ul>
		</div>
		<div class=" flex flex-row  justify-center gap-5">
			<button
			on:click={() => {OnSubmitQuestion()}}
				class=" m-5 flex h-12 w-fit  items-center justify-center rounded-md bg-blue-500 px-3 hover:bg-blue-600 active:bg-blue-800 "
			>
				<p class="my-auto text-xl font-semibold text-gray-200">Ask Question</p>
			</button>
			<button
			on:click={() => {OnResetQuestion()}}
				class=" m-5 flex h-12 w-fit  items-center justify-center rounded-md bg-blue-500 px-3 hover:bg-blue-600 active:bg-blue-800 "
			>
				<p class="my-auto text-xl font-semibold text-gray-200">Reset</p>
			</button>
			<button
				on:click={() => {OnCancelQuestion()}}
				class=" m-5 flex h-12 w-fit items-center justify-center  rounded-md border-2 border-red-600 bg-inherit px-3 text-red-600 hover:bg-red-500 hover:text-gray-200 active:bg-red-600 "
			>
				<p class="my-auto text-xl font-semibold  ">Discart</p>
			</button>
		</div>
	</div>
	<div class="max-h-full  min-h-screen w-3/12">
		<HotQuesCom />
	</div>
</div>
