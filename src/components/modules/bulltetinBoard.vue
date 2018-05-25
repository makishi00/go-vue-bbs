<template>
    <div class="field">
        <h1>掲示板</h1>
        <li v-for="item in items" :class="{ active: item.isActive }">
            <a :href=item.url>
                {{ item.name }}
            </a>
        </li>
        <h2>投稿フォーム</h2>
        <form action="">
            <label class="label">タイトル</label>
            <input class="input" type="text" v-model="title" required>
            <label class="label">コメント</label>
            <textarea class="textarea" v-model="body" required></textarea>
            <button class="button is-link" v-on:click="add">投稿</button>
        </form>
        <ul class="article-box">
            <li v-for="data in articleItems">
                タイトル:{{data.title}}<br>
                コメント:{{data.body}}<br>
                投稿日:{{data.created_at}}
                <form v-show="userId == data.UserID">
                    <button class="button is-link" v-on:click="del(data.id)">削除</button>
                </form>
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
                articleItems: [],
                title: "",
                body: "",
                userId: auth.getUserIDByToken(auth.getLoginToken()),
                id: 0
            }
        },
        created () {
            this.show()
        },
        methods: {
            async add() {
                await article.add(auth.getLoginToken(), this.title, this.body)
            },
            async del(id) {
                const response = await article.delete(auth.getLoginToken(), id)
                alert(await response.message)
            },
            async show() {
                if (auth.getLoginToken() !== null) {
                    const articleData = await article.show(auth.getLoginToken())
                    articleData.data.forEach((data) => this.articleItems.push(data))
                }
            },
            select(e) {
                for(let item of this.items) {
                    item.isActive = e.target.textContent === item.name
                }
            },
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
