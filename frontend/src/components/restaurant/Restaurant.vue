<script setup>
import { ref } from 'vue'
import { onMounted } from 'vue';
import { Suppliers } from '@/service/Suppliers';
import { Recipes } from '@/service/Recipes';

import axios from 'axios'


onMounted(() => {
  Suppliers.getFavouriteSuppliers().then((data) => suppliers.value = data);
  Recipes.getTodayRecipes().then((data) => recipes.value = data);
});

let suppliers = ref();
let recipes = ref();

// const getSeverity = (supplier) => {
//   switch (supplier.recommended) {
//     case 'INSTOCK':
//       return 'success';

//     case 'LOWSTOCK':
//       return 'warning';

//     case 'OUTOFSTOCK':
//       return 'danger';

//     default:
//       return null;
//   }
// };
</script>

<template>
  <div class="restaurant">
    <!-- {{ suppliers }} -->
    <!-- <transition-group name="fade" tag="ul" class="suppliers_list"> -->
    <!--   <recipe-item v-for="(item, index) in suppliers" :key="index" class="col-12" > -->
    <!--     <Rating :modelValue="item.rating" readonly :cancel="false" /> -->
    <!--   </recipe-item> -->
    <!-- </transition-group> -->

    <img class="logo" src="/images/restaurant.jpg" alt="restaurant">
    <div class="card">
      <DataTable :value="recipes" tableStyle="min-width: 50rem">
        <template #header>
          <div class="flex flex-wrap align-items-center justify-content-between gap-2">
            <span class="text-xl text-900 font-bold">Today Menu</span>
          </div>
        </template>
        <Column field="name" header="Name"></Column>
        <!-- <Column header="Image"> -->
        <!--   <template #body="suppliers"> -->
        <!--     <img v-if="suppliers.image" :src="`/images/${suppliers.image}`" :alt="suppliers.name" class="w-6rem border-round" /> -->
        <!--   </template> -->
        <!-- </Column> -->
        <!-- <Column field="category" header="Category"></Column> -->
        <!-- <Column field="rating" header="Reviews"> -->
        <!--   <template #body="suppliers"> -->
        <!--     <Rating :modelValue="suppliers.rating"  readonly :cancel="false" /> -->
        <!--   </template> -->
        <!-- </Column> -->
      </DataTable>
    </div>
    <div class="card">
      <DataTable :value="suppliers" tableStyle="min-width: 50rem">
        <template #header>
          <div class="flex flex-wrap align-items-center justify-content-between gap-2">
            <span class="text-xl text-900 font-bold">Top Suppliers</span>
          </div>
        </template>
        <Column field="name" header="Name"></Column>
        <!-- <Column header="Image"> -->
        <!--   <template #body="suppliers"> -->
        <!--     <img v-if="suppliers.image" :src="`/images/${suppliers.image}`" :alt="suppliers.name" class="w-6rem border-round" /> -->
        <!--   </template> -->
        <!-- </Column> -->
        <Column field="category" header="Category"></Column>
        <Column field="rating" header="Reviews">
          <template #body="suppliers">
            <Rating :modelValue="suppliers.rating" readonly :cancel="false" />
          </template>
        </Column>
      </DataTable>
    </div>

  </div>
</template>

<style scoped>
.restaurant {
  font-size: 10px;
}

h1 {
  text-align: center;
  text-decoration: underline;
  margin: 10px;
}

.logo {
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 50%;
  height: 50%;
  border-radius: 50%;
  border: 3px solid black;

  max-width: 400px;
  max-height: 400px;


}

h2 {
  text-align: center;
}

.card {
  background: var(--surface-card);
  padding: 2rem;
  border-radius: 10px;
  margin-bottom: 1rem;
  margin-top: 1rem;


  box-shadow: 0 2px 1px -1px rgb(0 0 0 / 20%), 0 1px 1px 0 rgb(0 0 0 / 14%), 0 1px 3px 0 rgb(0 0 0 / 12%);



  text-align: center;

  transition: all 0.3s ease;

  cursor: pointer;

}
</style>
