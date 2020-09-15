<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <div v-if="status">
      <p>Websocket {{ status }}</p>
      <button v-if="status === 'connected'" v-on:click="callAPI">Call API</button>
    </div>
    <div v-if="wsMessages.length">
      <p>Result from API:</p>
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
      connection: null,
      status: null,
    }
  },
  created() {
    this.startWS();
  },
  methods: {
    startWS: function () {
      this.connection = new WebSocket('ws://localhost:3000/ws');
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
      axios.post('http://localhost:8081/echo');
    }
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
