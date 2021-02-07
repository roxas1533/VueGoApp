<style scoped>
textarea{
    resize: none;
}
 .TalkComponent{
    display: block;
    margin-left: 0px;
 }
 #talkButton{
     border-radius: 10px;
     font-size: 1.6em;
     display: inline-block;
 }
 .Talk{
     display: inline-block;
     text-align: right;
 }
     ::-webkit-scrollbar{
      width: 10px;
    }
    ::-webkit-scrollbar-track{
      background: #4444;
      border-left: solid 1px #ececec;
    }
    ::-webkit-scrollbar-thumb{
      background: #888;
      border-radius: 10px;
      box-shadow: inset 0 0 0 2px #888;
    }
</style>
<template>
    <div class="Talk">
        <div class="Talk">
            <textarea class="TalkComponent" rows="8" cols="30" placeholder="Talk!"></textarea>
            <button class="TalkComponent" id="talkButton">Talk</button>
        </div>
    </div>
</template>
