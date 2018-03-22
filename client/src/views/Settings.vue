<template>
  <div class="settings">
    <h1>Settings</h1>
    <button class="primary hidden-md hidden-lg add-category"
            @click.prevent="addNewCategory">+ Add Category</button>
    <table class="table">
      <thead>
        <tr>
          <th>Categories</th>
          <th><a href=""
                @click.prevent="addNewCategory">Add</a></th>
        </tr>
      </thead>
      <tbody>
        <tr v-if="addingCategory">
          <td colspan="2">
            <input type="text"
                   v-model="newCategoryName"
                   @keydown.enter.prevent="newCategory"
                   @keydown.esc.prevent="cancelNewCategory"
                   @blur.prevent="cancelNewCategory"
                   id="new-category">
          </td>
        </tr>
        <tr v-for="category in categories" :key="category.id">
          <td data-label="Name">{{ category.name }}</td>
          <td><a href="" @click.prevent="deleteCategory(category.id)">Delete</a></td>
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
      if (!confirm('Are you sure you want to delete the category?')) { return }
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

.add-category {
  margin-left: 1rem;
}
</style>
