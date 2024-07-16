import axios from 'axios'

export let url 

export const Suppliers = {

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

  async getSuppliersData() {
    const response = await axios.get(url + '/suppliers')
    return response.data
  },

  getSuppliers() {
    return Promise.resolve(this.getSuppliersData());
  },

  async getFavouriteSuppliers() {
    try {
      const suppliers= await this.getSuppliers();
  
      // Check if recipes is an object
      if (typeof suppliers !== 'object' || suppliers === null) {
        throw new TypeError('Expected an object of suppliers');
      }
  
      // Convert object to array
      const suppliersArray = Object.values(suppliers);
  
      // Filter today's recipes
      return suppliersArray.filter(supplier=> supplier.isfavorite === true);
    } catch (error) {
      console.error('Error fetching fav suppliers:', error);
      throw error;
    }
  }, 
}

