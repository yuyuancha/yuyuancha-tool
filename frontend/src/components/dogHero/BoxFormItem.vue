<template>
    <FormItem>
        <Image :src="image" fit="fill" style="margin-right: 20px;"
               width="50px" height="50px" :alt="labelName" />
        <InputNumber :min="0" :max="10000" v-model="inputValue"
                     style="margin-right: 20px;"
                     @on-change="syncBoxNumber" />
        <Button v-for="number in boxNumberList" :key="number" shape="circle" icon="md-add"
                @click="addInputValue(number)" style="margin-right: 5px;">
            {{ number }}
        </Button>
        <Button @click="reset"><Icon type="md-refresh-circle" />重設</Button>
    </FormItem>
</template>

<script>
import {FormItem, Icon, Image, InputNumber} from "view-ui-plus";
import {ref} from "vue";

export default {
    name: "BoxFormItem",
    components: {Icon, InputNumber, Image, FormItem},
    emits: ['changeBoxNumber', 'calculatePoint'],
    props: {
        defaultValue: Number,
        labelName: String,
        boxKey: String,
        image: String,
    },
    setup(props, { emit }) {
        const boxNumberList = [10, 100, 1000];
        let inputValue = ref(props.defaultValue);

        function syncBoxNumber() {
            emit('changeBoxNumber', props.boxKey, inputValue.value);
            emit('calculatePoint');
        }

        function addInputValue(number) {
            inputValue.value += number;
            syncBoxNumber();
        }

        function reset() {
            inputValue.value = 0;
            syncBoxNumber();
        }

        return {
            props,
            inputValue,
            boxNumberList,
            addInputValue,
            syncBoxNumber,
            reset
        }
    }
}
</script>

<style scoped>

</style>