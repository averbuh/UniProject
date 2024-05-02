import axios from 'axios'

export const Suppliers = {
  
  async getSuppliersData() {
    const response = await axios.get('http://localhost:3000/suppliers')
    return response.data
  },

  getSuppliers() {
    return Promise.resolve(this.getSuppliersData());
  },

  getFavouriteSuppliers() {
    return Promise.resolve(this.getSuppliersData().then(suppliers => suppliers.filter(supplier => supplier.favourite === true)));
  }

}

