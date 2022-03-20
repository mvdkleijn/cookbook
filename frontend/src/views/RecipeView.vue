<template>
  <div>
    <h1>RecipeView</h1>
    <CreateRecipe />
    <RecipeTable :recipes="recipes"/>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { Recipes } from '@/lib/http/http';
import RecipeTable from '@/components/RecipeTable.vue';
import CreateRecipe from '@/views/CreateRecipe.vue';

@Options({
  components: {
    CreateRecipe,
    RecipeTable,
  },
  data() {
    return {
      recipes: {
        type: Object,
      },
    };
  },
  mounted() {
    Recipes.getAllRecipes().then((data) => { this.recipes = data; });
  },
})
export default class RecipeView extends Vue {}
</script>

<style lang="scss" scoped>

</style>
