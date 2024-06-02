import axios from 'axios'

const url_prod = 'http://recipes-service.apps.svc.cluster.local:80'
const url = 'http://localhost:8080'
let response
export const Recipes = {
  async getRecipesData() {
    response = await axios.get(url + '/recipes')
    return response.data
  },

  async addRecipe(recipe) {
    response = await axios.post(url + '/recipes', recipe)
    return response
  },

  async deleteRecipe(name) {
    response = await axios.delete(url + `/recipes/${name}`)
    return response
  },

  async updateRecipe(name, recipe) {
    response = await axios.put(url + `/recipes/${name}`, recipe)
    return response
  },

  getRecipes() {
    return Promise.resolve(this.getRecipesData())
  },

  getTodayRecipes() {
    return Promise.resolve(
      this.getRecipesData().then((recipes) => recipes.filter((recipe) => recipe.istoday === true))
    )
  },

  async getImageUrl(image) {
    response = await axios.get(url + `/recipes/image/${image}`)
    return response.data
  }
}
