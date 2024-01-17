import { createApp } from "vue";
import App from "./App.vue";
import ViewUIPlus from "view-ui-plus";
import "view-ui-plus/dist/styles/viewuiplus.css";
import { createRouter, createWebHistory } from "vue-router";
import Home from '@/components/Home.vue';
import DogHeroBoxCalculator from '@/components/dogHero/BoxCalculator.vue';
import DogHeroMonthlyTarget from '@/components/dogHero/MonthlyTarget.vue';
import MusicPianoSightSeeingPractice from '@/components/music/piano/SightSeeingPractice.vue';
import MusicGuitarChordsDisplay from '@/components/music/guitar/ChordsDisplay.vue';
import MusicGuitarDiatonicChordsDisplay from '@/components/music/guitar/DiatonicChordsDisplay.vue';
import PracticeThreeJs from '@/components/practice/threeJs/ThreeJsPractice.vue';

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
        },
        {
            path: "/music/piano/sightSeeingPractice",
            component: MusicPianoSightSeeingPractice,
        },
        {
            path: "/music/guitar/chordsDisplay",
            component: MusicGuitarChordsDisplay,
        },
        {
            path: "/music/guitar/diatonicChordsDisplay",
            component: MusicGuitarDiatonicChordsDisplay,
        },
        {
            path: "/practice/threeJs/threeJsPractice",
            component: PracticeThreeJs,
        }
    ],
});

createApp(App).use(ViewUIPlus).use(router).mount("#app");
