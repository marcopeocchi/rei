<template>
  <div class="flex items-center justify-between">
    <div>
      <span class="text-5xl text-neutral-700 font-semibold" :class="theme.textHeading">Welcome to </span>
      <span class="text-5xl font-bold" :class="theme.textServerName">{{ config?.servername ?? '...' }}</span>
    </div>
    <select @change="store.setTheme($event.target.value)" :value="themeName"
      class="bg-neutral-200 dark:bg-neutral-700 rounded py-2.5 px-2.5 capitalize opacity-20 hover:opacity-100 duration-100">
      <option v-for="key in Object.keys(themes).sort()">
        {{ key }}
      </option>
    </select>
  </div>
  <header class="bg-neutral-200 dark:bg-neutral-800 p-8 rounded mt-8 font-mono">
    <div v-if="loadingTop === false">
      <Entry title="Hostname" :value="top.hostname" pad-to="1" />
      <Entry title="OS" :value="top.os" pad-to="7" />
      <Entry title="Platform" :value="top.platform" pad-to="1" />
      <Entry title="CPU" :value="`${top.coreCount}  x  ${top.cpu}`" pad-to="6" />
      <Entry title="Free RAM" :value="`${(Number(top.ramFree) / 1_000_000).toFixed(0)}MB`" pad-to="1" />
      <Entry title="Uptime" :value="secondsToHms(uptime)" pad-to="3" />
      <br />
      <Entry v-if="loadingTemp === false" title="PKG Temperature"
        :value="`${(Number(temp.cpuTemp) / 1000).toFixed(2)}°C`" pad-to="1" />
    </div>
  </header>
  <main v-if="loadingCfg === false"
    class="mt-10 grid grid-cols-2 sm:grid-cols-4 md:grid-cols-5 xl:grid-cols-6 2xl:grid-cols-7 gap-4">
    <a v-for="service in config.services" :href="service.url">
      <div class="p-3 text-neutral-700 rounded-lg text-center cursor-pointer duration-100" :class="theme.button">
        {{ service.name }}
      </div>
    </a>
  </main>
</template>

<script>
import { useThemeStore } from '../stores/theme'
import { computed } from 'vue'
import { themes } from '../themes'
import Entry from './Entry.vue'

export default {
  data() {
    return {
      top: null,
      temp: null,
      config: null,
      loadingTemp: true,
      loadingCfg: true,
      loadingTop: true,
      loading: true,
      uptime: 0,
      themes,
    }
  },
  setup() {
    const store = useThemeStore()
    return {
      theme: computed(() => store.getTheme),
      themeName: computed(() => store.themeName),
      store,
    }
  },
  created() {
    fetch("/config")
      .then(res => res.json())
      .then(data => {
        this.config = data
        this.loadingCfg = false
      })
    fetch("/top")
      .then(res => res.json())
      .then(data => {
        this.top = data
        this.uptime = data.uptime
        this.loadingTop = false
      })
    const fetchTemp = () => fetch("/temp")
      .then(res => res.json())
      .then(data => {
        this.temp = data
        this.loadingTemp = false
      })
    fetchTemp()
    setInterval(() => this.uptime++, 1000)
    setInterval(() => fetchTemp(), 3000)
  },
  methods: {
    secondsToHms: (d) => {
      d = Number(d)
      const h = Math.floor(d / 3600)
      const m = Math.floor(d % 3600 / 60)
      const s = Math.floor(d % 3600 % 60)
      const hFmt = h > 0 ? h + (h == 1 ? " hour, " : " hours, ") : ""
      const mFmt = m > 0 ? m + (m == 1 ? " minute, " : " minutes, ") : ""
      const sFmt = s > 0 ? s + (s == 1 ? " second" : " seconds") : ""
      const uptime = `${hFmt}${mFmt}${sFmt}`.trim()
      return uptime
    },
  },
  components: { Entry }
}
</script>