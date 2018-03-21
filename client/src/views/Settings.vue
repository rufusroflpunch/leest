<template>
  <div class="settings">
    <h1>Settings</h1>
    <table class="table">
      <thead>
        <tr>
          <th>Categories</th>
          <th><a href=""
                @click.prevent="addNewCategory">Add</a></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="category in categories" :key="category.id">
          <td>{{ category.name }}</td>
          <td><a href="" @click.prevent="deleteCategory(category.id)">Delete</a></td>
        </tr>
        <tr>
          <td colspan="2" v-if="addingCategory">
            <input type="text"
                   v-model="newCategoryName"
                   @keydown.enter.prevent="newCategory"
                   @keydown.esc.prevent="cancelNewCategory"
                   @blur.prevent="cancelNewCategory"
                   id="new-category">
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import CategoryService from '@/services/Category'

export default {
  name: 'Settings',
  categoryService: new CategoryService(),
  data () {
    return {
      categories: [],
      addingCategory: false,
      newCategoryName: ''
    }
  },
  methods: {
    deleteCategory (id) {
      this.$options.categoryService.delete(id).then(_ => this.retrieveData())
    },
    addNewCategory () {
      this.addingCategory = true
      this.$nextTick(_ => document.getElementById('new-category').focus())
    },
    newCategory () {
      this.addingCategory = false
      this.$options.categoryService.create({name: this.newCategoryName}).then(_ => this.retrieveData())
      this.newCategoryName = ''
    },
    cancelNewCategory () {
      this.addingCategory = false
      this.newCategoryName = ''
    },
    retrieveData () {
      this.$options.categoryService.findAll().then(val => { this.categories = val.data })
    }
  },
  created () {
    this.retrieveData()
  }
}
</script>

<style lang="scss">
.table {
  margin-left: 1rem;
  width: 20rem;

  th {
    text-align: left;
  }
}
</style>
