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
    return response.data
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
    try {
      const recipes = await this.getRecipes();
  
      // Log the type and value of recipes for debugging
      console.log('Type of recipes:', typeof recipes);
      console.log('Recipes:', recipes);
  
      if (!Array.isArray(recipes)) {
        throw new TypeError('Expected an array of recipes');
      }
  
      return recipes.filter(recipe => recipe.istoday === true);
    } catch (error) {
      console.error('Error fetching today\'s recipes:', error);
      throw error;
    }
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
