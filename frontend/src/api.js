export const get = async () => {
  const response = await fetch(`/get?t=${Date.now()}`)
  if (!response.ok) return {}
  return await response.json()
}

export const put = balances => {
  fetch('/put', {
    method: 'POST',
    'Content-Type': 'application/json',
    body: JSON.stringify(balances),
  })
}

export const status = {
  WASNT_INVOLVED: 'wasnt_involved',
  ORDERED: 'ordered',
  WENT_DOWN: 'went_down',
}

export const calculate = (balances, statuses) => {
  let ordereds = 0
  let wentDowns = 0
  for (const name in balances) {
    if (statuses[name] === status.ORDERED) ordereds++
    if (statuses[name] === status.WENT_DOWN) wentDowns++
  }
  if (
    (ordereds === 0 && wentDowns === 0) ||
    (ordereds > 0 && wentDowns === 0) ||
    (ordereds === 0 && wentDowns > 0)
  ) {
    return
  }
  for (const [name, balance] of Object.entries(balances)) {
    if (statuses[name] === status.WENT_DOWN) balances[name] += 1 / wentDowns
    if (statuses[name] === status.ORDERED) balances[name] -= 1 / ordereds
    statuses[name] = 0
  }
  return [balances, statuses]
}
