import Vue from 'vue'
import Vuex from 'vuex'

import CategoryService from '@/services/Category'
import ListService from '@/services/List'
import ListItemService from '@/services/ListItem'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    categorySerice: new CategoryService(),
    listService: new ListService(),
    listItemService: new ListItemService(),
    categories: [],
    lists: [],
    listItems: []
  },
  mutations: {
    setCategories (state, categories) {
      state.categories = categories
    },
    setLists (state, lists) {
      state.lists = lists
    },
    setListItems (state, listItems) {
      state.listItems = listItems
    }
  },
  actions: {
    async retrieveCategories ({ commit, state }) {
      let resp = await state.categorySerice.findAll()
      commit('setCategories', resp.data)
    },
    async retrieveLists ({ state, commit }) {
      let resp = await state.listService.findAll()
      commit('setLists', resp.data)
    },
    async retrieveListItems ({ state, commit }) {
      let resp = await state.listItemService.findAll()
      commit('setListItems', resp.data)
    }
  }
})
