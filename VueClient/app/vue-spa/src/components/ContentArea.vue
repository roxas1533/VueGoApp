<style scoped>
    .contents{
        text-align: left;
        width: 100%;
        border-bottom: solid thin;
        border-color: black;

    }
    .string{
         margin-left: 2em;
        margin-top: 0.1em;
    }
    .username{
        font-weight: bold;
    }
    .id,.time{
        font-size: 0.7em;
    }
    .time{
        margin-left: auto;
    }
    .name,.time{
        display: inline-block;
        margin-bottom: 0em;
        margin-top: 0.5em;
        color: white;

    }
    .title{
        display: flex;
    }
    .content{
        margin-top: 0em;
        margin-bottom: 1em;
        color: white;

    }
</style>
<template>
    <div class="contents">
        <div class="string">
            <div class="title">
                <p class="name"><span class="username">{{username}} </span><span class="id">@{{id}}</span></p>
                <pre class="time">{{time}}</pre>
            </div>
            <p class="content">{{content}}</p>
        </div>
    </div>
</template>

<script>
export default {
  props: {
    username: String, // 追加
    content: String, // 追加
    id: Number, // 追加
    time: String,
  },
};
</script>
