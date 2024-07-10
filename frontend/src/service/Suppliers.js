import axios from 'axios'

const url = 'https://api.averbuchpro.com'

export const Suppliers = {
  
  async getSuppliersData() {
    const response = await axios.get(url + '/suppliers')
    return response.data
  },

  getSuppliers() {
    return Promise.resolve(this.getSuppliersData());
  },

  getFavouriteSuppliers() {
    return Promise.resolve(this.getSuppliersData().then(suppliers => suppliers.filter(supplier => supplier.favourite === true)));
  }

}

