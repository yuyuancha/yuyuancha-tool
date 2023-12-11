<template>
  <canvas :id="id" :width="canvasWidth" :height="canvasHeight"></canvas>
</template>

<script>
import {onMounted, ref} from "vue";
import {drawChords} from "@js/guitar/canvas";

export default {
    name: "ChordsItem",
    props: {
        id: {
            type: String,
        },
        chords: Object,
        canvasWidth: {
            type: Number,
            default: 180
        },
        canvasHeight: {
            type: Number,
            default: 250
        }
    },
    setup(props) {
        const context = ref({});

        onMounted(() => {
            const canvas = document.getElementById(props.id);
            context.value = canvas.getContext("2d");
            context.value.clearRect(0, 0, props.canvasWidth, props.canvasHeight);
            drawChords(context.value, 40, 80, props.chords);
        });

        return {
            props
        }
    }
}
</script>

<style scoped>

</style>