<script>
import { computed } from 'vue';
import { useThemeStore } from '../stores/theme'
export default {
  props: {
    title: String,
    value: String,
    padTo: String,
  },
  methods: {
    tab: (name, spaces) => name.concat(new Array(spaces)
      .fill(' ', 0, spaces)
      .concat('→ ')
      .reduce((c, n) => `${c}${n}`)
    )
  },
  setup() {
    const store = useThemeStore()
    return {
      theme: computed(() => store.getTheme)
    }
  }
}
</script>

<template>
  <div>
    <span class="font-bold whitespace-pre-wrap" :class="theme.textPrimary">
      {{ tab(title, Number(padTo)) }}
    </span>
    <span>{{ value }}</span>
  </div>
</template>