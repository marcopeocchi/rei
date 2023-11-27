<script lang="ts">
  import { getEndpoint } from './utils';

  let username: string;
  let password: string;
  let hasError: boolean;

  const login = () => {
    fetch(getEndpoint('/login'), {
      method: 'POST',
      body: JSON.stringify({
        username,
        password,
      }),
    })
      .then(() => window.location.reload())
      .catch((e) => (hasError = true));
  };
</script>

<div class="flex items-center justify-center h-screen">
  <div
    class="flex items-center flex-col gap-2 bg-neutral-50 border dark:border-0 dark:bg-neutral-800 p-12 rounded"
  >
    <h1 class="text-4xl font-extrabold text-red-300 dark:text-red-100 mb-6">
      Login
    </h1>
    <input
      placeholder="Username"
      class="appearance-none p-2 bg-neutral-200 dark:bg-neutral-700 rounded border dark:border-0"
      class:ring-2={hasError}
      class:ring-red-400={hasError}
      type="text"
      bind:value={username}
    />
    <input
      placeholder="Password"
      class="appearance-none p-2 bg-neutral-200 dark:bg-neutral-700 rounded border dark:border-0"
      class:ring-2={hasError}
      class:ring-red-400={hasError}
      type="password"
      bind:value={password}
    />
    <button
      type="submit"
      class="bg-red-300 hover:bg-red-400 duration-150 py-2 text-white rounded w-full"
      on:click={login}
    >
      Submit
    </button>
    {#if hasError}
      <p class="mt-6">Wrong username or password!</p>
    {/if}
  </div>
</div>
