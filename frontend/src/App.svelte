<script>
  import { onMount } from "svelte";
  let balances = {};
  let statuses = {};
  let newUser = "";
  onMount(async () => {
    const response = await fetch(`/get?t=${Date.now()}`);
    if (response.ok) {
      balances = await response.json();
    }
  });
  const statusClasses = [
    "status-not-involved",
    "status-ordered",
    "status-went-down"
  ];
  const put = () => {
    fetch("/put", {
      method: "POST",
      "Content-Type": "application/json",
      body: JSON.stringify(balances)
    });
  };
  const commit = () => {
    let ordereds = 0;
    let wentDowns = 0;
    for (const name in balances) {
      if (statuses[name] === 1) ordereds++;
      if (statuses[name] === 2) wentDowns++;
    }
    if (
      (ordereds === 0 && wentDowns === 0) ||
      (ordereds > 0 && wentDowns === 0) ||
      (ordereds === 0 && wentDowns > 0)
    ) {
      return;
    }
    for (const [name, balance] of Object.entries(balances)) {
      if (statuses[name] === 2) balances[name] += 1 / wentDowns;
      if (statuses[name] === 1) balances[name] -= 1 / ordereds;
      statuses[name] = 0;
    }
    put();
  };
  const newUserAdd = () => {
    balances[newUser] = 0;
    newUser = "";
    put();
  };
</script>

<style>
  table {
    table-layout: fixed;
    border-collapse: collapse;
    width: 100%;
  }
  th {
    text-align: left;
  }
  select,
  input,
  button {
    background-color: rgba(0, 0, 0, 0.2);
    border: 1px solid #aaa;
    box-sizing: border-box;
    height: 18px;
    margin: 2px 0;
    vertical-align: top;
    width: 100%;
  }
  button,
  input {
    background-color: white;
  }
  .status-not-involved td {
    background-color: white;
  }
  .status-ordered td {
    background-color: #ffbd45;
  }
  .status-went-down td {
    background-color: #aee571;
  }
</style>

<table>
  {#if Object.keys(balances).length}
    <tr>
      <th>name</th>
      <th>balance</th>
      <th>action</th>
    </tr>
    {#each Object.entries(balances) as [name, balance]}
      <tr class={statusClasses[statuses[name] || 0]}>
        <td>{name}</td>
        <td>{balance.toFixed(2)}</td>
        <td>
          <select
            value={(statuses[name] || 0).toString()}
            on:change={e => (statuses[name] = Number(e.target.value))}
          >
            <option value="0">wasn't involved</option>
            <option value="1">ordered</option>
            <option value="2">went down</option>
          </select>
        </td>
      </tr>
    {/each}
    <tr>
      <td colspan="2" />
      <td>
        <button type="button" on:click={commit}>commit</button>
      </td>
    </tr>
  {/if}
  <tr>
    <td colspan="2">
      <input bind:value={newUser} placeholder="add a new user" />
    </td>
    <td>
      <button on:click={newUserAdd}>create</button>
    </td>
  </tr>
</table>
