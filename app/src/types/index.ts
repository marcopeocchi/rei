export type Config = {
  servername: string
  services: Service[]
}

export type Service = {
  name: string
  url: string
}

export type Top = {
  cpu: string
  coreCount: number
  hostname: string
  os: string
  platform: string
  uptime: number
  cpuTemp: string
  ramFree: number
}

export type Thermals = {
  cpuTemp: string
}