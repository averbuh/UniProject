<script setup>
defineProps({
  ButtonAction: {
    type: Function,
    required: true

  }
})
</script>

<template>
  <div class="service-selector">
    <Button label="Recipes" icon="pi pi-apple" @click="ButtonAction('recipes')"/>
    <Button @click="ButtonAction('restaurant')" label="Restourant" icon="pi pi-briefcase" aria-label="Restourant"/>
    <Button @click="ButtonAction('suppliers')" label="Suppliers" icon="pi pi-truck" aria-label="Suppliers"/>
  </div>
</template>

<style scoped>
.service-selector {
  Button {
    margin: 10px;
  }

}
</style>
