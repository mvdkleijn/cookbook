<template>
  <div>
    <h1>Ingredients</h1>
    <DataTable
      :value="ingredients"
      :paginator="true"
      :rows="10"
      v-model:filters="filters"
      filterDisplay="menu"
      :loading="loading"
      responsiveLayout="scroll"
      :globalFilterFields="['id','name']"
      dataKey="id">
      <template #header>
        <div class="flex justify-content-between">
            <span>
              <CreateIngredient />
              <Button
                label="Delete"
                icon="pi pi-trash"
                class="p-button-danger mr-2"
                @click="confirmDeleteSelected"
                :disabled="!selectedProducts || !selectedProducts.length"
              />
              <Button
                type="button"
                icon="pi pi-filter-slash"
                label="Clear"
                class="p-button-outlined"
                @click="clearFilter()"
              />
            </span>
            <span class="p-input-icon-left">
                <i class="pi pi-search" />
                <InputText v-model="filters['global'].value" placeholder="Keyword Search" />
            </span>
        </div>
      </template>
      <template #empty>
        No ingredients found.
      </template>
      <template #loading>
        Loading ingredient data. Please wait.
      </template>
      <Column field="id" header="ID"></Column>
      <Column field="name" header="Name"></Column>
    </DataTable>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import ColumnGroup from 'primevue/columngroup';
import { FilterMatchMode, FilterOperator } from 'primevue/api';
import { Ingredients } from '@/lib/http/http';
import CreateIngredient from '@/components/CreateIngredient.vue';

@Options({
  components: {
    DataTable,
    Column,
    ColumnGroup,
    CreateIngredient,
    FilterMatchMode,
    FilterOperator,
  },
  data() {
    return {
      ingredients: [],
      loading: true,
      filters: null,
    };
  },
  created() {
    this.initFilters();
  },
  mounted() {
    Ingredients.getAllIngredients().then(
      (data) => { this.ingredients = data; this.loading = false; },
    );
  },
  methods: {
    clearFilter() {
      this.initFilters();
    },
    initFilters() {
      this.filters = {
        global: { value: null, matchMode: FilterMatchMode.CONTAINS },
        id: { value: null, matchMode: FilterMatchMode.CONTAINS },
        name: { value: null, matchMode: FilterMatchMode.CONTAINS },
      };
    },
  },
})
export default class IngredientView extends Vue {}
</script>
