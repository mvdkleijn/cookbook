<template>
  <div class="flex header">
    <div
    class="flex-initial flex align-items-center justify-content-center title"
    v-on:click="goHome()">
      Cookbook
    </div>
    <div class="flex-initial flex align-items-center justify-content-center nav">
      <router-link to="/app/recipes">Recipes</router-link>
      <router-link to="/app/ingredients">Ingredients</router-link>
      <router-link to="/app/shopping">Shopping Lists</router-link>
      <router-link to="/app/planner">Mealplanner</router-link>
    </div>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

@Options({
  components: {
  },
  data() {
    return {
      menuOptions: [],
    };
  },
})
export default class HeaderBar extends Vue {}
</script>

<style lang="scss" scoped>
  .header {
    height: 50px;
    box-shadow: 0px 1px 5px 0px;
    width: 100%;
    z-index: 10;
    position: relative;
    // padding-top: 1vh;
    // padding-left: 1vw;
    padding-top: 5px;
    padding-left: 5px;
    padding-right: 1vw;
  }
  .title {
    margin-right: 5px;
    margin-bottom: 5px;
  }
  .nav {
    padding-right: 5px;
    margin-bottom: 5px;
  a {
    margin-left: 5px;
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
