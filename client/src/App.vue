<template>
  <div id="app">
    <Navbar/>

    <div class="body">
      <router-view></router-view>
    </div>

    <footer class="hidden-sm">
      <p>&copy; 2017 Ryan Breen | <a href="mailto:ryan@breen.xyz">Contact</a></p>
    </footer>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar'
export default {
  name: 'App',
  components: {
    Navbar
  }
}
</script>

<style lang="scss">
@import '~mini.css/src/flavors/mini-nord.scss';

footer {
  position: fixed;
  bottom: 0;
  right: 0;
  left: 0;
}
</style>
