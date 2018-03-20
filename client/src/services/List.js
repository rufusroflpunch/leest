import Axios from './Axios'

export default class {
  findAll () {
    return Axios.get('/api/lists')
  }

  create (name, categoryId) {
    return Axios.post('/api/lists', {name, category_id: categoryId})
  }

  delete (id) {
    return Axios.delete(`/api/lists/${id}`)
  }

  update (list) {
    return Axios.put(`/api/lists/${list.id}`, list)
  }
}
