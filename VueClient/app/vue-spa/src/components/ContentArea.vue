<style scoped>
    .Wrapcontents{
        display: flex;
        flex-direction: column;
        padding: 8px;
        border-bottom: solid thin;
        border-color: black;
        cursor: pointer;

    }
    .contents{
        text-align: left;
        /* border-top: solid thin; */
        display: flex;
    }

    .string{
        display: block;
        margin: 0px;
        width: 100%;
    }
    .username{
        margin: 0px;
        font-weight: bold;
        cursor: pointer;
    }
    .username:hover{
        text-decoration: underline;
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
    .timeWrap{
         margin-left: auto;
    }
    .title{
        display: flex;
    }
    .content{

        color: white;
    }
    .pimg{
        width: 36px;
        height: 36px;
        float: left;
        cursor: pointer;
    }
    .buttonContent{
        text-align: right;
    }
    .reactionButton{
        display: inline-block;
        margin-right: 7px;
    }
    .far:hover{
        color: yellow;
    }
    .fas{
        color: yellow;
    }
</style>
<template>
    <div class="Wrapcontents">
        <div class="contents">
            <img  @click="showProfile" v-bind:src="profile+'/profile/'+id+'.png?'+(new Date()).getMinutes()" class="pimg">
            <div class="string" id="string">
                <div class="title">
                    <div class="wrappname">
                    <p class="name" id="name"><span @click="showProfile" class="username">{{username}} </span><span class="id">@{{id}}</span></p>
                    </div>
                    <div class="timeWrap">
                        <p class="time" id="time">{{time}}</p>
                    </div>
                </div>
                <p class="content">{{content}}</p>
            </div>
        </div>
        <div class="buttonContent">
            <div class="reactionButton">
                <i @click="favorite" class="far fa-star" ref="star"></i>
            </div>
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
    talkID: Number,
    isFavorite: Boolean,
  },
  data() {
    return {
      profile: C.APIserver,
    };
  },
  mounted() {
    if (this.isFavorite) {
      this.$refs.star.classList.toggle('far');
      this.$refs.star.classList.toggle('fas');
    }
  },
  computed: {
    fav() {
      return this.$store.getters.getFavlist;
    },
  },
  watch: {
    fav: {
      handler(fav) {
        if (fav.favid === this.talkID) {
          this.$refs.star.classList.toggle('far');
          this.$refs.star.classList.toggle('fas');
        }
      },
      deep: true,
    },
  },
  methods: {

    showProfile() {
      this.$emit('showProfile', this.username, this.id);
    },
    async favorite() {
      let url;
      if (this.isFavorite) url = `${this.$store.state.APIserver}/favorite/${this.talkID}`;
      else url = `${this.$store.state.APIserver}/favorite/${this.talkID}`;
      const r = await window.fetch(url, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${this.$store.state.JWTtoken}`,
          'Content-Type': 'text/plain',
        },
      }).then((res) => res.json());
      if (r.result) {
        this.isFavorite = !this.isFavorite;
        this.$store.state.favlist.favid = this.talkID;
        this.$store.state.favlist.state = this.isFavorite;
        console.log('test2', this.$store.getters.getFavlist);
      }
    },
  },
};
</script>
