import {commonGrayColor, lineWidth, pointRadius} from "@js/const";

function drawLine(context, x, y, width = lineWidth) {
    context.lineWidth = width;
    context.strokeStyle = commonGrayColor;
    context.beginPath();
    context.moveTo(x.start, y.start);
    context.lineTo(x.end, y.end);
    context.stroke();
}

function drawPoint(context, x, y, radius = pointRadius) {
    const circle = new Path2D();
    circle.moveTo(0, 0);
    circle.arc(x, y, radius, 0, 2 * Math.PI);

    context.fillStyle = commonGrayColor;
    context.fill(circle);
}

export {
    drawLine,
    drawPoint
}
