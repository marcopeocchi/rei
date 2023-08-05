import { readable } from 'svelte/store'
import type { Thermals, Top } from '../types'

const fetchTop = async () => {
  const res = await fetch('/top')
  const data: Top = await res.json()
  return data
}

const fetchThermals = async () => {
  const res = await fetch('/temp')
  const data: Thermals = await res.json()
  return {
    ...data,
    cpuTemp: data.cpuTemp.replace('/n', '')
  } as Thermals
}

export const topStore = readable<Promise<Top>>(fetchTop())

export const temperatureStore = readable<Promise<Thermals>>(
  fetchThermals(),
  (set) => {
    const interval = setInterval(() => {
      set(fetchThermals())
    }, 5000)

    return () => clearInterval(interval)
  }
)