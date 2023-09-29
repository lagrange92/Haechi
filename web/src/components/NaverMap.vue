<template>
  <div id="naver_map" style="width: 100%; height: 400px"></div>
</template>

<script>
export default {
  components: {},
  name: "NaverMap",
  mounted() {
    this.initMap();
  },
  methods: {
    initMap() {
      const naver = window.naver;
      const mapOptions = {
        center: new naver.maps.LatLng(37.3595704, 127.105399),
        zoom: 10,
      };
      new naver.maps.Map("naver_map", mapOptions);
    },
  },
};
</script>

<style></style>
