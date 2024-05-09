<script setup>
import ServiceSelector from './components/ServiceSelector.vue'
import { ref, onMounted, onUpdated } from 'vue';


//default component
let activeComp = ref('recipes');


function changeComp(value) {
  activeComp.value = value
}

// lifecycle hooks
// onMounted(() => {
// })


// onUpdated(() => {

// })
</script>
<template>
  <div class="service-selector">
    <ServiceSelector :ButtonAction="changeComp" />
  </div>
  <header>
  </header>

  <main>
    <component :is="activeComp"></component>
  </main>
</template>

<style scoped>
.service-selector {
  display: flex;
  justify-content: center;
}

/* @media (min-width: 1024px) { */
/*   header { */
/*     display: flex; */
/*     place-items: center; */
/*     padding-right: calc(var(--section-gap) / 2); */
/*   } */

/*   .logo { */
/*     margin: 0 2rem 0 0; */
/*   } */

/*   header .wrapper { */
/*     display: flex; */
/*     place-items: flex-start; */
/*     flex-wrap: wrap; */
/*   } */
/* } */
</style>
