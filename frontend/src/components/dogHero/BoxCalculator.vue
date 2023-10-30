<template>
    <Row justify="center" style="margin: 20px 0;">
        <Paragraph>
            月度達標累計寶箱積分 <Text type="danger" strong>{{ monthPoint }}</Text> 分
        </Paragraph>
    </Row>
    <Row v-for="item in pointList" :key="item.tag" justify="center">
        <Form v-model="formItem">
            <BoxFormItem :image="item.image"
                         :labelName="item.name" :boxKey="item.tag" :default-value="formItem[item.tag]"
                         @changeBoxNumber="changeBoxNumber" @calculatePoint="calculatePoint" />
        </Form>
    </Row>
    <Row justify="center">
        <Input size="large" prefix="md-game-controller-b" v-model="totalPoint"
               class="ivu-text-center" style="width: 100px; margin-bottom: 20px;" readonly />
    </Row>
</template>

<script>
import {Paragraph, Row, Text} from "view-ui-plus";
import {reactive, ref} from "vue";
import blueBoxImg from '@images/dogHero/blue-box.png';
import greenBoxImg from '@images/dogHero/green-box.png';
import orangeBoxImg from '@images/dogHero/orange-box.png';
import purpleBoxImg from '@images/dogHero/purple-box.png';
import BoxFormItem from "@/components/dogHero/BoxFormItem.vue";

export default {
    name: "BoxCalculator",
    components: {BoxFormItem, Text, Paragraph, Row},
    setup() {
        let formItem = reactive({
            box1: 0,
            box2: 0,
            box3: 0,
            box4: 0
        });
        const pointList = [
            {
                tag: 'box1',
                name: '綠寶箱',
                point: 1,
                image: greenBoxImg
            },
            {
                tag: 'box2',
                name: '藍寶箱',
                point: 10,
                image: blueBoxImg
            },
            {
                tag: 'box3',
                name: '紫寶箱',
                point: 20,
                image: purpleBoxImg
            },
            {
                tag: 'box4',
                name: '橙寶箱',
                point: 50,
                image: orangeBoxImg
            },
        ];
        let totalPoint = ref(0);
        const monthPoint = 80000;

        function changeBoxNumber(boxKey, number) {
            console.log("into change box number");
            console.log("boxKey", boxKey);
            console.log("number", number);
            formItem[`${boxKey}`] = number;
        }

        function calculatePoint() {
            let point = 0;

            for (let i = 1; i <= 4; i++) {
                const tag = `box${i}`;
                point += formItem[tag] * pointList.find((item) => {return item.tag === tag}).point;
            }

            totalPoint.value = point;
        }

        return {
            monthPoint,
            formItem,
            totalPoint,
            pointList,
            calculatePoint,
            changeBoxNumber
        }
    }
}
</script>

<style scoped>

</style>