<script lang="ts">
  import Home from "./lib/Home.svelte";
  import type { Config } from "./types";

  const fetcher = async () => {
    const res = await fetch("/config");
    const data: Config = await res.json();
    return data;
  };
</script>

<main
  class="bg-neutral-100 dark:bg-neutral-900 dark:text-neutral-100 min-h-screen"
>
  {#await fetcher()}
    <p class="flex items-center justify-center">Loading ...</p>
  {:then data}
    <Home config={data} />
  {/await}
</main>
