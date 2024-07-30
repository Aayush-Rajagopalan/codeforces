<script>
	import '../app.css';
</script>

<svelte:head>
	<link href="https://fonts.googleapis.com/css?family=Poppins" rel="stylesheet" />
	<title>codeforces</title>
	<meta name="title" content="codeforces" />
	<link rel="icon" href="/favicon.ico" />
	<meta name="description" content="a cleaner interface for codeforces problems" />

	<!-- Open Graph / Facebook -->
	<meta property="og:type" content="website" />
	<meta property="og:url" content="https://cf.aayus.me" />
	<meta property="og:title" content="codeforces" />
	<meta property="og:description" content="a cleaner interface for codeforces problems" />

	<!-- Twitter -->
	<meta property="twitter:card" content="summary_large_image" />
	<meta property="twitter:url" content="https://cf.aayus.me" />
	<meta property="twitter:title" content="codeforces" />
	<meta property="twitter:description" content="a cleaner interface for codeforces problems" />
</svelte:head>
<div class="flex flex-col items-center justify-center">
	<div class="flex h-full min-h-screen max-w-4xl flex-col items-center bg-zinc-950 text-white">
		<div
			class="mt-8 flex w-full min-w-[30%] max-w-4xl items-center justify-between px-4 text-zinc-400"
		>
			<a href="/">codeforces</a>

		</div>
		<slot></slot>
		<footer
			class="mt-auto flex w-full min-w-[30%] max-w-4xl items-center justify-between px-4  gap-1 pb-4 text-zinc-400"
		>
			<div>a project by <a href="https://aayushr.dev" target="_blank">aayush</a></div>
			<div><a href="https://github.com/aayush-rajagopalan/codeforces" target="_blank">view source</a></div>
		</footer>
	</div>
</div>
