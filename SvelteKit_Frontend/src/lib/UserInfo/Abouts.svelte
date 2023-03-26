<script lang="ts">
        import {UserData} from "../store/store"
        // export let UserData:any
        import {marked} from 'marked'

        import showdown from 'showdown';
        
        const converter = new showdown.Converter();
        const source =`# A Curious Learner, Full-Stack Web Developer, Security Researcher 
## Here are my skills and strengths:
✓  Expert in Golang <br/>
✓  Expert in Fiber framework (using Golang) <br/>
✓  Expert in WebAssembly (using Golang) <br/>
✓  Expert in database design, development, optimization, and migration <br/>
   (PostgreSQL, MySQL, MongoDB , Redis)<br/>
✓  Expert in ( Grpc, protocol buffer )<br/>
✓  Experienced in ( WebSocket, Socket.io, WebRTC ) for real-time client and server applications <br/>
✓  Experienced in ( Svelte.js ) and some knowledge in ( TypeScript )<br/>
✓  Good understanding of ( Docker, Bash, PowerShell, Git,<br/>
   Nginx, Kubernetes )<br/>


Github   : github.com/skshahriarahmedraka <br/>
Upwork   : upwork.com/o/profiles/users/~0107ef3405bffbe57e <br/>
Linkedin : linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/ <br/>
Telegram : t.me/shahriarraka <br/>
| Option | Description |
| ------ | ----------- |
| data   | path to data files to supply the data that will be passed into templates. |
| engine | engine to be used for processing templates. Handlebars is the default. |
| ext    | extension to be used for dest files. |

Right aligned columns

| Option | Description |
| ------:| -----------:|
| data   | path to data files to supply the data that will be passed into templates. |
| engine | engine to be used for processing templates. Handlebars is the default. |
| ext    | extension to be used for dest files. |
`

// let markdownText = '# Hello, world!';
let htmlText = converter.makeHtml(source);

import Markdown from 'svelte-markdown';
	import Badges from "./Badges.svelte";

const source2 = `
  # This is a header

This is a paragraph.

* This is a list
* With two items
  1. And a sublist
  2. That is ordered
    * With another
    * Sublist inside

| And this is | A table |
|-------------|---------|
| With two    | columns |`
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->


    <div class=" text-[#e7e9eb] text-2xl ml-4 ">
        Abouts
    </div>
    <div class=" p-6 text-gray-300 prose
     prose-h1:text-gray-300 prose-h2:text-gray-300 prose-h3:text-gray-300
      prose-h4:text-gray-300 prose-h5:text-gray-300 prose-h6:text-gray-300
       prose-headings:text-gray-300  


       prose-p:text-gray--300
prose-a:text-blue-500
prose-blockquote:blockquote
prose-figure:	figure
prose-figcaption:	figcaption
prose-strong:	strong
prose-em:	em
prose-code:	code
prose-pre:	pre
prose-ol:	ol
prose-ul:	ul
prose-li:	li
prose-table:	table
prose-thead:	thead
prose-tr:tr
prose-th:	th
prose-td:td
prose-img:	img
prose-video:	video
prose-hr:	hr
      ">
        <Markdown {source} />
        <!-- {@html htmlText} <br >  -->
        <!-- {@html marked(About)} -->
        <!-- <h1 id="acuriouslearnerfullstackwebdevelopersecurityresearcher">A Curious Learner, Full-Stack Web Developer, Security Researcher</h1> <h2 id="herearemyskillsandstrengths">Here are my skills and strengths:</h2> <p>✓ Expert in Golang <br/> ✓ Expert in Fiber framework (using Golang) <br/> ✓ Expert in WebAssembly (using Golang) <br/> ✓ Expert in database design, development, optimization, and migration <br/> (PostgreSQL, MySQL, MongoDB , Redis)<br/> ✓ Expert in ( Grpc, protocol buffer )<br/> ✓ Experienced in ( WebSocket, Socket.io, WebRTC ) for real-time client and server applications <br/> ✓ Experienced in ( Svelte.js ) and some knowledge in ( TypeScript )<br/> ✓ Good understanding of ( Docker, Bash, PowerShell, Git,<br/> Nginx, Kubernetes )<br/></p> <p>Github : github.com/skshahriarahmedraka <br/> Upwork : upwork.com/o/profiles/users/~0107ef3405bffbe57e <br/> Linkedin : linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/ <br/> Telegram : t.me/shahriarraka <br/></p> -->
        
    </div>
