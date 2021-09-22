<template>
  <n-page-header class="header">
    <template #avatar>
      <n-icon size="40"><Food20Regular /></n-icon>
    </template>
    <template #title>
      Cookbook
    </template>
    <template #subtitle>
      <span>
        <div class="nav">
          <router-link to="/app/recipes">Recipes</router-link>
          <router-link to="/app/ingredients">Ingredients</router-link>
          <router-link to="/app/shopping">Shopping Lists</router-link>
          <router-link to="/app/planner">Mealplanner</router-link>
        </div>
      </span>
    </template>
  </n-page-header>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { NIcon, NPageHeader } from 'naive-ui';
import { Food20Regular } from '@vicons/fluent';

@Options({
  components: {
    NIcon,
    NPageHeader,
    Food20Regular,
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
    height: 5vh;
    box-shadow: 0px 1px 5px 0px;
    width: 100%;
    z-index: 10;
    position: relative;
    padding-top: 1vh;
    padding-left: 1vw;
    padding-right: 1vw;
  }
  .nav {
    justify-content: space-evenly;
    align-self: center;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
