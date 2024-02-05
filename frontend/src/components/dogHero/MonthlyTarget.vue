<template>
    <Tabs :value="targets[0].tag">
        <TabPane v-for="item in targets" :key="item.tag"
                 :name="item.tag" :label="item.name">
            <MonthlyTargetTableVue :data="list[`${item.tag}`]" />
        </TabPane>
    </Tabs>
</template>

<script>
import MonthlyTargetTableVue from '@/components/dogHero/MonthlyTargetTable.vue';
import {ref} from "vue";
import {dogHeroApi} from "@/api";
import {TabPane, Tabs} from "view-ui-plus";

export default {
    name: "MonthlyTarget",
    components: {Tabs, TabPane, MonthlyTargetTableVue},
    setup() {
        const list = ref({});
        const targets = [
            {
                tag: "box",
                name: "累計寶箱",
            },
            {
                tag: "recruit",
                name: "累計招募"
            },
            {
                tag: "highCatch",
                name: "累計高級抓取",
            },
            {
                tag: "point",
                name: "累計花費點券"
            },
            {
                tag: "oil",
                name: "累計汽車燃油",
            },
            {
                tag: "food",
                name: "累計料理"
            },
            {
                tag: "arena",
                name: "累計競技場",
            },
            {
                tag: "lowCatch",
                name: "累計普通抓取"
            }
        ];

        async function getMonthlyTargetList() {
            await dogHeroApi.post("/monthlyTarget/list").then((response) => {
                list.value = response.data.data;
            }).catch(() => { console.log("error api"); });
        }

        getMonthlyTargetList();

        return {
            list,
            targets,
        }
    }
}
</script>

<style scoped>

</style>