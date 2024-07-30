<script>
	import '../app.css';
</script>

<svelte:head>
	<link href="https://fonts.googleapis.com/css?family=Poppins" rel="stylesheet" />
	<title>codeforces - a project by aayush</title>
	<meta name="title" content="codeforces - a project by aayush" />
	<link rel="icon" href="/favicon.ico" />
	<meta name="description" content="a cleaner interface for codeforces problems" />

	<!-- Open Graph / Facebook -->
	<meta property="og:type" content="website" />
	<meta property="og:url" content="https://cf.aayus.me" />
	<meta property="og:title" content="codeforces - a project by aayush" />
	<meta property="og:description" content="a cleaner interface for codeforces problems" />

	<!-- Twitter -->
	<meta property="twitter:card" content="summary_large_image" />
	<meta property="twitter:url" content="https://cf.aayus.me" />
	<meta property="twitter:title" content="codeforces - a project by aayush" />
	<meta property="twitter:description" content="a cleaner interface for codeforces problems" />
</svelte:head>
<div class="flex flex-col items-center justify-center">
	<div class="flex h-full min-h-screen max-w-4xl flex-col items-center bg-zinc-950 text-white">
		<div class="w min-w-3xl mt-8 flex w-full max-w-3xl items-center justify-between text-zinc-400">
			<div>cf.aayus.me</div>
			<div class="">
				a project by <a href="https://aayushr.dev" target="_blank">aayush</a>
			</div>
		</div>
		<slot></slot>
	</div>
</div>
