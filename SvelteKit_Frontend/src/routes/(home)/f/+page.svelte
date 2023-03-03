<script lang="ts" context="module">
	export async function load({ params, fetch, session, stuff }) {
		// const UserURL = params.user
		// console.log("🚀 ~ file: index.svelte ~ line 4 ~ load ~ UserURL", UserURL)
		// console.log("🚀 ~ file: index.svelte ~ line 4 ~ load ~ UserURL type", typeof( UserURL))
		if (!session.user.authenticated) {
			return {
				status: 302,
				redirect: '/login'
			};
		}
        return {}
	}
</script>

<script lang="ts">
	// your script goes here
</script>

<!-- markup (zero or more items) goes here -->


<h1 class=" flex flex-col text-white items-center justify-center ">Following
</h1>



