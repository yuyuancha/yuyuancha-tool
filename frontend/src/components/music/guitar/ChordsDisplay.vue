<template>
    <h3>和弦展示</h3>
    <Space direction="vertical" :width="canvasMaxWidth">
        <Row :width="canvasMaxWidth" justify="space-around">
            <canvas id="canvas" :width="canvasMaxWidth" :height="canvasMaxHeight"></canvas>
        </Row>
    </Space>
</template>

<script>
import {Row, Space} from "view-ui-plus";
import {onMounted, ref} from "vue";
import {drawChords} from "@js/guitar/canvas";
import chords from "@js/guitar/chords";

export default {
    name: "ChordsDisplay",
    components: {Row, Space},
    setup() {
        const context = ref({});
        const canvasMaxWidth = 800;
        const canvasMaxHeight = 1200;

        const chordsList = [
            chords["C"],
            chords["Am"],
            chords["G"],
            chords["Em"],
            chords["F"],
            chords["D"],
        ];

        onMounted(() => {
            const canvas = document.getElementById('canvas');
            context.value = canvas.getContext("2d");
            chordsList.forEach((chords, index) => {
                const modX = index % 4;
                const floorY = Math.floor(index / 4);
                drawChords(context.value, 40 + modX * 160, 80 + floorY * 240, chords);
            });
        });

        return {
            canvasMaxWidth,
            canvasMaxHeight,
        }
    }
}
</script>

<style scoped>

</style>