import axios from 'axios'



export let url
// const url_dev = 'http://localhost:8080'
let response
export const Recipes = {

  changeURL() {
    if (url === 'https://prod.api.averbuchpro.com') {
      url = 'https://stage.api.averbuchpro.com'
    } 
    else if (url === 'http://stage.api.averbuchpro.com') {
      url = 'https://prod.api.averbuchpro.com'
    }
    else {
      url = 'https://prod.api.averbuchpro.com'
    } 
  },

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
  
      // Check if recipes is an object
      if (typeof recipes !== 'object' || recipes === null) {
        throw new TypeError('Expected an object of recipes');
      }
  
      // Convert object to array
      const recipesArray = Object.values(recipes);
  
      // Filter today's recipes
      return recipesArray.filter(recipe => recipe.istoday === true);
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
