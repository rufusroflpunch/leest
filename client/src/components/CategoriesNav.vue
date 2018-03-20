<template>
  <div class="categories-nav">
    <nav class="nav">
      <button class="primary" @click.prevent="addList">+ Add List</button>
      <div v-for="list in lists" :key="list.id">
        <a href="" class="close-button" @click.prevent="$emit('deleteList', list.id)">&times;</a>
        &nbsp;&nbsp;&nbsp;&nbsp;
        <a href="#" @click.prevent="selectList(list.id)">{{ list.name }}</a>
      </div>
      <input id="newList"
             type="text"
             class="new-item"
             v-model="newListName"
             v-show="newList"
             @keydown.esc.prevent="newList = false"
             @keydown.enter.prevent="createList">
    </nav>
  </div>
</template>

<script>
export default {
  name: 'CategoriesNav',
  data () {
    return {
      newListName: '',
      newList: false
    }
  },
  methods: {
    addList () {
      this.newList = true
      this.newListName = ''
      this.$nextTick(_ => document.getElementById('newList').focus())
    },
    createList () {
      this.$emit('createList', this.newListName)
      this.newList = false
      this.newListName = ''
    },
    selectList (id) {
      this.$emit('selectList', id)
    }
  },
  props: {
    lists: {
      type: Array,
      default: _ => []
    }
  }
}
</script>

<style lang="scss" scoped>
.nav {
  width: 80%;
}

.new-item {
  border: none;
  background-color: #ECEFF4;
}

.nav {
  a {
    display: inline;
  }

  .close-button {
    font-size: 1.5rem;
    color: rgb(131, 0, 0);
  }
}
</style>
