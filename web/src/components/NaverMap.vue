<template>
  <div id="naver_map" style="width: 100%; height: 100vh"></div>
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
        center: new naver.maps.LatLng(37.541, 126.986),
        zoom: 12,
      };
      const naverMap = new naver.maps.Map("naver_map", mapOptions);

      var heatmapData = [
        new naver.maps.visualization.WeightedLocation(37.541, 126.986, 0.1),
        new naver.maps.visualization.WeightedLocation(37.551, 126.976, 0.2),
        new naver.maps.visualization.WeightedLocation(37.561, 126.966, 0.3),
        new naver.maps.visualization.WeightedLocation(37.571, 126.956, 0.4),
        new naver.maps.visualization.WeightedLocation(37.581, 126.946, 0.5),
        new naver.maps.visualization.WeightedLocation(37.531, 126.936, 0.6),
        new naver.maps.visualization.WeightedLocation(37.521, 126.926, 0.7),
        new naver.maps.visualization.WeightedLocation(37.511, 126.916, 0.8),
        new naver.maps.visualization.WeightedLocation(37.501, 126.906, 0.9),
        new naver.maps.visualization.WeightedLocation(37.541, 126.996, 1.0),
        // new naver.maps.visualization.WeightedLocation(37.541, 126.986, 1),
        // new naver.maps.visualization.WeightedLocation(37.551, 126.976, 1),
        // new naver.maps.visualization.WeightedLocation(37.561, 126.966, 1),
        // new naver.maps.visualization.WeightedLocation(37.571, 126.956, 0.8),
        // new naver.maps.visualization.WeightedLocation(37.581, 126.946, 0.7),
        // new naver.maps.visualization.WeightedLocation(37.531, 126.936, 1),
        // new naver.maps.visualization.WeightedLocation(37.521, 126.926, 0.9),
        // new naver.maps.visualization.WeightedLocation(37.511, 126.916, 0.8),
        // new naver.maps.visualization.WeightedLocation(37.501, 126.906, 0.8),
        // new naver.maps.visualization.WeightedLocation(37.541, 126.996, 1.0),
      ];

      naver.maps.onJSContentLoaded = function () {
        new naver.maps.visualization.HeatMap({
          map: naverMap,
          data: heatmapData,
          opacity: 1,
          radius: 50,
          colorMap: naver.maps.visualization.SpectrumStyle.RAINBOW,
        });
      };
    },
  },
};
</script>

<style></style>
