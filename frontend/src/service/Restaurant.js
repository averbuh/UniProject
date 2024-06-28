import axios from 'axios'

const url = 'https://api.averbuchpro.com'
// const url = 'http://localhost:8080'

export const Restaurant = {
  async getRestaurantsData() {
    const response = await axios.get(url + '/restaurant')
    return response.data
  },

  getRestaurants() {
    return Promise.resolve(this.getRestaurantsData());
  }
}
