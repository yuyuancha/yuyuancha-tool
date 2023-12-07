<template>
    <router-link to="/">
        <Row justify="center" align="middle">
            <Image :src="Logo" fit="fill" width="80px" height="80px" alt="logo" style="margin: 20px 20px;" />
            <h1 style="text-align: center;">芋圓茶工具箱</h1>
        </Row>
    </router-link>
    <hr class="ivu-mb" />
    <Row>
        <Col span="5" style="margin-right: 20px;">
            <Menu :active-name="route.path" :open-names="openNames">
                <SubMenuComponent :subMenus="subMenus" />
            </Menu>
        </Col>
        <Col span="18">
            <router-view></router-view>
        </Col>
    </Row>
</template>

<script>
import {useRoute} from "vue-router";
import {MenuItem,Menu,Submenu,Icon,Row,Image} from "view-ui-plus";
import Logo from '@images/yuyuancha.png';
import {ref} from "vue";
import SubMenuComponent from "@/components/SubMenuComponent.vue";

export default {
  name: 'App',
    components: {SubMenuComponent, MenuItem,Menu,Submenu,Icon,Row,Image},
    setup() {
      const route = useRoute();
      const subMenus = ref([
          {
              name: "dogHero",
              iconType: "md-game-controller-b",
              nameCN: "英雄你好狗",
              menuItems: [
                  { path: "/dogHero/boxCalculator", iconType: "", nameCN: "寶箱計算器" },
                  { path: "/dogHero/monthlyTarget", iconType: "", nameCN: "月度達標表" }
              ],
              subMenus: [],
          },
          {
              name: "music",
              iconType: "md-game-controller-b",
              nameCN: "音樂",
              menuItems: [],
              subMenus: [
                  {
                      name: "piano",
                      iconType: "md-game-controller-b",
                      nameCN: "鋼琴",
                      menuItems: [
                          { path: "/music/piano/sightSeeingPractice", iconType: "", nameCN: "視譜練習" }
                      ],
                      subMenus: [],
                  },
                  {
                      name: "guitar",
                      iconType: "md-game-controller-b",
                      nameCN: "吉他",
                      menuItems: [
                          { path: "/music/guitar/chordsDisplay", iconType: "", nameCN: "和弦展示" },
                          { path: "/music/guitar/diatonicChordsDisplay", iconType: "", nameCN: "順階和弦展示" }
                      ],
                      subMenus: [],
                  }
              ],
          }
      ]);

      const openNames = ref([]);

      pushOpenNames(subMenus.value);

      function pushOpenNames(subMenus) {
          subMenus.forEach((subMenu) => {
              openNames.value.push(subMenu.name);
              if (subMenu.subMenus.length > 0) {
                  pushOpenNames(subMenu.subMenus);
              }
          });
      }

      return {
          Logo,
          route,
          subMenus,
          openNames
      }
    }
}
</script>

<style>
</style>
