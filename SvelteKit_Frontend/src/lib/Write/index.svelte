<script lang="ts">

    import {marked} from 'marked'
    import TextArea from './TextArea.svelte'
    

    
  
    let markdown:string=""

    
   let minRows = 1;
	 let maxRows=0;
    	
	$: minHeight = `${1 + minRows * 1.2}em`;
	$: maxHeight = maxRows ? `${1 + maxRows * 1.2}em` : `auto`;
        
   
  </script>
  
  <!-- <h1>Markdown Editor</h1> -->
  <TextArea bind:value={markdown} minRows={4} maxRows={40} />
  
  
  <!-- <textarea bind:value={markdown}  /> -->
  <!-- <textarea class="textarea textarea-info bg-inherit" placeholder="Bio"></textarea> -->
  {#if markdown.trim() !=""}

    <div class=" prose 
    prose-headings:text-[#e7e9eb] prose-p:text-[#e7e9eb] prose-a:text-blue-500 prose-blockquote:text-[#e7e9eb]
     prose-figure:text-white 
    prose-figcaption:text-[#e7e9eb]
    prose-strong:text-[#e7e9eb]
    prose-em:text-[#e7e9eb]
    prose-code:text-[#e7e9eb]
    prose-pre:text-[#e7e9eb]
    prose-ol:text-[#e7e9eb]
    prose-ul:text-[#e7e9eb]
    prose-li:text-[#e7e9eb]

    prose-table:text-[#e7e9eb]
    prose-table:border-2
    prose-table:border-gray-300
    prose-thead:text-[#e7e9eb]
  
    prose-tr:text-[#e7e9eb]
    prose-tr:border-2
    prose-tr:border-gray-300

    prose-th:text-[#e7e9eb]
    prose-th:border-2
    prose-th:border-gray-300
    
    prose-td:text-[#e7e9eb]
    prose-td:border-2
    prose-td:border-gray-300
    

    prose-img:text-[#e7e9eb]
    prose-video:text-[#e7e9eb]
    prose-hr:text-[#e7e9eb]
    prose-hr:bg-gray-500
     prose-blockquote:border-sky-400
     prose-li:marker:text-white
                 mt-5 p-5 bg-inherit min-w-full max-w-full  overflow-hidden break-words  ">{@html marked(markdown)}</div>
    <div class=" h-12 w-fit px-3 my-5   rounded-md bg-sky-500 hover:bg-blue-600 flex justify-center items-center ">
      <p class="  my-auto text-gray-200 font-semibold text-xl ">Post Your Answer</p>
    </div>
  {/if}
  
 