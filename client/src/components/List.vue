<template>
  <div class="list" v-if="list.id">
    <h1>{{ list.name }}</h1>
    <button class="primary" @click.prevent="newItem">+ Add Item</button>
    <h2 v-if="incompleteItems.length > 0">Incomplete</h2>
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
    <transition-group name="pop">
      <div v-for="item in incompleteItems" :key="item.id" class="item-list">
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
        <div class="delete-button"
             @click.prevent="deleteItem(item)">&times;</div>
      </div>
    </transition-group>
    <h2 v-if="completeItems.length > 0">Complete</h2>
    <transition-group name="pop">
      <div v-for="item in completeItems" :key="item.id" class="item-list">
        <VCheckbox v-model="item.done" @input="completeItem(item.id)" />
        <div class="item shadow-small"
            :id="`item_${item.id}`"
            :class="{'done': item.done}">{{ item.description }}</div>
        <div class="delete-button"
             @click.prevent="deleteItem(item)">&times;</div>
      </div>
    </transition-group>
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
    },
    deleteItem (item) {
      this.$emit('deleteListItem', item.id)
    }
  },
  computed: {
    incompleteItems () {
      return this.listItems.filter(li => !li.done)
    },
    completeItems () {
      return this.listItems.filter(li => li.done)
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

.item.done {
  text-decoration: line-through;
  color: rgba(#2e3440, 0.4);
}

.pop-enter-active {
  animation: pop 0.2s;
}

.pop-leave-active {
  animation: pop 0.2s reverse;
}

@keyframes pop {
  0% {
    opacity: 0;
    max-height: 0rem;
  }
  100% {
    opacity: 1;
    max-height: 5rem;
  }
}

.delete-button {
  font-size: 1.5rem;
  color: rgb(131, 0, 0);
  text-decoration: none;
  font-size: 2rem;
}

.delete-button:hover {
  color: rgb(200, 0, 0);
  cursor: pointer;
}
</style>
