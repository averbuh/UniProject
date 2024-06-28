import axios from 'axios'

const url = 'https://api.averbuchpro.com'
// const url_dev = 'http://localhost:8080'
let response
export const Recipes = {

  getUrl(){
    return url
  },
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
