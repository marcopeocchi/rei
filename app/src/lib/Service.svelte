<script lang="ts">
  export let url: string;
  export let name: string;
</script>

<a href={url}>
  <div
    class="
      p-3
      text-neutral-700
      rounded-lg text-center cursor-pointer duration-100
      bg-blue-300 hover:bg-blue-400 dark:hover:text-blue-100
    "
  >
    {name}
  </div>
</a>
