<template>
  <div>
    <h1>RecipeView</h1>
    <CreateRecipe />
    <RecipeTable :recipes="recipes"/>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { getRecipes } from '@/lib/http/http';
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
    getRecipes().then((res) => { this.recipes = res.data; });
  },
})
export default class RecipeView extends Vue {}
</script>

<style lang="scss" scoped>

</style>
