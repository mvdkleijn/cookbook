<template>
  <!-- <div class="p-d-flex p-flex-wrap p-jc-center">
    <div class="card p-mr-4 p-mb-4" v-for="recipe in recipes" :key="recipe.title" v-on:click="onClick(recipe.id)">
      <img alt="Header image" :src="require(`@/assets/images/placeholder/${image}`)" class="card card-image">
      <span class="card-title">{{recipe.title}}</span>
    </div>
  </div> -->
  <div class="container">
    <div class="grid">   
      <div class="card" v-for="recipe in recipes" :key="recipe.title" v-on:click="onClick(recipe.id)">
          <img alt="Header image" :src="require(`@/assets/images/placeholder/${image}`)">
          <span class="title">{{recipe.title}}</span>
      </div>
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
  // .card {
  //   min-width: 200px;
  //   min-height: 134px;
  //   max-width: 15vw;
  //   max-height: 10vw;
  //   object-fit: contain;
  //   border-radius: 5px;
  //   box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
  //   transition: 0.3s;
  //   position: relative;
  //   text-align: center;
  //   color: rgb(255, 255, 255);
  // }

  // .card:hover {
  //   box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2);
  // }

  // .card-image {
  //   filter: brightness(50%);
  // }

  // .card-title {
  //   position: absolute;
  //   top: 50%;
  //   left: 50%;
  //   transform: translate(-50%, -50%);
  //   font-weight:normal;
  //   font-size: 20px;
  // }
  .grid {
  display: grid;
  grid-gap: 1rem;

  .card {
    &:hover {
      img {
        transform: scale(1.1);
        filter: blur(0px);
      }
    }
    cursor: pointer;
    border-radius: 0.5rem;
    position: relative;
    overflow: hidden;
    border: 0px solid #ccc;
    box-sizing: border-box;
    img {
      transition: transform 1s ease-in-out, filter 1s ease-in-out;
      filter: blur(3px);
      display: block;
      width: 100%;
      box-sizing: border-box;
    }
    span.title {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      text-align: center;
      width: 100%;
      padding: 1rem;
      color: #fff;
      font-size: 1.7rem;
      text-shadow: 0px 0px 7px rgba(0, 0, 0, 0.9);
      box-sizing: border-box;
    }
  }
}

.container {
  margin: 0 auto;
}

@media (min-width: 480px) {
  .container {
    max-width: 100%;
  }
  .grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
@media (min-width: 600px) {
  .container {
    max-width: 100%;
  }
  .grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
@media (min-width: 1025px) {
  .container {
    max-width: 1024px;
  }
  .grid {
    grid-template-columns: repeat(3, 1fr);
  }
}
@media (min-width: 1400px) {
  .container {
    max-width: 1400px;
  }
  .grid {
    grid-template-columns: repeat(4, 1fr);
  }
}

body,
html {
  margin: 0;
  font-family: Arial;
}

body {
  padding-top: 6rem;
  padding-bottom: 6rem;
}

</style>
