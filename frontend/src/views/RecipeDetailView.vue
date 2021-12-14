<template>
  <div class="recipeDetail">

    <div class="grid grid-nogutter surface-section text-800">
        <div class="col-12 md:col-6 p-6 text-center md:text-left flex align-items-center ">
            <section>
                <span class="block text-6xl font-bold mb-1">{{recipe.title}}</span>
                <p class="mt-0 mb-4 text-700 line-height-3">{{recipe.description}}</p>
            </section>
        </div>
        <div class="col-12 md:col-6 overflow-hidden">
            <img :src="require(`@/assets/images/placeholder/${image}`)" alt="Image" class="md:ml-auto block md:h-full" style="clip-path: polygon(8% 0, 100% 0%, 100% 100%, 0 100%)">
        </div>
    </div>
    {{recipe}}
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { getRecipe } from '@/lib/http/http';
import Image from 'primevue/image';
// import RecipeLayout from '@/components/RecipeLayout.vue';

@Options({
  components: {
    Image,
  },
  props: {
    recipeID: 0,
  },
  data() {
    return {
      recipe: {},
      image: 'chicken.jpg'
    };
  },
  mounted() {
    getRecipe(this.$route.params.id).then((res:any) => { this.recipe = res.data; });
  },
})
export default class RecipeView extends Vue {}
</script>

<style lang="scss" scoped>
  .recipeDetail {
    padding-top: 3vh;
    padding-left: 3vw;
    padding-right: 3vw;
  }
  .headerTitle {
    color: white;
    position: absolute;
    top: 50%;
    left: 2%;
  }
</style>
