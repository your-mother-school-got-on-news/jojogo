var info = []

const app = Vue.createApp({
  data() {
    axios
      .get('http://localhost:8081/group/view')
      .then(response => {
        this.info = response["data"];
      })
    return{
      info,
    } 
  },
  // mounted () {
  //   axios
  //     .get('http://localhost:8081/group/view') // 'http://localhost:8081/group/view'
  //     .then(response => (this.info = response["data"])) // (this.info = response["data"]) this.info.push(response)
  // },
})

app.mount('#app')
