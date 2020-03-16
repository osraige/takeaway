<script>
    let people = {
        0: {"id": 0, "status": 0, "balance": 0, "name": 'senan'},
        1: {"id": 1, "status": 1, "balance": 0, "name": 'fao'},
        2: {"id": 2, "status": 1, "balance": 0, "name": 'ed'},
        3: {"id": 3, "status": 2, "balance": 0, "name": 'oskar'}
    };
    const updateStatus = (id, status) => {
        people[id].status = status
    }
    const statusClasses = [
        "status-not-involved",
        "status-ordered",
        "status-went-down"
    ]
    const commit = () => {
        let ordereds = 0
        let wentDowns = 0
        for (const [id, person] of Object.entries(people)) {
            if (person.status === 1) ordereds++
            if (person.status === 2) wentDowns++
        }
        if ((ordereds === 0 && wentDowns === 0) ||
            (ordereds > 0 && wentDowns === 0) || 
            (ordereds === 0 && wentDowns > 0)) {
            return
        }
        for (const [id, person] of Object.entries(people)) {
            if (person.status === 2) person.balance += 1/wentDowns
            if (person.status === 1) person.balance -= 1/ordereds
            people[person.id] = person
        }
    }
</script>

<style>
    table {
      table-layout: fixed;
      border-collapse: collapse;
      width: 100%;
    }
    td {
        border-left: solid 1px black;
    }
    th {
        text-align: left;
    }
    select, button {
        border: 0;
    }
    button {
        padding: 0 40px;
    }
    select {
        width: 100%;
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
    .button-row {
        text-align: right;
    }
</style>

<table>
    <tr>
        <th>name</th>
        <th>balance</th>
        <th>action</th>
    </tr>
    {#each Object.entries(people) as [id, person], id}
    <tr class="{statusClasses[person.status]}">
        <td>{person.name}</td>
        <td>{person.balance.toFixed(2)}</td>
        <td>
            <select
                value="{person.status.toString()}"
                on:change="{(e) => updateStatus(id, Number(e.target.value))}"
            >
                <option value="0">wasn't involved</option>
                <option value="1">ordered</option>
                <option value="2">went down</option>
            </select>
        </td>
    </tr>
    {/each}
</table>
<div class="button-row">
    <button type="button" on:click="{commit}">commit</button>
</div>
