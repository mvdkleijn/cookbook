<template>
  <div class="surface-ground px-4 py-8 md:px-6 lg:px-8">
    <div class="grid grid-nogutter surface-section text-800">
      <div class="col-12 md:col-6 p-6 text-center md:text-left flex align-items-center ">
        <section>
          <span class="block text-6xl font-bold mb-1">{{recipe.title}}</span>
          <p class="mt-0 mb-4 text-700 line-height-3">{{recipe.description}}</p>
          <div class="p-grid p-formgrid p-fluid mt-0 mb-4 text-700 line-height-3">
            <span class="p-col-1">
              <img :src="require(`@/assets/icons/knife1.svg`)" class="p-mr-5" style="height: 20px"/>
              {{recipe.preptime}} Minutes
            </span>
            <span class="p-col-1">
              <img :src="require(`@/assets/icons/pot.svg`)" class="p-mr-5" style="height: 20px"/>
              {{recipe.cooktime}} Minutes
            </span>
            <span class="p-col-1">
              <i class="pi pi-clock p-mr-5"/>
              {{recipe.totaltime}} Minutes
            </span>
          </div>
          {{recipe}}
        </section>
      </div>
      <div class="col-12 md:col-6 overflow-hidden">
        <img :src="require(`@/assets/images/placeholder/${image}`)" alt="Image" class="md:ml-auto block md:h-full" style="clip-path: polygon(8% 0, 100% 0%, 100% 100%, 0 100%)">
      </div>
    </div>

    <div class="grid">
      <div class="col-12 lg:col-4">
        <div class="p-3 h-full">
          <div class="shadow-2 p-3 h-full flex flex-column surface-card" style="border-radius: 6px">
            <div class="text-900 font-medium text-xl mb-2">Ingredients</div>
            <hr class="my-3 mx-0 border-top-1 border-none surface-border" />
            <table>
              <div v-for="section in recipe.sections" :key="section.id">
                <thead>
                  <tr>
                    <th class="name" colspan="2">{{section.name}}</th>
                  </tr>
                </thead>
                <tr v-for="ingredient in filterIngredients(key)" :key="ingredient.ingredientid">
                  <td class="name">{{ingredient.name}}</td> <td class="amount">{{ingredient.amount}} {{ingredient.unit}}</td>
                </tr>
              </div>
            </table>
          </div>
        </div>
      </div>

      <div class="col-12 lg:col-8">
        <div class="p-3 h-full">
          <div class="shadow-2 p-3 h-full flex flex-column surface-card" style="border-radius: 6px">
            <div class="text-900 font-medium text-xl mb-2">Method</div>
            <hr class="my-3 mx-0 border-top-1 border-none surface-border" />
            {{recipe.method}}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import Image from 'primevue/image';
import {
  Recipes,
  Ingredients,
  Sections,
  RecipeIngredientResponse,
} from '@/lib/http/http';

interface ExtendedIngredients extends RecipeIngredientResponse {
  name?: string;
}

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
      recipeIngredients: new Array<ExtendedIngredients>(),
      ingredientSections: new Map<number, string>(),
      image: 'chicken.jpg',
    };
  },
  mounted() {
    Recipes.getSingleRecipe(this.$route.params.id).then((data) => { this.recipe = data; });
    Ingredients.getRecipeIngredients(this.$route.params.id).then((data) => {
      this.recipeIngredients = data;
      this.resolveIngredientNames();
      this.resolveIngredientSections().then((sections:Map<number, string>) => {this.ingredientSections = sections});
    });
  },
  methods: {
    resolveIngredientNames() {
      for (const i in this.recipeIngredients) {
        Ingredients.getSingleIngredient(this.recipeIngredients[i].ingredientid).then((data) => {this.recipeIngredients[i].name = data[0].name; });
      }
    },

    resolveIngredientSections(this: {recipeIngredients:Array<ExtendedIngredients>}) {
      return new Promise((success, error) => {
            const sections = new Map<number, string>();
            const promises:Array<any> = [];
            this.recipeIngredients.forEach(ingredient => {
                promises.push(Sections.getSingleIngredientSection(ingredient.sectionid));
            });
          Promise.all(promises).then((recipeDatas) => {
              recipeDatas.forEach((data) => {
                if (!sections.has(data[0].id)) {
                    sections.set(data[0].id, data[0].name);
                };
              })
              success(sections);
          });
      });
    },

    filterIngredients(sectionID: number) {
      return this.recipeIngredients.filter(function(recipe:any) {
        return recipe.sectionid === sectionID;
      });
    },
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
  .name {
    text-align:left;
  }
  .amount {
    text-align: right;
  }
</style>
