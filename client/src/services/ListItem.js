import Axios from './Axios'

export default class {
  findAll () {
    return Axios.get('/api/list_items')
  }

  update (listItem) {
    return Axios.put(`/api/list_items/${listItem.id}`, listItem)
  }

  create (listItem) {
    return Axios.post('/api/list_items', listItem)
  }

  toggleDone (id) {
    return Axios.put(`/api/list_items/${id}/toggle_done`)
  }

  delete (id) {
    return Axios.delete(`/api/list_items/${id}`)
  }
}
