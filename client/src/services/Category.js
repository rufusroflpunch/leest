import Axios from './Axios'

export default class {
  findAll () {
    return Axios.get('/api/categories')
  }

  create (category) {
    return Axios.post('/api/categories', category)
  }

  delete (id) {
    return Axios.delete(`/api/categories/${id}`)
  }
}
