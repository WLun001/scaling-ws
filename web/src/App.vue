<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <p>Using Websocket {{ usingServer }}</p>
    <div v-if="status">
      <p>Websocket {{ status }}</p>
      <div v-if="status === 'connected'">
        <button v-on:click="callAPI">Call API</button>
      </div>
    </div>
    <div v-if="wsMessages.length">
      <p>Result from API: ({{ wsMessages.length }})</p>
      <p v-for="(message, index) in wsMessages" v-bind:key="index">{{ index }} - {{ message }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'App',
  data() {
    return {
      wsMessages: [],
      connection: undefined,
      status: undefined,
      usingServer: 'server 1',
    }
  },
  created() {
    const useServer2 = confirm('switch to ws sever 2');
    this.usingServer = useServer2 ? 'server 2' : 'server 1';
    this.startWS();
  },
  methods: {
    startWS: function () {
      const port = this.usingServer === 'server 1' ? 3000 : 4000;
      this.connection = new WebSocket(`ws://localhost:${port}/ws`);
      this.connection.onmessage = event => {
        this.wsMessages.push(event.data)
      }

      this.connection.onopen = () => {
        this.status = 'connected';
      }

      this.connection.onerror = () => {
        this.status = 'error';
      }

      this.connection.onclose = () => {
        this.status = 'disconnected';
        setTimeout(() => this.startWS(), 1000)
      }
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
