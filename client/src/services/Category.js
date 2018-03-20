import Axios from './Axios'

export default class {
  findAll () {
    return Axios.get('/api/categories')
  }
}
