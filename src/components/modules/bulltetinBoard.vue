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
        <ul class="article-box">
            <li v-for="data in articleItems">
                タイトル:{{data.title}}<br>
                本文:{{data.body}}
            </li>
        </ul>
    </div>
</template>

<script>
    import article from '../../service/article'
    import auth from"../../service/auth"

    export default {
        data () {
            return {
                items: [
                    {name: "編集", url: "", isActive: false},
                    {name: "削除", url: "", isActive: false}
                ],
                articleItems: []
            }
        },
        methods: {
            async show() {
                if (auth.getLoginToken() !== null) {
                    const articleData = await article.show(auth.getLoginToken())
                    articleData.data.forEach((data) => this.articleItems.push(data))
                }
                console.log(this.articleItems)
            },
            select(e) {
                for(let item of this.items) {
                    item.isActive = e.target.textContent === item.name
                }
            },
        },
        created () {
            this.show()
        }
    }
</script>

<style>
div.field {
    width: 80%;
    margin-left: 2.5%;
    float: left;
}
ul.article-box {
    margin-top: 20px;
}
ul.article-box li {
    margin-bottom: 10px;
}
</style>
