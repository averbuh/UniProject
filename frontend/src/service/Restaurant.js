import axios from 'axios'

const url = 'http://afd1df950f69b4a29b47d82a5a1cd9f4-856604395.eu-central-1.elb.amazonaws.com'
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
