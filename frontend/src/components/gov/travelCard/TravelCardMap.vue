<template>
    <div>
        <h3>國民旅遊卡地圖</h3>
        <Space direction="vertical" style="margin: 20px 0;">
            <h4>選擇顯示行業別</h4>
            <Select v-model="selectCategory" @on-change="changeSelectCategory" size="default" style="width: 300px;">
                <Option v-for="category in categories" :value="category.id" :key="category.id">{{ category.name }}</Option>
            </Select>
        </Space>
        <div class="map-container border rounded">
            <ul class="nav justify-content-center border-bottom">
            </ul>
            <div class="google-map" id="map"></div>
        </div>
    </div>
</template>

<script>
import {onBeforeMount, onMounted, ref} from "vue";
import {govApi} from "@/api";
import {Space} from "view-ui-plus";

export default {
    name: "TravelCardMap",
    components: {Space},
    setup() {
        let lat = 25.0325917;
        let lng = 121.5624999;
        const shops = ref([]);
        const categories = ref([]);
        const selectCategory = ref(0);

        let map;

        onMounted(async () => {
            console.log("最後：", lat, lng);

            initGoogleMap();

            await getTravelCardCategories();
            selectCategory.value = categories.value[0].id;

            await getTravelCardShops();
            setShopMarkers();
        });

        async function changeSelectCategory() {
            shops.value = [];
            initGoogleMap();
            await getTravelCardShops();
            setShopMarkers();
        }

        async function getTravelCardShops() {
            const data = {
                categoryId: selectCategory.value,
            };

            await govApi.post("/travelCard/shops", data).then((response) => {
                response.data.data.list.forEach((item) => {
                    shops.value.push(item);
                });
            }).catch(() => { console.log("error api"); });
        }

        async function getTravelCardCategories() {
            await govApi.post("/travelCard/categories").then((response) => {
                response.data.data.forEach((item) => {
                    categories.value.push({
                        id: item.id,
                        name: item.name,
                    });
                });
            }).catch(() => { console.log("error api"); });
        }

        function initGoogleMap() {
            navigator.geolocation.getCurrentPosition((position) => {
                console.log("偵測：", position.coords.latitude, position.coords.longitude)
                lat = position.coords.latitude;
                lng = position.coords.longitude;
                map = new google.maps.Map(document.getElementById("map"), {
                    center: { lat: lat, lng: lng },
                    zoom: 15,
                    maxZoom: 20,
                    minZoom: 3,
                    streetViewControl: false,
                    mapTypeControl: false
                });
            });

        }

        function setShopMarkers() {
            shops.value.forEach((shop) => {
                setGoogleMapMaker(shop);
            });
        }

        function setGoogleMapMaker(shop) {
            const marker = new google.maps.Marker({
                position: { lat: shop.latitude, lng: shop.longitude },
                map: map,
            });

            const categoryName = categories.value.find((category) => category.id === shop.category_id).name;
            const infoWindow = new google.maps.InfoWindow({
                content: `
                  <div id="content">
                    <h3>${shop.name}</h3>
                    <p>行業別: ${categoryName}</p>
                    <p>地址: ${shop.address}</p>
                    <p>電話: ${shop.phone_number}</p>
                  </div>
                `,
                maxWidth: 300
            });

            marker.addListener("click", () => {
                infoWindow.open(map, marker);
            });
        }

        return {
            categories,
            selectCategory,
            changeSelectCategory,
        };
    }
}
</script>

<style scoped>
.google-map {
    width: 100%;
    height: 50%;
}
</style>