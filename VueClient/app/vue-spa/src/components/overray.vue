<style scoped>
    .overray{
        position: block;
        width: 100%;
        height: 100%;
        background-color: rgba(0, 0, 0, 0.7);
        z-index: 299;
        pointer-events: auto;
    }
</style>
<template>
    <div class="overray" id="overray" @click="close"></div>
</template>
<script>
export default {
  name: 'overray',
  methods: {
    close() {
      this.$emit('close');
    },
  },
};
</script>
