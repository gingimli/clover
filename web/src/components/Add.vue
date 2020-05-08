<template>
  <div>
    <form id="add" method="post" @submit="addItem" v-on:submit.prevent>
      <p><label>
        <input v-model="item">
      </label></p>
      <p>Add item: {{ item }}</p>
      <p><input type="submit" value="Submit"></p>
    </form>
  </div>
</template>

<script>
    import axios from 'axios'
    export default {
        name: 'Add',
        data() {
            return {
                item: '',
            }
        },
        methods: {
            addItem(payload) {
                const path = `http://localhost:3000/add?task=${this.item}`
                axios.post(path, payload)
                    .catch((error) => {
                        console.error(error)
                    })
                this.item = ''
            },
        }
    }
</script>