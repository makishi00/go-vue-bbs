<template>
  <section id="sidebar">
    <h1>Title</h1>
    <ul @click="select">
      <li v-for="item in items" :class="{ active: item.isActive }">{{item.name}}</li>
    </ul>
    <div class="container">
      <div class="profile">
        <div class="img">
          <img src="" alt="icon">
        </div>
        <p class="username">User Name</p>
      </div>
      <p class="user-info">ユーザー情報</p>
      <p class="logout">ログアウト</p>
    </div>
  </section>
</template>

<script>
export default {
  data() {
    return {
      items: [
        { name: "content-1", isActive: true },
        { name: "content-2", isActive: false },
        { name: "content-3", isActive: false },
        { name: "content-4", isActive: false },
        { name: "content-5", isActive: false }
      ]
    }
  },
  components: {
  },
  methods: {
    select(e) {
      for(let item of this.items) {
        item.isActive = e.target.textContent === item.name
      }
    }
  }
}
</script>

<style lang="scss" scoped>
#sidebar {
  background-color: #272727;
  width: 15%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  box-shadow: 5px 0px 15px black;
  h1 {
    font-size: 1.2rem;
    font-weight: 100;
    letter-spacing: 1.5px;
    color: white;
    padding: 15px 30px;
  }
  ul {
    width: 100%;
    color: white;
    font-size: 13px;
    cursor: pointer;
    li {
      padding: 15px 30px;
      list-style: none;
      &.active {
        background-color: rgba(255,255,255,0.15);
      }
    }
  }
  .container {
    margin-top: auto;
    color: white;
    font-size: 13px;
    padding: 0 30px;
    .profile {
      width: 100%;
      .img {
        width: 50%;
        overflow: hidden;
        border-radius: 50%;
        border: 1px solid #979797;
        img {
          display: block;
          width: 100%;
        }
      }
      .username {
        margin: 20px 0;
      }
    }
    .user-info {
      cursor: pointer;
    }
    .logout {
      cursor: pointer;
      margin: 5px 0;
    }
  }
}
</style>

