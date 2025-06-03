document.addEventListener('alpine:init', () => {
  Alpine.data('context', () => ({
    init() {
      this.getMap();
    },
    loading: false,
    select: 'start',
    map: null,
    getMap() {
      this.loading = true
      fetch("/api/getMap").then(res => res.json()).then(data => {
        this.map = data
        this.loading = false
      })
    },
    clickCell(x, y) {
      console.info(`Clicked cell at (${x}, ${y})`)
      fetch("/api/click", {
        method: "POST",
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ x, y })
      }).then(res => res.json()).then(data => {
        console.info("Cell clicked:", data)
      }).catch(err => {
        console.error("Error clicking cell:", err)
      })

    }

  }))
})
