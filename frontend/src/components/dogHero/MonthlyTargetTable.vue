<template>
    <Table border :columns="columns" :data="data" style="width: 500px; margin-left: 10%;" />
</template>

<script>
export default {
    name: "MonthlyTargetTable",
    props: {
        data: Array,
    },
    setup(props) {
        const columns = [
            {
                title: '進度表',
                key: 'progressing'
            },
            {
                title: '幸運寶箱數',
                key: 'luckyBoxNumber'
            }
        ];

        return {
            props,
            columns
        }
    }
}
</script>
