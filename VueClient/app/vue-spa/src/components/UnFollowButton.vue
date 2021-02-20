<style scoped>
    .follow{
      display:flex;
      align-items: center;
      padding: 12px;
      cursor: pointer;
    }
    .followText{
      color: white;
      background-color: rgb(0, 183, 255);;
      border-radius: 100vh;
      font-size: 0.9em;
      width: 80px;
      padding-top: 7px;
      padding-bottom: 7px;
      font-weight: bold;
      transition: all 200ms 0s ease;
    }
    .followText:hover{
      background-color: rgb(238, 56, 56);
    }
    .UnFollowText{
      display: none;
    }
    .followText:hover .FollowingText{
      display: none;
    }
    .followText:hover .UnFollowText{
      display: block;
    }
</style>

<template>
    <div class="follow" @click="unfollow">
        <div class="followText">
            <span class="UnFollowText">UnFollow</span>
            <span class="FollowingText">Following</span>
        </div>
    </div>
</template>

<script>
export default {
  methods: {
    unfollow() {
      this.$emit('unfollow');
      this.$destroy();
      if (this.$el.parentNode) this.$el.parentNode.removeChild(this.$el);
    },
  },
};
</script>
