<template>
  <div class="top">
    <SideBar/>
    <BulletinBoard>
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
