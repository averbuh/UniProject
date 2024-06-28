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
    return response
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

  async getTodayRecipes() {
    const recipes = await this.getRecipesData();

    console.log(typeof recipes)

    // if (!Array.isArray(recipes)) {
    //   throw new Error('Recipes data is not an array.');
    // }
    const todayRecipes = recipes.filter((recipe) => recipe.istoday === true);
    return todayRecipes;
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
