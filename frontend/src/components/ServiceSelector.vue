<script setup>
import { Recipes } from '@/service/Recipes';
import { Suppliers } from '@/service/Suppliers';
import { ref } from 'vue';

defineProps({
  ButtonAction: {
    type: Function,
    required: true

  }
})

const ButtonLabel = ref("Switch Backend Endpoint");

function ChangeApiUrl() {
  Recipes.changeURL();
  Suppliers.changeURL();
    
  ButtonLabel.value = Recipes.getUrl();

}

</script>

<template>
  <div class="service-selector">
    <Button label="Recipes" icon="pi pi-apple" @click="ButtonAction('recipes')"/>
    <Button @click="ButtonAction('restaurant')" label="Restourant" icon="pi pi-briefcase" aria-label="Restourant"/>
    <Button @click="ButtonAction('suppliers')" label="Suppliers" icon="pi pi-truck" aria-label="Suppliers"/>
    <Button @click="ChangeApiUrl()" label="Switch Backend Endpoint" icon="pi pi-refresh" aria-label="Switch Backend"/>
  </div>
</template>

<style scoped>
.service-selector {
  Button {
    margin: 10px;
  }

}
</style>
