<template>
  <Button label="New" icon="pi pi-plus" @click="openDialog()"/>
  <Dialog
    header="New recipe"
    :breakpoints="{'960px': '75vw', '640px': '100vw'}"
    :style="{width: '50vw'}"
    :modal="true"
    v-model:visible="visible"
  >
    <div class="p-fluid">
      <div class="p-field">
        <label for="recipetitle">Title</label>
        <InputText
          id="recipetitle"
          type="text"
          v-model="recipe.title"
          :class="{'p-invalid': validationErrors.title && submitted}"
          autofocus
        />
        <small
          v-show="validationErrors.title && submitted"
          class="p-error">
            Recipe title is required.
        </small>
      </div>

      <div class="p-field">
        <label for="recipedescription">Description</label>
        <br/>
        <Textarea
          id="recipedescription"
          type="text"
          v-model="recipe.description"
          :autoResize="true"
          :class="{'p-invalid': validationErrors.description && submitted}"
        />
        <small
          v-show="validationErrors.description && submitted"
          class="p-error">
            Recipe description is required.
        </small>
      </div>

      <div class="p-field">
        <label for="recipepreptime">Preparation duration in minutes</label>
        <InputNumber
          id="recipepreptime"
          v-model="recipe.prepTime"
          mode="decimal" showButtons
          :class="{'p-invalid': validationErrors.prepTime && submitted}"
        />
        <small
          v-show="validationErrors.prepTime && submitted"
          class="p-error">
            Preparation time is required.
        </small>
      </div>

      <div class="p-field">
        <label for="recipecooktime">Cooking duration in minutes</label>
        <InputNumber
          id="recipecooktime"
          v-model="recipe.cookTime"
          mode="decimal" showButtons
          :class="{'p-invalid': validationErrors.cookTime && submitted}"
        />
        <small
          v-show="validationErrors.cookTime && submitted"
          class="p-error">
            Cooking duration is required.
        </small>
      </div>

    </div>

    <template #footer>
      <Button label="Cancel" icon="pi pi-cross" @click="closeDialog()"/>
      <Button label="Create" icon="pi pi-check" @click="closeDialog()"/>
    </template>

  </Dialog>

</template>

<script lang="ts">
/* eslint-disable */
import { Options, Vue } from 'vue-class-component';
import Dialog from 'primevue/dialog';

@Options({
  components: {
    Dialog
  },
  data() {
    return {
      visible: false,
      recipe: {
        type: Object
      },
      validationErrors: {
        type: Object
      },
    };
  },
  methods: {
    openDialog() {
      this.visible = true
    },
    closeDialog() {
      this.visible = false
    },
    validateForm() {
      if (!this.recipe.title.trim()) {
        this.validationErrors['title'] = true;
      } else { delete this.validationErrors['title']; }

      if (!this.recipe.description.trim()) {
        this.validationErrors['description'] = true;
      } else { delete this.validationErrors['description']; }

      if (!this.recipe.prepTime) {
        this.validationErrors['prepTime'] = true;
      } else { delete this.validationErrors['prepTime']; }

      if (!this.recipe.cookTime) {
        this.validationErrors['cookTime'] = true;
      } else { delete this.validationErrors['cookTime']; }

      return !Object.keys(this.validationErrors).length;
    },
    

  }
})
export default class CreateRecipe extends Vue {}
</script>

<style lang="scss" scoped>
</style>
