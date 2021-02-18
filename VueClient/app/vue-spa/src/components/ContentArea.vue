<style scoped>
    .contents{
        text-align: left;
        border-bottom: solid thin;
        border-color: black;
        display: flex;
        padding: 8px;
    }
    .string{
        display: block;
        margin: 0px;
        width: 100%;
        /* width: clac(100%-50px); */
        /* margin-left: 50px; */
    }
    .username{
        margin: 0px;
        font-weight: bold;
    }
    .id,.time{
        font-size: 0.7em;
        color: rgb(95, 107, 113);
    }

    .name,.time{
        margin: 0px;
        display: inline-block;
    }
    .name{
        color: white;
        width: 146px;
    }
    .time{
        padding-left: 5px;
        margin-right: 0px;
    }
    .Wrapname{
        width: 100%;
    }
    .title{
        display: flex;
        /* width: 100%; */
    }
    .content{
        margin-top: 0em;
        margin-bottom: 1em;
        color: white;
        /* display: none; */
    }
    .pimg{
        width: 36px;
        height: 36px;
        float: left;
    }

</style>
<template>
    <div class="contents">
        <img v-bind:src="profile+'/profile/'+id+'.png?'+(new Date()).getMinutes()" class="pimg">
        <div class="string" id="string">
            <div class="title">
                <div class="wrappname">
                  <p class="name" id="name"><span class="username">{{username}} </span><span class="id">@{{id}}</span></p>
                </div>
                <div class="timeWrap">
                    <p class="time" id="time">{{time}}</p>
                </div>
            </div>
        <p class="content">{{content}}</p>
        </div>

    </div>
</template>

<script>
import C from '../store/const';

export default {
  props: {
    username: String, // 追加
    content: String, // 追加
    id: Number, // 追加
    time: String,
  },
  data() {
    return {
      profile: C.APIserver,
    };
  },
};
</script>
