<template>
  <div class="list-nav">
    <nav class="nav">
      <a href="" @click.prevent="$emit('hideListNav')"><icon name="arrow-left" scale="2" class="hidden-md hidden-lg back-button" /></a>
      <h3 class="header">Lists</h3>
      <button class="primary add-list-button" @click.prevent="addList">+ Add List</button>
      <div v-show="newList">
        <input id="newList"
              type="text"
              class="new-item"
              v-model="newListName"
              @keydown.esc.prevent="newList = false"
              @keydown.enter.prevent="createList">
      </div>
      <div v-for="list in lists" :key="list.id">
        <a href="" class="close-button" @click.prevent="$emit('deleteList', list.id)">&times;</a>
        &nbsp;&nbsp;&nbsp;&nbsp;
        <a href="#" @click.prevent="selectList(list.id)">{{ list.name }}</a>
      </div>
    </nav>
  </div>
</template>

<script>
import 'vue-awesome/icons/arrow-left'
export default {
  name: 'ListNav',
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

.back-button {
  color: lighten(#2e3440, 20%);
}

.back-button:hover {
  color: #2e3440;
}

.header {
  margin-left: 0;
}

.add-list-button {
  margin-left: 0;
}
</style>
