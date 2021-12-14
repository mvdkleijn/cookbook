<template>
  <div>
    <h1>Logged Out</h1>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

@Options({
  // mounted() {
  //   this.$oidc.signOut();
  // },
})
export default class Home extends Vue {}
</script>
