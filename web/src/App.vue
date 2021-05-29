<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <p>NATS server {{ this.server }}</p>
    <button v-on:click="callAPI">Call API</button>
    <div v-if="messages.length">
      <p>Result from API: ({{ messages.length }})</p>
      <p v-for="(message, index) in messages" v-bind:key="index">{{ index }} - {{ message }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import { connect, StringCodec } from 'nats.ws/nats.cjs'

export default {
  name: 'App',
  data() {
    return {
      messages: [],
      connection: undefined,
      status: undefined,
      server: 'ws://127.0.0.1:5000',
    }
  },
  created() {
    void this.connectToNATS();
  },
  methods: {
    connectToNATS: async function () {
      const nc = await connect({servers: this.server});
      console.log(nc.status());
      const sc = StringCodec();
      const sub = nc.subscribe('com.scaling-ws.updates')
      void await (async () => {
        for await (const m of sub) {
          this.messages.push(sc.decode(m.data));
          console.log(`[${sub.getProcessed()}]: ${sc.decode(m.data)}`);
        }
        console.log("subscription closed");
      })();
    },
    callAPI: function () {
      axios.post('http://localhost:3000/ping');
    },
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
