<template>
  <div class="categories-nav">
    <nav class="nav">
      <h3 class="header">Categories</h3>
      <div v-for="category in categories" :key="category.id">
        <a href="#"
           @click.prevent="selectCategory(category.id)"
           :class="{'selected': currentCategory.id === category.id}">{{ category.name }}</a>
      </div>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'CategoriesNav',
  data () {
    return {
      newCategoryName: '',
      newCategory: false
    }
  },
  methods: {
    selectCategory (id) {
      this.$emit('selectCategory', id)
    }
  },
  props: {
    categories: {
      type: Array,
      default: _ => []
    },
    currentCategory: {
      type: Object,
      default: _ => {}
    }
  }
}
</script>

<style lang="scss" scoped>
.nav {
  a {
    display: inline;
  }

  .close-button {
    font-size: 1.5rem;
    color: rgb(131, 0, 0);
  }
}

.header {
  margin-left: 0;
}

.selected {
  font-size: 1.5rem;
  font-weight: bold;
  color: darkslategrey;
}
</style>
