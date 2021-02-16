<style scoped>
    i{
        font-size: 2em;
        color: rgb(0, 140, 255);
    }
    h2{
        margin: 0px;
        max-width: 100%;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    .TitleBar{
        display: flex;
        background-color: rgb(20, 32, 43);
        border-radius: 15px;
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
    .editprofile{
        align-items:center;
        /* justify-content:center; */
        display:flex;
        width:100% ;
        color: #000;
    }
    button{
        height: 30px;
        border: none;
        border-radius: 100vh;
        font-size: 1.2em;
        width: 80px;
        background-color: rgb(0, 183, 255);;
    }
    button:focus{
        outline: none;
    }
    .save{
        margin-left:0px;
         align-items:center;
        display:flex;
    }
</style>
<template>
    <div>
    <div class="TitleBar">
        <div class="close">
            <div class="t-middle">
                <i class="tims fas fa-times"></i>
            </div>
        </div>
        <div class="editprofile">
            <h2>プロフィールを編集</h2>
        </div>
        <div class="save">
            <button>保存</button>
        </div>

    </div>
    <div edit></div>
    </div>
</template>
