<template>
    <Submenu v-for="subMenu in subMenus" :key="subMenu.name" :name="subMenu.name">
        <template #title>
            <Icon :type="subMenu.iconType" />
            {{subMenu.nameCN}}
        </template>
        <MenuItem v-for="menuItem in subMenu.menuItems" :key="menuItem.path" :name="menuItem.path" :to="menuItem.path">
            <Icon v-if="menuItem.iconType !== ''" :type="subMenu.iconType" />
            {{menuItem.nameCN}}
        </MenuItem>
        <SubMenuComponent :subMenus="subMenu.subMenus" />
    </Submenu>
</template>

<script>
import {Icon, MenuItem, Submenu} from "view-ui-plus";

export default {
    name: "SubMenuComponent",
    components: {MenuItem, Icon, Submenu},
    props: {
        subMenus: Array
    },
    setup(props) {
        return {
            props
        }
    }
}
</script>

<style scoped>

</style>