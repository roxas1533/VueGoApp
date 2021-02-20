<style scoped>
    .wrap{
        padding: 12px;
        display: flex;
        flex-direction:column;
        text-align: left;
    }
    .title{
        font-size: 11px;
    }
    .number{
        font-weight: bold;
    }
</style>
<template>
    <div class="wrap">
        <span class="title">{{title}}</span>
        <span class="number">{{num}}</span>
    </div>
</template>

<script>
export default {
  props: {
    title: String,
    num: Number,
  },
};
</script>
