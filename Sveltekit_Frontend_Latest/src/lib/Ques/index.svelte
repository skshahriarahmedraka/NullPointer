
<script lang="ts">
    // import {QuestionData} from "../store/store"
    import RelatedQues from "$lib/RelatedQues/index.svelte"
    // import Ans from "$lib/Ans/index.svelte"
    import MarkDownWriter from "$lib/Write/index.svelte"
	import AvatarDefault from "$lib/icons/avatarDefault.svg";
	import type { QuestionDataType } from "$lib/store/types";

    export let QuestionData:QuestionDataType
    // export let RelatedQuestionList:any
    // export let RelatedQuestionListLoading:boolean
    // export let QuestionData:any
    let DropDownData={
        "Show":false,
        "a-z":true ,
        "List":["Newest","Recent Activity","Highest score","Most frequent","bounty"]
    }
    function DropDownClick(){
        DropDownData["Show"]=!DropDownData["Show"]
    }
    function DropDownBlur(){
        if (DropDownData["Show"]){
            DropDownData["Show"]=false
        }
    }
    let writeAns=false
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
	let AskedBy = {
		UserID: 'skraka',
		UserName: 'Sk Shahriar Ahmed Raka',

		UserImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
		Badges: { Reputation: 66813, Gold: 999025, Silver: 8880324, Bronze: 77789423 }
	};
	let ModifiedBy = {
		UserID: 'skraka',
		UserName: ' sheikh Ahmed Raka',

		UserImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
		Badges: { Reputation: 681385285, Gold: 8893785, Silver: 646234, Bronze: 77455345 }
	};

//     var isoTimestamp = new Date().toISOString();
// var localTimestamp = new Date(isoTimestamp).toLocaleString();

// console.log(localTimestamp);

</script>

<style>
</style>


<div class=" w-[1100px] min-h-screen max-h-full  bg-[#2d2d2d] mt-2 pl-5 ">        
    <div class=" h-[150px] w-[1050px]  ml-3 mt-3  flex flex-col text-[#e7e9eb] ">
        <!-- QUESTION TITLE -->
        <div class=" flex flex-row ">
            <p class=" text-2xl line-clamp-3 basis-10/12 font-raleway ">{QuestionData.QuesTitle}</p>
            <div class="ml-3  h-12  basis-2/12  rounded-md bg-[#0964aa] hover:bg-blue-600 flex justify-center items-center ">
                <p class="  my-auto text-gray-200 font-semibold text-xl ">Ask Question</p>
            </div>

        </div>
        <div class="flex flex-row space-x-3 mt-2 ml-3">
            <div class=""><p class="text-[#959ba0] inline-flex ">Asked</p> {QuestionData.QuesAskedTime}</div>
            <div class=""><p class="text-[#959ba0] inline-flex ">Modified</p> {QuestionData.QuesEditedTime}</div>
            <div class=""><p class="text-[#959ba0] inline-flex ">Viewed</p> {QuestionData.QuesViewed}</div>
        </div>
        
    </div>
    <!-- HORIZONTAL BAR -->
    <div class=" h-[1px] w-[1050px] bg-stone-400 ml-3 overflow-hidden "></div>
    <div class=" flex flex-row  ">
        <!-- QUESTION DESCRIPTION CONTAINER  -->
        <div class="w-[850px]  ">
            <!-- QUESTION DESCRIPTION --> 
           <div class="">
                <div class=" w-full   my-3 flex flex-row pt-2 ">
                    <!-- Ratings -->
                    <div class=" w-14 flex flex-col ">
                        <!-- UP -->
                        <svg  class="  h-8 w-8 mx-2 fill-[#696f75] hover:fill-gray-400" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                            <path d="m7.247 4.86-4.796 5.481c-.566.647-.106 1.659.753 1.659h9.592a1 1 0 0 0 .753-1.659l-4.796-5.48a1 1 0 0 0-1.506 0z"/>
                        </svg>
                        <div class=" text-xl mx-2 text-center {(QuestionData.QuesUpvote - QuestionData.QuesDownvote)<0 ? " text-red-500" :" text-green-500" }">
                            {QuestionData.QuesUpvote - QuestionData.QuesDownvote}
                        </div>
                        <!-- DOWN -->
                        <svg class=" h-8 w-8 mx-2 fill-[#696f75] hover:fill-gray-400 " xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"  viewBox="0 0 16 16">
                            <path d="M7.247 11.14 2.451 5.658C1.885 5.013 2.345 4 3.204 4h9.592a1 1 0 0 1 .753 1.659l-4.796 5.48a1 1 0 0 1-1.506 0z"/>
                        </svg>
                    </div>
                    <!-- Question Detail  -->
                    <div class="w-full text-[#e7e9eb] font-sf-pro ">
                        <div class="">
                            {QuestionData.QuesDescription}
                        </div>
                        <!-- Question Tag -->
                        <div class="  min-h-10 max-h-20  w-full mt-2 flex flex-row flex-wrap">
                            {#each QuestionData.QuesTags as tag}
                                <div class="text-[#9bc0da] bg-[#3d4951] hover:text-teal-200 hover:bg-slate-600 hover:cursor-pointer m-1 px-2 py-1 rounded-md">
                                    {tag}
                                </div>
                            {/each}   
                        </div>
                        <!-- share edit askedBy modifiedBy -->
                        <div class="flex flex-row mt-2  ">
                            <svg class="w-8 h-8 m-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M15 8a3 3 0 10-2.977-2.63l-4.94 2.47a3 3 0 100 4.319l4.94 2.47a3 3 0 10.895-1.789l-4.94-2.47a3.027 3.027 0 000-.74l4.94-2.47C13.456 7.68 14.19 8 15 8z"></path></svg>
                            <svg class="w-8 h-8 m-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M17.414 2.586a2 2 0 00-2.828 0L7 10.172V13h2.828l7.586-7.586a2 2 0 000-2.828z"></path><path fill-rule="evenodd" d="M2 6a2 2 0 012-2h4a1 1 0 010 2H4v10h10v-4a1 1 0 112 0v4a2 2 0 01-2 2H4a2 2 0 01-2-2V6z" clip-rule="evenodd"></path></svg>
                            <svg class="w-8 h-8 m-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M5 3a1 1 0 000 2c5.523 0 10 4.477 10 10a1 1 0 102 0C17 8.373 11.627 3 5 3z"></path><path d="M4 9a1 1 0 011-1 7 7 0 017 7 1 1 0 11-2 0 5 5 0 00-5-5 1 1 0 01-1-1zM3 15a2 2 0 114 0 2 2 0 01-4 0z"></path></svg>
                            <!-- Gap -->
                            <div class="grow"></div>
                            <!-- MODIFIED BY -->
                            {#if QuestionData.QuesEditedTime != ''}
				<!-- content here -->
				<div class=" flex h-20  w-60 flex-col overflow-hidden px-2 ">
					<p class="text-sm text-white ">Edited : {QuestionData.QuesEditedTime}</p>
					<div class="flex  h-12 w-full flex-row  ">
                        {#if ModifiedBy.UserImage != "" }
                             <img
                                 src="{ModifiedBy.UserImage}"
                                 alt="Modified By UserImage"
                                 class=" aspect-square  active:ring-offset-base-50 mt-1 h-10  w-10 cursor-pointer rounded-3xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
                             />
                        {:else}
                        <img
                        src="{AvatarDefault}"
                        alt="Modified By UserImage"
                        class=" aspect-square  active:ring-offset-base-50 mt-1 h-10  w-10 cursor-pointer rounded-full object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
                    />
                             <!-- else content here -->
                        {/if}
						<div class=" flex flex-col ">
							<div class=" ml-2  line-clamp-1">{ModifiedBy.UserName}</div>
							<div class=" ml-2  flex flex-row ">
								{#if ModifiedBy.Badges.Reputation != 0}
									<svg
										class="mt-2 ml-1 h-2 w-2 place-content-center  fill-white  "
										viewBox="0 0 512 512"
										xmlns="http://www.w3.org/2000/svg"
										><path
											d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
										/></svg
									>
									<p class="mx-1   text-white">{ RoundNum(ModifiedBy.Badges.Reputation)}</p>
								{/if}
                                {#if ModifiedBy.Badges.Gold != 0}
									<svg
										class="mt-2 ml-1 h-2 w-2 place-content-center  fill-[#ffcc01]  "
										viewBox="0 0 512 512"
										xmlns="http://www.w3.org/2000/svg"
										><path
											d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
										/></svg
									>
									<p class="mx-1   text-[#ffcc01]">{ RoundNum(ModifiedBy.Badges.Gold)}</p>
								{/if}
								{#if  ModifiedBy.Badges.Silver != 0}
									<svg
										class="mt-2 ml-1 h-2 w-2  place-content-center fill-[#b4b8bc]  "
										viewBox="0 0 512 512"
										xmlns="http://www.w3.org/2000/svg"
										><path
											d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
										/></svg
									>
									<p class="mx-1  text-[#b4b8bc]">{ RoundNum(ModifiedBy.Badges.Silver)}</p>
								{/if}
								{#if  ModifiedBy.Badges.Bronze != 0}
									<svg
										class="mt-2 ml-1 h-2 w-2 place-content-center fill-[#d1a684]  "
										viewBox="0 0 512 512"
										xmlns="http://www.w3.org/2000/svg"
										><path
											d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
										/></svg
									>
									<p class="mx-1  text-[#d1a684]">{ RoundNum(ModifiedBy.Badges.Bronze)}</p>
								{/if}
							</div>
						</div>
					</div>
				</div>
			{/if}
			<!-- ANSWERED BY -->
			<div class=" ml-2 flex h-20  w-56 flex-col overflow-hidden px-2">
				<p class="text-sm text-white ">Answered : {"answer time"}</p>
				<div class="flex h-12 w-full flex-row ">
                    {#if AskedBy.UserImage != ""}
                    <!-- svelte-ignore a11y-img-redundant-alt -->
                         <img
                             src="{AskedBy.UserImage}"
                             alt="Answered By Image"
                             class=" aspect-square active:ring-offset-base-50 mt-1 h-10  w-10 cursor-pointer rounded-3xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
                         />
                    {:else}
                    <img
                    src="{AvatarDefault}"
                    alt="Modified By UserImage"
                    class=" aspect-square  active:ring-offset-base-50 mt-1 h-10  w-10 cursor-pointer rounded-full object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
                />
                    {/if}
					<div class=" flex flex-col ">
						<div class=" ml-2 line-clamp-1">{AskedBy.UserName}</div>
						<div class=" flex flex-row">
							{#if AskedBy.Badges.Gold != 0}
								<svg
									class="mt-2 ml-1 h-2 w-2 place-content-center  fill-[#ffcc01]  "
									viewBox="0 0 512 512"
									xmlns="http://www.w3.org/2000/svg"
									><path
										d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
									/></svg
								>
								<p class="mx-1   text-[#ffcc01]">{ RoundNum(AskedBy.Badges.Gold) }</p>
							{/if}
							{#if AskedBy.Badges.Silver  != 0}
								<svg
									class="mt-2 ml-1 h-2 w-2  place-content-center fill-[#b4b8bc]  "
									viewBox="0 0 512 512"
									xmlns="http://www.w3.org/2000/svg"
									><path
										d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
									/></svg
								>
								<p class="mx-1  text-[#b4b8bc]">{ RoundNum(AskedBy.Badges.Silver)}</p>
							{/if}
							{#if  AskedBy.Badges.Bronze != 0}
								<svg
									class="mt-2 ml-1 h-2 w-2 place-content-center fill-[#d1a684]  "
									viewBox="0 0 512 512"
									xmlns="http://www.w3.org/2000/svg"
									><path
										d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
									/></svg
								>
								<p class="mx-1  text-[#d1a684]">{ RoundNum(AskedBy.Badges.Bronze)}</p>
							{/if}
						</div>
					</div>
				</div>
			</div>
                        </div>
                    </div>
                </div>
                <!-- Number of ANSWER -->
                <div class=" h-14  rounded-lg m-4 mr-2 border-[1px] border-gray-600 text-[#e7e9eb] text-2xl flex flex-row place-content-center place-items-center">
                    <div class="ml-3">{QuestionData.Answers.length} Answers</div>
                    <!-- write answer button -->
                    <div on:click={()=>{writeAns=!writeAns}} on:keypress={()=>{}} class="ml-3 px-2 cursor-pointer rounded-md  border-2 border-gray-600 bg-gray-500 bg-opacity-25 hover:bg-gray-800">Write Answer</div>
                    <div class="grow"></div>
                    <!-- drop down -->
                    <div class="m-2 mr-4 ">
                        <div class=" inline-block relative ">
                          <button  on:blur="{DropDownBlur }"  on:click="{DropDownClick}" class="text-[#e7e9eb] text-lg font-medium border-2 border-[#688fac]  rounded inline-flex items-center p-1 pl-2">
                            <svg class="w-6 h-6 mr-2 fill-[#688fac]" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M5 4a1 1 0 00-2 0v7.268a2 2 0 000 3.464V16a1 1 0 102 0v-1.268a2 2 0 000-3.464V4zM11 4a1 1 0 10-2 0v1.268a2 2 0 000 3.464V16a1 1 0 102 0V8.732a2 2 0 000-3.464V4zM16 3a1 1 0 011 1v7.268a2 2 0 010 3.464V16a1 1 0 11-2 0v-1.268a2 2 0 010-3.464V4a1 1 0 011-1z"></path></svg>
                            <span class="mr-1">Most Populer</span>
                          </button>
                          {#if DropDownData["Show"]}
                               <ul class=" absolute  w-full text-[#e7e9eb] bg-slate-700 text-lg font-medium border-2 border-[#688fac]  rounded    ">
                                 <li class=" hover:bg-slate-500 p-1 pl-2">Most Populer</li>
                                 <li class="hover:bg-slate-500 p-1 pl-2">Most Shared</li>
                                 <li class="hover:bg-slate-500 p-1 pl-2">Most commented</li>
                               </ul>
                          {/if}
                        </div>
                      
                    </div>
                    <!-- <div class="w-14 h-8 bg-slate-200 "></div> -->
                </div>
                <div class=" m-5  w-[95%] ">
                    {#if writeAns}
                     <MarkDownWriter/>
                    {/if}
                </div>
                <!-- {#each QuestionData.Answers as ans }
                    <Ans {ans} />
                {/each} -->
            
            </div>
        
        </div>
        
       
        <!--List of Answers -->
        <!-- RELATED QUESTION -->
        {#if true}
             <!-- content here -->
             <RelatedQues/>
        {/if}


    </div>
    <!-- <Ans/> -->
</div>