<script lang="ts">
	import { fetchPublicQuestionDataArr } from "$lib/store/fetch";
	import type { QuestionDataType } from "$lib/store/types";

	// import { toggle_class } from "svelte/internal";
	// import {PublicQuesData} from "../store/store"

	// export let PublicQuesData:any

	let PublicQuesData = [
		// [voteNumber(int),AnsNum(int),Viewnum(num),AnsAccepted(bool),askedby(array),tag(Array),QuesTitle(string),QuesDes(string) ]
		[
			235,
			999,
			1626842,
			true,
			['yuji', 'Yuji Itadori', 999, 'Jan 12, 2010 at 16:18'],
			['go', 'rust', 'backend', 'linux'],
			'How to check if a map contains a key in Go?',
			"185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"
		],
		[
			235,
			9,
			1626842,
			true,
			['yuji', 'Yuji Itadori', 999, 'Jan 12, 2010 at 16:18'],
			['go', 'rust', 'backend', 'linux'],
			'How to check if a map contains a key in Go?',
			"185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"
		],
		[
			0,
			0,
			0,
			false,
			['yuji', 'Yuji Itadori', 999, 'Jan 12, 2010 at 16:18'],
			['go', 'rust', 'backend', 'linux'],
			'How to check if a map contains a key in Go? How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?How to check if a map contains a key in Go?',
			"185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec 185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"
		],
		[
			786,
			999,
			1626842,
			false,
			['yuji', 'Yuji Itadori', 999, 'Jan 12, 2010 at 16:18'],
			['go', 'rust', 'backend', 'linux'],
			'How to check if a map contains a key in Go?',
			"185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"
		],
		[
			235,
			999,
			1626842,
			true,
			['yuji', 'Yuji Itadori', 999, 'Jan 12, 2010 at 16:18'],
			['go', 'rust', 'backend', 'linux'],
			'How to check if a map contains a key in Go?',
			"185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"
		],
		[
            235,
			0,
			1626842,
			false,
			['yuji', 'Yuji Itadori', 999, 'Jan 12, 2010 at 16:18'],
			['go', 'rust', 'backend', 'linux'],
			'How to check if a map contains a key in Go?',
			"185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a key's existence in a map?I couldn't find the answer in the language spec"
		]
	];
    
    let d = new Date()
    // console.log("🚀 ~ file: index.svelte:121 ~ d:", d)
	let QuestionList2: {
		ID: string;
		QuesTitle: string;
		QuesitonAskedTime: string | Date;
		QuesModified: string;
		QuesViewed: number;
		QuesUpvote: number;
		QuesDownvote: number;
		QuesBookmark: number;
		QuesTags: string[];
		QuesAskedBy: string;
		QuesAnsAccepted: string;

		QuesAskedTimeExact: string;
		QuesAskedByImage: string;

		QuesEditedBy: string;

		QuesEditedTime: string | Date;
		QuesDescription: string;
		QuesComment: string[][];
		Answers: string[];
	}[] = [
		{
			ID: '1',
			QuesTitle: 'How to check if a map contains a key in Go?',
			QuesitonAskedTime: new Date(),
			QuesModified: 'youKnowWho',
			QuesViewed: 2,
			QuesUpvote: 5,
			QuesDownvote: 6,
			QuesBookmark: 23,
			QuesTags: ['go', 'rust',"python"],
			QuesAskedBy: 'Sk Shahriar Ahmed raka',
			QuesAnsAccepted: '', // Id of excepted ans

			QuesAskedTimeExact: '',
			QuesAskedByImage: '',
			QuesEditedBy: '',

			QuesEditedTime: new Date(),
			QuesDescription: '185I know I can iterate over a map m by,for k, v := range m and look for a key but is there a more efficient way of testing a keys existence in a map?I couldnt find the answer in the language spec',
			QuesComment: [
				['raka', 'no comment'],
				['ssar', 'hello']
			],
			Answers: ['4215', '5524']
		}
	];
	let QuestionList:QuestionDataType[] = [] as QuestionDataType[]
	let fetchData = async () => {
		
		QuestionList = await fetchPublicQuestionDataArr("all",20)
	}
	fetchData()
	console.log("🚀 ~ file: index.svelte:126 ~ QuestionList:", QuestionList)
	
</script>

<div class="h-full w-full ">
	{#if QuestionList.length != 0}
		{#each QuestionList as i}
			<div class=" min-h-32 mx-2 flex max-h-40 w-full flex-row border-t-2 border-[#404245] py-2 ">
				<!-- quesVote -->
				<div class="my-2 flex h-full w-60 flex-col items-center justify-center gap-1 ">
					<!-- VoteNumber -->
					<div class="  {i.QuesUpvote-i.QuesDownvote > 0 ? 'text-[#e7e9eb]' : 'text-[#959ba0]'} ">{i.QuesUpvote-i.QuesDownvote} Votes</div>
					<!-- Answer Number -->
					{#if i.Answers.length ==0}
						<!-- acceped ans -->
						<div class=" rounded-md bg-green-500 px-2 text-white">{i.Answers.length} Answers</div>
					{:else if i.QuesAnsAccepted != ""}
						<!-- Many ans but No accepted -->
						<div class="rounded-md border-2 border-green-500 px-2 text-[#e7e9eb]">
							{i.Answers.length} Answers
						</div>
					{:else}
						<!-- No answer -->
						<div class="text-[#959ba0]">{i.Answers.length} Answers</div>
					{/if}
					<!-- View Number -->
					<div class="{i.QuesViewed > 0 ? 'text-[#e7e9eb]' : 'text-[#959ba0]'} ">{i.QuesViewed} Views</div>
				</div>
				<!-- Qtitle &detail -->
				<div class=" flex h-full w-[1200px] flex-col">
					<!-- Title -->
					<div class=" text-xl text-sky-600 line-clamp-2 hover:cursor-pointer hover:text-blue-600">
						{i.QuesTitle}
					</div>
					<!-- description -->
					<div class=" text-md text-[#e7e9eb] line-clamp-2">{i.QuesDescription}</div>
					<!-- tag ans user -->
					<div class="flex flex-row  items-center justify-center ">
						<div class="  min-h-10 mt-2  flex max-h-20 w-full flex-row flex-wrap">
							<!-- {#each i[5] as tag} -->
                                  <div class="text-[#9bc0da] hover:text-teal-200 hover:bg-slate-600 bg-[#3d4951] m-1 px-2 py-1 rounded-md hover:cursor-pointer">
                                      {"tag"}
                                  </div>
                              <!-- {/each}    -->
						</div>
						<div class="" />

						<p class="h-6 w-36 text-center mt-2 mr-1 text-sky-600 hover:text-blue-600 hover:cursor-pointer line-clamp-1">{"Ques Asked By Name"}</p>
						<p class="text-[#e7e9eb] mt-2 mr-1">{"146k"}</p>
						<p class=" w-80 text-[#959ba0] mt-2 ">{d.toDateString()+" "+d.toLocaleTimeString()}</p>
					</div>
				</div>
			</div>
		{/each}
	{:else}
		<!-- else content here -->
		<div class=" h-full w-full text-xl text-white">NO Question Found</div>
	{/if}
</div>


<style>
	/* your styles go here */
</style>
