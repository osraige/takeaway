<script>
  import { onMount } from 'svelte'
  import * as api from './api.js'
  import Tailwind from './Tailwind.svelte'
  import BlockBalances from './BlockBalances.svelte'
  import BlockAddUser from './BlockAddUser.svelte'

  let balances = {}
  let statuses = {}
  onMount(async () => {
    balances = await api.get()
  })

  const addUser = name => {
    if (!name) return
    balances[name] = 0
    api.put(balances)
  }

  const commit = () => {
    ;[balances, statuses] = api.calculate(balances, statuses)
    api.put(balances)
  }
</script>

<Tailwind />

<div class="bg-gray-300 min-h-screen flex items-center justify-center">
  <div class="flex flex-col gap-8 p-2">
    {#if Object.keys(balances).length}
      <div class="block grid grid-cols-3 items-center gap-2">
        <BlockBalances
          {balances}
          {statuses}
          on:select={e => (statuses[e.detail.name] = e.detail.status)}
          on:commit={e => commit()}
        />
      </div>
    {/if}
    <div class="block flex flex-col gap-2">
      <BlockAddUser on:add={e => addUser(e.detail)} />
    </div>
  </div>
</div>
