<script lang="ts">
	import { goto } from '$app/navigation';
	import K8s from '$lib/Navbar/svgs/k8s.svelte';
	import Following from '$lib/Navbar/svgs/following.svelte';
	import Spaces from './svgs/Spaces.svelte';
	import QuestionForYou from './svgs/QuestionForYou.svelte';
	import Search from './svgs/Search.svelte';
	import GoldDot from './svgs/GoldDot.svelte';
	import SilverDot from './svgs/SilverDot.svelte';
	import BronzeDot from './svgs/BronzeDot.svelte';
	import NotificationIcon from './svgs/NotificationIcon.svelte';
	import Svelte from './svgs/svelte.svelte';
	import Flutter from './svgs/Flutter.svelte';
	import Vue from './svgs/vue.svelte';
	import Menubar from './svgs/Menubar.svelte';
	import WhiteDot from './svgs/WhiteDot.svelte';
	import GPodcast from '$lib/icons/gPodcast.svelte';
	import Nullpointer from '$lib/icons/null pointer.svelte';
	import Stedia from '$lib/icons/stedia.svelte';
	import StediaRed from '$lib/icons/stediaRed.svelte';
	import Stediablue from '$lib/icons/stediablue.svelte';

	let UserdataLoading: boolean = true;

	let UserData: {
		UserID: string;
		UserName: string;
		Email: string;
		Password: string;
		UserImage: string;
		Badges: {
			Reputation: number;
			Gold: number;
			Silver: number;
			Bronze: number;
		};
		Follower: string[];
		Following: string[];
		Location: string;
		MembershipTime: string;
		LastSeen: string;
		Aboutme: string;
		Mysite: string;
		Github: string;
		Twitter: string;
		Linkedin: string;
		TopTags: {
			Name: string;
			Score: number;
			NumberOfPost: number;
		}[];

		SelectedPanel: string;
		AccountType: string;
	} = {
		UserID: 'skraka',
		UserName: 'Sk Shahriar Ahmed Raka',
		Email: 'skshahra@gmail.com',
		Password: '123456',
		// UserTitle string `json:"UserTitle"`
		UserImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
		Badges: {
			Reputation: 13452543,
			Gold: 999,
			Silver: 888,
			Bronze: 777
		},
		Follower: ['RKA', 'SHAHRIAR	'],
		Following: ['RKA', 'SHAHRIAR'],
		// Badges map[string]int
		Location: 'Dhaka, Bangladesh',
		MembershipTime: '3 year 5 Month',
		LastSeen: 'This Week',
		Aboutme:
			'A Curious Learner, Full-Stack Web Developer, Security Researcher\nHere are my skills and strengths:\n✓ Expert in Golang\n ✓ Expert in Fiber framework (using Golang) \n ✓ Expert in WebAssembly (using Golang)  Expert in Golang     ✓ Expert in Fiber framework (using Golang)    ✓ Expert in WebAssembly (using Golang) ✓ Expert in database design, development, optimization, and migration    (PostgreSQL, MySQL, MongoDB , Redis) ✓ Expert in ( Grpc, protocol buffer ) ✓ Experienced in ( WebSocket, Socket.io, WebRTC ) for real-time client and server applications ✓ Experienced in ( Svelte.js ) and some knowledge in ( TypeScript ) ✓ Good understanding of ( Docker, Bash, PowerShell, Git,    Nginx, Kubernetes )        Github : github.com/skshahriarahmedraka    Upwork : upwork.com/o/profiles/users/~0107ef3405bffbe57e    Linkedin : linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/  Telegram : t.me/shahriarraka ',
		Mysite: 'www.shahriarraka.me',
		Github: 'www.github.com/skshahriarahmedraka',
		Twitter: 'www.twitter.com/shahriarraka7',
		Linkedin: 'www.linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/',
		// "TopTags"    : ["Go","Rust","Python","Svelte","PostgreSQL"],
		TopTags: [
			{ Name: 'Go', Score: 12, NumberOfPost: 4 },
			{ Name: 'Rust', Score: 10, NumberOfPost: 6 },
			{ Name: 'Python', Score: 7, NumberOfPost: 4 },
			{ Name: 'Svelte', Score: 12, NumberOfPost: 4 },
			{ Name: 'PostgreSQL', Score: 12, NumberOfPost: 4 }
		],
		SelectedPanel: 'Profile',
		AccountType: 'regular'
	};

	let NavbarOpts = ['Profile', 'Followers', 'Followings', 'Other Networks', 'settings', 'logout'];
	let NotificationOpts = [
		'thanks for your ans',
		'new Question new Question new Question new Question new Question',
		'reply of your question ',
		'thanks for your ans',
		'new Question new Question new Question new Question new Question',
		'reply of your question ',
		'thanks for your ans',
		'new Question new Question new Question new Question new Question',
		'reply of your question ',
		'thanks for your ans',
		'new Question new Question new Question new Question new Question',
		'reply of your question ',
		'thanks for your ans',
		'new Question new Question new Question new Question new Question',
		'reply of your question ',
		'thanks for your ans',
		'new Question new Question new Question new Question new Question',
		'reply of your question ',
		'thanks for your ans',
		'new Question new Question new Question new Question new Question',
		'reply of your question ',
		'thanks for your ans',
		'new Question new Question new Question new Question new Question',
		'reply of your question '
	];

	let NotificationOpts2 = [
		[
			'answer',
			'5 jul at 1:23',
			'How to ensure translation with I18n is working in production?',
			'either add the pass keyword to an empty function, also function args cannot have .s in them only valid names and * or ** followed by a name'
		],
		[
			'comment',
			'29 aug at 3:12',
			'How build lib for 386 arch with cgo on windows?',
			" have a golang library that builds and works well on Linux, MacOs and Windows. The problem comes when I'm trying to build it for 386 on the amd64 Windows VM"
		],
		['reward', '29 aug at 3:12', 20, 'Golang pass sensitive data as argument using docker'],
		[
			'answer',
			'30 jan at 22:23',
			'How to access a method from the extended abstract class when using it as a type constraint in a method',
			"Except this doesn't work because it isn't assumed that T extends Animal so there I can't access it.What can I do make this getAnimalSound method working?"
		],
		[
			'comment',
			'29 aug at 3:12',
			'What I would do in your place is to use overriding. However, it also only works on non static methods. (untested but should work)',
			"Yeah, sorry I was a bit out of practice. However, you don't need a type parameter at all! You just say that the thing you want to have the sound from is an animal, doesn't matter which. If you give it a dog, dart will call the getSound method inside of the dog. I've updated my answer accordingly"
		],
		['reward', '29 aug at 3:12', -5, 'https and websocket handler with different deadlines'],
		[
			'reward',
			'29 aug at 3:12',
			-20,
			'How can I add entire Array of structures into mysql database using GORM'
		],
		[
			'comment',
			'24 mar at 2:43',
			'Encoding with image/jpeg cause image saturation / wrong pixels',
			'I don\'t speak "Go" at all, but you appear to be trying to get and set RGBA values on JPEG images '
		],
		[
			'comment',
			'23 sep at 23;54',
			"Can't install golangci-lint locally",
			'I am dockerized a golang application and i am trying to access the application.'
		],
		[
			'answer',
			'5 jul at 1:23',
			'How to ensure translation with I18n is working in production?',
			'either add the pass keyword to an empty function, also function args cannot have .s in them only valid names and * or ** followed by a name'
		],
		[
			'comment',
			'29 aug at 3:12',
			'How build lib for 386 arch with cgo on windows?',
			" have a golang library that builds and works well on Linux, MacOs and Windows. The problem comes when I'm trying to build it for 386 on the amd64 Windows VM"
		],
		['reward', '29 aug at 3:12', 20, 'Golang pass sensitive data as argument using docker'],
		[
			'answer',
			'30 jan at 22:23',
			'How to access a method from the extended abstract class when using it as a type constraint in a method',
			"Except this doesn't work because it isn't assumed that T extends Animal so there I can't access it.What can I do make this getAnimalSound method working?"
		],
		[
			'comment',
			'29 aug at 3:12',
			'What I would do in your place is to use overriding. However, it also only works on non static methods. (untested but should work)',
			"Yeah, sorry I was a bit out of practice. However, you don't need a type parameter at all! You just say that the thing you want to have the sound from is an animal, doesn't matter which. If you give it a dog, dart will call the getSound method inside of the dog. I've updated my answer accordingly"
		],
		['reward', '29 aug at 3:12', -5, 'https and websocket handler with different deadlines'],
		[
			'reward',
			'29 aug at 3:12',
			-20,
			'How can I add entire Array of structures into mysql database using GORM'
		],
		[
			'comment',
			'24 mar at 2:43',
			'Encoding with image/jpeg cause image saturation / wrong pixels',
			'I don\'t speak "Go" at all, but you appear to be trying to get and set RGBA values on JPEG images '
		],
		[
			'comment',
			'23 sep at 23;54',
			"Can't install golangci-lint locally",
			'I am dockerized a golang application and i am trying to access the application.'
		]
	];

	let NavbarOptsOpen = false;
	let NotificationOptsOpen = false;

	//   $: (NavbarOptsOpen ? (NotificationOptsOpen=false) : (NotificationOptsOpen=true))
	function AllOtherFalse() {
		NavbarOptsOpen = false;
		NotificationOptsOpen = false;
	}

	function ShortenNumber(i: number) {
		let bigFig: string[] = ['k', 'M', 'B', 'T'];
		let s: string = '';
		for (let x = 0; i > 999; x++) {
			if (i > 999) {
				i = i / 1000;
				s = bigFig[x];
			}
		}
		return String(i.toPrecision(3) + s);
	}
</script>

{#if UserdataLoading}
	<div
		class=" top-0 flex h-[53px] w-screen  flex-row items-center justify-center overflow-hidden border-b-2 border-solid border-[#32353a] bg-[#262626]  "
	>
		<div class=" grow" />
		<!-- Company logo -->
		<button
			class="flex h-[53px] w-fit flex-row gap-1 items-center justify-center rounded-lg px-3  hover:bg-slate-800 "
			on:click={() => {
				goto('/');
			}}
		>
			<Stediablue class=" h-14  w-20  " />
			<!-- <GPodcast  /> -->
			<Nullpointer  class="h-7" />
			<!-- <p class="font-serif text-xl leading-5 text-sky-500">Null Pointer</p> -->
		</button>

		<!-- following -->
		<button class="rounded-lg px-2 hover:bg-slate-800" on:click={() => goto('/f')}>
			<Following />
			<!-- <svg class="h-16 w-16 -mt-1 " width="210mm" height="297mm" viewBox="0 0 210 297" version="1.1" id="svg5" inkscape:version="1.1.2 (1:1.1+202205011109+08b2f3d93c)" sodipodi:docname="following2.svg" xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape" xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd" xmlns="http://www.w3.org/2000/svg" xmlns:svg="http://www.w3.org/2000/svg"> <sodipodi:namedview id="namedview7" pagecolor="#ffffff" bordercolor="#666666" borderopacity="1.0" inkscape:pageshadow="2" inkscape:pageopacity="0.0" inkscape:pagecheckerboard="0" inkscape:document-units="mm" showgrid="false" inkscape:zoom="0.60221661" inkscape:cx="396.0369" inkscape:cy="552.95718" inkscape:window-width="1920" inkscape:window-height="1022" inkscape:window-x="0" inkscape:window-y="27" inkscape:window-maximized="1" inkscape:current-layer="layer1" /> <defs id="defs2" /> <g inkscape:label="Layer 1" inkscape:groupmode="layer" id="layer1"> <path clip-rule="evenodd" d="m 166.14754,181.72922 c 0,9.34958 -7.5819,16.92963 -16.93334,16.92963 h -8.46693 c -9.3554,0 -16.93545,-7.58005 -16.93545,-16.92963 0,-2.07936 -0.10954,-4.12988 -0.27067,-6.16638 -0.0288,-0.34105 -0.0413,-0.68421 -0.0743,-1.02103 -0.18203,-1.93066 -0.44238,-3.84466 -0.76491,-5.73431 -0.0619,-0.36989 -0.13229,-0.72972 -0.19844,-1.09564 -0.35771,-1.90368 -0.76888,-3.78672 -1.26497,-5.64304 -0.0704,-0.27067 -0.15081,-0.52705 -0.22728,-0.79375 -0.53128,-1.91241 -1.11628,-3.80153 -1.79017,-5.65362 -0.008,-0.027 -0.0228,-0.0601 -0.0331,-0.0889 -0.67786,-1.86028 -1.45097,-3.67956 -2.27383,-5.46973 -0.11589,-0.26035 -0.22331,-0.52096 -0.34528,-0.77734 -0.80857,-1.70948 -1.6809,-3.37556 -2.60694,-5.01254 -0.16113,-0.28945 -0.32649,-0.57467 -0.49186,-0.86439 -0.95091,-1.62666 -1.95553,-3.22262 -3.01995,-4.7707 -0.14685,-0.21299 -0.29978,-0.42572 -0.4445,-0.63447 -1.10808,-1.57718 -2.2614,-3.11944 -3.48086,-4.6056 -0.0331,-0.0413 -0.0743,-0.0889 -0.10742,-0.1323 -1.22582,-1.48008 -2.51355,-2.90009 -3.8407,-4.28492 -0.21907,-0.23363 -0.43418,-0.46514 -0.66145,-0.69453 -1.29408,-1.31869 -2.646896,-2.58578 -4.038075,-3.80551 -0.251089,-0.2233 -0.507471,-0.44661 -0.759619,-0.65934 -1.424252,-1.21549 -2.876285,-2.38733 -4.385204,-3.49329 -0.214048,-0.15716 -0.437091,-0.30401 -0.651139,-0.46117 -1.571096,-1.13057 -3.169973,-2.218 -4.831821,-3.22871 -0.02778,-0.0145 -0.05159,-0.0331 -0.0807,-0.0516 -1.63724,-0.99643 -3.332161,-1.91215 -5.044544,-2.79057 -0.309033,-0.15716 -0.607748,-0.31829 -0.916781,-0.4691 -1.646502,-0.81227 -3.326871,-1.55258 -5.040577,-2.24711 -0.360627,-0.14684 -0.721519,-0.29342 -1.087173,-0.43206 -1.741488,-0.6739 -3.507846,-1.28985 -5.311511,-1.83568 -0.280193,-0.0847 -0.569383,-0.16113 -0.849577,-0.24183 -6.763279,-1.96585 -13.907293,-3.04271 -21.307425,-3.04271 -4.674658,0 -8.466666,3.78698 -8.466666,8.4627 v 8.46905 c 0,4.67386 3.792273,8.46693 8.466666,8.46693 4.647671,0 9.132359,0.67786 13.409084,1.846 1.310481,0.36168 2.596091,0.76888 3.863445,1.22978 0.07647,0.0228 0.142611,0.0579 0.214048,0.0807 6.270361,2.30293 11.951759,5.80443 16.78358,10.22429 0.208756,0.19421 0.427831,0.37624 0.63156,0.57468 0.759619,0.71728 1.490398,1.46129 2.202392,2.2262 0.313266,0.33682 0.626268,0.6739 0.930275,1.02103 0.635529,0.7112 1.243277,1.43245 1.836473,2.18281 0.351366,0.44662 0.688445,0.91149 1.02526,1.37266 0.532342,0.72575 1.072885,1.44304 1.566862,2.19128 0.360628,0.55192 0.678921,1.12024 1.019969,1.68063 0.664634,1.10173 1.290902,2.22197 1.870604,3.37556 0.426773,0.84534 0.853811,1.6891 1.234017,2.56328 0.265642,0.6014 0.48895,1.2237 0.730779,1.84177 0.318294,0.82073 0.631561,1.64121 0.906463,2.48258 0.208756,0.62627 0.3937,1.25677 0.574675,1.88939 0.256381,0.8972 0.48895,1.80684 0.697706,2.71833 0.137583,0.59955 0.274902,1.19671 0.388673,1.80869 0.195262,1.01706 0.337873,2.05052 0.470164,3.08002 0.06721,0.54768 0.151871,1.0832 0.204523,1.63301 0.09393,0.99853 0.113771,2.01347 0.151871,3.02445 0.06085,0.57891 0.08996,1.16364 0.08996,1.7571 0,9.34958 -7.579784,16.92962 -16.93545,16.92962 h -6.351059 -2.116666 c -9.350375,0 -16.931217,-7.58004 -16.931217,-16.92962 0,-4.67599 -3.792273,-8.46905 -8.467725,-8.46905 -4.674658,0 -8.466667,3.79306 -8.466667,8.46905 0,4.67386 3.792273,8.46269 8.466667,8.46269 2.341033,0 4.234392,1.8976 4.234392,4.2336 0,2.33998 -1.893359,4.23333 -4.234392,4.23333 -9.350641,-0.001 -16.930424,-7.5811 -16.930424,-16.93068 0,-9.35567 7.579783,-16.93598 16.930423,-16.93598 9.351433,0 16.936508,7.58005 16.936508,16.93598 0,4.67386 3.786717,8.4627 8.462434,8.4627 h 2.116666 6.351059 c 4.674658,0 8.467725,-3.78884 8.467725,-8.4627 0,-6.4143 -1.466586,-12.47087 -4.020609,-17.91997 -0.218016,-0.46513 -0.432064,-0.93425 -0.66966,-1.39515 -0.313267,-0.61172 -0.650081,-1.2155 -0.991129,-1.81292 -0.256382,-0.45085 -0.532342,-0.89324 -0.807244,-1.32927 -0.34634,-0.55404 -0.697706,-1.11416 -1.067594,-1.65153 -0.569383,-0.82074 -1.173162,-1.61846 -1.794139,-2.39765 -0.370946,-0.45482 -0.759619,-0.89086 -1.144059,-1.33324 -0.455877,-0.52096 -0.92075,-1.04378 -1.400439,-1.54199 -0.3937,-0.40323 -0.787665,-0.79825 -1.195917,-1.18666 -0.51779,-0.49821 -1.058333,-0.97366 -1.599935,-1.447 -0.34634,-0.29978 -0.688446,-0.60775 -1.043782,-0.89721 -0.902229,-0.7276 -1.827477,-1.43457 -2.786591,-2.08968 -0.280194,-0.19023 -0.574675,-0.35136 -0.858838,-0.53525 -0.745067,-0.47942 -1.50495,-0.95488 -2.283089,-1.39118 -0.398992,-0.2233 -0.801953,-0.43206 -1.206236,-0.64082 -0.735806,-0.38444 -1.490398,-0.7403 -2.250017,-1.08347 -0.37412,-0.16536 -0.740039,-0.34104 -1.119187,-0.49821 -1.125538,-0.45878 -2.269596,-0.88688 -3.4417,-1.25677 -0.14261,-0.0434 -0.289454,-0.0765 -0.436033,-0.11985 -1.0541,-0.31618 -2.132277,-0.59743 -3.218392,-0.82894 -0.404019,-0.0868 -0.821796,-0.16113 -1.234281,-0.2376 -0.88265,-0.16748 -1.779852,-0.29977 -2.682082,-0.40931 -0.389731,-0.0476 -0.768879,-0.095 -1.162843,-0.13229 -1.295929,-0.11774 -2.596092,-0.19844 -3.921125,-0.19844 -9.350111,-2.6e-4 -16.929894,-7.58243 -16.929894,-16.93624 v -8.46905 c 0,-9.3517 7.579783,-16.931746 16.930423,-16.931746 1.376627,0 2.738967,0.03916 4.101042,0.105304 0.493977,0.02276 0.987954,0.07038 1.486164,0.09922 0.906463,0.06191 1.812925,0.11774 2.709863,0.204523 0.470429,0.04763 0.930275,0.103452 1.395412,0.161131 0.897202,0.103452 1.79414,0.208757 2.682081,0.341048 0.531284,0.07646 1.058334,0.167482 1.585384,0.256382 0.754591,0.124089 1.499658,0.256381 2.250016,0.398991 0.602457,0.119856 1.195917,0.241829 1.789113,0.369888 0.778139,0.167481 1.55231,0.347398 2.316162,0.537633 0.512763,0.128058 1.021027,0.256381 1.53379,0.388673 0.754592,0.204523 1.50495,0.421743 2.25425,0.640823 0.541602,0.16113 1.091406,0.31829 1.633008,0.49001 0.692415,0.21695 1.367367,0.45481 2.055813,0.69664 0.593196,0.20453 1.190625,0.40296 1.77456,0.62627 0.583936,0.21908 1.167871,0.46091 1.746779,0.68845 0.631561,0.25215 1.272382,0.49424 1.89865,0.76491 0.541602,0.23151 1.063625,0.47942 1.599936,0.72152 1.343554,0.60775 2.667529,1.24857 3.976952,1.92246 0.408252,0.20876 0.821531,0.41328 1.224756,0.6305 0.688446,0.36989 1.362075,0.76068 2.035969,1.14935 0.437356,0.24606 0.883708,0.48789 1.315773,0.75036 0.697706,0.41751 1.384829,0.85778 2.073275,1.29593 0.375179,0.24183 0.755385,0.47334 1.129771,0.72152 0.66966,0.44026 1.329002,0.8972 1.979083,1.35599 0.418571,0.29554 0.840317,0.59319 1.257829,0.8972 0.598488,0.43815 1.191683,0.88265 1.775619,1.329 0.412485,0.31433 0.817561,0.63659 1.218411,0.95912 0.62415,0.50218 1.24433,1.00674 1.85631,1.52347 0.32253,0.2749 0.64082,0.55192 0.95912,0.83105 0.65722,0.57468 1.31471,1.14935 1.95553,1.74678 0.21484,0.19421 0.41328,0.39476 0.6223,0.5932 5.9862,5.67214 11.15616,12.2084 15.28419,19.41433 0.0516,0.0847 0.10133,0.17568 0.15081,0.26458 0.55192,0.96944 1.0832,1.95131 1.5957,2.94349 0.0661,0.12806 0.12806,0.25215 0.19421,0.38047 0.51673,1.01098 1.01705,2.02592 1.49251,3.0615 0.0206,0.0516 0.0455,0.10345 0.0704,0.15716 0.49821,1.08532 0.97764,2.17858 1.42822,3.28877 0,0.004 0,0.008 0.004,0.0185 2.32966,5.71976 4.05342,11.75173 5.09324,18.01919 0.75036,4.53125 1.16363,9.17786 1.16363,13.92634 0,4.67387 3.79095,8.4627 8.46667,8.4627 h 8.4664 c 4.67572,0 8.46455,-3.78883 8.46455,-8.4627 0,-6.85456 -0.65537,-13.55222 -1.85631,-20.05912 -0.0146,-0.0722 -0.0228,-0.14261 -0.0392,-0.21908 -0.29766,-1.58988 -0.6305,-3.17103 -0.99642,-4.73578 -0.0124,-0.0619 -0.0228,-0.11985 -0.0373,-0.17568 -0.3556,-1.50495 -0.74004,-3.00143 -1.15332,-4.48151 -0.0537,-0.19421 -0.10742,-0.38867 -0.16536,-0.58499 -0.42175,-1.47585 -0.8681,-2.92683 -1.34779,-4.38018 -0.0103,-0.0288 -0.0206,-0.0661 -0.0288,-0.0992 -3.62162,-10.89766 -8.89688,-21.03305 -15.53633,-30.12202 -0.0185,-0.0331 -0.0373,-0.0619 -0.0619,-0.091 -0.83105,-1.13506 -1.68883,-2.24499 -2.56328,-3.34672 -0.17357,-0.21695 -0.34105,-0.4445 -0.51673,-0.66357 -0.81015,-1.01097 -1.65153,-2.00316 -2.50111,-2.98503 -0.24395,-0.28099 -0.48789,-0.56012 -0.73184,-0.84137 -0.84746,-0.96335 -1.72799,-1.91215 -2.60853,-2.84851 -0.24394,-0.25638 -0.49,-0.51673 -0.7358,-0.77311 -0.8599,-0.89297 -1.73831,-1.76953 -2.6252,-2.62943 -0.33284,-0.31829 -0.66146,-0.63659 -0.99642,-0.95488 -0.87445,-0.82471 -1.7653,-1.64121 -2.66859,-2.443164 -0.32861,-0.293687 -0.66357,-0.578908 -0.99827,-0.868362 -0.89508,-0.779198 -1.80234,-1.548342 -2.72838,-2.29235 -0.42572,-0.34316 -0.85381,-0.673889 -1.28588,-1.01282 -0.86809,-0.678128 -1.74042,-1.366573 -2.63339,-2.021682 -0.44239,-0.322527 -0.89297,-0.634471 -1.34356,-0.959114 -0.92604,-0.659342 -1.85208,-1.322917 -2.79453,-1.955536 -0.20664,-0.136525 -0.40931,-0.264583 -0.61383,-0.398991 -2.39157,-1.575065 -4.84108,-3.06123 -7.35674,-4.450557 -0.13229,-0.07646 -0.26458,-0.157162 -0.39899,-0.229393 -0.75036,-0.407194 -1.518181,-0.791634 -2.2778,-1.186392 -0.783431,-0.407194 -1.561571,-0.816504 -2.354262,-1.200944 -0.745067,-0.363802 -1.509977,-0.706966 -2.264569,-1.052248 -0.768879,-0.351366 -1.52744,-0.702733 -2.306902,-1.039812 -0.820738,-0.3556 -1.646502,-0.684213 -2.477559,-1.021027 -0.787664,-0.318294 -1.565804,-0.630502 -2.363787,-0.930275 -0.759619,-0.285221 -1.528498,-0.556154 -2.29235,-0.820738 -0.897202,-0.314325 -1.799431,-0.6223 -2.705894,-0.915723 -0.725487,-0.233627 -1.461294,-0.460904 -2.197364,-0.679979 -0.91149,-0.274902 -1.827213,-0.531283 -2.748227,-0.777346 -0.792692,-0.214841 -1.585384,-0.423862 -2.383367,-0.618066 -0.886883,-0.216959 -1.779852,-0.417513 -2.677054,-0.61595 -0.816504,-0.182034 -1.627717,-0.357717 -2.453481,-0.518848 -0.935303,-0.179917 -1.884098,-0.351367 -2.828661,-0.506413 -0.768879,-0.134409 -1.53379,-0.2667 -2.311929,-0.380471 -1.120246,-0.161131 -2.25425,-0.293423 -3.383756,-0.427831 -0.617009,-0.07038 -1.220523,-0.150813 -1.837532,-0.21299 -1.271323,-0.12409 -2.557991,-0.21299 -3.839633,-0.293423 -0.531283,-0.03519 -1.053042,-0.08678 -1.589617,-0.109537 -0.417512,-0.01852 -0.840316,-0.01852 -1.262856,-0.03307 v -0.04762 c -1.067594,-0.01032 -2.131219,-0.417513 -2.947723,-1.234017 -1.651529,-1.651794 -1.651529,-4.334669 0,-5.986462 0.902229,-0.901171 2.107407,-1.262857 3.284538,-1.182423 C 115.49596,65.693546 166.14754,117.78286 166.14754,181.72922 Z M 34.909966,71.659107 c -2.335742,0 -4.2291,-1.897592 -4.2291,-4.233598 0,-2.339975 1.893358,-4.233333 4.2291,-4.233333 2.339975,0 4.234392,1.893358 4.234392,4.233333 0,2.336006 -1.894417,4.233598 -4.234392,4.233598 z" fill="#e67e22" fill-rule="evenodd" id="path824" style="stroke-width:0.264583;fill:#0ea5e9;fill-opacity:1" /> </g> </svg> -->
		</button>
		<!-- <svg class="h-10 w-10 m-2 " enable-background="new 0 0 512 512" height="512px" id="Layer_5" version="1.1" viewBox="0 0 512 512" width="512px" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><path clip-rule="evenodd" d="M512,448.014C512,483.351,483.344,512,448,512h-32.001  c-35.359,0-64.008-28.649-64.008-63.986c0-7.859-0.414-15.609-1.023-23.306c-0.109-1.289-0.156-2.586-0.281-3.859  c-0.688-7.297-1.672-14.531-2.891-21.673c-0.234-1.398-0.5-2.758-0.75-4.141c-1.352-7.195-2.906-14.312-4.781-21.328  c-0.266-1.023-0.57-1.992-0.859-3c-2.008-7.228-4.219-14.368-6.766-21.368c-0.031-0.102-0.086-0.227-0.125-0.336  c-2.562-7.031-5.484-13.907-8.594-20.673c-0.438-0.984-0.844-1.969-1.305-2.938c-3.056-6.461-6.353-12.758-9.853-18.945  c-0.609-1.094-1.234-2.172-1.859-3.267c-3.594-6.148-7.391-12.18-11.414-18.031c-0.555-0.805-1.133-1.609-1.68-2.398  c-4.188-5.961-8.547-11.79-13.156-17.407c-0.125-0.156-0.281-0.336-0.406-0.5c-4.633-5.594-9.5-10.961-14.516-16.195  c-0.828-0.883-1.641-1.758-2.5-2.625c-4.891-4.984-10.004-9.773-15.262-14.383c-0.949-0.844-1.918-1.688-2.871-2.492  c-5.383-4.594-10.871-9.023-16.574-13.203c-0.809-0.594-1.652-1.149-2.461-1.743c-5.938-4.273-11.981-8.383-18.262-12.203  c-0.105-0.055-0.195-0.125-0.305-0.195c-6.188-3.766-12.594-7.227-19.066-10.547c-1.168-0.594-2.297-1.203-3.465-1.773  c-6.223-3.07-12.574-5.868-19.051-8.493c-1.363-0.555-2.727-1.109-4.109-1.633c-6.582-2.547-13.258-4.875-20.075-6.938  c-1.059-0.32-2.152-0.609-3.211-0.914c-25.562-7.43-52.563-11.5-80.532-11.5c-17.668,0-32,14.313-32,31.985v32.009  c0,17.665,14.333,32.001,32,32.001c17.566,0,34.516,2.562,50.68,6.977c4.953,1.367,9.812,2.906,14.602,4.648  c0.289,0.086,0.539,0.219,0.809,0.305c23.699,8.704,45.172,21.938,63.434,38.643c0.789,0.734,1.617,1.422,2.387,2.172  c2.871,2.711,5.633,5.523,8.324,8.414c1.184,1.273,2.367,2.547,3.516,3.859c2.402,2.688,4.699,5.414,6.941,8.25  c1.328,1.688,2.602,3.445,3.875,5.188c2.012,2.743,4.055,5.454,5.922,8.282c1.363,2.086,2.566,4.234,3.855,6.352  c2.512,4.164,4.879,8.398,7.07,12.758c1.613,3.195,3.227,6.384,4.664,9.688c1.004,2.273,1.848,4.625,2.762,6.961  c1.203,3.102,2.387,6.203,3.426,9.383c0.789,2.367,1.488,4.75,2.172,7.141c0.969,3.391,1.848,6.829,2.637,10.274  c0.52,2.266,1.039,4.523,1.469,6.836c0.738,3.844,1.277,7.75,1.777,11.641c0.254,2.07,0.574,4.094,0.773,6.172  c0.355,3.774,0.43,7.61,0.574,11.431c0.23,2.188,0.34,4.398,0.34,6.641c0,35.337-28.648,63.986-64.008,63.986h-24.004h-8  c-35.34,0-63.992-28.649-63.992-63.986c0-17.673-14.333-32.009-32.004-32.009c-17.668,0-32,14.336-32,32.009  c0,17.665,14.333,31.985,32,31.985c8.848,0,16.004,7.172,16.004,16.001c0,8.844-7.156,16-16.004,16C28.648,512,0,483.351,0,448.014  c0-35.36,28.648-64.01,63.989-64.01c35.344,0,64.012,28.649,64.012,64.01c0,17.665,14.312,31.985,31.984,31.985h8h24.004  c17.668,0,32.004-14.32,32.004-31.985c0-24.243-5.543-47.134-15.196-67.729c-0.824-1.758-1.633-3.531-2.531-5.273  c-1.184-2.312-2.457-4.594-3.746-6.852c-0.969-1.704-2.012-3.376-3.051-5.024c-1.309-2.094-2.637-4.211-4.035-6.242  c-2.152-3.102-4.434-6.117-6.781-9.062c-1.402-1.719-2.871-3.367-4.324-5.039c-1.723-1.969-3.48-3.945-5.293-5.828  c-1.488-1.524-2.977-3.017-4.52-4.485c-1.957-1.883-4-3.68-6.047-5.469c-1.309-1.133-2.602-2.297-3.945-3.391  c-3.41-2.75-6.907-5.422-10.532-7.898c-1.059-0.719-2.172-1.328-3.246-2.023c-2.816-1.812-5.688-3.609-8.629-5.258  c-1.508-0.844-3.031-1.633-4.559-2.422c-2.781-1.453-5.633-2.798-8.504-4.095c-1.414-0.625-2.797-1.289-4.23-1.883  c-4.254-1.734-8.578-3.352-13.008-4.75c-0.539-0.164-1.094-0.289-1.648-0.453c-3.984-1.195-8.059-2.258-12.164-3.133  c-1.527-0.328-3.106-0.609-4.665-0.898c-3.336-0.633-6.727-1.133-10.137-1.547c-1.473-0.18-2.906-0.359-4.395-0.5  c-4.898-0.445-9.812-0.75-14.82-0.75C28.648,288.009,0,259.352,0,223.999V191.99c0-35.345,28.648-63.994,63.989-63.994  c5.203,0,10.352,0.148,15.5,0.398c1.867,0.086,3.734,0.266,5.617,0.375c3.426,0.234,6.852,0.445,10.242,0.773  c1.778,0.18,3.516,0.391,5.274,0.609c3.391,0.391,6.781,0.789,10.137,1.289c2.008,0.289,4,0.633,5.992,0.969  c2.852,0.469,5.668,0.969,8.504,1.508c2.277,0.453,4.52,0.914,6.762,1.398c2.941,0.633,5.867,1.313,8.754,2.032  c1.938,0.484,3.859,0.969,5.797,1.469c2.852,0.773,5.688,1.594,8.52,2.422c2.047,0.609,4.125,1.203,6.172,1.852  c2.617,0.82,5.168,1.719,7.77,2.633c2.242,0.773,4.5,1.523,6.707,2.367c2.207,0.828,4.414,1.742,6.602,2.602  c2.387,0.953,4.809,1.868,7.176,2.891c2.047,0.875,4.02,1.812,6.047,2.727c5.078,2.297,10.082,4.719,15.031,7.266  c1.543,0.789,3.105,1.562,4.629,2.383c2.602,1.398,5.148,2.875,7.695,4.344c1.653,0.93,3.34,1.844,4.973,2.836  c2.637,1.578,5.234,3.242,7.836,4.898c1.418,0.914,2.855,1.789,4.27,2.727c2.531,1.664,5.023,3.391,7.48,5.125  c1.582,1.117,3.176,2.242,4.754,3.391c2.262,1.656,4.504,3.336,6.711,5.023c1.559,1.188,3.09,2.406,4.605,3.625  c2.359,1.898,4.703,3.805,7.016,5.758c1.219,1.039,2.422,2.086,3.625,3.141c2.484,2.172,4.969,4.344,7.391,6.602  c0.812,0.734,1.562,1.492,2.352,2.242c22.625,21.438,42.165,46.142,57.767,73.377c0.195,0.32,0.383,0.664,0.57,1  c2.086,3.664,4.094,7.375,6.031,11.125c0.25,0.484,0.484,0.953,0.734,1.438c1.953,3.821,3.844,7.657,5.641,11.571  c0.078,0.195,0.172,0.391,0.266,0.594c1.883,4.102,3.695,8.234,5.398,12.43c0,0.016,0,0.031,0.016,0.07  c8.805,21.618,15.32,44.416,19.25,68.104c2.836,17.126,4.398,34.688,4.398,52.635c0,17.665,14.328,31.985,32,31.985H448  c17.672,0,31.992-14.32,31.992-31.985c0-25.907-2.477-51.221-7.016-75.814c-0.055-0.273-0.086-0.539-0.148-0.828  c-1.125-6.009-2.383-11.985-3.766-17.899c-0.047-0.234-0.086-0.453-0.141-0.664c-1.344-5.688-2.797-11.344-4.359-16.938  c-0.203-0.734-0.406-1.469-0.625-2.211c-1.594-5.578-3.281-11.062-5.094-16.555c-0.039-0.109-0.078-0.25-0.109-0.375  c-13.688-41.188-33.626-79.495-58.72-113.847c-0.07-0.125-0.141-0.234-0.234-0.344c-3.141-4.29-6.383-8.485-9.688-12.649  c-0.656-0.82-1.289-1.68-1.953-2.508c-3.062-3.821-6.242-7.571-9.453-11.282c-0.922-1.062-1.844-2.117-2.766-3.18  c-3.203-3.641-6.531-7.227-9.859-10.766c-0.922-0.969-1.852-1.953-2.781-2.922c-3.25-3.375-6.57-6.688-9.922-9.938  c-1.258-1.203-2.5-2.406-3.766-3.609c-3.305-3.117-6.672-6.203-10.086-9.234c-1.242-1.11-2.508-2.188-3.773-3.282  c-3.383-2.945-6.812-5.852-10.312-8.664c-1.609-1.297-3.227-2.547-4.86-3.828c-3.281-2.563-6.578-5.165-9.953-7.641  c-1.672-1.219-3.375-2.398-5.078-3.625c-3.5-2.492-7-5-10.562-7.391c-0.781-0.516-1.547-1-2.32-1.508  c-9.039-5.953-18.297-11.57-27.805-16.821c-0.5-0.289-1-0.594-1.508-0.867c-2.836-1.539-5.738-2.992-8.609-4.484  c-2.961-1.539-5.902-3.086-8.898-4.539c-2.816-1.375-5.707-2.672-8.559-3.977c-2.906-1.328-5.773-2.656-8.719-3.93  c-3.102-1.344-6.223-2.586-9.364-3.859c-2.977-1.203-5.918-2.383-8.934-3.516c-2.871-1.078-5.777-2.102-8.664-3.102  c-3.391-1.188-6.801-2.352-10.227-3.461c-2.742-0.883-5.523-1.742-8.305-2.57c-3.445-1.039-6.906-2.008-10.387-2.938  c-2.996-0.812-5.992-1.602-9.008-2.336c-3.352-0.82-6.727-1.578-10.118-2.328c-3.086-0.688-6.152-1.352-9.273-1.961  c-3.535-0.68-7.121-1.328-10.691-1.914c-2.906-0.508-5.797-1.008-8.738-1.438c-4.234-0.609-8.52-1.109-12.789-1.617  c-2.332-0.266-4.613-0.57-6.945-0.805c-4.805-0.469-9.668-0.805-14.512-1.109c-2.008-0.133-3.98-0.328-6.008-0.414  c-1.578-0.07-3.176-0.07-4.773-0.125v-0.18c-4.035-0.039-8.055-1.578-11.141-4.664c-6.242-6.243-6.242-16.383,0-22.626  c3.41-3.406,7.965-4.773,12.414-4.469C320.561,9.454,512,206.327,512,448.014L512,448.014z M15.984,32.001  C7.156,32.001,0,24.829,0,16C0,7.156,7.156,0,15.984,0c8.844,0,16.004,7.156,16.004,16C31.988,24.829,24.828,32.001,15.984,32.001  L15.984,32.001z" fill="#E67E22" fill-rule="evenodd"/></svg> -->
		<!-- spaces -->
		<button class="rounded-lg hover:bg-slate-800" on:click={() => goto('/g')}>
			<Spaces />
		</button>
		<!-- question for you -->
		<button
			class="rounded-lg hover:bg-slate-800"
			on:click={() => {
				goto('/fy');
			}}
		>
			<QuestionForYou />
		</button>
		<!-- search bar  -->
		<div class="mx-2 flex justify-center  text-gray-200">
			<div class=" xl:w-[600px]">
				<div class="input-group relative flex w-full flex-row items-stretch ">
					<!-- search input -->
					<input
						type="search"
						class="form-control  relative m-0 block w-full min-w-0 flex-auto rounded border border-solid  border-gray-800 bg-gray-600 bg-clip-padding px-3 py-1.5 text-base font-normal transition ease-in-out  focus:border-blue-600 focus:bg-gray-600 focus:outline-none"
						placeholder="Search"
						aria-label="Search"
						aria-describedby="button-addon2"
					/>
					<!-- search button -->
					<button
						class="  btn inline-block items-center rounded bg-sky-500 px-6 py-2.5 text-xs font-medium uppercase leading-tight text-white shadow-md transition duration-150  ease-in-out hover:bg-blue-600 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800  active:shadow-lg"
						type="button"
						id="button-addon2"
					>
						<Search />
					</button>
				</div>
			</div>
		</div>
		<div
			class="flex flex-row items-center justify-center rounded-lg hover:cursor-pointer hover:bg-gray-700 hover:bg-opacity-50 "
			on:click={() => {
				goto(`${UserData['UserID']}`);
			}}
			on:keydown={(e) => {}}
		>
			<!-- User -->
			<img
				src={UserData['UserImage']}
				alt=""
				class=" aspect-square active:ring-offset-base-50 mx-2  h-10 w-10  cursor-pointer rounded-xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
			/>
			<!-- user Credit -->
			<div class="max-w-36   flex h-fit  min-w-min flex-row items-center justify-center">
				<p class="mx-1 font-semibold text-white">
					{ShortenNumber(UserData['Badges']['Reputation'])}
				</p>

				{#if UserData['Badges']['Gold'] != 0}
					<GoldDot />
					<p class="mx-1   text-[#ffcc01]">{ShortenNumber(UserData['Badges']['Gold'])}</p>
				{/if}
				{#if UserData['Badges']['Silver'] != 0}
					<SilverDot />
					<p class="mx-1 my-4 text-[#b4b8bc]">{ShortenNumber(UserData['Badges']['Silver'])}</p>
				{/if}
				{#if UserData['Badges']['Bronze'] != 0}
					<BronzeDot />
					<p class="mx-1 my-4 text-[#d1a684]">{ShortenNumber(UserData['Badges']['Bronze'])}</p>
				{/if}
			</div>
		</div>
		<!-- notification -->
		<div class="   ">
			<button
				class=" flex  h-full w-10 items-center justify-center  rounded-lg hover:bg-slate-800  focus:outline-none"
				on:blur={() => {
					NotificationOptsOpen = false;
				}}
				on:click={() => {
					if (NotificationOptsOpen) {
						NotificationOptsOpen = false;
					} else {
						AllOtherFalse();
						NotificationOptsOpen = true;
					}
				}}
			>
				<NotificationIcon />
			</button>
			{#if NotificationOptsOpen}
				<div
					class=" fixed z-30 float-left -ml-[340px]  mt-2 flex max-h-screen min-h-fit w-96  flex-col overflow-scroll   rounded-lg border-2 border-gray-600 bg-gray-500 text-gray-200"
				>
					{#each NotificationOpts2 as i}
						<button
							class="flex h-fit w-full  flex-row border-0 border-b-[1px] border-gray-700 bg-[#2d2d2d]  py-2 px-4 text-[#3196e3] transition-all duration-150  ease-linear  hover:bg-[#404245] "
						>
							<!-- <svg class="h-6 w-6" version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 98.1 118" style="enable-background:new 0 0 98.1 118;" xml:space="preserve" > <style type="text/css"> .st0 { fill: #ff3e00; } .st1 { fill: #ffffff; } </style> <path class="st0" d="M91.8,15.6C80.9-0.1,59.2-4.7,43.6,5.2L16.1,22.8C8.6,27.5,3.4,35.2,1.9,43.9c-1.3,7.3-0.2,14.8,3.3,21.3  c-2.4,3.6-4,7.6-4.7,11.8c-1.6,8.9,0.5,18.1,5.7,25.4c11,15.7,32.6,20.3,48.2,10.4l27.5-17.5c7.5-4.7,12.7-12.4,14.2-21.1  c1.3-7.3,0.2-14.8-3.3-21.3c2.4-3.6,4-7.6,4.7-11.8C99.2,32.1,97.1,22.9,91.8,15.6" /> <path class="st1" d="M40.9,103.9c-8.9,2.3-18.2-1.2-23.4-8.7c-3.2-4.4-4.4-9.9-3.5-15.3c0.2-0.9,0.4-1.7,0.6-2.6l0.5-1.6l1.4,1  c3.3,2.4,6.9,4.2,10.8,5.4l1,0.3l-0.1,1c-0.1,1.4,0.3,2.9,1.1,4.1c1.6,2.3,4.4,3.4,7.1,2.7c0.6-0.2,1.2-0.4,1.7-0.7L65.5,72  c1.4-0.9,2.3-2.2,2.6-3.8c0.3-1.6-0.1-3.3-1-4.6c-1.6-2.3-4.4-3.3-7.1-2.6c-0.6,0.2-1.2,0.4-1.7,0.7l-10.5,6.7  c-1.7,1.1-3.6,1.9-5.6,2.4c-8.9,2.3-18.2-1.2-23.4-8.7c-3.1-4.4-4.4-9.9-3.4-15.3c0.9-5.2,4.1-9.9,8.6-12.7l27.5-17.5  c1.7-1.1,3.6-1.9,5.6-2.5c8.9-2.3,18.2,1.2,23.4,8.7c3.2,4.4,4.4,9.9,3.5,15.3c-0.2,0.9-0.4,1.7-0.7,2.6l-0.5,1.6l-1.4-1  c-3.3-2.4-6.9-4.2-10.8-5.4l-1-0.3l0.1-1c0.1-1.4-0.3-2.9-1.1-4.1c-1.6-2.3-4.4-3.3-7.1-2.6c-0.6,0.2-1.2,0.4-1.7,0.7L32.4,46.1  c-1.4,0.9-2.3,2.2-2.6,3.8s0.1,3.3,1,4.6c1.6,2.3,4.4,3.3,7.1,2.6c0.6-0.2,1.2-0.4,1.7-0.7l10.5-6.7c1.7-1.1,3.6-1.9,5.6-2.5  c8.9-2.3,18.2,1.2,23.4,8.7c3.2,4.4,4.4,9.9,3.5,15.3c-0.9,5.2-4.1,9.9-8.6,12.7l-27.5,17.5C44.8,102.5,42.9,103.3,40.9,103.9" /> </svg>
							<div class="">
								<div
									class=" h-fit  w-full border-0 py-2 px-4 text-[#3196e3] transition-all duration-150  ease-linear   "
								>
									{i}
								</div>
							</div> -->

							{#if i[0] === 'answer'}
								<Svelte />
								<div class=" flex w-80 flex-col gap-1">
									<div class="StackTextColor flex flex-row  font-mono">
										<div class="">{i[0]}</div>
										<div class=" grow" />
										<div class="">{i[1]}</div>
									</div>
									<div class=" text-left line-clamp-1">
										{i[2]}
									</div>
									<div class=" StackTextColor text-left leading-tight line-clamp-2">
										{i[3]}
									</div>
								</div>
							{:else if i[0] === 'comment'}
								<Flutter />
								<div class=" flex w-80 flex-col gap-1">
									<div class="StackTextColor flex flex-row  font-mono">
										<div class="">{i[0]}</div>
										<div class=" grow" />
										<div class="">{i[1]}</div>
									</div>
									<div class=" text-left line-clamp-1">
										{i[2]}
									</div>
									<div class=" StackTextColor text-left leading-tight line-clamp-2">
										{i[3]}
									</div>
								</div>
							{:else if i[0] === 'reward'}
								<Vue />
								<div class=" flex w-80 flex-col gap-1">
									<div class="StackTextColor flex flex-row  font-mono ">
										{#if Number(i[2]) < 0}
											<div class=" mx-1 text-xs font-semibold text-red-500 ">{i[2]}</div>
										{:else}
											<div class=" mx-1 font-semibold text-green-400">+{i[2]}</div>
										{/if}
										<div class="">{i[0]}</div>
										<div class=" grow" />
										<div class="">{i[1]}</div>
									</div>
									<div class=" text-left line-clamp-1">
										{i[3]}
									</div>
								</div>
							{/if}
						</button>
					{/each}
				</div>
			{/if}
		</div>

		<!-- menu -->
		<div class="   ">
			<button
				class=" flex h-full w-10  items-center justify-center overflow-hidden rounded-lg hover:bg-slate-800 focus:outline-none"
				on:blur={() => {
					NavbarOptsOpen = false;
				}}
				on:click={() => {
					if (NavbarOptsOpen) {
						NavbarOptsOpen = false;
					} else {
						AllOtherFalse();
						NavbarOptsOpen = true;
					}
				}}
			>
				<Menubar />
			</button>
			{#if NavbarOptsOpen}
				<div
					class=" fixed z-30 float-left  -ml-[340px] mt-2 flex max-h-screen min-h-fit w-96  flex-col overflow-scroll rounded-lg  border-2 border-gray-600 bg-[#2d2d2d]  text-gray-200"
				>
					<div class="mt-3 ml-2 mb-1 flex flex-row">
						<img
							src={UserData['UserImage']}
							alt=""
							class=" aspect-square active:ring-offset-base-50 mx-2  h-16 w-16  cursor-pointer rounded-xl object-cover transition-all duration-150 ease-linear hover:rounded-xl hover:ring hover:ring-cyan-500  active:rounded-md  active:ring  active:ring-blue-600"
						/>
						<div class="flex flex-col items-center justify-center">
							<div class=" text-2xl font-light line-clamp-1 ">{UserData.UserName}</div>
							<p class=" font-semibold">
								{ShortenNumber(UserData.Badges.Reputation)} Reputation <WhiteDot />
								{ShortenNumber(UserData.Follower.length)} Follower
							</p>
						</div>
					</div>
					<!-- <button
						class="flex h-fit w-full  flex-row border-0 border-b-[1px] border-gray-700 bg-[#2d2d2d]  py-2 px-4 text-[#3196e3] transition-all duration-150  ease-linear  hover:bg-[#404245] "
					>
						profile
					</button> -->
					<button
						class=" h-fit w-full border-0  border-b-[1px] border-gray-700 bg-[#2d2d2d] py-2  px-4 text-xl text-[#3196e3] transition-all duration-150  ease-linear  hover:bg-[#404245]  "
					>
						My followers
					</button>
					<button
						class=" h-fit w-full border-0  border-b-[1px] border-gray-700 bg-[#2d2d2d] py-2  px-4 text-xl text-[#3196e3] transition-all duration-150  ease-linear  hover:bg-[#404245]  "
					>
						My Following
					</button>
					<button
						class=" h-fit w-full border-0  border-b-[1px] border-gray-700 bg-[#2d2d2d] py-2  px-4 text-xl text-[#3196e3] transition-all duration-150  ease-linear  hover:bg-[#404245]  "
					>
						My Stats
					</button>
					<button
						class=" h-fit w-full border-0  border-b-[1px] border-gray-700 bg-[#2d2d2d] py-2  px-4 text-xl text-[#3196e3] transition-all duration-150  ease-linear  hover:bg-[#404245]  "
					>
						HeapOverflow Stats
					</button>
					<button
						class=" h-fit w-full border-0  border-b-[1px] border-gray-700 bg-[#2d2d2d] py-2  px-4 text-xl text-[#3196e3] transition-all duration-150  ease-linear  hover:bg-[#404245]  "
					>
						Settings
					</button>

					<button
						class=" h-fit w-full border-0  border-b-[1px] border-gray-700 bg-[#2d2d2d] py-2  px-4 text-xl text-red-500 transition-all duration-150  ease-linear  hover:bg-[#404245]  "
					>
						Logout
					</button>
				</div>
			{/if}
		</div>

		<div class=" grow" />
	</div>
{:else}
	<!-- else content here -->
{/if}
