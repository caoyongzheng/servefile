<template lang="html">
  <svg
    viewBox="0 0 1024 1024"
    width="12"
    height="10"
    :class="{ down }"
  >
    <path d="M779.334025 457.424422 347.133822 25.224244C314.616547-7.293052 261.89556-7.293052 229.378287 25.224244 196.860957 57.741546 196.860957 110.462501 229.378287 142.979801L598.612762 511.996785 237.10385 877.866763C202.418446 908.060688 198.777402 960.655766 228.971339 995.341158 259.165276 1030.026539 311.760365 1033.667581 346.445714 1003.473651L775.127163 579.105645C811.520124 547.425245 813.452168 491.542573 779.334025 457.424422Z" fill="#6B717D" />
  </svg>
</template>

<script>
export default {
  name: 'ArrowIcon',
  props: {
    down: {
      type: Boolean,
    }
  }
}
</script>

<style lang="css" scoped>
svg{
  transition: transform 150ms ease;
  transform: rotateZ(0deg);
}
.down {
  transform: rotateZ(90deg);
}
</style>
