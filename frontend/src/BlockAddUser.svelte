<script>
  import { createEventDispatcher } from 'svelte'

  const dispatch = createEventDispatcher()

  let name = ''
</script>

<input
  class="inp border-gray-300"
  type="text"
  placeholder="Robert Wyatt"
  bind:value={name}
/>

<button
  class="inp bg-blue-400 text-white"
  on:click={() => {
    dispatch('add', name)
    name = ''
  }}
>
  add new user
</button>
