<template lang="html">
  <div id="main">
    <div class="tree-wrap">
      <Tree />
    </div>
  </div>
</template>

<script>
import merge from 'lodash/merge'
import set from 'lodash/set'
import Tree from './Tree.vue'

export default {
  name: 'Main',
  components: { Tree },
  data() {
    return {
      tree: {}
    }
  },
  mounted() {
    this.$store.dispatch('getNode')
  },
}
</script>

<style lang="css" scoped>
  #main {
    display: flex;
  }
</style>
<style lang="css">
  html,body {
    margin: 0;
  }
  .tree-wrap {
    width: 260px;
    overflow: auto;
    height: 100vh;
    position: relative;
    background-color: #21252B;
  }
</style>
