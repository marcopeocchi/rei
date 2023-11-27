<script lang="ts">
  import { tabTo } from './utils';

  export let title: string;
  export let value: string;
  export let tabs: number;
</script>

<div>
  <span class="font-bold whitespace-pre-wrap text-red-400 dark:text-red-100">
    {tabTo(title, tabs)}
  </span>
  <span>{value}</span>
</div>

<style scoped>
  div {
    font-family: 'JetBrains Mono', monospace;
  }
</style>
