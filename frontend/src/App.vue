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
            <Menu :active-name="route.path" :open-names="['dogHero']">
                <Submenu name="dogHero">
                    <template #title>
                        <Icon type="md-game-controller-b" />
                        英雄你好狗
                    </template>
                    <MenuItem name="/dogHero/boxCalculator" to="/dogHero/boxCalculator">
                        寶箱計算器
                    </MenuItem>
                    <MenuItem name="/dogHero/monthlyTarget" to="/dogHero/monthlyTarget">
                        月度達標表
                    </MenuItem>
                </Submenu>
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

export default {
  name: 'App',
    components: {MenuItem,Menu,Submenu,Icon,Row,Image},
    setup() {
      const route = useRoute();

      return {
          Logo,
          route
      }
    }
}
</script>

<style>
</style>
