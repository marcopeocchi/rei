<script lang="ts">
  import { get } from "svelte/store";
  import { temperatureStore, topStore } from "./store";
  import Entry from "./Entry.svelte";
  import { formatHHMMSS } from "./utils";
  import { onDestroy } from "svelte";

  const top = get(topStore);

  let interval;
  let uptime = 0;
  let temperature = 0;

  temperatureStore.subscribe((thermals) =>
    thermals.then((resolved) => (temperature = Number(resolved.cpuTemp)))
  );

  topStore.subscribe((t) =>
    t.then((resolved) => {
      uptime = Number(resolved.uptime);
      interval = setInterval(() => uptime++, 1000);
    })
  );

  onDestroy(() => clearInterval(interval));
</script>

<header class="bg-neutral-200 dark:bg-neutral-800 p-8 rounded mt-8 font-mono">
  {#await top}
    <p class="flex items-center justify-center">Loading top...</p>
  {:then data}
    <Entry title="Hostname" value={data.hostname} tabs={1} />
    <Entry title="OS" value={data.os} tabs={7} />
    <Entry title="Platform" value={data.platform} tabs={1} />
    <Entry title="CPU" value={`${data.coreCount} x ${data.cpu}`} tabs={6} />
    <Entry
      title="Free RAM"
      value={`${(Number(data.ramFree) / 1_000_000).toFixed(0)}MB`}
      tabs={1}
    />
  {/await}
  <Entry title="Uptime" value={formatHHMMSS(uptime)} tabs={3} />
  <br />
  <Entry
    title="PKG Temperature"
    value={`${(Number(temperature) / 1000).toFixed(2)}°C`}
    tabs={1}
  />
</header>
