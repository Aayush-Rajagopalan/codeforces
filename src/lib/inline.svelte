<script lang="ts">
    import katex from 'katex';
    import 'katex/dist/katex.min.css';
    export let content: string;
    import { onMount } from 'svelte';

    onMount(() => {
      renderKaTeX();
    });
    // Function to replace LaTeX patterns with Katex component
    function renderLatex(content: string) {
      const regex = /\$\$\$([^$]+)\$\$\$/g;
      return content.replace(regex, (_, latex) => `<span class="math"> ${latex} </span>`);
    }
  
    let processedContent = '';
  
    $: processedContent = renderLatex(content);
    function renderKaTeX() {
        const elements = document.querySelectorAll('.math');
        elements.forEach(el => {
          // @ts-ignore
            katex.render(el.textContent, el, {
                throwOnError: false
            });
        });
    }
  </script>
  
  <div>
    {@html processedContent}
  </div>
     

  <style>
    .latex {
      display: inline-block;
    }
  </style>
  