export const tabTo = (name: string, spaces: number) =>
  name.concat(new Array(spaces)
    .fill(' ', 0, spaces)
    .concat('→ ')
    .reduce((c, n) => `${c}${n}`)
  )

export const formatHHMMSS = (d: number) => {
  d = Number(d)
  const h = Math.floor(d / 3600)
  const m = Math.floor(d % 3600 / 60)
  const s = Math.floor(d % 3600 % 60)
  const hFmt = h > 0 ? h + (h == 1 ? ' hour, ' : ' hours, ') : ''
  const mFmt = m > 0 ? m + (m == 1 ? ' minute, ' : ' minutes, ') : ''
  const sFmt = s > 0 ? s + (s == 1 ? ' second' : ' seconds') : ''
  const uptime = `${hFmt}${mFmt}${sFmt}`.trim()

  return uptime.endsWith(',')
    ? uptime.substring(-1)
    : uptime
}