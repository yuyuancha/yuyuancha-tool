<template>
    <h3>視譜練習</h3>
    <Space direction="vertical" :width="canvasMaxWidth">
        <Row :width="canvasMaxWidth" justify="space-around">
            <canvas id="canvas" :width="canvasMaxWidth" :height="canvasMaxHeight"></canvas>
        </Row>
        <Row :width="canvasMaxWidth" justify="space-around">
            <Text style="font-size: 30px;">
                {{message}}
            </Text>
        </Row>
        <Row :width="canvasMaxWidth" justify="space-around">
            <Grid :col="7" :border="false" padding="0" :style="`width: ${canvasMaxWidth}px;`">
                <GridItem v-for="note in notes" :key="note.name">
                    <Space style="width: 100%; height: 50px;">
                        <Button @click="clickNote(note.name)">{{note.name}}</Button>
                    </Space>
                </GridItem>
            </Grid>
        </Row>
    </Space>
</template>

<script>
import {onMounted, ref} from "vue";
import {Grid, GridItem, Row, Space, Text} from "view-ui-plus";

export default {
    name: "SightSeeingPractice",
    components: {Grid, GridItem, Text, Row, Space},
    setup() {
        const canvasMaxWidth = 500;
        const canvasMaxHeight = 150;
        const commonGrayColor = '#6F6F6F'; // 通用灰色

        const noteRadius = 10; // 音符的半徑

        const lineSpacing = 20; // 線的間距
        const linePadding = 50; // 線的左右間距
        const sightLineWidth = 2; // 線的寬度
        const topLineY = canvasMaxHeight / 2 - 2 * lineSpacing; // 五線譜第一條的 y 座標
        const middleX = canvasMaxWidth / 2; // 五線譜中間的 x 座標

        const message = ref("");
        const questionNote = ref({});
        const context = ref({});

        const pointY = {
            line6: topLineY - lineSpacing,
            grid5: topLineY - 0.5 * lineSpacing,
            line5: topLineY,
            grid4: topLineY + 0.5 * lineSpacing,
            line4: topLineY + lineSpacing,
            grid3: topLineY + 1.5 * lineSpacing,
            line3: topLineY + 2 * lineSpacing,
            grid2: topLineY + 2.5 * lineSpacing,
            line2: topLineY + 3 * lineSpacing,
            grid1: topLineY + 3.5 * lineSpacing,
            line1: topLineY + 4 * lineSpacing,
            grid0: topLineY + 4.5 * lineSpacing,
            line0: topLineY + 5 * lineSpacing,
        };

        const notes = {
            12: { name: "高Sol", note: "高音G", y: pointY.grid5 },
            11: { name: "高Fa", note: "高音F", y: pointY.line5 },
            10: { name: "高Mi", note: "高音E", y: pointY.grid4 },
            9: { name: "高Re", note: "高音D", y: pointY.line4 },
            8: { name: "高Do", note: "高音C", y: pointY.grid3 },
            7: { name: "Si", note: "B", y: pointY.line3 },
            6: { name: "La", note: "A", y: pointY.grid2 },
            5: { name: "Sol", note: "G", y: pointY.line2 },
            4: { name: "Fa", note: "F", y: pointY.grid1 },
            3: { name: "Mi", note: "E", y: pointY.line1 },
            2: { name: "Re", note: "D", y: pointY.grid0 },
            1: { name: "Do", note: "C", y: pointY.line0 }
        };

        onMounted(() => {
            const canvas = document.getElementById('canvas');
            context.value = canvas.getContext("2d");

            drawQuestion();
        });

        function drawQuestion() {
            context.value.clearRect(0, 0, canvasMaxWidth, canvasMaxHeight);

            drawSightLine();
            drawTrebleClef();

            randomQuestion();
            drawNote(questionNote.value, middleX);
        }

        function drawSightLine() {
            for (let i = 0; i < 5; i++) {
                drawLine(
                    {start: linePadding, end: canvasMaxWidth - linePadding},
                    {start: topLineY + i * lineSpacing, end: topLineY + i * lineSpacing}
                );
            }
        }

        function drawTrebleClef() {
            context.value.font = "160px Arial";
            context.value.fillStyle = commonGrayColor;
            context.value.fillText("𝄞", 60, pointY.grid0);
        }

        function drawLine(x, y) {
            context.value.lineWidth = sightLineWidth;
            context.value.strokeStyle = commonGrayColor;
            context.value.beginPath();
            context.value.moveTo(x.start, y.start);
            context.value.lineTo(x.end, y.end);
            context.value.stroke();
        }

        function drawNote(note, x) {
            drawPoint(x, note.y);

            if (note.y === pointY.line0) {
                drawLine(
                    {start: x - 15, end: x + 15},
                    {start: pointY.line0, end: pointY.line0}
                );
            }
        }

        function drawPoint(x, y) {
            const circle = new Path2D();
            circle.moveTo(0, 0);
            circle.arc(x, y, noteRadius, 0, 2 * Math.PI);

            context.value.fillStyle = commonGrayColor;
            context.value.fill(circle);
        }

        function clickNote(note) {
            if (note === questionNote.value.name) {
                drawQuestion();
                return;
            }

            message.value = `答錯了！這個音是 ${questionNote.value.name} 喔！`;
        }

        function randomQuestion() {
            const newNote = notes[Math.floor(Math.random() * Object.keys(notes).length) + 1];
            if (newNote.name === questionNote.value.name) {
                randomQuestion();
                return;
            }
            questionNote.value = newNote;
            message.value = `請問這個音符是?`;
        }

        return {
            notes,
            message,
            canvasMaxWidth,
            canvasMaxHeight,
            clickNote
        }
    }
}
</script>

<style scoped>

</style>