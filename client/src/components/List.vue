<template>
  <div class="list">
    <button class="primary" @click.prevent="newItem">+ Add Item</button>
    <div v-for="item in listItems" :key="item.id" class="item-list">
      <VCheckbox v-model="item.done" @input="completeItem(item.id)" />
      <div class="item shadow-small"
           v-show="editItem !== item.id"
           @click.prevent="startEdit(item)"
           :id="`item_${item.id}`">{{ item.description }}</div>
      <input class="edit-item"
             v-show="editItem === item.id"
             :value="item.description"
             @keydown.esc.prevent="cancelEdit(item)"
             @keydown.enter.prevent="saveEdit(item.id)"
             @blur.prevent="cancelEdit(item)"
             :id="`editItem${item.id}`">
    </div>
    <div>
      <input type="text"
             id="new-item"
             class="edit-item"
             v-show="addItem"
             v-model="addItemText"
             @keydown.esc.prevent="cancelNew"
             @keydown.enter.prevent="saveNew"
             @blur.prevent="cancelNew">
    </div>
  </div>
</template>

<script>
import VCheckbox from '@/components/VCheckbox'

export default {
  name: 'List',
  data () {
    return {
      editItem: -1,
      preEdit: '',
      addItem: false,
      addItemText: ''
    }
  },
  methods: {
    startEdit (item) {
      this.preEdit = item.description
      this.editItem = item.id
      this.$nextTick(_ => document.getElementById(`editItem${item.id}`).focus())
    },
    cancelEdit (item) {
      item.description = this.preEdit
      this.preEdit = ''
      this.editItem = -1
    },
    saveEdit (id) {
      let description = document.getElementById(`editItem${id}`).value
      this.preEdit = ''
      this.editItem = -1
      this.$emit('saveListItem', {id, description})
    },
    newItem () {
      this.addItem = true
      this.$nextTick(_ => document.getElementById('new-item').focus())
    },
    cancelNew () {
      this.addItem = false
      this.addItemText = ''
    },
    saveNew () {
      this.addItem = false
      this.$emit('createNewItem', {
        description: this.addItemText
      })
      this.addItemText = false
    },
    completeItem (id) {
      this.$emit('toggleDone', id)
    }
  },
  props: {
    list: {
      type: Object,
      default: _ => {}
    },
    listItems: {
      type: Array,
      default: _ => []
    }
  },
  components: {
    VCheckbox
  }
}
</script>

<style lang="scss" scoped>
.item {
  margin: 0.7rem;
  border: 1px solid rgba(#2E3440, 0.2);
  padding: 0.3rem;
}

.edit-item {
  margin: 0.7rem;
  width: 100%;
}

.item:hover {
  border: 1px solid #2E3440;
}

.input-group {
  display: inline;
}

.item-list {
  display: flex;
  align-items: center;
}

.item, .edit-item {
  flex-grow: 1;
}
</style>
