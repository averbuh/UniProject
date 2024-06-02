<script setup>
import { Recipes } from '@/service/Recipes'
import { ref, onMounted, onUpdated } from 'vue'
const props = defineProps({
  recipe: Object
})

const emits = defineEmits(['change'])

const visibleDescription = ref(false)

const selected = ref(false)

const tempURL = ref('')

onMounted(() => {})

//change first letter to uppercase
function adaptString(string) {
  string = string.charAt(0).toUpperCase() + string.slice(1).split('-').join(' ')
  return string
}

function GetImageUrl() {
  console.log(typeof props.recipe.image)
  console.log(props.recipe.image)
  if (props.recipe.image == '') {
    tempURL.value = {
      url: 'https://st3.depositphotos.com/14514162/16904/i/450/depositphotos_169040394-stock-photo-isolated-blank-note-with-clip.jpg'
    }
  } else Recipes.getImageUrl(props.recipe.image).then((data) => (tempURL.value = data))
  console.log(tempURL.value)
}
</script>

<template id="recipe-item">
  <li @click.stop="visibleDescription = true" @click="GetImageUrl" class="recipe_item">
    <h1>{{ adaptString(recipe.name) }}</h1>
    <p>{{ props.tempURL }}</p>
    <h2>Ingredients:</h2>
    <ol>
      <li v-for="ingredient in recipe.ingredients" :key="ingredient">{{ ingredient }}</li>
    </ol>
    <!-- <input type="checkbox" id="checkbox" v-model="checked" /> -->
  </li>
  <label id="select_label" for="select">Select</label>
  <label id="today_label" for="checkbox">Today</label>
  <br />
  <Checkbox
    id="select"
    v-model="selected"
    :binary="true"
    @change="$emit('change', recipe.name, selected)"
  />
  <Checkbox
    id="today"
    v-model="recipe.istoday"
    :binary="true"
    @change="Recipes.updateRecipe(recipe.name, recipe)"
  />

  <Dialog
    v-model:visible="visibleDescription"
    modal
    header="Description"
    :style="{ width: '25rem' }"
    :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
  >
    <span class="p-text-secondary block mb-5">How to cook</span>
    <Textarea
      @change="Recipes.updateRecipe(recipe.name, recipe)"
      v-model="recipe.description"
      autoResize
      rows="5"
      cols="30"
    />
    <div class="card flex justify-content-center">
      <Image
        v-if="tempURL.url != ''"
        :key="tempURL"
        :src="tempURL.url"
        alt="Image"
        width="250"
        preview
      />
    </div>

    <!-- <Galleria :value="images" :responsiveOptions="responsiveOptions" :numVisible="5" containerStyle="max-width: 640px"> -->
    <!--   <template #item="slotProps"> -->
    <!--     <img :src="`@/assets/images`" :alt="'image'" style="width: 100%" /> -->
    <!--   </template> -->
    <!--   <template #thumbnail="slotProps"> -->
    <!--     <img :src="`@/assets/images`" :alt="'images'" /> -->
    <!--   </template> -->
    <!-- </Galleria> -->

    <div class="flex justify-content-end gap-2">
      <Button
        type="button"
        label="Cancel"
        severity="secondary"
        @click="visibleDescription = false"
      ></Button>
    </div>
  </Dialog>
</template>

<style scoped>
.recipe_item {
  font-size: 10px;
  border: 3px solid #73ad21;
  gap: 10px;
  padding: 10px;
  margin: 10px;
  list-style-type: none;
  flex-direction: column;
  align-items: center;
  display: list-item;
  margin: 10px 10px;
  line-height: 20px;

  border-radius: 10px;
  background-color: white;
  margin-top: 30px;
  padding: 15px;
  border-radius: 10px;
  cursor: pointer;

  ul {
    list-style-type: square;
    line-height: 20px;
  }
}

#today {
  margin-left: 10px;
  margin-right: 10px;
}

#select {
  margin-left: 10px;
  margin-right: 10px;
}

#select_label {
  margin-left: 6px;
  margin-right: 12px;
}

Button {
  margin-left: 10px;
  margin-right: 10px;
  padding: 3px;
  font-size: 15px;
  width: 100px;
}
</style>
