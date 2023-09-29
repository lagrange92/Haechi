<template>
  <div id="naver_map" style="width: 100%; height: 100vh"></div>
</template>

<script>
import axios from "axios";

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
        minZoom: 12,
        zoom: 12,
      };
      const naverMap = new naver.maps.Map("naver_map", mapOptions);

      //   var heatmapData = [
      //   new naver.maps.visualization.WeightedLocation(37.541, 126.986, 0.1),
      //     new naver.maps.visualization.WeightedLocation(37.551, 126.976, 0.2),
      //     new naver.maps.visualization.WeightedLocation(37.561, 126.966, 0.3),
      //     new naver.maps.visualization.WeightedLocation(37.571, 126.956, 0.4),
      //     new naver.maps.visualization.WeightedLocation(37.581, 126.946, 0.5),
      //     new naver.maps.visualization.WeightedLocation(37.531, 126.936, 0.6),
      //     new naver.maps.visualization.WeightedLocation(37.521, 126.926, 0.7),
      //     new naver.maps.visualization.WeightedLocation(37.511, 126.916, 0.8),
      //     new naver.maps.visualization.WeightedLocation(37.501, 126.906, 0.9),
      //     new naver.maps.visualization.WeightedLocation(37.541, 126.996, 1.0),
      //     // new naver.maps.visualization.WeightedLocation(37.541, 126.986, 1),
      //     // new naver.maps.visualization.WeightedLocation(37.551, 126.976, 1),
      //     // new naver.maps.visualization.WeightedLocation(37.561, 126.966, 1),
      //     // new naver.maps.visualization.WeightedLocation(37.571, 126.956, 0.8),
      //     // new naver.maps.visualization.WeightedLocation(37.581, 126.946, 0.7),
      //     // new naver.maps.visualization.WeightedLocation(37.531, 126.936, 1),
      //     // new naver.maps.visualization.WeightedLocation(37.521, 126.926, 0.9),
      //     // new naver.maps.visualization.WeightedLocation(37.511, 126.916, 0.8),
      //     // new naver.maps.visualization.WeightedLocation(37.501, 126.906, 0.8),
      //     // new naver.maps.visualization.WeightedLocation(37.541, 126.996, 1.0),
      //   ];

      var heatmapData = [];

      naver.maps.onJSContentLoaded = async function () {
        try {
          const response = await axios.get("http://localhost:1323/ppltn");
          //   console.log(response.data);

          for (let i = 0; i < response.data.length; i++) {
            heatmapData.push(
              new naver.maps.visualization.WeightedLocation(
                parseFloat(response.data[i].areaLatitude),
                parseFloat(response.data[i].areaLongitude),
                parseFloat(response.data[i].areaAvgPpltn)
              )
            );
          }

          console.log(heatmapData);

          new naver.maps.visualization.HeatMap({
            map: naverMap,
            data: heatmapData,
            opacity: 0.7,
            radius: 70,
            colorMap: naver.maps.visualization.SpectrumStyle.RdBu,
          });
        } catch (error) {
          console.error("API 호출 중 오류 발생:", error);
        }
      };
    },
  },
};
</script>

<style></style>
