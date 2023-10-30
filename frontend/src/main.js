import { createApp } from "vue";
import App from "./App.vue";
import ViewUIPlus from "view-ui-plus";
import "view-ui-plus/dist/styles/viewuiplus.css";
import { createRouter, createWebHistory } from "vue-router";
import Home from '@/components/Home.vue';
import DogHeroBoxCalculator from '@/components/dogHero/BoxCalculator.vue';
import DogHeroMonthlyTarget from '@/components/dogHero/MonthlyTarget.vue';

const router = createRouter({
    history: createWebHistory(),
    base: import.meta.env.BASE_URL,
    routes: [
        {
            path: "/",
            component: Home,
        },
        {
            path: "/dogHero/boxCalculator",
            component: DogHeroBoxCalculator,
        },
        {
            path: "/dogHero/monthlyTarget",
            component: DogHeroMonthlyTarget,
        }
    ],
});

createApp(App).use(ViewUIPlus).use(router).mount("#app");
