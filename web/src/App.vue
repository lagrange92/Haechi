<template>
  <div id="app">
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
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
</style>
