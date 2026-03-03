import axios from "axios"

const BASE_URL = "https://jsonplaceholder.typicode.com";

export async function getPosts(){
    const response = await axios.get(`${BASE_URL}/posts`,{
        params:{_limit:10 },
    })
    return response.data;
}

export async function createPost(title,body){
    const response = await axios.post(`${BASE_URL}/posts`, {
        title, body, userId:1, 
    })
    return response.data;
}