<template>
    <div class="container mt-4">
        <h2 class="text-center text-secondary pb-2">台北市營運餐廳</h2>
        <div class="map-container border rounded">
            <ul class="nav justify-content-center border-bottom">
                <!--營運地區 nav-->
            </ul>
            <!--地圖呈現在此-->
            <div class="google-map" id="map"></div>
        </div>
    </div>
</template>

<script>
import {onMounted} from "vue";

export default {
    name: "GoogleMapPractice",
    setup() {
        const lat = 25.0325917;
        const lng = 121.5624999;
        let map;

        onMounted(() => {
            initMap();
            setMarker();
        })

        function initMap() {
            // 透過 Map 物件建構子建立新地圖 map 物件實例，並將地圖呈現在 id 為 map 的元素中
            map = new google.maps.Map(document.getElementById("map"), {
                // 設定地圖的中心點經緯度位置
                center: { lat: lat, lng: lng },
                // 設定地圖縮放比例 0-20
                zoom: 15,
                // 限制使用者能縮放地圖的最大比例
                maxZoom: 20,
                // 限制使用者能縮放地圖的最小比例
                minZoom: 3,
                // 設定是否呈現右下角街景小人
                streetViewControl: false,
                // 設定是否讓使用者可以切換地圖樣式：一般、衛星圖等
                mapTypeControl: false
            });
        }

        function setMarker() {
            const marker = new google.maps.Marker({
                position: { lat: lat, lng: lng },
                map: map
            });
            // 透過 InfoWindow 物件建構子建立新訊息視窗
            const infowindow = new google.maps.InfoWindow({
                // 設定想要顯示的內容
                content: `
          <div id="content">
            <h3>看屁啊餐廳</h3>
            <p>一家專門在罵你看屁啊的餐廳唷<3</p>
          </div>
        `,
                // 設定訊息視窗最大寬度
                maxWidth: 200
            });
            // 在地標上監聽點擊事件
            marker.addListener("click", () => {
                // 指定在哪個地圖和地標上開啟訊息視窗
                infowindow.open(map, marker);
            });
        }
    }
}
</script>

<style scoped>
.google-map {
    width: 100%;
    height: 400px;
}
</style>