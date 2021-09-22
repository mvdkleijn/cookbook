<template>
  <div class="background">  
    <div class="splash">
      <h1>Welcome to your cookbook.</h1>
      <n-button style="color: white;" size="large" round @click="$router.push('App')">Enter</n-button>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue } from 'vue-class-component';

export default class Splashscreen extends Vue {}
</script>

<style lang="scss" scoped>
  .splash {
    color: white;
    position: fixed;
    top: 50%;
    left: 50%;
    /* bring your own prefixes */
    transform: translate(-50%, -50%);
  }
  .background {
    width: 100vw;
    height: 100vh;
    background: url('~@/assets/images/splash/one.jpg') no-repeat center center fixed;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
  }
</style>