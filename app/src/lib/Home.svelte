<script lang="ts">
  import type { Config } from "../types";
  import Service from "./Service.svelte";
  import Stats from "./Stats.svelte";
  import Title from "./Title.svelte.svelte";

  export let config: Config;
</script>

<div class="container mx-auto py-8 2xl:px-28 xl:px-28 lg:px-28 md:px-10 px-8">
  <Title title={config.servername} />
  <Stats />
  <div
    class="mt-10 grid grid-cols-2 sm:grid-cols-4 md:grid-cols-5 xl:grid-cols-6 2xl:grid-cols-7 gap-4"
  >
    {#each config.services as service}
      <Service name={service.name} url={service.url} />
    {/each}
  </div>
</div>
