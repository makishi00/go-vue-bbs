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
}

let article;
export default article = new Article();