<template>
  <div>
    <h1>AuthedView</h1>
    {{ingredients}}
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { getIngredients } from '@/lib/http/http';

@Options({
  data() {
    return {
      ingredients: [],
    };
  },
  mounted() {
    getIngredients().then((res) => { this.ingredients = res.data; });
  },
})
export default class Home extends Vue {}
</script>
