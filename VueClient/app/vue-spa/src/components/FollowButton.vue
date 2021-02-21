<style scoped>
    .follow{
      display:flex;
      align-items: center;
      padding: 12px;
      cursor: pointer;
    }
    .followText{
      color: rgb(0, 183, 255);
      border-radius: 100vh;
      font-size: 0.9em;
      border: solid thin rgb(0, 183, 255);
      padding-top: 7px;
      padding-bottom: 7px;
      font-weight: bold;
      padding-left: 5px;
      padding-right: 5px;
      min-width: 80px;

    }
    .followText:hover{
      background-color: rgb(2, 23, 56);
    }
</style>

<template>
    <div class="follow" @click="follow">
        <span class="followText">{{Follow}}</span>
    </div>
</template>

<script>
export default {
  props: {
    Follow: String,
  },
  methods: {
    follow() {
      this.$emit('follow');
      this.$destroy();
      if (this.$el.parentNode) this.$el.parentNode.removeChild(this.$el);
    },
  },
};
</script>
