document.addEventListener('alpine:init', () => {
  Alpine.data('context', () => ({
    init() {
      this.getMap();
    },
    loading: false,
    map: null,
    getMap() {
      this.loading = true
      fetch("/api/getMap").then(res => res.json()).then(data => {
        this.map = data
        this.loading = false
      })
    },

  }))
})
