<template>
    <div class="field">
        <h1 id="ss">掲示板</h1>
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
            <li :id="index" v-show="data.UserID" v-for="(data, index) in articleItems">
                投稿日:{{data.created_at}}<br>
                <p id="article-title">タイトル:{{data.title}}</p>
                <p id="article-comment">コメント:{{data.body}}</p>
                <form action="" v-show="userId == data.UserID" class="edit">
                    <label class="label">タイトル</label>
                    <input class="input" type="text" :value="data.title" required>
                    <label class="label">コメント</label>
                    <textarea class="textarea" :value="data.body" required></textarea>
                    <button class="button is-link" :id="index" v-on:click="edit(index, data.id)">編集する</button>
                </form>
                <form v-show="userId == data.UserID">
                    <button class="button is-link" v-on:click="del(data.id)">削除</button>
                </form>
                <button v-show="userId == data.UserID" class="button is-link open" v-on:click="editForm(index, true)" :data-num="data.id">編集</button>
                <button v-show="userId == data.UserID" class="button is-link close" v-on:click="editForm(index, false)" >閉じる</button>
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
                articleItems: [{flg:false}],
                title: "",
                body: "",
                userId: auth.getUserIDByToken(auth.getLoginToken()),
                id: 0,
            }
        },
        created () {
            this.show()
        },
        methods: {
            async add() {
                await article.add(auth.getLoginToken(), this.title, this.body)
            },
            async show() {
                if (auth.getLoginToken() !== null) {
                    const articleData = await article.show(auth.getLoginToken())
                    articleData.data.forEach((data) => this.articleItems.push(data))
                }
            },
            async del(id) {
                const response = await article.delete(auth.getLoginToken(), id)
                alert(await response.message)
            },
            async edit(index, id) {
                const tag = document.getElementById(index)
                const title = tag.children[3].getElementsByTagName('input')[0].value
                const body = tag.children[3].getElementsByTagName('textarea')[0].value
                const response = await article.edit(auth.getLoginToken(), id, title, body)
                alert(await response.message)
            },
            select(e) {
                for(let item of this.items) {
                    item.isActive = e.target.textContent === item.name
                }
            },
            editForm: function (index, flg) {
                const tag = document.getElementById(index)
                if (flg) {
                    tag.children[1].style.display = 'none'
                    tag.children[2].style.display = 'none'
                    tag.getElementsByClassName('edit')[0].style.display = 'block'
                    tag.getElementsByClassName('open')[0].style.display = 'none'
                    tag.getElementsByClassName('close')[0].style.display = 'block'
                } else {
                    tag.children[1].style.display = 'block'
                    tag.children[2].style.display = 'block'
                    tag.getElementsByClassName('edit')[0].style.display = 'none'
                    tag.getElementsByClassName('open')[0].style.display = 'block'
                    tag.getElementsByClassName('close')[0].style.display = 'none'
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
ul.article-box {
    margin-top: 20px;
}
ul.article-box li {
    margin-bottom: 10px;
}
.close {
    display: none;
}
.edit {
    display: none;
}
</style>
