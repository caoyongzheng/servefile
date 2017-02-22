<template lang="html">
  <div class="tree">
    <div
      class="header"
      :class="{ dir: node.isDir }"
      :style="{ paddingLeft: level * 15 + 10 + 'px' }"
      v-on:click="toggle"
    >
      <ArrowIcon v-if="node.isDir" /> {{ node.name }}
    </div>
    <Tree v-for="n in node.children" v-if="node.open" :path="n.path"></Tree>
  </div>
</template>

<script>
import get from 'lodash/get'
import ArrowIcon from './ArrowIcon.vue'

export default {
  name: 'Tree',
  components: { ArrowIcon },
  props: {
    path: {
      type: String,
      default: ''
    }
  },
  methods: {
    getNodeByPath(path = '') {
      if (!path || path === '' || path === '.') {
        return this.$store.state.tree
      }
      const splitpaths = path.split('/')
      const paths = []
      for (let i = 0; i < splitpaths.length; i++) {
        if (splitpaths[i]) {
          paths.push('children', splitpaths[i])
        }
      }
      return get(this.$store.state.tree, paths)
    },
    toggle() {
      const node = this.node
      node.open = !node.open
      if (node.isDir && !node.loaded) {
        node.loaded = true
        this.$store.dispatch('getNode', node.path)
      }
      this.$store.commit('set', node)
    },
  },
  computed: {
    node() {
      return this.getNodeByPath(this.path)
    },
    level() {
      const path = this.path
      if (!path || path === '' || path === '.') {
        return 0
      }
      return path.split('/').filter(s => s).length
    }
  },
}
</script>

<style lang="css" scoped>
  .tree {
    box-sizing: border-box;
    width: 100%;
  }
  .header {
    cursor: pointer;
    color: #9DA5B3;
    height: 2em;
    line-height: 2em;
  }
  .header:hover {
    background-color: #2C313A;
  }
  .header.dir {

  }
</style>
