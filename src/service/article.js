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
            await axios.post("/api/bbs/add",
                {
                    title: title,
                    body: body
                },
                {
                    headers: {
                        "Authorization": token
                    }
                }
            )
        } catch (e) {
            console.log(e)
        }
    }
}

export default new Article();