<template>
  <div id="app">
    <p>Welcome to Haechi!!!</p>
    <NaverMap />
  </div>
</template>

<script>
import NaverMap from "./components/NaverMap.vue";

export default {
  name: "App",
  components: {
    NaverMap,
  },
};
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
