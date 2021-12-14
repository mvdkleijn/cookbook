<template>
  <div class="p-d-flex p-flex-wrap p-jc-center">
    <div class="card p-mr-4 p-mb-4" v-for="recipe in recipes" :key="recipe.title" v-on:click="onClick(recipe.id)">
      <img alt="Header image" :src="require(`@/assets/images/placeholder/${image}`)" class="card card-image">
      <span class="card-title">{{recipe.title}}</span>
    </div>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

@Options({
  props: {
    recipes: {
      type: Object,
    },
  },
  data() {
    return {
      image: 'chicken.jpg',
      value: [],
    };
  },
  methods: {
    onClick(recipeID:number) {
      console.log(recipeID);
      this.$router.push({ name: 'RecipeDetailView', params: { id: recipeID } });
    },
  },
})
export default class RecipeTable extends Vue {}
</script>

<style lang="scss" scoped>
  .card {
    min-width: 200px;
    min-height: 134px;
    max-width: 15vw;
    max-height: 10vw;
    object-fit: contain;
    border-radius: 5px;
    box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
    transition: 0.3s;
    position: relative;
    text-align: center;
    color: rgb(255, 255, 255);
  }

  .card:hover {
    box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2);
  }

  .card-image {
    filter: brightness(50%);
  }

  .card-title {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    font-weight:normal;
    font-size: 20px;
  }
</style>
