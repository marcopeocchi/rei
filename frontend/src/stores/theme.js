import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { themes } from '../themes'

export const useThemeStore = defineStore('theme', () => {
  const theme = ref(localStorage.getItem("theme") ?? "sakura")
  const themeName = computed(() => theme.value)
  const getTheme = computed(() => themes[theme.value] ?? themes.sakura)

  function setTheme(name) {
    theme.value = name
    localStorage.setItem("theme", name)
  }

  return { theme, getTheme, themeName, setTheme }
})
