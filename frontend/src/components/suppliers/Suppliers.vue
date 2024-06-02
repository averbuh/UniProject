<script setup>
import { ref } from 'vue'
import { onMounted } from 'vue';
import { Suppliers } from '@/service/Suppliers';

import axios from 'axios'


onMounted(() => {
  Suppliers.getSuppliers().then((data) => suppliers.value = data);
});

const suppliers = ref();

const getSeverity = (supplier) => {
  switch (supplier.recommended) {
    case 'INSTOCK':
      return 'success';

    case 'LOWSTOCK':
      return 'warning';

    case 'OUTOFSTOCK':
      return 'danger';

    default:
      return null;
  }
};

</script>

<template>
  <div class="card">
    <DataView :value="suppliers">
      <template #list="slotProps">
        <div class="grid grid-nogutter">
          <div v-for="(item, index) in suppliers" :key="index" class="col-12">
            <div class="flex flex-column sm:flex-row sm:align-items-center p-4 gap-3"
              :class="{ 'border-top-1 surface-border': index !== 0 }">
              <div class="md:w-10rem relative">
                <img class="block xl:block mx-auto border-round w-full" :src="`/images/${item.image}`"
                  :alt="item.name" />
                <!-- <Tag :value="item.inventoryStatus" :severity="getSeverity(item)" class="absolute" -->
                <!--   style="left: 4px; top: 4px"></Tag> -->
              </div>
              <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">
                <div class="flex flex-row md:flex-column justify-content-between align-items-start gap-2">
                  <div>
                    <span class="font-medium text-secondary text-sm">{{ item.category }}</span>
                    <div class="text-lg font-medium text-900 mt-2">{{ item.name }}</div>
                  </div>
                  <div class="surface-100 p-1" style="border-radius: 30px">
                    <div class="surface-0 flex align-items-center gap-2 justify-content-center py-1 px-2"
                      style="border-radius: 30px; box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.04), 0px 1px 2px 0px rgba(0, 0, 0, 0.06)">
                      <span class="text-900 font-medium text-sm">{{ item.rating }}</span>
                      <i class="pi pi-star-fill text-yellow-500"></i>
                    </div>
                  </div>
                </div>
                <div class="flex flex-column md:align-items-end gap-5">
                  <div class="flex flex-row-reverse md:flex-row gap-2">
                    <Button icon="pi pi-heart" outlined></Button>
                    <Button icon="pi pi-shopping-cart" label="Catalog"
                      class="flex-auto md:flex-initial white-space-nowrap"></Button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>
    </DataView>
  </div>


</template>

<style scoped>
img {
  width: 75%;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
  text-decoration: underline;
}

.card {
  background: var(--surface-card);
  padding: 2rem;
  border-radius: 10px;
  margin-bottom: 1rem;

  box-shadow: 0 2px 1px -1px rgb(0 0 0 / 20%), 0 1px 1px 0 rgb(0 0 0 / 14%), 0 1px 3px 0 rgb(0 0 0 / 12%);

  display: flexbox;
  flex-direction: column;
  gap: 1rem;

  max-width: 600px;
  margin: 0 auto;

  text-align: center;

  transition: all 0.3s ease;

  cursor: pointer;

}
</style>
