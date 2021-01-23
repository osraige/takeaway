<script>
  import { onMount } from "svelte";

  const get = async () => {
    const response = await fetch(`/get?t=${Date.now()}`);
    if (!response.ok) return {};
    return await response.json();
  };

  const put = () => {
    fetch("/put", {
      method: "POST",
      "Content-Type": "application/json",
      body: JSON.stringify(balances),
    });
  };

  const status = {
    WASNT_INVOLVED: "wasnt_involved",
    ORDERED: "ordered",
    WENT_DOWN: "went_down",
  };

  let balances = {};
  let statuses = {};
  onMount(async () => {
    balances = await get();
  });

  let newUser = "";
  const newUserAdd = () => {
    if (!newUser) return;
    balances[newUser] = 0;
    newUser = "";
    put();
  };

  const calculate = () => {
    let ordereds = 0;
    let wentDowns = 0;
    for (const name in balances) {
      if (statuses[name] === status.ORDERED) ordereds++;
      if (statuses[name] === status.WENT_DOWN) wentDowns++;
    }
    if (
      (ordereds === 0 && wentDowns === 0) ||
      (ordereds > 0 && wentDowns === 0) ||
      (ordereds === 0 && wentDowns > 0)
    ) {
      return;
    }
    for (const [name, balance] of Object.entries(balances)) {
      if (statuses[name] === status.WENT_DOWN) balances[name] += 1 / wentDowns;
      if (statuses[name] === status.ORDERED) balances[name] -= 1 / ordereds;
      statuses[name] = 0;
    }
  };

  const commit = () => {
    calculate();
    put();
  };
</script>

<div class="bg-gray-300 min-h-screen flex items-center justify-center">
  <div class="flex flex-col gap-8 p-2">
    {#if Object.keys(balances).length}
      <div
        class="block bg-gray-50 p-4 grid grid-cols-3 items-center gap-y-2 gap-x-4"
      >
        <div class="font-bold">name</div>
        <div class="font-bold text-right">balance</div>
        <div class="font-bold">action</div>
        {#each Object.entries(balances) as [name, balance]}
          <div class="truncate">{name}</div>
          <div class="text-right font-mono">{balance.toFixed(2)}</div>
          <!-- svelte-ignore a11y-no-onchange -->
          <select
            class="inp pr-8 border-gray-300"
            value={statuses[name] || status.WASNT_INVOLVED}
            on:change={(e) => {
              statuses[name] = e.target.value;
            }}
            class:bg-green-200={statuses[name] === status.ORDERED}
            class:bg-yellow-200={statuses[name] === status.WENT_DOWN}
          >
            <option value={status.WASNT_INVOLVED}>wasn't involved</option>
            <option value={status.ORDERED}>ordered</option>
            <option value={status.WENT_DOWN}>went down</option>
          </select>
        {/each}
        <button
          class="inp col-span-full bg-blue-400 text-white"
          on:click={commit}
        >
          commit
        </button>
      </div>
    {/if}
    <div class="block bg-gray-50 flex flex-col gap-2">
      <input
        class="inp border-gray-300"
        type="text"
        placeholder="Robert Wyatt"
        bind:value={newUser}
      />
      <button class="inp bg-blue-400 text-white" on:click={newUserAdd}>
        add new user
      </button>
    </div>
  </div>
</div>

<style global>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;

  @layer base {
    body {
      @apply m-0 p-0;
    }
  }

  @layer components {
    .inp {
      @apply rounded m-0 p-0 px-2;
    }
    .block {
      @apply p-4 rounded-lg shadow-lg;
    }
  }
</style>
