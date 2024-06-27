import axios from 'axios'

const url = 'http://afd1df950f69b4a29b47d82a5a1cd9f4-856604395.eu-central-1.elb.amazonaws.com'
// const url_dev = 'http://localhost:8080'
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
  // async uploadImage(body) {
  //   response =await axios.post(url +'/recipes/upload', body, {
  //     headers: {
  //       'Content-Type': 'multipart/form-data'
  //     },
  //     onUploadProgress: (progressEvent) => {
  //       const { loaded, total } = progressEvent
  //     }
  //   })
  //   return response
  // }
}
