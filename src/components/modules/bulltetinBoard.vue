<template>
    <div class="field">
        <h1>掲示板</h1>
        <li v-for="item in items" :class="{ active: item.isActive }">
            <a :href=item.url>
                {{ item.name }}
            </a>
        </li>
        <p>投稿フォーム</p>
        <label class="label">Name</label>
        <input class="input" type="text">
        <label class="label">コメント</label>
        <textarea class="textarea"></textarea>
        <button class="button is-link">Submit</button>
    </div>
</template>

<script>
export default {
  data () {
      return {
        items: [
            {name: "編集", url: "", isActive: false},
            {name: "削除", url: "", isActive: false}
        ]
    }
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

<style>
div.field {
    width: 80%;
    margin-left: 2.5%;
    float: left;
}
</style>
