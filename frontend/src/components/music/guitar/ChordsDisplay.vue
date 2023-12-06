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

export default {
    name: "ChordsDisplay",
    components: {Row, Space},
    setup() {
        const context = ref({});
        const canvasMaxWidth = 800;
        const canvasMaxHeight = 1200;
        const commonGrayColor = '#6F6F6F'; // 通用灰色

        const pointRadius = 10; // 音符的半徑
        const lineWidth = 2; // 線的寬度

        const chordsWidthSpace = 20; // 和弦的左右間距
        const chordsHeightSpace = 30; // 和弦的上下間距

        const lineNum = 6; // 弦的數量

        const chordsList = [
            {name: "C", points: [null, 3, 2, 0, 1, 0], baseLine: 0},
            {name: "Am", points: [null, 0, 2, 2, 1, 0], baseLine: 0},
            {name: "G", points: [3, 2, 0, 0, 0, 3], baseLine: 0},
            {name: "Em", points: [0, 2, 2, 0, 0, 0], baseLine: 0},
            {name: "F", points: [null, null, 3, 2, 1, 1], baseLine: 0},
            {name: "D", points: [null, null, 0, 2, 3, 2], baseLine: 0},
        ];

        onMounted(() => {
            const canvas = document.getElementById('canvas');
            context.value = canvas.getContext("2d");
            chordsList.forEach((chords, index) => {
                const modX = index % 4;
                const floorY = Math.floor(index / 4);
                drawChords(40 + modX * 160, 80 + floorY * 240, chords);
            });
        });

        function drawChords(startX, startY, chords) {
            context.value.font = "25px Arial";
            context.value.fillStyle = commonGrayColor;
            context.value.fillText(chords.name, startX + 1.5 * chordsWidthSpace, startY - 30);

            drawZeroChords(startX, startY);
            drawBaseLine(startX, startY, chords.baseLine);
            drawPoints(startX, startY, chords.points);
        }

        function drawPoints(startX, startY, points) {
            points.forEach((point, index) => {
                const x = startX + index * chordsWidthSpace;
                const y = startY + (point - 0.5) * chordsHeightSpace;
                if (point === null) {
                    context.value.font = "15px Arial";
                    context.value.fillStyle = commonGrayColor;
                    context.value.fillText("X", x - chordsWidthSpace/4, startY - 5);
                    return;
                }
                if (point === 0) {
                    return;
                }
                drawPoint(x, y);
            });
        }

        function drawBaseLine(startX, startY, baseLine = 0) {
            if (baseLine === 0) {
                return;
            }
            context.value.font = "18px Arial";
            context.value.fillStyle = commonGrayColor;
            context.value.fillText(baseLine, startX - 15, startY + 10);
        }

        function drawZeroChords(startX, startY) {
            const halfLineWidth = lineWidth / 2;
            const endX = startX + (lineNum - 1) * chordsWidthSpace + lineWidth;
            const endY = startY + (lineNum - 1) * chordsHeightSpace + lineWidth;
            let width = 3 * lineWidth;

            for (let i = 0; i < lineNum; i++) {
                const x = startX + i * chordsWidthSpace + halfLineWidth;
                const y = startY + i * chordsHeightSpace + halfLineWidth;

                if (i !== 0) {
                    width = lineWidth;
                }

                drawLine({start: startX, end: endX}, {start: y, end: y}, width);
                drawLine({start: x, end: x}, {start: startY, end: endY});
            }
        }

        function drawLine(x, y, width = lineWidth) {
            context.value.lineWidth = width;
            context.value.strokeStyle = commonGrayColor;
            context.value.beginPath();
            context.value.moveTo(x.start, y.start);
            context.value.lineTo(x.end, y.end);
            context.value.stroke();
        }

        function drawPoint(x, y) {
            const circle = new Path2D();
            circle.moveTo(0, 0);
            circle.arc(x, y, pointRadius, 0, 2 * Math.PI);

            context.value.fillStyle = commonGrayColor;
            context.value.fill(circle);
        }

        return {
            canvasMaxWidth,
            canvasMaxHeight,
        }
    }
}
</script>

<style scoped>

</style>