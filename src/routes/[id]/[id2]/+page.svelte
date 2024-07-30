<script lang="ts">
	import { onMount } from 'svelte';
	/** @type {import('./$types').PageData} */
	import Katex from '$lib/Katex.svelte';
	export let data;

	let sanitizedContent = '';

	function sanitizeContent(content: any) {
		const wrapperElements = ['p', 'div'];
		let tempElement = document.createElement('div');
		tempElement.innerHTML = content;

		const childNodes = tempElement.childNodes;
		let insideSampleTest = false;
		let foundSectionTitle = false;

		const sanitizedContent = Array.from(childNodes)
			.filter((node) => {
				if (node.nodeType === Node.ELEMENT_NODE) {
					const element = node as Element;
					const nodeName = element.nodeName.toLowerCase();

					if (nodeName === 'div' && element.classList.contains('sample-test')) {
						insideSampleTest = true;
					}

					if (
						insideSampleTest &&
						nodeName === 'div' &&
						element.classList.contains('section-title')
					) {
						foundSectionTitle = true;
					}

					if (insideSampleTest && !foundSectionTitle) {
						if (
							element.classList.contains('input') ||
							element.classList.contains('output') ||
							element.classList.contains('title')
						) {
							return false;
						}
					}

					if (
						nodeName === 'div' &&
						element.classList.contains('sample-test') &&
						foundSectionTitle
					) {
						insideSampleTest = false;
						foundSectionTitle = false;
					}

					return wrapperElements.includes(nodeName);
				}
				return false;
			})
			.map((node) => (node as Element).outerHTML)
			.join('');

		return sanitizedContent;
	}

	onMount(() => {
		sanitizedContent = sanitizeContent(data.content);
	});
</script>

{#if sanitizedContent}
	<div class="flex h-full flex-col bg-zinc-950 text-white">
		<h1 class="py-8 text-center text-3xl font-semibold">{data.title}</h1>
		<div>{@html sanitizedContent}</div>
	</div>
{:else}
{/if}
