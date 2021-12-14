<template>
  <div>
    <HeaderBar/>
    <router-view class="app-rv" />
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import HeaderBar from '@/components/HeaderBar.vue';

@Options({
  components: {
    HeaderBar,
  },
  // mounted() {
  //   const devtools = document.createElement('script')
  //   devtools.setAttribute('src', 'http://10.0.0.54:8098')
  //   document.head.appendChild(devtools)
  // },
})
export default class cookbook extends Vue {}
</script>

<style lang="scss">
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    // text-align: center;
    color: #2c3e50;
  }
  .app-rv {
    margin-top: 1vh;
  }
</style>
