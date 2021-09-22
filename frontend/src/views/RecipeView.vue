<template>
  <div>
    <h1>RecipeView</h1>
    {{recipes}}
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { getRecipes } from '@/lib/http/http';

@Options({
  data() {
    return {
      recipes: [],
    };
  },
  mounted() {
    getRecipes().then((res) => { this.recipes = res.data; });
  },
})
export default class RecipeView extends Vue {}
</script>
