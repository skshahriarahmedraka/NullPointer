package handler

import (
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (H *DatabaseCollections)HotQuestion(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Content-Type", "application/json")

	x:= model.HotQuestions{
	{" How can I convert a zero-terminated byte array to string? " ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" ><path d="M15.9.476a2.14 2.14 0 0 0-.823.218L3.932 6.01c-.582.277-1.005.804-1.15 1.432L.054 19.373c-.13.56-.025 1.147.3 1.627q.057.087.12.168l7.7 9.574c.407.5 1.018.787 1.662.784h12.35c.646.001 1.258-.3 1.664-.793l7.696-9.576c.404-.5.555-1.16.4-1.786L29.2 7.43c-.145-.628-.57-1.155-1.15-1.432L16.923.695A2.14 2.14 0 0 0 15.89.476z" fill="#326ce5"/><path d="M16.002 4.542c-.384.027-.675.356-.655.74v.188c.018.213.05.424.092.633a6.22 6.22 0 0 1 .066 1.21c-.038.133-.114.253-.218.345l-.015.282c-.405.034-.807.096-1.203.186-1.666.376-3.183 1.24-4.354 2.485l-.24-.17c-.132.04-.274.025-.395-.04a6.22 6.22 0 0 1-.897-.81 5.55 5.55 0 0 0-.437-.465l-.148-.118c-.132-.106-.294-.167-.463-.175a.64.64 0 0 0-.531.236c-.226.317-.152.756.164.983l.138.11a5.55 5.55 0 0 0 .552.323c.354.197.688.428.998.7a.74.74 0 0 1 .133.384l.218.2c-1.177 1.766-1.66 3.905-1.358 6.006l-.28.08c-.073.116-.17.215-.286.288a6.22 6.22 0 0 1-1.194.197 5.57 5.57 0 0 0-.64.05l-.177.04h-.02a.67.67 0 0 0-.387 1.132.67.67 0 0 0 .684.165h.013l.18-.02c.203-.06.403-.134.598-.218.375-.15.764-.265 1.162-.34.138.008.27.055.382.135l.3-.05c.65 2.017 2.016 3.726 3.84 4.803l-.122.255c.056.117.077.247.06.376-.165.382-.367.748-.603 1.092a5.58 5.58 0 0 0-.358.533l-.085.18a.67.67 0 0 0 .65 1.001.67.67 0 0 0 .553-.432l.083-.17c.076-.2.14-.404.192-.61.177-.437.273-.906.515-1.196a.54.54 0 0 1 .286-.14l.15-.273a8.62 8.62 0 0 0 6.146.015l.133.255c.136.02.258.095.34.205.188.358.34.733.456 1.12a5.57 5.57 0 0 0 .194.611l.083.17a.67.67 0 0 0 1.187.131.67.67 0 0 0 .016-.701l-.087-.18a5.55 5.55 0 0 0-.358-.531c-.23-.332-.428-.686-.6-1.057a.52.52 0 0 1 .068-.4 2.29 2.29 0 0 1-.111-.269c1.82-1.085 3.18-2.8 3.823-4.82l.284.05c.102-.093.236-.142.373-.138.397.076.786.2 1.162.34.195.09.395.166.598.23.048.013.118.024.172.037h.013a.67.67 0 0 0 .841-.851.67.67 0 0 0-.544-.446l-.194-.046a5.57 5.57 0 0 0-.64-.05c-.404-.026-.804-.092-1.194-.197-.12-.067-.22-.167-.288-.288l-.27-.08a8.65 8.65 0 0 0-1.386-5.993l.236-.218c-.01-.137.035-.273.124-.378.307-.264.64-.497.99-.696a5.57 5.57 0 0 0 .552-.323l.146-.118a.67.67 0 0 0-.133-1.202.67.67 0 0 0-.696.161l-.148.118a5.57 5.57 0 0 0-.437.465c-.264.302-.556.577-.873.823a.74.74 0 0 1-.404.044l-.253.18c-1.46-1.53-3.427-2.48-5.535-2.67 0-.1-.013-.25-.015-.297-.113-.078-.192-.197-.218-.332a6.23 6.23 0 0 1 .076-1.207c.043-.21.073-.42.092-.633v-.2c.02-.384-.27-.713-.655-.74zm-.834 5.166l-.2 3.493h-.015c-.01.216-.137.4-.332.504s-.426.073-.6-.054l-2.865-2.03a6.86 6.86 0 0 1 3.303-1.799c.234-.05.47-.088.707-.114zm1.668 0c1.505.187 2.906.863 3.99 1.924l-2.838 2.017c-.175.14-.415.168-.618.072s-.333-.3-.336-.524zm-6.72 3.227l2.62 2.338v.015c.163.142.234.363.186.574s-.21.378-.417.435v.01l-3.362.967a6.86 6.86 0 0 1 .974-4.34zm11.753 0c.796 1.295 1.148 2.814 1.002 4.327l-3.367-.97v-.013c-.21-.057-.37-.224-.417-.435s.023-.43.186-.574l2.6-2.327zm-6.404 2.52h1.072l.655.832-.238 1.04-.963.463-.965-.463-.227-1.04zm3.434 2.838c.045-.005.1-.005.135 0l3.467.585c-.5 1.44-1.487 2.67-2.775 3.493l-1.34-3.244a.59.59 0 0 1 .509-.819zm-5.823.015c.196.003.377.104.484.268s.124.37.047.55v.013l-1.332 3.218C11 21.54 10.032 20.325 9.517 18.9l3.437-.583c.038-.004.077-.004.116 0zm2.904 1.4a.59.59 0 0 1 .537.308h.013l1.694 3.057-.677.2c-1.246.285-2.547.218-3.758-.194l1.7-3.057c.103-.18.293-.29.5-.295z" fill="#fff" stroke="#fff" stroke-width=".055"/></svg>`,"https://stackoverflow.com/"},
    {" go get results in 'terminal prompts disabled' error for github private repo " ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_vue</title><path d="M24.4,3.925H30L16,28.075,2,3.925H12.71L16,9.525l3.22-5.6Z" style="fill:#41b883"/><path d="M2,3.925l14,24.15L30,3.925H24.4L16,18.415,7.53,3.925Z" style="fill:#41b883"/><path d="M7.53,3.925,16,18.485l8.4-14.56H19.22L16,9.525l-3.29-5.6Z" style="fill:#35495e"/></svg>`,"https://stackoverflow.com/"},
    {"s it possible to capture a Ctrl+C signal (SIGINT) and run a cleanup function, in a \"defer\" fashion?" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_svelte</title><path d="M26.47,5.7A8.973,8.973,0,0,0,14.677,3.246L7.96,7.4a7.461,7.461,0,0,0-3.481,5.009,7.686,7.686,0,0,0,.8,5.058,7.358,7.358,0,0,0-1.151,2.8,7.789,7.789,0,0,0,1.4,6.028,8.977,8.977,0,0,0,11.794,2.458L24.04,24.6a7.468,7.468,0,0,0,3.481-5.009,7.673,7.673,0,0,0-.8-5.062,7.348,7.348,0,0,0,1.152-2.8A7.785,7.785,0,0,0,26.47,5.7" style="fill:#ff3e00"/><path d="M14.022,26.64A5.413,5.413,0,0,1,8.3,24.581a4.678,4.678,0,0,1-.848-3.625,4.307,4.307,0,0,1,.159-.61l.127-.375.344.238a8.76,8.76,0,0,0,2.628,1.274l.245.073-.025.237a1.441,1.441,0,0,0,.271.968,1.63,1.63,0,0,0,1.743.636,1.512,1.512,0,0,0,.411-.175l6.7-4.154a1.366,1.366,0,0,0,.633-.909,1.407,1.407,0,0,0-.244-1.091,1.634,1.634,0,0,0-1.726-.622,1.509,1.509,0,0,0-.413.176l-2.572,1.584a4.934,4.934,0,0,1-1.364.582,5.415,5.415,0,0,1-5.727-2.06A4.678,4.678,0,0,1,7.811,13.1,4.507,4.507,0,0,1,9.9,10.09l6.708-4.154a4.932,4.932,0,0,1,1.364-.581A5.413,5.413,0,0,1,23.7,7.414a4.679,4.679,0,0,1,.848,3.625,4.272,4.272,0,0,1-.159.61l-.127.375-.344-.237a8.713,8.713,0,0,0-2.628-1.274l-.245-.074.025-.237a1.438,1.438,0,0,0-.272-.968,1.629,1.629,0,0,0-1.725-.622,1.484,1.484,0,0,0-.411.176l-6.722,4.14a1.353,1.353,0,0,0-.631.908,1.394,1.394,0,0,0,.244,1.092,1.634,1.634,0,0,0,1.726.621,1.538,1.538,0,0,0,.413-.175l2.562-1.585a4.9,4.9,0,0,1,1.364-.581,5.417,5.417,0,0,1,5.728,2.059,4.681,4.681,0,0,1,.843,3.625A4.5,4.5,0,0,1,22.1,21.905l-6.707,4.154a4.9,4.9,0,0,1-1.364.581" style="fill:#fff"/></svg>`,"https://stackoverflow.com/"},
    {"What is the shortest way to simply sort an array of structs by (arbitrary) field names?" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_nginx</title><path d="M15.948,2h.065a10.418,10.418,0,0,1,.972.528Q22.414,5.65,27.843,8.774a.792.792,0,0,1,.414.788c-.008,4.389,0,8.777-.005,13.164a.813.813,0,0,1-.356.507q-5.773,3.324-11.547,6.644a.587.587,0,0,1-.657.037Q9.912,26.6,4.143,23.274a.7.7,0,0,1-.4-.666q0-6.582,0-13.163a.693.693,0,0,1,.387-.67Q9.552,5.657,14.974,2.535c.322-.184.638-.379.974-.535" style="fill:#019639"/><path d="M8.767,10.538q0,5.429,0,10.859a1.509,1.509,0,0,0,.427,1.087,1.647,1.647,0,0,0,2.06.206,1.564,1.564,0,0,0,.685-1.293c0-2.62-.005-5.24,0-7.86q3.583,4.29,7.181,8.568a2.833,2.833,0,0,0,2.6.782,1.561,1.561,0,0,0,1.251-1.371q.008-5.541,0-11.081a1.582,1.582,0,0,0-3.152,0c0,2.662-.016,5.321,0,7.982-2.346-2.766-4.663-5.556-7-8.332A2.817,2.817,0,0,0,10.17,9.033,1.579,1.579,0,0,0,8.767,10.538Z" style="fill:#fff"/></svg>`,"https://stackoverflow.com/"},
    {"exec: \"gcc\": executable file not found in %PATH% when trying go build" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_html</title><polygon points="5.902 27.201 3.655 2 28.345 2 26.095 27.197 15.985 30 5.902 27.201" style="fill:#e44f26"/><polygon points="16 27.858 24.17 25.593 26.092 4.061 16 4.061 16 27.858" style="fill:#f1662a"/><polygon points="16 13.407 11.91 13.407 11.628 10.242 16 10.242 16 7.151 15.989 7.151 8.25 7.151 8.324 7.981 9.083 16.498 16 16.498 16 13.407" style="fill:#ebebeb"/><polygon points="16 21.434 15.986 21.438 12.544 20.509 12.324 18.044 10.651 18.044 9.221 18.044 9.654 22.896 15.986 24.654 16 24.65 16 21.434" style="fill:#ebebeb"/><polygon points="15.989 13.407 15.989 16.498 19.795 16.498 19.437 20.507 15.989 21.437 15.989 24.653 22.326 22.896 22.372 22.374 23.098 14.237 23.174 13.407 22.341 13.407 15.989 13.407" style="fill:#fff"/><polygon points="15.989 7.151 15.989 9.071 15.989 10.235 15.989 10.242 23.445 10.242 23.445 10.242 23.455 10.242 23.517 9.548 23.658 7.981 23.732 7.151 15.989 7.151" style="fill:#fff"/></svg>`,"https://stackoverflow.com/"},
    {"Go install fails with error: no install location for directory xxx outside GOPATH" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_flutter</title><polyline points="15.383 18.316 18.744 15.042 27.093 15.042 19.697 22.438 15.383 18.316 15.383 18.316 15.383 18.316 15.383 18.316 15.383 18.316" style="fill:#40d0fd"/><polygon points="4.907 16.125 9.106 20.424 27.093 2.287 18.744 2.287 4.907 16.125" style="fill:#41d0fd;isolation:isolate"/><polygon points="11.176 22.479 15.435 26.675 19.697 22.438 15.383 18.316 11.176 22.479" style="fill:#1fbcfd"/><polygon points="15.435 26.675 19.697 22.438 26.989 29.813 18.593 29.813 15.435 26.675" style="fill:#095a9d"/><polygon points="15.435 26.675 19.406 25.354 18.068 24.057 15.435 26.675" style="fill:#0e5199"/></svg>`,"https://stackoverflow.com/"},
    {"Function declaration syntax: things in parenthesis before function name" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_docker</title><path d="M18.191,13.071H20.7v2.566H21.97a5.5,5.5,0,0,0,1.744-.292,4.462,4.462,0,0,0,.848-.383,3.149,3.149,0,0,1-.589-1.623,3.427,3.427,0,0,1,.616-2.416l.264-.305.314.253a4,4,0,0,1,1.575,2.538,3.837,3.837,0,0,1,2.913.271l.345.2-.181.354a3.629,3.629,0,0,1-3.648,1.74c-2.173,5.413-6.9,7.976-12.642,7.976A7.958,7.958,0,0,1,6.3,20.211l-.025-.043-.226-.459a7.28,7.28,0,0,1-.579-3.693l.035-.38H7.648V13.071h2.51v-2.51h5.02V8.051h3.012v5.02Z" style="fill:#3a4e55"/><path d="M26.324,14.021A3.311,3.311,0,0,0,24.906,11.2a3.072,3.072,0,0,0,.289,3.821,5.279,5.279,0,0,1-3.225,1.037H5.883a6.779,6.779,0,0,0,.667,3.737l.183.335a6.2,6.2,0,0,0,.379.569h0q.992.064,1.829.045h0a8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.4,8.4,0,0,1-1.845.3c.044,0-.046.007-.046.007l-.082.007c-.291.016-.6.02-.925.02-.351,0-.7-.007-1.083-.026l-.01.007a7.882,7.882,0,0,0,6.063,2.41c5.56,0,10.276-2.465,12.365-8,1.482.152,2.906-.226,3.553-1.49a3.5,3.5,0,0,0-3.122-.022" style="fill:#00aada"/><path d="M26.324,14.021A3.311,3.311,0,0,0,24.906,11.2a3.072,3.072,0,0,0,.289,3.821,5.279,5.279,0,0,1-3.225,1.037H6.836a5.223,5.223,0,0,0,2.106,4.686h0a8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.83,8.83,0,0,1-1.894.314L9.543,21.1c1.892.971,4.636.967,7.782-.241a21.868,21.868,0,0,0,9.1-6.889l-.1.048" style="fill:#27b9ec"/><path d="M5.913,17.732a6.431,6.431,0,0,0,.637,2.061l.183.335a6.2,6.2,0,0,0,.379.569q.992.064,1.829.045a8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.826,8.826,0,0,1-1.891.307l-.1,0c-.291.016-.6.026-.922.026-.351,0-.709-.007-1.1-.026a7.913,7.913,0,0,0,6.076,2.413c4.76,0,8.9-1.807,11.3-5.8Z" style="fill:#088cb9"/><path d="M6.98,17.732a4.832,4.832,0,0,0,1.961,3.01,8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.959,8.959,0,0,1-1.9.307c1.892.971,4.628.957,7.773-.252a20.545,20.545,0,0,0,5.377-3.13Z" style="fill:#039cc7"/><path d="M9.889,13.671h.172v1.813H9.889V13.671Zm-.33,0h.179v1.813H9.559V13.671Zm-.33,0h.179v1.813H9.23V13.671Zm-.33,0h.179v1.813H8.9V13.671Zm-.33,0h.179v1.813H8.57V13.671Zm-.323,0h.172v1.813H8.248V13.671Zm-.181-.181h2.175v2.176H8.066V13.49Z" style="fill:#00acd3"/><path d="M12.4,11.161h.172v1.813H12.4V11.161Zm-.33,0h.179v1.813H12.07V11.161Zm-.33,0h.179v1.813H11.74V11.161Zm-.33,0h.179v1.813H11.41V11.161Zm-.33,0h.178v1.813h-.178V11.161Zm-.323,0h.172v1.813h-.172V11.161Zm-.181-.181h2.176v2.176H10.577V10.979Z" style="fill:#00acd3"/><path d="M12.4,13.671h.172v1.813H12.4V13.671Zm-.33,0h.179v1.813H12.07V13.671Zm-.33,0h.179v1.813H11.74V13.671Zm-.33,0h.179v1.813H11.41V13.671Zm-.33,0h.178v1.813h-.178V13.671Zm-.323,0h.172v1.813h-.172V13.671Zm-.181-.181h2.176v2.176H10.577V13.49Z" style="fill:#26c2ee"/><path d="M14.909,13.671h.172v1.813h-.172V13.671Zm-.33,0h.179v1.813H14.58V13.671Zm-.33,0h.179v1.813H14.25V13.671Zm-.33,0H14.1v1.813h-.179V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.323,0h.172v1.813h-.172V13.671Zm-.181-.181h2.176v2.176H13.087V13.49Z" style="fill:#00acd3"/><path d="M14.909,11.161h.172v1.813h-.172V11.161Zm-.33,0h.179v1.813H14.58V11.161Zm-.33,0h.179v1.813H14.25V11.161Zm-.33,0H14.1v1.813h-.179V11.161Zm-.33,0h.179v1.813h-.179V11.161Zm-.323,0h.172v1.813h-.172V11.161Zm-.181-.181h2.176v2.176H13.087V10.979Z" style="fill:#26c2ee"/><path d="M17.42,13.671h.172v1.813H17.42V13.671Zm-.33,0h.179v1.813H17.09V13.671Zm-.33,0h.179v1.813H16.76V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.33,0h.179v1.813H16.1V13.671Zm-.323,0h.172v1.813h-.172V13.671ZM15.6,13.49h2.176v2.176H15.6V13.49Z" style="fill:#26c2ee"/><path d="M17.42,11.161h.172v1.813H17.42V11.161Zm-.33,0h.179v1.813H17.09V11.161Zm-.33,0h.179v1.813H16.76V11.161Zm-.33,0h.179v1.813h-.179V11.161Zm-.33,0h.179v1.813H16.1V11.161Zm-.323,0h.172v1.813h-.172V11.161Zm-.181-.181h2.176v2.176H15.6V10.979Z" style="fill:#00acd3"/><path d="M17.42,8.65h.172v1.813H17.42V8.65Zm-.33,0h.179v1.813H17.09V8.65Zm-.33,0h.179v1.813H16.76V8.65Zm-.33,0h.179v1.813h-.179V8.65Zm-.33,0h.179v1.813H16.1V8.65Zm-.323,0h.172v1.813h-.172V8.65ZM15.6,8.469h2.176v2.176H15.6V8.469Z" style="fill:#26c2ee"/><path d="M19.93,13.671H20.1v1.813H19.93V13.671Zm-.33,0h.178v1.813H19.6V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.323,0h.172v1.813h-.172V13.671Zm-.181-.181h2.176v2.176H18.107V13.49Z" style="fill:#00acd3"/><path d="M12.616,19.193a.6.6,0,1,1-.6.6.6.6,0,0,1,.6-.6" style="fill:#d5eef2"/><path d="M12.616,19.363a.431.431,0,0,1,.156.029.175.175,0,1,0,.241.236.43.43,0,1,1-.4-.265" style="fill:#3a4e55"/><path d="M2,17.949H29.92c-.608-.154-1.923-.362-1.707-1.159-1.105,1.279-3.771.9-4.444.267-.749,1.087-5.111.674-5.415-.173-.939,1.1-3.85,1.1-4.789,0-.3.847-4.666,1.26-5.415.173-.673.631-3.338,1.012-4.444-.267.217.8-1.1,1.005-1.707,1.159" style="fill:#3a4e55"/><path d="M14.211,23.518a5.287,5.287,0,0,1-2.756-2.711,9.2,9.2,0,0,1-1.987.3q-.436.024-.917.025-.554,0-1.168-.033a7.942,7.942,0,0,0,6.145,2.43q.344,0,.683-.013" style="fill:#c0dbe1"/><path d="M12.007,21.773a5.206,5.206,0,0,1-.552-.966,9.2,9.2,0,0,1-1.987.3,6.325,6.325,0,0,0,2.539.664" style="fill:#d5eef2"/></svg>`,"https://stackoverflow.com/"},
	{" How can I convert a zero-terminated byte array to string? " ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" ><path d="M15.9.476a2.14 2.14 0 0 0-.823.218L3.932 6.01c-.582.277-1.005.804-1.15 1.432L.054 19.373c-.13.56-.025 1.147.3 1.627q.057.087.12.168l7.7 9.574c.407.5 1.018.787 1.662.784h12.35c.646.001 1.258-.3 1.664-.793l7.696-9.576c.404-.5.555-1.16.4-1.786L29.2 7.43c-.145-.628-.57-1.155-1.15-1.432L16.923.695A2.14 2.14 0 0 0 15.89.476z" fill="#326ce5"/><path d="M16.002 4.542c-.384.027-.675.356-.655.74v.188c.018.213.05.424.092.633a6.22 6.22 0 0 1 .066 1.21c-.038.133-.114.253-.218.345l-.015.282c-.405.034-.807.096-1.203.186-1.666.376-3.183 1.24-4.354 2.485l-.24-.17c-.132.04-.274.025-.395-.04a6.22 6.22 0 0 1-.897-.81 5.55 5.55 0 0 0-.437-.465l-.148-.118c-.132-.106-.294-.167-.463-.175a.64.64 0 0 0-.531.236c-.226.317-.152.756.164.983l.138.11a5.55 5.55 0 0 0 .552.323c.354.197.688.428.998.7a.74.74 0 0 1 .133.384l.218.2c-1.177 1.766-1.66 3.905-1.358 6.006l-.28.08c-.073.116-.17.215-.286.288a6.22 6.22 0 0 1-1.194.197 5.57 5.57 0 0 0-.64.05l-.177.04h-.02a.67.67 0 0 0-.387 1.132.67.67 0 0 0 .684.165h.013l.18-.02c.203-.06.403-.134.598-.218.375-.15.764-.265 1.162-.34.138.008.27.055.382.135l.3-.05c.65 2.017 2.016 3.726 3.84 4.803l-.122.255c.056.117.077.247.06.376-.165.382-.367.748-.603 1.092a5.58 5.58 0 0 0-.358.533l-.085.18a.67.67 0 0 0 .65 1.001.67.67 0 0 0 .553-.432l.083-.17c.076-.2.14-.404.192-.61.177-.437.273-.906.515-1.196a.54.54 0 0 1 .286-.14l.15-.273a8.62 8.62 0 0 0 6.146.015l.133.255c.136.02.258.095.34.205.188.358.34.733.456 1.12a5.57 5.57 0 0 0 .194.611l.083.17a.67.67 0 0 0 1.187.131.67.67 0 0 0 .016-.701l-.087-.18a5.55 5.55 0 0 0-.358-.531c-.23-.332-.428-.686-.6-1.057a.52.52 0 0 1 .068-.4 2.29 2.29 0 0 1-.111-.269c1.82-1.085 3.18-2.8 3.823-4.82l.284.05c.102-.093.236-.142.373-.138.397.076.786.2 1.162.34.195.09.395.166.598.23.048.013.118.024.172.037h.013a.67.67 0 0 0 .841-.851.67.67 0 0 0-.544-.446l-.194-.046a5.57 5.57 0 0 0-.64-.05c-.404-.026-.804-.092-1.194-.197-.12-.067-.22-.167-.288-.288l-.27-.08a8.65 8.65 0 0 0-1.386-5.993l.236-.218c-.01-.137.035-.273.124-.378.307-.264.64-.497.99-.696a5.57 5.57 0 0 0 .552-.323l.146-.118a.67.67 0 0 0-.133-1.202.67.67 0 0 0-.696.161l-.148.118a5.57 5.57 0 0 0-.437.465c-.264.302-.556.577-.873.823a.74.74 0 0 1-.404.044l-.253.18c-1.46-1.53-3.427-2.48-5.535-2.67 0-.1-.013-.25-.015-.297-.113-.078-.192-.197-.218-.332a6.23 6.23 0 0 1 .076-1.207c.043-.21.073-.42.092-.633v-.2c.02-.384-.27-.713-.655-.74zm-.834 5.166l-.2 3.493h-.015c-.01.216-.137.4-.332.504s-.426.073-.6-.054l-2.865-2.03a6.86 6.86 0 0 1 3.303-1.799c.234-.05.47-.088.707-.114zm1.668 0c1.505.187 2.906.863 3.99 1.924l-2.838 2.017c-.175.14-.415.168-.618.072s-.333-.3-.336-.524zm-6.72 3.227l2.62 2.338v.015c.163.142.234.363.186.574s-.21.378-.417.435v.01l-3.362.967a6.86 6.86 0 0 1 .974-4.34zm11.753 0c.796 1.295 1.148 2.814 1.002 4.327l-3.367-.97v-.013c-.21-.057-.37-.224-.417-.435s.023-.43.186-.574l2.6-2.327zm-6.404 2.52h1.072l.655.832-.238 1.04-.963.463-.965-.463-.227-1.04zm3.434 2.838c.045-.005.1-.005.135 0l3.467.585c-.5 1.44-1.487 2.67-2.775 3.493l-1.34-3.244a.59.59 0 0 1 .509-.819zm-5.823.015c.196.003.377.104.484.268s.124.37.047.55v.013l-1.332 3.218C11 21.54 10.032 20.325 9.517 18.9l3.437-.583c.038-.004.077-.004.116 0zm2.904 1.4a.59.59 0 0 1 .537.308h.013l1.694 3.057-.677.2c-1.246.285-2.547.218-3.758-.194l1.7-3.057c.103-.18.293-.29.5-.295z" fill="#fff" stroke="#fff" stroke-width=".055"/></svg>`,"https://stackoverflow.com/"},
    {" go get results in 'terminal prompts disabled' error for github private repo " ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_vue</title><path d="M24.4,3.925H30L16,28.075,2,3.925H12.71L16,9.525l3.22-5.6Z" style="fill:#41b883"/><path d="M2,3.925l14,24.15L30,3.925H24.4L16,18.415,7.53,3.925Z" style="fill:#41b883"/><path d="M7.53,3.925,16,18.485l8.4-14.56H19.22L16,9.525l-3.29-5.6Z" style="fill:#35495e"/></svg>`,"https://stackoverflow.com/"},
    {"s it possible to capture a Ctrl+C signal (SIGINT) and run a cleanup function, in a \"defer\" fashion?" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_svelte</title><path d="M26.47,5.7A8.973,8.973,0,0,0,14.677,3.246L7.96,7.4a7.461,7.461,0,0,0-3.481,5.009,7.686,7.686,0,0,0,.8,5.058,7.358,7.358,0,0,0-1.151,2.8,7.789,7.789,0,0,0,1.4,6.028,8.977,8.977,0,0,0,11.794,2.458L24.04,24.6a7.468,7.468,0,0,0,3.481-5.009,7.673,7.673,0,0,0-.8-5.062,7.348,7.348,0,0,0,1.152-2.8A7.785,7.785,0,0,0,26.47,5.7" style="fill:#ff3e00"/><path d="M14.022,26.64A5.413,5.413,0,0,1,8.3,24.581a4.678,4.678,0,0,1-.848-3.625,4.307,4.307,0,0,1,.159-.61l.127-.375.344.238a8.76,8.76,0,0,0,2.628,1.274l.245.073-.025.237a1.441,1.441,0,0,0,.271.968,1.63,1.63,0,0,0,1.743.636,1.512,1.512,0,0,0,.411-.175l6.7-4.154a1.366,1.366,0,0,0,.633-.909,1.407,1.407,0,0,0-.244-1.091,1.634,1.634,0,0,0-1.726-.622,1.509,1.509,0,0,0-.413.176l-2.572,1.584a4.934,4.934,0,0,1-1.364.582,5.415,5.415,0,0,1-5.727-2.06A4.678,4.678,0,0,1,7.811,13.1,4.507,4.507,0,0,1,9.9,10.09l6.708-4.154a4.932,4.932,0,0,1,1.364-.581A5.413,5.413,0,0,1,23.7,7.414a4.679,4.679,0,0,1,.848,3.625,4.272,4.272,0,0,1-.159.61l-.127.375-.344-.237a8.713,8.713,0,0,0-2.628-1.274l-.245-.074.025-.237a1.438,1.438,0,0,0-.272-.968,1.629,1.629,0,0,0-1.725-.622,1.484,1.484,0,0,0-.411.176l-6.722,4.14a1.353,1.353,0,0,0-.631.908,1.394,1.394,0,0,0,.244,1.092,1.634,1.634,0,0,0,1.726.621,1.538,1.538,0,0,0,.413-.175l2.562-1.585a4.9,4.9,0,0,1,1.364-.581,5.417,5.417,0,0,1,5.728,2.059,4.681,4.681,0,0,1,.843,3.625A4.5,4.5,0,0,1,22.1,21.905l-6.707,4.154a4.9,4.9,0,0,1-1.364.581" style="fill:#fff"/></svg>`,"https://stackoverflow.com/"},
    {"What is the shortest way to simply sort an array of structs by (arbitrary) field names?" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_nginx</title><path d="M15.948,2h.065a10.418,10.418,0,0,1,.972.528Q22.414,5.65,27.843,8.774a.792.792,0,0,1,.414.788c-.008,4.389,0,8.777-.005,13.164a.813.813,0,0,1-.356.507q-5.773,3.324-11.547,6.644a.587.587,0,0,1-.657.037Q9.912,26.6,4.143,23.274a.7.7,0,0,1-.4-.666q0-6.582,0-13.163a.693.693,0,0,1,.387-.67Q9.552,5.657,14.974,2.535c.322-.184.638-.379.974-.535" style="fill:#019639"/><path d="M8.767,10.538q0,5.429,0,10.859a1.509,1.509,0,0,0,.427,1.087,1.647,1.647,0,0,0,2.06.206,1.564,1.564,0,0,0,.685-1.293c0-2.62-.005-5.24,0-7.86q3.583,4.29,7.181,8.568a2.833,2.833,0,0,0,2.6.782,1.561,1.561,0,0,0,1.251-1.371q.008-5.541,0-11.081a1.582,1.582,0,0,0-3.152,0c0,2.662-.016,5.321,0,7.982-2.346-2.766-4.663-5.556-7-8.332A2.817,2.817,0,0,0,10.17,9.033,1.579,1.579,0,0,0,8.767,10.538Z" style="fill:#fff"/></svg>`,"https://stackoverflow.com/"},
    {"exec: \"gcc\": executable file not found in %PATH% when trying go build" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_html</title><polygon points="5.902 27.201 3.655 2 28.345 2 26.095 27.197 15.985 30 5.902 27.201" style="fill:#e44f26"/><polygon points="16 27.858 24.17 25.593 26.092 4.061 16 4.061 16 27.858" style="fill:#f1662a"/><polygon points="16 13.407 11.91 13.407 11.628 10.242 16 10.242 16 7.151 15.989 7.151 8.25 7.151 8.324 7.981 9.083 16.498 16 16.498 16 13.407" style="fill:#ebebeb"/><polygon points="16 21.434 15.986 21.438 12.544 20.509 12.324 18.044 10.651 18.044 9.221 18.044 9.654 22.896 15.986 24.654 16 24.65 16 21.434" style="fill:#ebebeb"/><polygon points="15.989 13.407 15.989 16.498 19.795 16.498 19.437 20.507 15.989 21.437 15.989 24.653 22.326 22.896 22.372 22.374 23.098 14.237 23.174 13.407 22.341 13.407 15.989 13.407" style="fill:#fff"/><polygon points="15.989 7.151 15.989 9.071 15.989 10.235 15.989 10.242 23.445 10.242 23.445 10.242 23.455 10.242 23.517 9.548 23.658 7.981 23.732 7.151 15.989 7.151" style="fill:#fff"/></svg>`,"https://stackoverflow.com/"},
    {"Go install fails with error: no install location for directory xxx outside GOPATH" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_flutter</title><polyline points="15.383 18.316 18.744 15.042 27.093 15.042 19.697 22.438 15.383 18.316 15.383 18.316 15.383 18.316 15.383 18.316 15.383 18.316" style="fill:#40d0fd"/><polygon points="4.907 16.125 9.106 20.424 27.093 2.287 18.744 2.287 4.907 16.125" style="fill:#41d0fd;isolation:isolate"/><polygon points="11.176 22.479 15.435 26.675 19.697 22.438 15.383 18.316 11.176 22.479" style="fill:#1fbcfd"/><polygon points="15.435 26.675 19.697 22.438 26.989 29.813 18.593 29.813 15.435 26.675" style="fill:#095a9d"/><polygon points="15.435 26.675 19.406 25.354 18.068 24.057 15.435 26.675" style="fill:#0e5199"/></svg>`,"https://stackoverflow.com/"},
    {"Function declaration syntax: things in parenthesis before function name" ,`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><title>file_type_docker</title><path d="M18.191,13.071H20.7v2.566H21.97a5.5,5.5,0,0,0,1.744-.292,4.462,4.462,0,0,0,.848-.383,3.149,3.149,0,0,1-.589-1.623,3.427,3.427,0,0,1,.616-2.416l.264-.305.314.253a4,4,0,0,1,1.575,2.538,3.837,3.837,0,0,1,2.913.271l.345.2-.181.354a3.629,3.629,0,0,1-3.648,1.74c-2.173,5.413-6.9,7.976-12.642,7.976A7.958,7.958,0,0,1,6.3,20.211l-.025-.043-.226-.459a7.28,7.28,0,0,1-.579-3.693l.035-.38H7.648V13.071h2.51v-2.51h5.02V8.051h3.012v5.02Z" style="fill:#3a4e55"/><path d="M26.324,14.021A3.311,3.311,0,0,0,24.906,11.2a3.072,3.072,0,0,0,.289,3.821,5.279,5.279,0,0,1-3.225,1.037H5.883a6.779,6.779,0,0,0,.667,3.737l.183.335a6.2,6.2,0,0,0,.379.569h0q.992.064,1.829.045h0a8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.4,8.4,0,0,1-1.845.3c.044,0-.046.007-.046.007l-.082.007c-.291.016-.6.02-.925.02-.351,0-.7-.007-1.083-.026l-.01.007a7.882,7.882,0,0,0,6.063,2.41c5.56,0,10.276-2.465,12.365-8,1.482.152,2.906-.226,3.553-1.49a3.5,3.5,0,0,0-3.122-.022" style="fill:#00aada"/><path d="M26.324,14.021A3.311,3.311,0,0,0,24.906,11.2a3.072,3.072,0,0,0,.289,3.821,5.279,5.279,0,0,1-3.225,1.037H6.836a5.223,5.223,0,0,0,2.106,4.686h0a8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.83,8.83,0,0,1-1.894.314L9.543,21.1c1.892.971,4.636.967,7.782-.241a21.868,21.868,0,0,0,9.1-6.889l-.1.048" style="fill:#27b9ec"/><path d="M5.913,17.732a6.431,6.431,0,0,0,.637,2.061l.183.335a6.2,6.2,0,0,0,.379.569q.992.064,1.829.045a8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.826,8.826,0,0,1-1.891.307l-.1,0c-.291.016-.6.026-.922.026-.351,0-.709-.007-1.1-.026a7.913,7.913,0,0,0,6.076,2.413c4.76,0,8.9-1.807,11.3-5.8Z" style="fill:#088cb9"/><path d="M6.98,17.732a4.832,4.832,0,0,0,1.961,3.01,8.972,8.972,0,0,0,2.669-.389.193.193,0,1,1,.126.365c-.09.031-.184.061-.281.088h0a8.959,8.959,0,0,1-1.9.307c1.892.971,4.628.957,7.773-.252a20.545,20.545,0,0,0,5.377-3.13Z" style="fill:#039cc7"/><path d="M9.889,13.671h.172v1.813H9.889V13.671Zm-.33,0h.179v1.813H9.559V13.671Zm-.33,0h.179v1.813H9.23V13.671Zm-.33,0h.179v1.813H8.9V13.671Zm-.33,0h.179v1.813H8.57V13.671Zm-.323,0h.172v1.813H8.248V13.671Zm-.181-.181h2.175v2.176H8.066V13.49Z" style="fill:#00acd3"/><path d="M12.4,11.161h.172v1.813H12.4V11.161Zm-.33,0h.179v1.813H12.07V11.161Zm-.33,0h.179v1.813H11.74V11.161Zm-.33,0h.179v1.813H11.41V11.161Zm-.33,0h.178v1.813h-.178V11.161Zm-.323,0h.172v1.813h-.172V11.161Zm-.181-.181h2.176v2.176H10.577V10.979Z" style="fill:#00acd3"/><path d="M12.4,13.671h.172v1.813H12.4V13.671Zm-.33,0h.179v1.813H12.07V13.671Zm-.33,0h.179v1.813H11.74V13.671Zm-.33,0h.179v1.813H11.41V13.671Zm-.33,0h.178v1.813h-.178V13.671Zm-.323,0h.172v1.813h-.172V13.671Zm-.181-.181h2.176v2.176H10.577V13.49Z" style="fill:#26c2ee"/><path d="M14.909,13.671h.172v1.813h-.172V13.671Zm-.33,0h.179v1.813H14.58V13.671Zm-.33,0h.179v1.813H14.25V13.671Zm-.33,0H14.1v1.813h-.179V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.323,0h.172v1.813h-.172V13.671Zm-.181-.181h2.176v2.176H13.087V13.49Z" style="fill:#00acd3"/><path d="M14.909,11.161h.172v1.813h-.172V11.161Zm-.33,0h.179v1.813H14.58V11.161Zm-.33,0h.179v1.813H14.25V11.161Zm-.33,0H14.1v1.813h-.179V11.161Zm-.33,0h.179v1.813h-.179V11.161Zm-.323,0h.172v1.813h-.172V11.161Zm-.181-.181h2.176v2.176H13.087V10.979Z" style="fill:#26c2ee"/><path d="M17.42,13.671h.172v1.813H17.42V13.671Zm-.33,0h.179v1.813H17.09V13.671Zm-.33,0h.179v1.813H16.76V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.33,0h.179v1.813H16.1V13.671Zm-.323,0h.172v1.813h-.172V13.671ZM15.6,13.49h2.176v2.176H15.6V13.49Z" style="fill:#26c2ee"/><path d="M17.42,11.161h.172v1.813H17.42V11.161Zm-.33,0h.179v1.813H17.09V11.161Zm-.33,0h.179v1.813H16.76V11.161Zm-.33,0h.179v1.813h-.179V11.161Zm-.33,0h.179v1.813H16.1V11.161Zm-.323,0h.172v1.813h-.172V11.161Zm-.181-.181h2.176v2.176H15.6V10.979Z" style="fill:#00acd3"/><path d="M17.42,8.65h.172v1.813H17.42V8.65Zm-.33,0h.179v1.813H17.09V8.65Zm-.33,0h.179v1.813H16.76V8.65Zm-.33,0h.179v1.813h-.179V8.65Zm-.33,0h.179v1.813H16.1V8.65Zm-.323,0h.172v1.813h-.172V8.65ZM15.6,8.469h2.176v2.176H15.6V8.469Z" style="fill:#26c2ee"/><path d="M19.93,13.671H20.1v1.813H19.93V13.671Zm-.33,0h.178v1.813H19.6V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.33,0h.179v1.813h-.179V13.671Zm-.323,0h.172v1.813h-.172V13.671Zm-.181-.181h2.176v2.176H18.107V13.49Z" style="fill:#00acd3"/><path d="M12.616,19.193a.6.6,0,1,1-.6.6.6.6,0,0,1,.6-.6" style="fill:#d5eef2"/><path d="M12.616,19.363a.431.431,0,0,1,.156.029.175.175,0,1,0,.241.236.43.43,0,1,1-.4-.265" style="fill:#3a4e55"/><path d="M2,17.949H29.92c-.608-.154-1.923-.362-1.707-1.159-1.105,1.279-3.771.9-4.444.267-.749,1.087-5.111.674-5.415-.173-.939,1.1-3.85,1.1-4.789,0-.3.847-4.666,1.26-5.415.173-.673.631-3.338,1.012-4.444-.267.217.8-1.1,1.005-1.707,1.159" style="fill:#3a4e55"/><path d="M14.211,23.518a5.287,5.287,0,0,1-2.756-2.711,9.2,9.2,0,0,1-1.987.3q-.436.024-.917.025-.554,0-1.168-.033a7.942,7.942,0,0,0,6.145,2.43q.344,0,.683-.013" style="fill:#c0dbe1"/><path d="M12.007,21.773a5.206,5.206,0,0,1-.552-.966,9.2,9.2,0,0,1-1.987.3,6.325,6.325,0,0,0,2.539.664" style="fill:#d5eef2"/></svg>`,"https://stackoverflow.com/"},

	}
		c.JSON(http.StatusOK,x)
}