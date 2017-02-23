<template lang="html">
  <div id="main">
    <div class="tree-wrap" :style="{ width: treeWidth + 'px' }">
      <Tree />
      <div class="drag-line" v-on:mousedown="onMouseDown" />
    </div>
    <div class="editor-wrap">
      <textarea class="textarea" v-model="$store.state.textarea" v-on:keydown.meta.83="onSave"></textarea>
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
      treeWidth: 260,
      minWidth: 260,
      moving: false,
    }
  },
  mounted() {
    this.$store.dispatch('getNode')
    document.addEventListener('mousemove', (e) => {
      if (this.moving) {
        this.treeWidth = Math.max(this.minWidth, e.clientX)
      }
    }, true)
    document.addEventListener('mouseup', () => {
      if (this.moving) {
        this.moving = false
      }
    })
  },
  methods: {
    onMouseDown(e) {
      this.moving = true
    },
    onSave(e) {
      e.stopPropagation()
      e.preventDefault()
      const curnode = this.$store.getters.curnode
      if (curnode.isDir) {
        alert('curnode is Folder')
        return
      } else {
        fetch(`${process.env.BaseURL}/treeleaf`, {
          method: 'POST',
          header: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ path: curnode.path, content: this.$store.state.textarea })
        })
        .then(r => r.json())
        .then(({ success, err }) => {
          if (success) {
            curnode.content = this.$store.state.textarea
            this.$store.commit('set', curnode)
            alert('save success')
          } else {
            alert(err)
          }
        })
      }
    }
  }
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
  * {
    box-sizing: border-box;
  }
</style>
<style lang="css" scoped>
  .tree-wrap {
    position: relative;
    box-sizing: border-box;
    overflow: auto;
    height: 100vh;
    position: relative;
    background-color: #21252B;
    border-right: solid 2px #181A1F;
  }
  .drag-line {
    height: 100%;
    width: 4px;
    position: absolute;
    top: 0;
    bottom: 0;
    right: 0;
    cursor: col-resize;
  }
  .editor-wrap {
    flex: 1;
    position: relative;
  }
  .textarea {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    width: 100%;
    border: none;
    outline: none;
    background-color: #282C34;
    color: #ABB2BF;
    font-size: 20px;
    line-height: 1.5;
    resize: none;
  }
</style>
