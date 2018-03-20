<template>
  <div class="categories container rounded shadow-small">
    <div class="row">
      <div class="col-sm-12 col-md-3 col-lg-2" v-for="cat in categories" :key="cat.id" :class="{'active': cat === value}">
        <icon name="angle-right" v-if="cat.id === value.id" class="bigger"/>
        <a href="" class="category-button" :class="{'active': cat === value}" @click.prevent="$emit('input', cat)">
          {{ cat.name }}
        </a>
        <icon name="angle-left" v-if="cat.id === value.id" class="bigger"/>
      </div>
    </div>
  </div>
</template>

<script>
import 'vue-awesome/icons/angle-right'
import 'vue-awesome/icons/angle-left'

export default {
  name: 'Categories',
  props: {
    categories: {
      type: Array,
      default: _ => []
    },
    value: Object
  }
}
</script>

<style lang="scss" scoped>
.categories {
  background-color: darken(#647faf, 20%);
}
.category-button {
  font-size: 120%;
  font-weight: bold;
  text-decoration: none;
}
.active {
  // border: 1px solid white;
  color: white;
}
.bigger {
  color: white;
}

</style>
