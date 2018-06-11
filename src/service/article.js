import axios from 'axios';

class Article {
    async show(token) {
        try {
            const response = await axios.get('/api/bbs/show', { headers: { Authorization: token } })
            console.log(response)
            return response
        } catch (e) {
            console.log(e)
        }
    }
    async add(token, title, body) {
        try {
            await axios.post("/api/bbs/add", { title: title, body: body }, { headers: { "Authorization": token } })
        } catch (e) {
            console.log(e)
        }
    }
    async delete(token, id) {
        try {
            const response = await axios.post("/api/bbs/delete", { id: id }, { headers: { "Authorization": token } })
            response.message = "投稿を削除しました"
            return response
        } catch (e) {
            return {message: "他人の投稿は削除できません"}
        }
    }
    async edit(token, id, title, body) {
        try {
            const response = await axios.post("/api/bbs/edit", { id: id, title:title, body:body }, { headers: { "Authorization": token } })
            response.message = "投稿を編集しました"
            return response
        } catch (e) {
            return {message: "他人の投稿は編集できません"}
        }
    }
}

export default new Article();