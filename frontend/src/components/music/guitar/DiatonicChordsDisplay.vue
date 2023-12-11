<template>
    <h3>順階和弦展示</h3>
    <Space direction="vertical" :width="canvasMaxWidth">
        <Space direction="horizontal" style="margin-top: 30px;">
            <Button v-for="chords in diatonicChords" :key="chords.name" @click="selectDiatonicChords(chords)">{{chords.name}}</Button>
        </Space>
        <Space direction="horizontal" style="margin-top: 30px;">
            <Text style="font-size: 20px;">{{diatonicChordsName}}</Text>
        </Space>
        <Row :width="canvasMaxWidth" justify="space-around">
            <canvas id="canvas" :width="canvasMaxWidth" :height="canvasMaxHeight"></canvas>
        </Row>
    </Space>
</template>

<script>
import {onMounted, ref} from "vue";
import {Row, Space, Text} from "view-ui-plus";
import {drawChords} from "@js/guitar/canvas";
import diatonicChords from "@js/guitar/diatonicChords";
import ChordsItem from "@/components/music/guitar/ChordsItem.vue";

export default {
    name: "DiatonicChordsDisplay",
    components: {ChordsItem, Text, Row, Space},
    setup() {
        const context = ref({});
        const canvasMaxWidth = 800;
        const canvasMaxHeight = 1200;

        const diatonicChordsName = ref("");

        function selectDiatonicChords(diatonicChords) {
            context.value.clearRect(0, 0, canvasMaxWidth, canvasMaxHeight);
            diatonicChordsName.value = diatonicChords.name + "順階和弦展示";
            diatonicChords.chords.forEach((chords, index) => {
                const modX = index % 4;
                const floorY = Math.floor(index / 4);
                drawChords(context.value, 40 + modX * 160, 80 + floorY * 240, chords);
            });
        }

        onMounted(() => {
            const canvas = document.getElementById('canvas');
            context.value = canvas.getContext("2d");
            selectDiatonicChords(diatonicChords[0]);
        });

        return {
            canvasMaxWidth,
            canvasMaxHeight,
            diatonicChords,
            diatonicChordsName,
            selectDiatonicChords
        }
    }
}
</script>

<style scoped>

</style>