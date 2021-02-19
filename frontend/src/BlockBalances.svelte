<script>
  export let balances = {}
  export let statuses = {}

  import { createEventDispatcher } from 'svelte'
  import { status } from './api.js'

  const dispatch = createEventDispatcher()
</script>

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
    on:change={e => {
      dispatch('select', { name, status: e.target.value })
    }}
    class:bg-yellow-200={statuses[name] === status.ORDERED}
    class:bg-green-200={statuses[name] === status.WENT_DOWN}
  >
    <option value={status.WASNT_INVOLVED}>wasn't involved</option>
    <option value={status.ORDERED}>ordered</option>
    <option value={status.WENT_DOWN}>went down</option>
  </select>
{/each}

<button
  class="inp col-span-full bg-blue-400 text-white"
  on:click={() => dispatch('commit', null)}
>
  commit
</button>
