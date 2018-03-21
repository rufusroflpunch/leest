<template>
  <div class="home">
    <div class="row">
      <div class="category col-sm-12">
        <Categories :categories="categories" v-model="currentCategory" />
      </div>
    </div>
    <div class="row">
      <div class="col-sm-12 col-md-4">
        <CategoriesNav :lists="categoryLists"
                       @createList="createList"
                       @deleteList="deleteList"
                       @selectList="selectList" />
      </div>
      <div class="middle-spacer hidden-sm"></div>
      <div class="col-sm-12 col-md">
        <List :list="currentList"
              :listItems="currentListItems"
              @saveListItem="saveListItem"
              @createNewItem="createNewItem"
              @toggleDone="toggleDone"
              @deleteListItem="deleteListItem" />
      </div>
    </div>
    <div class="row hidden-md bottom-spacer"></div>
  </div>
</template>

<script>
import Categories from '@/components/Categories'
import List from '@/components/List'
import CategoriesNav from '@/components/CategoriesNav'
import CategoryService from '@/services/Category'
import ListService from '@/services/List'
import ListItemService from '@/services/ListItem'

export default {
  name: 'Lists',
  categoryService: new CategoryService(),
  listService: new ListService(),
  listItemService: new ListItemService(),
  data () {
    return {
      currentCategory: {},
      currentListId: -1,
      currentList: {},
      currentListItems: [{}],
      categories: [{initial: true}],
      lists: [{}],
      listItems: [{}],
    }
  },
  computed: {
    categoryLists () {
      return this.lists.filter(l => l.category_id === this.currentCategory.id)
    },
  },
  methods: {
    async retrieveData () {
      await this.retrieveCategories()
      await this.retrieveLists()
      await this.retrieveListItems()
    },
    async retrieveCategories () {
      let resp = await this.$options.categoryService.findAll()
      this.categories = resp.data
    },
    async retrieveLists () {
      let resp = await this.$options.listService.findAll()
      this.lists = resp.data
    },
    async retrieveListItems () {
      let resp = await this.$options.listItemService.findAll()
      this.listItems = resp.data
    },
    async createList (newListName) {
      let resp = await this.$options.listService.create(newListName, this.currentCategory.id)
      this.retrieveData()
    },
    async deleteList (id) {
      await this.$options.listService.delete(id)
      this.retrieveData()
      this.$router.push({name: 'lists'})
    },
    async saveListItem (listItem) {
      await this.$options.listItemService.update(listItem)
      await this.retrieveData()
      this.selectList(this.currentListId)
    },
    selectList (id) {
      this.currentListId = id
      this.currentList = this.lists.filter(l => l.id === id)[0] || {}
      this.currentListItems = this.listItems.filter(li => li.list_id === id)
    },
    async createNewItem (item) {
      await this.$options.listItemService.create(Object.assign(item, {list_id: this.currentListId}))
      await this.retrieveData()
      this.selectList(this.currentListId)
    },
    async toggleDone (id) {
      await this.$options.listItemService.toggleDone(id)
      await this.retrieveData()
      this.selectList(this.currentListId)
    },
    async deleteListItem (id) {
      await this.$options.listItemService.delete(id)
      await this.retrieveData()
      this.selectList(this.currentListId)
    }
  },
  mounted () {
    this.retrieveData()
  },
  watch: {
    categories (newVal, oldVal) {
      // If the new old value is an initial value, then that means the page is just loaded and we can
      // safely set the active category to the first category available.
      if (oldVal[0].initial) {
        this.currentCategory = newVal[0]
      }
    }
  },
  components: {
    Categories,
    List,
    CategoriesNav
  }
}
</script>

<style lang="scss" scoped>
.bottom-spacer {
  height: 5rem;
}

.middle-spacer {
  width: 2rem;
}
</style>
