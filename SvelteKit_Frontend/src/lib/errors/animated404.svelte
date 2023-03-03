

<script lang="ts">
    // your script goes here
</script>

<style lang="css">
    /* your styles go here */

    /* @import url(https://fonts.googleapis.com/css?family=Open+Sans:800); */

#g-usetag {
	stroke-width: 6;
	stroke: #f2385a;
	animation: gstrokeanim 2s infinite linear;
}
@keyframes gstrokeanim {
	0% {
		stroke-width: 6;
	}
	50% {
		stroke-width: 0;
	}
	100% {
		stroke-width: 6;
	}
}

.text {
	fill: none;
	stroke-width: 6;
	stroke-linejoin: round;
	stroke-dasharray: 70 330;
	stroke-dashoffset: 0;
	-webkit-animation: stroke 6s infinite linear;
	animation: stroke 6s infinite linear;
}

.text:nth-child(5n + 1) {
	stroke: #f2385a;
	-webkit-animation-delay: -1.2s;
	animation-delay: -1.2s;
}

.text:nth-child(5n + 2) {
	stroke: #f5a503;
	-webkit-animation-delay: -2.4s;
	animation-delay: -2.4s;
}

.text:nth-child(5n + 3) {
	stroke: #e9f1df;
	-webkit-animation-delay: -3.6s;
	animation-delay: -3.6s;
}

.text:nth-child(5n + 4) {
	stroke: #56d9cd;
	-webkit-animation-delay: -4.8s;
	animation-delay: -4.8s;
}

.text:nth-child(5n + 5) {
	stroke: #3aa1bf;
	-webkit-animation-delay: -6s;
	animation-delay: -6s;
}

@-webkit-keyframes stroke {
	100% {
		stroke-dashoffset: -400;
	}
}

@keyframes stroke {
	100% {
		stroke-dashoffset: -400;
	}
}

/*////////////////////v animation//////////////////*/

.text1 {
	fill: none;

	stroke-linejoin: round;
	stroke-dasharray: 1160;
	stroke-dashoffset: 1160;
	-webkit-animation: vstroke 6s infinite linear;
	animation: vstroke 6s infinite ease-in-out;
}
.text1:nth-child(3n + 3) {
	stroke-width: 4;
	stroke: #3aa1bf;
	-webkit-animation-delay: 0s;
	animation-delay: 00s;
}
.text1:nth-child(3n + 1) {
	stroke: yellow;
	stroke-width: 1;
	-webkit-animation-delay: 0.2s;
	animation-delay: -0.2s;
}
.text1:nth-child(3n + 2) {
	stroke: red;
	stroke-width: 2;
	-webkit-animation-delay: -0.1s;
	animation-delay: -0.1s;
}

@keyframes vstroke {
	100% {
		stroke-dashoffset: 0;
	}
}

/*///////////////// Other styles /////////////////////////*/
html,
body {
	height: 100%;
}

body {
	background: #212121;
	background-size: 0.2em 100%;
	font: 14.5em/1 Open Sans, Impact;
	text-transform: uppercase;
	margin: 0;
}

svg {
	position: absolute;
	width: 100%;
	height: 100%;
}

</style>

<!-- markup (zero or more items) goes here -->


<svg viewBox="0 0 1320 300">
	<!--pattern-->
	<defs>
		<pattern id="GPattern" x="0" y="0" width="20" height="20" patternUnits="userSpaceOnUse" patternTransform="rotate(35)">
			<animateTransform attributeType="xml" attributeName="patternTransform" type="rotate" from="35" to="395" begin="0" dur="160s" repeatCount="indefinite" />
			<circle cx="10" cy="10" r="10" stroke="none" fill="yellow">
				<animate attributeName="r" type="xml" from="1" to="1" values="1; 10; 1" begin="0s" dur="2s" repeatCount="indefinite" />
			</circle>
		</pattern>
	</defs>

	<!-- Symbol -->
	<symbol id="s-text">
		<text text-anchor="middle" x="35%" y="50%" dy=".35em">
			4
		</text>
	</symbol>
	<symbol id="v-text">
		<text text-anchor="middle" x="50%" y="50%" dy=".35em">
			0
		</text>
	</symbol>
	<symbol id="g-text">
		<text text-anchor="middle" x="65%" y="50%" dy=".35em">
			4
		</text>
	</symbol>
	<!-- Duplicate symbols -->
	<use xlink:href="#s-text" class="text"></use>
	<use xlink:href="#s-text" class="text"></use>
	<use xlink:href="#s-text" class="text"></use>
	<use xlink:href="#s-text" class="text"></use>
	<use xlink:href="#s-text" class="text"></use>
	<use xlink:href="#v-text" class="text1">

	</use>
	<use xlink:href="#v-text" class="text1"></use>
	<use xlink:href="#v-text" class="text1"></use>

	<use id="g-usetag" xlink:href="#g-text" class="text2" style="fill: url(#GPattern)" />
	<!-- </use> -->

</svg>