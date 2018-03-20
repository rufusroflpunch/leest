<template>
  <div class="checkbox"
       :class="{'done': value}"
       @click.prevent="$emit('input', !value)">
    <Icon name="check" />
  </div>
</template>

<script>
import 'vue-awesome/icons/check'

export default {
  name: 'VCheckbox',
  props: {
    value: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style lang="scss" scoped>
.checkbox {
  color: rgba(#2E3440, 0.3);

  svg {
    width: 1.5rem;
    height: 1.5rem;
  }
}

.checkbox.done {
  color: #2e3440;
}

.checkbox:hover:not(.done) {
  color: rgba(#2e3440, 0.6);
}
</style>
