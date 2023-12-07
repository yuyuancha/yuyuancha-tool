import {drawLine, drawPoint} from "@js/common/canvas";
import {commonGrayColor, lineWidth} from "@js/const";

const lineNum = 6; // 弦的數量
const chordsWidthSpace = 20; // 和弦的左右間距
const chordsHeightSpace = 30; // 和弦的上下間距

function drawChords(context, startX, startY, chords) {
    context.font = "25px Arial";
    context.fillStyle = commonGrayColor;
    context.fillText(chords.name, startX + 1.5 * chordsWidthSpace, startY - 30);

    drawZeroChords(context, startX, startY);
    drawBaseLine(context, startX, startY, chords.baseLine);
    drawPoints(context, startX, startY, chords.points);
}

function drawPoints(context, startX, startY, points) {
    points.forEach((point, index) => {
        const x = startX + index * chordsWidthSpace;
        const y = startY + (point - 0.5) * chordsHeightSpace;
        if (point === null) {
            context.font = "15px Arial";
            context.fillStyle = commonGrayColor;
            context.fillText("X", x - chordsWidthSpace/4, startY - 5);
            return;
        }
        if (point === 0) {
            return;
        }
        drawPoint(context, x, y);
    });
}

function drawBaseLine(context, startX, startY, baseLine = 0) {
    if (baseLine === 0) {
        return;
    }
    context.font = "18px Arial";
    context.fillStyle = commonGrayColor;
    context.fillText(baseLine.toString(), startX - 15, startY + 10);
}

function drawZeroChords(context, startX, startY) {
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

        drawLine(context, {start: startX, end: endX}, {start: y, end: y}, width);
        drawLine(context, {start: x, end: x}, {start: startY, end: endY});
    }
}

export {
    drawChords
}
