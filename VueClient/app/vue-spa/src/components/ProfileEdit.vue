<style scoped>
    i{
        font-size: 2em;
        color: rgb(0, 140, 255);
    }
    i:hover{
        cursor: pointer;
    }
    h2,.detail{
        margin: 0px;
        max-width: 100%;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    input:focus{
        outline: none;
    }
    .TitleBar{
        border-radius: 15px 15px 0px 0px;
    }
    .edit{
        border-radius: 0px 0px 15px 15px;
        align-items:center;
        justify-content:center;
    }
    .TitleBar,.edit{
        display: flex;
        background-color: rgb(20, 32, 43);
        padding-left: 15px;
        padding-right: 15px;
        position: relative;
        width: 100%;
    }
    .close{
        width: 80px ;
        height: 50px ;
        text-align: left;
        align-items:center;
        display:flex;
    }
   .t-middle{
        width: 40px ;
        height: 40px ;
        text-align: left;
        align-items:center;
        justify-content:center;
        display:flex;
    }
    .t-middle:hover{
        background-color: rgba(0, 183, 255,0.1);
        border-radius: 200px;
    }
    .editprofile{
        align-items:center;
        /* justify-content:center; */
        display:flex;
        width:100% ;
    }
    span{
        color: white;
        height: 30px;
        border: none;
        border-radius: 100vh;
        font-size: 1.2em;
        width: 80px;
        background-color: rgb(0, 183, 255);
    }
    span:hover{
        cursor: pointer;
        background-color: rgb(0, 124, 173);
    }
    button:focus{
        outline: none;
    }
    .save{
        margin-left:0px;
         align-items:center;
        display:flex;
    }
    .name{
        position: absolute;
        font-size: 1.35em;
        height: 1.35em;
        width: 99%;
        bottom: 0px;
        box-sizing: border-box;
        background: none;
        border: none;
        left: 0;
        right: 0;
        margin: auto;
        pointer-events:auto;
    }

    .nameBox{
        position: relative;
        height: 3.5em;
        width: 75%;
        display: flex;
        border: solid 2px black;
        border-radius: 5px;
        background-color: rgb(18, 18, 18);
    }
    .nameBox:focus-within {
        border:solid rgb(0, 183, 255) 2px;
    }
    .nameBox:focus-within .detail{
        color: rgb(0, 183, 255);
    }
</style>
<template>
    <div>
        <div class="TitleBar">
            <div class="close">
                <div class="t-middle">
                    <div class="t-middle-middle">
                        <i class="tims fas fa-times"></i>
                    </div>
                </div>
            </div>
            <div class="editprofile">
                <h2>プロフィールを編集</h2>
            </div>
            <div class="save">
                <span>保存</span>
            </div>
        </div>
        <div class=edit>
            <div class="nameBox" v-on:click="focus">
                <p class="detail">名前</p>
                <input type="text" class="name">
            </div>
        </div>
    </div>
</template>

<script>
export default {
  methods: {
    focus(e) {
      const c = e.target.lastElementChild;
      if (c) c.focus();
    },
  },
};
</script>
