<script setup>
import RecipeItem from './RecipeItem.vue'
import { useToast } from 'primevue/usetoast'
import { ref, onUpdated, onBeforeMount } from 'vue'
import { Recipes } from '@/service/Recipes'
import axios from 'axios'

const recipes = ref([])

onBeforeMount(() => {
  Recipes.getRecipes().then((data) => (recipes.value = data))
})

onUpdated(() => {
  console.log('updated')
})

const status = ref(null)
const name = ref(null)
const ingredients_list = ref([])
const ingredient = ref(null)
const selections = ref([])
const description = ref(null)

function addIngredient() {
  ingredients_list.value.push(ingredient.value)
  ingredient.value = ''
}

// function GetImagesUrl(recipes) {
//   Recipes.getTempImagesUrl(recipes).then((data) => {
//     tempURL.value = data
//   })

// }

function addRecipe() {
  const data = {
    name: name.value,
    istoday: false,
    ingredients: ingredients_list.value,
    description: description.value,
    image: fileName.value
  }
  Recipes.addRecipe(data).then((response) => {
    if (response.status == 200) {
      toast.add({
        severity: 'info',
        summary: 'Success',
        detail: 'Recipe Added',
        life: 10000
      })
    } else {
      toast.add({
        severity: 'error',
        summary: 'Error',
        detail: 'Recipe Not Added',
        life: 10000
      })
    }
  })

  ingredients_list.value = []
  name.value = ''
  fileName.value = ''
  description.value = ''
  // window.location.reload();

  setTimeout(() => {
    Recipes.getRecipes().then((data) => (recipes.value = data))
  }, 200)
}

function deleteRecipe() {
  for (const name of selections.value) {
    setTimeout(() => {
      Recipes.deleteRecipe(name).then((response) => (response.status = status))
    })
  }
  selections.value = []
  setTimeout(() => {
    Recipes.getRecipes().then((data) => (recipes.value = data))
  }, 200)
}

const handleSelection = (name, selected) => {
  if (selected) {
    selections.value.push(name)
  }
}

const toast = useToast()
const fileName = ref('')

const onUpload = async (event) => {
  const fileUp = event.files[0]

  console.log(fileUp)

  const body = new FormData()
  body.append('file', fileUp)

  console.log(body)
  const imgUpload = await axios.post(Recipes.getUrl() + '/recipes/upload', body, {
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    onUploadProgress: (progressEvent) => {
      const { loaded, total } = progressEvent
    }
  })

  if (imgUpload) {
    console.log(imgUpload)
    toast.add({
      severity: 'info',
      summary: 'Success',
      detail: 'File Uploaded',
      life: 10000
    })
  } else {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'File Not Uploaded',
      life: 10000
    })
  }
  fileName.value = fileUp.name
}

// async handleFileUploading(file: any) {
//   // add your fileuploading logic to server and set data to the uploadedFile
//   this.uploadedFile.fileType = 'image' // uploaded file type
//   this.uploadedFile.fileUrl = 'http://localhost:8080/recipes/upload' // uploaded file url
//   this.uploadedFile.fileName = file.name // uploaded file name

//   await new Promise((resolve) => setTimeout(resolve, 2000))
// };
</script>

<template>
  <transition-group name="fade" tag="ul" class="recipes_list">
    <li v-for="recipe in recipes" :key="recipe">
      <recipe-item v-on:change="handleSelection" :recipe="recipe" />
    </li>
  </transition-group>
  <br />
  <Button id="delete-recipe" label="Delete" icon="pi pi-trash" @click="deleteRecipe()" />
  <br />
  recipes: {{ recipes }}
  <br />
  <div class="add-recipe">
    <Toast />
    <label for="name">Recipe's Name</label>
    <br />
    <InputText id="name" v-model="name" />
    <br />
    <label for="ingredients">Recipe's Ingredients</label>
    <br />
    <InputText id="ingredient" v-model="ingredient" @keyup.enter="addIngredient" />
    <br />
    {{ ingredients_list }}
    <br />
    <Button id="add-ingredient" label="Add Ingredient" icon="pi pi-plus" @click="addIngredient" />
    <br />
    <Textarea id="description" v-model="description" autoResize rows="5" cols="30" />
    <br />
    <div class="card">
      <FileUpload
        mode="basic"
        name="model[]"
        accept="image/*"
        :max-file-size="5000000"
        :auto="true"
        :custom-upload="true"
        choose-label="Upload image"
        @uploader="onUpload"
      />
    </div>
    <br />
    <Button id="add-recipe" label="Add Recipe" icon="pi pi-plus" @click="addRecipe" />
  </div>
</template>

<style scoped>
.recipes-top {
  margin-top: 15px;
  margin-bottom: 15px;
}

h1 {
  text-align: center;
}

#image-upload {
  margin: 10px 10px;
}

.add-recipe {
  margin: auto auto;

  max-width: 660px;

  min-width: 660px;

  background-color: var(--surface-card);
  padding: 2rem;
  border-radius: 10px;

  box-shadow:
    0 2px 1px -1px rgb(0 0 0 / 20%),
    0 1px 1px 0 rgb(0 0 0 / 14%),
    0 1px 3px 0 rgb(0 0 0 / 12%);

  #ingredient {
    margin-top: 15px;
    margin-bottom: 15px;
  }
}

.fade-enter-active,
.fade-leave-active {
  transition: all 1s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}

#description {
  font-size: 20px;
}

.recipes_list {
  /* list-style-type: none; */
  /* flex-direction: column; */
  /* align-items: center; */
  /* display: list-item; */
  /* margin: 10px 10px; */
  /* padding: 10px 10px; */
  /* border-radius: 10px; */
  /* background-color: white; */
  padding: 2rem;
  border-radius: 10px;
  margin-bottom: 1rem;
  list-style-type: none;

  box-shadow:
    0 2px 1px -1px rgb(0 0 0 / 20%),
    0 1px 1px 0 rgb(0 0 0 / 14%),
    0 1px 3px 0 rgb(0 0 0 / 12%);

  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  gap: 1rem;

  margin: 20px auto;

  text-align: left;

  transition: all 0.3s ease;
}

#delete-recipe {
  margin: 20px auto;

  text-align: left;
}

.add-recipe {
  border-radius: 10px;
  margin: 0 auto;
  float: center;
}

#add-recipe {
  margin-top: 15px;
  margin-bottom: 15px;
  margin-left: 10px;
  margin-right: 10px;
}

#add-ingredient {
  margin-top: 15px;
  margin-bottom: 15px;
  margin-left: 10px;
  margin-right: 10px;
}
</style>
