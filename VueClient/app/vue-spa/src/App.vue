<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style>
  @import url("https://use.fontawesome.com/releases/v5.6.1/css/all.css");

*{
  margin: 0ex;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #1b2a37;
  font-family: "Helvetica Neue" , Helvetica , Arial , Verdana , Roboto , "游ゴシック" , "Yu Gothic" , "游ゴシック体" ,
   "YuGothic" , "ヒラギノ角ゴ Pro W3" , "Hiragino Kaku Gothic Pro" , "Meiryo UI" , "メイリオ" , Meiryo , "ＭＳ Ｐゴシック" , "MS PGothic" , sans-serif;
   color:#888888;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
