<template>
  <div>
    <div v-for="item in msg" :value="item.value" :key="item.value">{{ item.Value }}</div>
  </div>
</template>

<script>
    import axios from 'axios'
    export default {
        name: 'List',
        data() {
            return {
                msg: '',
                Value: '',
            }
        },
        methods: {
            getMessage() {
                const path = 'http://localhost:3000'
                axios.get(path)
                    .then((res) => {
                        this.msg = res.data
                    })
                    .catch((error) => {
                        console.error(error)
                    })
            }
        },
        created() {
            this.getMessage()
        }
    };
</script>