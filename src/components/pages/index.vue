<template>
  <div class="top">
    <SideBar/>
    <BulltetinBoard/>
  </div>
</template>

<script>
import SideBar from '../modules/sidebar.vue'
import BulltetinBoard from '../modules/bulltetinBoard.vue'

export default {
  components: {
    SideBar,
    BulltetinBoard
  }
}
</script>

<style>
div.top {
  overflow: hidden;
}
</style>
